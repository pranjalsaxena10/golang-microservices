package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pranjalsaxena10/golang-microservices/mvc/service"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {

		// Returning Bad Request to client
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("user_id must be a number"))

		return 
	}

	user, err := service.GetUser(userId)

	if err != nil {
		// Handle the error here
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))

		return
	}

	userJson, _ := json.Marshal(user)

	response.Write(userJson)



}