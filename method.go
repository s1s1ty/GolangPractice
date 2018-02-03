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

type Account struct {
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
	a1 := Account{account_name: "Facebook", time_spent_daily: 2}
	a2 := Account{"Youtube", 4}

	fmt.Printf("%s will take %d hours in this month from your life.\n", a1.account_name, a1.SpentTime())
	fmt.Printf("%s will take %d hours in this month from your life.\n", a2.account_name, a2.SpentTime())
}
