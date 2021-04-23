package service

import (
	"errors"
	"log"
	"time"

	"github.com/Fadhelbulloh/Management-Asset/model"
	"github.com/Fadhelbulloh/Management-Asset/repository"

	"github.com/renstrom/shortuuid"
	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	Login(username, password string) (map[string]interface{}, error)
	Register(user model.User) (interface{}, error)
	Update(user model.User) (interface{}, error)
	Delete(id string) (interface{}, error)
	FindAll(sortType, sortBy, search string, from, limit int) ([]model.User, error)
	FindByID(id string) (model.User, error)
}

type service struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) Services {
	return &service{repo: repo}
}
func (srv *service) Register(user model.User) (interface{}, error) {
	if user.Password == "" || user.Email == "" || user.Username == "" {
		return nil, errors.New("username, password, and email cannont be empty")
	}

	if srv.repo.UsernameChecker(user.Username) {
		return nil, errors.New("username already exists")
	}

	if srv.repo.EmailChecker(user.Email) {
		return nil, errors.New("email already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	currentTime := time.Now()

	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.ID = shortuuid.New()

	return srv.repo.Insert(user)
}

func (srv *service) Login(username, password string) (map[string]interface{}, error) {
	user, err := srv.repo.FindByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Username and password doesn't match ", err)
		return nil, err
	}
	return map[string]interface{}{"user": user}, nil
}

func (srv *service) Update(user model.User) (interface{}, error) {
	user.UpdatedAt = time.Now()
	return srv.repo.Update(user, user.ID)
}

func (srv *service) Delete(id string) (interface{}, error) {
	return srv.repo.Delete(id)
}

func (srv *service) FindAll(sortType, sortBy, search string, from, limit int) ([]model.User, error) {
	users, err := srv.repo.FindAll(sortBy, sortType, search)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	users = paginate(users, from, limit)

	return users, nil
}

func (srv *service) FindByID(id string) (model.User, error) {
	return srv.repo.FindByID(id)
}

func paginate(users []model.User, from, limit int) []model.User {
	pg := from * limit
	lmtPg := pg + limit

	if len(users) > lmtPg {
		users = users[pg:lmtPg]
	} else if len(users) >= pg {
		users = users[pg:]
	}

	return users
}
