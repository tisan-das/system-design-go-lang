package main

import (
	"log"
	"net/http"

	"github.com/tisan-das/system-design-go-lang/repository"
	"github.com/tisan-das/system-design-go-lang/router"
	"github.com/tisan-das/system-design-go-lang/service"
)

func main() {
	log.Print("hello world")

	dbConn, err := repository.NewPostgresDBConnection("localhost", "tisan", "tisan", "tisan_db", 5432)
	if err != nil {
		log.Fatalf("Error occurred connecting to DB: %s", err)
		return
	}

	groceryRepo := repository.NewGroceryRepo(dbConn)
	groceryService := service.NewGroceryService(groceryRepo)

	router := router.NewWebRouter(groceryService)
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
