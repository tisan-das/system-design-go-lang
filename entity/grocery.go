package entity

type Grocery struct {
	Id       int    `json:"id" gorm:"primaryKey;column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Quantity int    `json:"quantity" gorm:"column:quantity"`
}

func (grocery *Grocery) TableName() string {
	return "sample.grocery"
}
