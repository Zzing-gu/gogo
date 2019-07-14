package main

import "fmt"

func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Printf(string(from), "-->", string(to))
	Move(n-1, via, to, from)
}

func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
}

func main() {
	Hanoi(3)
}
