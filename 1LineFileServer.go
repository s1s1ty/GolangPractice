package main

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
)

func main() {
	fmt.Println("Server Starting at 8001 port...")
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8001", http.FileServer(http.Dir(user.HomeDir)))
}
