package Services

import (
	"fmt"
	"studentAPI/Models"
	"testing"
)

func TestValidateStudent(t *testing.T) {
	cases := []struct {
		in   Models.Student
		want error
	}{
		{
			Models.Student{
				Id:         "1",
				First_name: "abcd",
				Last_name:  "efgh",
				DOB:        "2005-12-10",
				Address:    "Street 5",
			},
			nil,
		},
		{
			Models.Student{
				Id:         "-1",
				First_name: "abcd",
				Last_name:  "efgh",
				DOB:        "2005-12-10",
				Address:    "Street 5",
			},
			ErrorInvalidId,
		},
		{
			Models.Student{
				Id:         "2",
				First_name: "abc12",
				Last_name:  "efgh",
				DOB:        "2005-12-10",
				Address:    "Street 5",
			},
			ErrorInvalidName,
		},
	}

	for _, tc := range cases {
		got := ValidateStudent(&tc.in)
		if got != tc.want {
			t.Errorf("ValidateStudent failed \n got -> %v \n wanted -> %v", got, tc.want)
			fmt.Println(tc.want, "\n", got)
		}
	}
}
