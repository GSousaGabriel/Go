package main

import (
	"fmt"
)

func main() {
	as := [2]string{"tesy", "test 2"} //array
	fmt.Println(as)
	fmt.Println(len(as), cap(as))
	fmt.Printf("Type is %T\n", as)

	for i, v := range as {
		fmt.Printf("Value %v is on index %v\n", v, i)
	}

	fmt.Println("--------")

	ss := []string{"hello", "world"} //slice
	fmt.Println(ss)

	for _, v := range ss {
		fmt.Println("Value", v)
	}

	ss = append(ss, "test append", "roger that")

	fmt.Println("Value", ss[2])

	fmt.Println("Sliced", ss[2:])

	ss = append(ss[:1], ss[2:]...)

	fmt.Println("Deleted", ss)

	fmt.Println(ss)

	mss := make([]int, 0, 10)

	fmt.Println(cap(mss), len(mss))

	mss = append(mss, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)

	fmt.Println(cap(mss), len(mss))

	fmt.Println("--------")

	jb := []string{"test", "tested"}
	km := []string{"aprrove", "approved"}

	jb2 := make([]string, len(jb))
	copy(jb2, jb)
	fmt.Println(jb2)

	jjs := [][]string{jb, km}

	fmt.Println(jb, km, jjs)
}
