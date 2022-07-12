package Models

type Marks struct {
	StudentId uint   `json:"student_id" gorm:"foreign_key"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
}

func (m *Marks) Tablename() string {
	return "marks"
}
