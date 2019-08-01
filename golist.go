package main

import "fmt"

import "github.com/chenjr15/golist/linkedlist"

func main() {
	a := linkedlist.MakeLinkedList([]int{1, 2, 3})

	fmt.Println("golist")
	fmt.Println(a)
}
