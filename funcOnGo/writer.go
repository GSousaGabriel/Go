package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))

	return err
}

func writerTest() {
	p := person{
		first: "Jenny",
	}

	f, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
