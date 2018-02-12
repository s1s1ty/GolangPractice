package main

import (
	"fmt"
)

type unknownType interface{}

type SalaryCalculator interface {
	CalculateSalary() (unknownType, int)
}

type Employee struct {
	name  string
	empId unknownType
}

type Permanent struct {
	Employee
	basicpay int
	pf       int
}

type Contract struct {
	Employee
	basicpay int
}

func (p Permanent) CalculateSalary() (unknownType, int) {
	return p.empId, p.basicpay + p.pf
}

func (c Contract) CalculateSalary() (unknownType, int) {
	return c.empId, c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	fmt.Println("For empId: ")
	for _, v := range s {
		empId, exp := v.CalculateSalary()
		expense += exp
		fmt.Println(empId)
	}
	fmt.Printf("Total Expense Per Month %d Taka.\n", expense)

}

func main() {
	emp1 := Employee{"Shaon", 1}
	emp2 := Employee{"Tropa", 2}
	emp3 := Employee{"Sujan", "con-1"}

	per_e1 := Permanent{emp1, 5000, 20}
	per_e2 := Permanent{emp2, 6000, 30}
	con_e1 := Contract{emp3, 3000}

	employees := []SalaryCalculator{per_e1, per_e2, con_e1}
	// fmt.Println(employees.name)
	totalExpense(employees)

}
