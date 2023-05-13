package repository

import (
	"github.com/tisan-das/system-design-go-lang/entity"
	"gorm.io/gorm"
)

type GroceryRepository interface {
	Save(entity.Grocery) error
	Delete(entity.Grocery) error
	FindAll() ([]entity.Grocery, error)
	Find(groceryId int) (entity.Grocery, error)
}

type groceryRepo struct {
	DB *gorm.DB
}

func NewGroceryRepo(db *gorm.DB) GroceryRepository {
	return &groceryRepo{
		DB: db,
	}
}

func (repo *groceryRepo) Save(grocery entity.Grocery) error {
	result := repo.DB.Save(&grocery)
	return result.Error
}

func (repo *groceryRepo) Delete(grocery entity.Grocery) error {
	result := repo.DB.Delete(&grocery)
	return result.Error
}

func (repo *groceryRepo) FindAll() ([]entity.Grocery, error) {

	var groceryList []entity.Grocery
	result := repo.DB.Find(&groceryList)
	return groceryList, result.Error
}

func (repo *groceryRepo) Find(groceryId int) (entity.Grocery, error) {
	var grocery entity.Grocery
	result := repo.DB.Where("id = ?", groceryId).First(&grocery)
	return grocery, result.Error
}
