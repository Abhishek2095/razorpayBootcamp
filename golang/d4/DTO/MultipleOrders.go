package DTO

type MultipleOrder struct {
	CustomerID string   `json:"customer_id"`
	ProductID  []string ` json:"product_id"`
	Quantity   []uint   `json:"quantity"`
}
