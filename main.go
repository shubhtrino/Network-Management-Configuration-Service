package main

import (
	"anchor_exercise/rest_api"
	"fmt"
)

func main() {
	server := rest_api.HandleRequests()
	fmt.Println("Starting Server on port 8080")
	server.ListenAndServe(":8080")
}
