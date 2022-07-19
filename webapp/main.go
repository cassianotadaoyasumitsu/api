package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/utils"
)

func main() {
	r := router.Gerar()
	utils.CarregarTemplates()

	fmt.Println("Rodando Webapp {port:5000}")
	log.Fatal(http.ListenAndServe(":5000", r))
}
