package Customer

type Customer struct {
	ID      string `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	Address string `json:"address" gorm:"address"`
}

func (customer Customer) Tablename() string {
	return "customers"
}
