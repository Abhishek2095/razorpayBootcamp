package Product

type Product struct {
	ID       string `json:"id" gorm:"id"`
	Name     string `json:"name" gorm:"name"`
	Quantity uint   `json:"quantity" gorm:"quantity"`
	Price    uint   `json:"price" gorm:"price"`
}

func (product Product) TableName() string {
	return "products"
}
