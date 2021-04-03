package main

import "net/http"

func main() {
	
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		
		writer.Write([]byte("Welcome to First Web Server!"))
	})

	if error:= http.ListenAndServe(":8080", nil); error != nil {
		panic(error)
	}
	
}