package main

import "fmt"

const daysInMonth = 28
const fullTimeBasicPay = 500
const contractorBasicPay = 100
const freelancerBasicPay = 10

type SalaryCalulator interface {
	calculateSalary() int
}

type FullTime struct {
	id           int
	monthsWorked int
}

type Contractor struct {
	id           int
	monthsWorked int
}

type Freelancer struct {
	id          int
	hoursWorked int
}

func (F *FullTime) calculateSalary() int {
	return F.monthsWorked * daysInMonth * fullTimeBasicPay
}

func (C *Contractor) calculateSalary() int {
	return C.monthsWorked * daysInMonth * contractorBasicPay
}

func (F *Freelancer) calculateSalary() int {
	return F.hoursWorked * freelancerBasicPay
}

func main() {
	fullTimeEmp := FullTime{
		id:           1,
		monthsWorked: 1,
	}
	contractorEmp := Contractor{
		id:           2,
		monthsWorked: 3,
	}
	freelanceEmp := Freelancer{
		id:          3,
		hoursWorked: 36,
	}

	fmt.Println(fullTimeEmp.calculateSalary())
	fmt.Println(contractorEmp.calculateSalary())
	fmt.Println(freelanceEmp.calculateSalary())
}
