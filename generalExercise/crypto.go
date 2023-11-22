package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./SNOWY-EVENING.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := sha256.New()
	if _, err := io.Copy(sum, file); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", sum.Sum(nil))

	loop()
}
