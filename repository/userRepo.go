package repository

import "github.com/Fadhelbulloh/Management-Asset/model"

type Repository interface {
	FindAll(sortBy, sortType, search string) ([]model.User, error)
	FindByID(id string) (model.User, error)
	FindByUsername(username string) (model.User, error)
	UsernameChecker(username string) bool
	EmailChecker(email string) bool
	Insert(doc interface{}) (interface{}, error)
	Update(doc interface{}, id string) (interface{}, error)
	Delete(id string) (interface{}, error)
}
