package Models


type Marks struct {
	Id			string  `json:"id"`
	StudentId	string	`json:"student_id"`
	Subject    	string `json:"subject"`
	Marks		int	`json:"marks" binding:"gte=0,lte=100"`
}

func (m *Marks) TableName() string {
	return "marks"
}
