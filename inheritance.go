package main

import (
	"fmt"
	"time"
)

var months = map[string]int{
	"January":   31,
	"February":  28,
	"March":     31,
	"April":     30,
	"May":       31,
	"June":      30,
	"July":      31,
	"August":    31,
	"September": 30,
	"October":   31,
	"November":  30,
	"December":  31,
}

type Person struct {
	name string
	age  int
}

type Account struct {
	Person           //anonymous field
	account_name     string
	time_spent_daily int
}

func (a1 Account) SpentTime() int {
	current_month := time.Now().Month()
	days := months[current_month.String()]
	total_time := days * a1.time_spent_daily
	return total_time

}

func main() {
	a1 := Account{Person{"Shaon", 100}, "Facebook", 2}
	a2 := Account{Person{"Munni", 100}, "Youtube", 4}

	fmt.Printf("%s will take %d hours in this month from %s's life.\n",
		a1.account_name, a1.SpentTime(), a1.name)

	fmt.Printf("%s will take %d hours in this month from %s's life.\n",
		a2.account_name, a2.SpentTime(), a2.name)
}
