package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chirag018/mongoapi/router"
)

func main() {
	fmt.Println("mongodb api")
	r := router.Router()
	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening to port 4000..")
}
