package service

import (
	"github.com/pranjalsaxena10/golang-microservices/mvc/dao"
	"github.com/pranjalsaxena10/golang-microservices/mvc/model"
)

func GetUser(userId int64) (*model.User, error) {
	return dao.GetUser(userId)
}