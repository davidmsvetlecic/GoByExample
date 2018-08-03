package main

import "fmt"

func intSeq() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()
	fmt.Println(nextInt())

}
