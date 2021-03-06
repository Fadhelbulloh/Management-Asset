package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Fadhelbulloh/Management-Asset/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	col = "user"
)

func ConnectMongo(host, username, password string) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s/?connect=direct", host)
	opt := options.Client().ApplyURI(url)

	if username != "" && password != "" {
		opt.SetAuth(options.Credential{Username: username, Password: password})
	}

	return mongo.Connect(context.Background(), opt)
}

type repo struct {
	db *mongo.Database
}

func NewMongoRepo(client *mongo.Client) *repo {
	return &repo{db: client.Database(os.Getenv("MONGO_DB"))}
}

func (rep *repo) FindAll(sortBy, sortType, search string) ([]model.User, error) {
	if sortBy == "" {
		sortBy = "_id"
	}

	if sortType == "" {
		sortType = "desc"
	}

	mapOrder := map[string]int{"asc": 1, "desc": -1}

	opt := options.Find().SetSort(bson.M{sortBy: mapOrder[strings.ToLower(sortType)]}).SetProjection(bson.M{"password": 0})

	var filters []bson.M
	filter := bson.M{}

	if search != "" {
		pattern := fmt.Sprintf(".*%s.*", search)
		filters = append(filters, bson.M{"username": primitive.Regex{Pattern: pattern, Options: "i"}})
	}

	if len(filters) > 0 {
		filter = bson.M{"$and": filters}
	}

	res, err := rep.db.Collection(col).Find(context.Background(), filter, opt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []model.User
	if err = res.All(context.Background(), &users); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (rep *repo) FindByID(id string) (model.User, error) {
	var user model.User
	err := rep.db.Collection(col).FindOne(context.Background(), bson.M{"_id": id}, options.FindOne().SetProjection(bson.M{"password": 0})).
		Decode(&user)
	return user, err
}

func (rep *repo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := rep.db.Collection(col).FindOne(context.Background(), bson.M{"email": email}, options.FindOne().SetProjection(bson.M{"password": 0})).
		Decode(&user)
	return user, err
}

func (rep *repo) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := rep.db.Collection(col).FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func (rep *repo) FindByUsernameAndPassword(username, password string) (model.User, error) {
	var user model.User
	err := rep.db.Collection(col).
		FindOne(context.Background(), bson.M{"$and": bson.A{bson.M{"username": username}, bson.M{"password": password}}}).
		Decode(&user)
	return user, err
}

func (rep *repo) UsernameChecker(username string) bool {
	err := rep.db.Collection(col).FindOne(context.Background(), bson.M{"username": username}).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (rep *repo) EmailChecker(email string) bool {
	err := rep.db.Collection(col).FindOne(context.Background(), bson.M{"email": email}).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (rep *repo) Insert(doc interface{}) (interface{}, error) {
	return rep.db.Collection(col).InsertOne(context.Background(), doc)
}

func (rep *repo) Update(doc interface{}, id string) (interface{}, error) {
	return rep.db.Collection(col).UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": doc})
}

func (rep *repo) Delete(id string) (interface{}, error) {
	return rep.db.Collection(col).DeleteOne(context.Background(), bson.M{"_id": id})
}
