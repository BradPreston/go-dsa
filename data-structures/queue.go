package dataStructures

import "fmt"

type node struct {
    Value   any
    next    *node
}

type Queue struct {
    First   *node
    Last    *node
    Length  int
}

// Print loops through all the nodes in the queue and prints them to the console.
func (q *Queue) Print() {
    current := q.First

    for {
        if current != nil {
            if current.next != nil {
                fmt.Printf("current: %v, next: %v\n", current.Value, current.next.Value)
                current = current.next
            } else {
                fmt.Printf("current: %v, next: nil\n", current.Value)
                break
            }
        }
    }
}

// Enqueue adds a new node to the beginning of the queue.
func (q *Queue) Enqueue(val any) {
    newNode := node{
        Value: val,
        next: nil,
    }

    if q.Length == 0 {
        q.First = &newNode
        q.Last = &newNode
    } else {
        q.Last.next = &newNode
        q.Last = &newNode
    }

    q.Length += 1
}

// Dequeue removes the first node at the beginning of the queue and returns the Value of that node.
func (q *Queue) Dequeue() (any, error) {
    if q.Length == 0 {
        return nil, fmt.Errorf("%s", "queue already has length of zero.")
    }

    temp := q.First

    if q.First == q.Last {
        q.Last = nil
    }
    
    q.First = q.First.next

    q.Length -= 1

    return temp.Value, nil
}
