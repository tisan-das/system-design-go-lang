package router

import (
	"github.com/gorilla/mux"
	"github.com/tisan-das/system-design-go-lang/handler"
	"github.com/tisan-das/system-design-go-lang/service"
)

func NewWebRouter(groceryService service.GroceryService) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")

	groceryHandler := handler.NewGroceryHandler(groceryService)
	router.HandleFunc("/grocery/{grocery-id}", groceryHandler.GetGrocery).Methods("GET")
	router.HandleFunc("/grocery/", groceryHandler.GetAllGrocery).Methods("GET")

	router.HandleFunc("/grocery/", groceryHandler.AddGrocery).Methods("POST")
	router.HandleFunc("/grocery/{grocery-id}", groceryHandler.DeleteGrocery).Methods("DELETE")
	router.HandleFunc("/grocery/", groceryHandler.UpdateGrocery).Methods("PUT")

	return router
}
