package dataStructures

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
    Head    *node
    Tail    *node
    Length  int
}

var indexOutOfBoundsErr = "index must be greater than zero and less than the length of the list."
var emptyListErr = "list already has length of zero."

// Print loops through all the nodes in the queue and prints them to the console.
func (s *SinglyLinkedList) Print() {
    current := s.Head

    for {
        if current == nil {
            break;
        }

        if current.next == nil {
            fmt.Printf("current: %v, next: nil\n", current.Value)
            break
        } 
        
        fmt.Printf("current: %v, next: %v\n", current.Value, current.next.Value)
        current = current.next
    }
}

// Push appends a new node to the end of the linked list.
func (s *SinglyLinkedList) Push(val any) {
    newNode := node{
        Value: val,
        next: nil,
    }

    if s.Length == 0 {
        s.Head = &newNode
        s.Tail = &newNode
    } else {
        s.Tail.next = &newNode
        s.Tail = &newNode
    }

    s.Length += 1
}

// Pop removes a node from the end of the list and returns it.
func (s *SinglyLinkedList) Pop() (*node, error) {
    if s.Length == 0 {
        return nil, errors.New(emptyListErr)
    }

    current := s.Head
    newTail := current

    for {
        if current.next == nil {
            break;
        }
        newTail = current
        current = current.next
    }

    s.Tail = newTail
    s.Tail.next = nil
    s.Length -= 1

    if s.Length == 0 {
        s.Head = nil
        s.Tail = nil
    }

    return current, nil
}

// Shift removes the first item from the list and returns it.
func (s *SinglyLinkedList) Shift() (*node, error) {
    if s.Length == 0 {
        return nil, errors.New(emptyListErr)
    }

    current := s.Head
    s.Head = current.next
    current.next = nil
    s.Length -= 1

    if s.Length == 0 {
        s.Head = nil
        s.Tail = nil
    }

    return current, nil
}

// Unshift appends a node to beginning of the list.
func (s *SinglyLinkedList) Unshift(val any) {
    newNode := node{
        Value: val,
        next: nil,
    }

    if s.Length == 0 {
        s.Head = &newNode
        s.Tail = &newNode
    } else {
        newNode.next = s.Head
        s.Head = &newNode
    }

    s.Length += 1
}

// Get finds a node by index and returns it.
func (s *SinglyLinkedList) Get(index int) (*node, error) {
    if index < 0 || index >= s.Length {
        return nil, errors.New(indexOutOfBoundsErr)
    }

    counter := 0
    current := s.Head

    for {
        if counter >= index {
            break
        }

        current = current.next
        counter += 1
    }

    return current, nil
}

// Set finds a node by index and replaces its's Value with a new Value.
func (s *SinglyLinkedList) Set(index int, val any) error {
    node, err := s.Get(index)
    if err != nil {
        return err
    }

    node.Value = val

    return nil
}

// Insert adds a node to the list at the specified index.
func (s *SinglyLinkedList) Insert(index int, val any) error {
    if index < 0 || index >= s.Length {
        return errors.New(indexOutOfBoundsErr)
    }

    if index == s.Length {
        s.Push(val)
        return nil
    }

    if index == 0 {
        s.Unshift(val)
        return nil
    }

    newNode := node{
        Value: val,
        next: nil,
    }

    nodeBeforeInsert, err := s.Get(index - 1)
    if err != nil {
        return err
    }

    newNode.next = nodeBeforeInsert.next
    nodeBeforeInsert.next = &newNode

    s.Length += 1

    return nil
}

// Remove deletes a node by index.
func (s *SinglyLinkedList) Remove(index int) (*node, error) {
    if index < 0 || index >= s.Length {
        return nil, errors.New(indexOutOfBoundsErr)
    }

    if index == s.Length - 1 {
        node, _ := s.Pop()
        return node, nil       
    }

    if index == 0 {
        node, _ := s.Shift()
        return node, nil
    }

    nodeBeforeRemove, _ := s.Get(index - 1)
    removedNode := nodeBeforeRemove.next
    nodeBeforeRemove.next = removedNode.next

    s.Length -= 1

    return removedNode, nil
}

// Reverse reverses the list.
func (s *SinglyLinkedList) Reverse() {
    tempHead := s.Head
    s.Head = s.Tail
    s.Tail = tempHead

    var next *node
    var prev *node

    count := 0

    for {
        if count >= s.Length {
            break
        }

        next = tempHead.next
        tempHead.next = prev
        prev = tempHead
        tempHead = next
        
        count += 1
    }
}
