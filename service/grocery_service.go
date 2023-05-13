package service

import (
	"github.com/tisan-das/system-design-go-lang/entity"
	"github.com/tisan-das/system-design-go-lang/repository"
)

type GroceryService interface {
	GetGrocery(int) (entity.Grocery, error)
	GetAllGrocery() ([]entity.Grocery, error)
	UpdateGrocery()
	DeleteGrocery(int) error
	AddGrocery(entity.Grocery) error
}

type groceryService struct {
	groceryRepo repository.GroceryRepository
}

func NewGroceryService(groceryRepository repository.GroceryRepository) GroceryService {
	return &groceryService{
		groceryRepo: groceryRepository,
	}
}

func (svc *groceryService) GetGrocery(id int) (entity.Grocery, error) {
	return svc.groceryRepo.Find(id)
}

func (svc *groceryService) GetAllGrocery() ([]entity.Grocery, error) {
	return svc.groceryRepo.FindAll()
}

func (svc *groceryService) UpdateGrocery() {
	// return nil
}

func (svc *groceryService) DeleteGrocery(id int) error {

	groceryItem, err := svc.groceryRepo.Find(id)
	if err != nil {
		return err
	}
	err = svc.groceryRepo.Delete(groceryItem)
	return err
}

func (svc *groceryService) AddGrocery(grocery entity.Grocery) error {
	return svc.groceryRepo.Save(grocery)
}
