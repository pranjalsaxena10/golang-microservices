package dao

import (
	"fmt"

	"github.com/pranjalsaxena10/golang-microservices/mvc/model"
)

var (
	usersMap = map[int64]*model.User {
		123: {Id: 123, FirstName: "Ted", LastName: "Mosby", Email: "ted_evelyn_mosby@outlook.com"},
	}
)


func GetUser(userId int64) (*model.User, error) {

	if user := usersMap[userId] ; user != nil {
		return user, nil
	}

	return nil, fmt.Errorf("user %v is not present", userId)
}