// package for go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var testErrorMessage = errors.New("Norgae math: square root of negative number")

// function that works for tests
func main() {
	fmt.Printf("%T\n", testErrorMessage)
	defer foo()
	defer recoverExp()

	_, erro := sqrt(-10)
	if erro != nil {
		log.Fatalln(erro)
	}

	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Fatalln(err) //do not execute defer
		log.Panicln("Error here")
		panic(err)
	}
}

func foo() {
	fmt.Println("TEST")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, testErrorMessage
	}
	return 42, nil
}
