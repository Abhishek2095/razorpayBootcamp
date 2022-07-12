package Models

type Student struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	MarksList   []Marks
}

func (s *Student) Tablename() string {
	return "students"
}
