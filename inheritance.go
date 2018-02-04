package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullname() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type book struct {
	name   string
	author // embeded type
}

func (b book) bookInfo() {
	fmt.Println("Book Name:", b.name)
	fmt.Println("Author Name:", b.fullname())
	fmt.Println("Author Bio:", b.bio)
}

func main() {
	author1 := author{
		"shaon",
		"shaonty",
		"no bio",
	}
	book1 := book{
		"Go is awesome",
		author1,
	}
	book1.bookInfo()
}
