package main

import (
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.InitRouter()

	fmt.Println("API running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
