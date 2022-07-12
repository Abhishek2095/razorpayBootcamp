package Models

type Student struct {
	Id         string   `json:"id" gorm:"primary_key"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	DOB        string `json:"DOB"`
	Address    string `json:"address"`
}

func (s *Student) TableName() string {
	return "student"
}
