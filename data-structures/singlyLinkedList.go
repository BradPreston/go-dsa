package dataStructures

import "fmt"

type SinglyLinkedList struct {
    Head    *node
    Tail    *node
    Length  int
}

// Print loops through all the nodes in the queue and prints them to the console.
func (s *SinglyLinkedList) Print() {
    current := s.Head

    for {
        if current == nil {
            break;
        }

        if current.next == nil {
            fmt.Printf("current: %v, next: nil\n", current.value)
            break
        } 
        
        fmt.Printf("current: %v, next: %v\n", current.value, current.next.value)
        current = current.next
    }
}

// Push appends a new node to the end of the linked list.
func (s *SinglyLinkedList) Push(val any) {
    newNode := node{
        value: val,
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

// Pop removes a node from the end of the list and returns it's value.
func (s *SinglyLinkedList) Pop() (any, error) {
    if s.Length == 0 {
        return nil, fmt.Errorf("%s", "list already has length of zero.")
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

    return current.value, nil
}

// Shift removes the first item from the list and returns it's value.
func (s *SinglyLinkedList) Shift() (any, error) {
    if s.Length == 0 {
        return nil, fmt.Errorf("%s", "list already has length of zero.")
    }

    current := s.Head
    s.Head = current.next
    current.next = nil
    s.Length -= 1

    if s.Length == 0 {
        s.Head = nil
        s.Tail = nil
    }

    return current.value, nil
}
