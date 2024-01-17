package main

import (
	"devbook/src/application/router"
	"devbook/src/infra/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	r := router.InitRouter()

	fmt.Printf("API running on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
