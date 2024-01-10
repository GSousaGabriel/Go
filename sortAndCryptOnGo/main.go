package main

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	sort.Ints(xi)
	fmt.Println(xi)

	p1 := Person{"James", 6}
	p2 := Person{"Maria", 89}
	p4 := Person{"Silver", 2}
	p3 := Person{"Roger", 65}

	people := []Person{p3, p1, p4, p2}

	fmt.Println(people)
	//sort.Sort(ByAge(people))
	sort.Sort(ByName(people))
	fmt.Println(people)

	fmt.Println("-----------------------------------")

	pass := "pass"
	cPass, errorPass := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

	if errorPass != nil {
		fmt.Println(errorPass)
	}
	fmt.Println(string(cPass))

	errCompare := bcrypt.CompareHashAndPassword(cPass, []byte("fasdasda"))

	if errCompare != nil {
		fmt.Println("Login Error")
		return
	}
	fmt.Println("Logged")
}
