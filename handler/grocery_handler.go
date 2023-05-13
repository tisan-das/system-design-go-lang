package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tisan-das/system-design-go-lang/entity"
	"github.com/tisan-das/system-design-go-lang/service"
)

type GroceryHandler interface {
	GetGrocery(http.ResponseWriter, *http.Request)
	GetAllGrocery(http.ResponseWriter, *http.Request)
	AddGrocery(http.ResponseWriter, *http.Request)
	DeleteGrocery(http.ResponseWriter, *http.Request)
	UpdateGrocery(http.ResponseWriter, *http.Request)
}

type groceryHandler struct {
	svc service.GroceryService
}

func NewGroceryHandler(groceryService service.GroceryService) GroceryHandler {
	return &groceryHandler{
		svc: groceryService,
	}
}

func (handler *groceryHandler) GetGrocery(response http.ResponseWriter, request *http.Request) {

	contextParams := mux.Vars(request)
	groceryId, exists := contextParams["grocery-id"]
	itemId, err := strconv.Atoi(groceryId)

	if !exists || err != nil {
		msg := "The parameter grocery-id is not found"
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	log.Printf("Request received to fetch details for the grocery Id: %d", itemId)
	groceryItem, err := handler.svc.GetGrocery(itemId)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching grocery with item id %d: %s", itemId, err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}
	log.Printf("Found grocery item: %+v", groceryItem)

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	var responseBytes bytes.Buffer
	json.NewEncoder(&responseBytes).Encode(groceryItem)
	response.Write(responseBytes.Bytes())
	log.Printf("Response created against grocery item id %d: %+v", itemId, groceryItem)
}

func (handler *groceryHandler) GetAllGrocery(response http.ResponseWriter, request *http.Request) {
	log.Print("Request received to fetch all the grocery items available")
	groceryItems, err := handler.svc.GetAllGrocery()
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching all grocery items: %s", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}
	log.Printf("Found all grocery items: %+v", groceryItems)

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	var responseBytes bytes.Buffer
	json.NewEncoder(&responseBytes).Encode(groceryItems)
	response.Write(responseBytes.Bytes())
	log.Printf("Response created for all the grocery items: %+v", groceryItems)
}

func (handler *groceryHandler) AddGrocery(response http.ResponseWriter, request *http.Request) {
	var grocery entity.Grocery
	err := json.NewDecoder(request.Body).Decode(&grocery)
	if err != nil {
		msg := fmt.Sprintf("Error occurred parsing grocery item: %s", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}
	log.Printf("Request recieved to add grocery item %+v", grocery)

	// Use validation module
	if grocery.Id <= 0 || grocery.Name == "" || grocery.Quantity < 0 {
		msg := "Invalid grocery attribute vaule provided!"
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	groceryItem, err := handler.svc.GetGrocery(grocery.Id)
	if err == nil && groceryItem.Id > 0 {
		msg := "A grocery already exists with this id"
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	log.Printf("Adding grocery item %+v", grocery)
	err = handler.svc.AddGrocery(grocery)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while adding grocery %+v: %s", grocery, err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

}

func (handler *groceryHandler) DeleteGrocery(response http.ResponseWriter, request *http.Request) {
	contextParams := mux.Vars(request)
	groceryId, exists := contextParams["grocery-id"]
	itemId, err := strconv.Atoi(groceryId)

	if !exists || err != nil {
		msg := "The parameter grocery-id is not found"
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	log.Printf("Request received to delete the grocery with Id: %d", itemId)
	err = handler.svc.DeleteGrocery(itemId)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while deleting grocery with item id %d: %s", itemId, err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func (handler *groceryHandler) UpdateGrocery(response http.ResponseWriter, request *http.Request) {
	var grocery entity.Grocery
	err := json.NewDecoder(request.Body).Decode(&grocery)
	if err != nil {
		msg := fmt.Sprintf("Error occurred parsing grocery item: %s", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}
	log.Printf("Request recieved to update grocery item %+v", grocery)

	// Use validation module
	if grocery.Id <= 0 || grocery.Name == "" || grocery.Quantity < 0 {
		msg := "Invalid grocery attribute vaule provided!"
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	groceryItem, err := handler.svc.GetGrocery(grocery.Id)
	if err != nil {
		msg := fmt.Sprintf("Error occurred fetching grocery details for grocery %+v: %s", groceryItem, err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

	log.Printf("Updating grocery item %+v", grocery)
	err = handler.svc.AddGrocery(grocery)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while adding grocery %+v: %s", grocery, err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(msg))
		return
	}

}
