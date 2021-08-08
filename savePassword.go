package main

import (
	"fmt"
	"log"
	"os"
)

func savePassword(password string) {
	log.SetPrefix("File Error: ")
	log.SetFlags(0)
	f, err := os.OpenFile("password.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(password + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Password saved to password.txt")
}
