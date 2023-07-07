package doublyLinkedList

import (
	"errors"
	"fmt"
)

type node struct {
	Value any
	next  *node
	prev  *node
}

type DoublyLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

// var indexOutOfBoundsErr = "index must be greater than zero and less than the length of the list."
var emptyListErr = "list already has length of zero."

// Print loops over each node in the list and prints it to the console.
func (d *DoublyLinkedList) Print() {
	current := d.Head

	for {
        if current.next == nil && current.prev == nil {
			fmt.Printf("current: %v, prev: nil, next: nil\n", current.Value)
            break;
        }
        
		if current.next == nil {
			fmt.Printf("current: %v, prev: %v, next: nil\n", current.Value, current.prev.Value)
			break
		}

	    fmt.Printf("current: %v, prev: %v, next: %v\n", current.Value, current.prev.Value, current.next.Value)

		current = current.next
	}
}

// Push adds a node to the end of the list.
func (d *DoublyLinkedList) Push(val any) {
	newNode := node{
		Value: val,
		next:  nil,
		prev:  nil,
	}

	if d.Length == 0 {
		d.Head = &newNode
		d.Tail = &newNode
	} else {
		d.Tail.next = &newNode
		newNode.prev = d.Tail
		d.Tail = &newNode
	}

	d.Length += 1
}

// Pop removes the last node from the list and returns it.
func (d *DoublyLinkedList) Pop() (*node, error) {
	if d.Length == 0 {
		return nil, errors.New(emptyListErr)
	}

	currentTail := d.Tail

	if d.Length == 1 {
		d.Head = nil
		d.Tail = nil
	} else {
		d.Tail = currentTail.prev
		d.Tail.next = nil
		currentTail.prev = nil
	}

	d.Length -= 1

	return currentTail, nil
}

// Shift removes the first node from the list and returns it
func (d *DoublyLinkedList) Shift() (*node, error) {
    if d.Length == 0 {
        return nil, errors.New(emptyListErr)
    }

    currentHead := d.Head

    if d.Length == 1 {
        d.Head = nil
        d.Tail = nil
    } else {
        d.Head = currentHead.next
        d.Head.prev = nil
        currentHead.next = nil
    }

    d.Length -= 1

    return currentHead, nil
}
