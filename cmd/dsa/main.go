package main

import (
	"fmt"
	"log"

	dataStructures "github.com/BradPreston/go-dsa/data-structures"
)

func main() {
    q := dataStructures.Queue{}
    q.Enqueue("first")
    q.Enqueue("second")
    q.Enqueue("third")
    q.Print()

    removed, err := q.Dequeue()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(removed)
    fmt.Println("after remove\n=================")

    q.Print()
}
