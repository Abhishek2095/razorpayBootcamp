package Order

type Order struct {
	ID         string `json:"id" gorm:"id"`
	CustomerId string `json:"customer_id" gorm:"customer_id"`
	ProductId  string `json:"product_id" gorm:"product_id"`
	Quantity   uint   `json:"quantity" gorm:"quantity"`
	Amount     uint   `json:"amount" gorm:"amount"`
}

func (order *Order) TableName() string {
	return "orders"
}
