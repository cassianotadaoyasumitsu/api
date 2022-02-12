package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Listening :%d", config.APIDoor)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIDoor), r))
}
