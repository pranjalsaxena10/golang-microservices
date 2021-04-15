package routes

import (
	"net/http"
	"github.com/pranjalsaxena10/golang-microservices/mvc/controller"
)

func StartApplicationRoutes() {
	http.HandleFunc("/users", controller.GetUser)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}