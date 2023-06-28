package main

import (
	"fmt"

	dataStructures "github.com/BradPreston/go-dsa/data-structures"
)

func main() {
    list := dataStructures.SinglyLinkedList{}
    list.Push("one")
    list.Push("two")
    list.Print()

    popped, _ := list.Pop()
    fmt.Printf("popped: %v\n", popped)
    list.Print()
    shift, _ := list.Shift()
    fmt.Printf("shifted: %v\n", shift)
    list.Print()
}
