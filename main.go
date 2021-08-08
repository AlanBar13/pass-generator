package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)

	fmt.Println("Create password (length, hasNumbers, hasSymbols)(8, true, true)")
	var length = flag.Int("l", 8, "Length of password")
	var symbols = flag.Bool("s", true, "Length of password")
	var numbers = flag.Bool("n", true, "Length of password")
	flag.Parse()
	fmt.Printf("Length: %v, Numbers: %v, Symbols: %v \n", *length, *numbers, *symbols)
	password := createPassword(*length, *numbers, *symbols)
	fmt.Println("Password created: ", password)
	error := clipboard.WriteAll(password)
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("Password copied to the clipboard")
	}
}
