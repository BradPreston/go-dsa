package dataStructures

import (
	"fmt"
	"testing"
)

var incorrectLengthError = "%s: incorrect list length. got %d, but wanted %d"
var incorrectNodeValueError = "%s: node contains incorrect value. got %s, but wanted %s"
var incorrectErrError = "%s: incorrect error. got %s, but wanted %s"
var outOfBoundsError = "index must be greater than zero and less than the length of the list."

func Test_Push(t *testing.T) {
    tests := []struct{
        name    string
        values  []string
        length  int
    }{
        { "adds one node to end of list", []string{"test"}, 1 },
        { "adds three nodes to the end of list",  []string{"one","two","three"}, 3 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, node := range test.values {
            l.Push(node)
        }

        if l.Length != test.length {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.length)
        }
    }
}

func Test_Pop(t *testing.T) {
    tests := []struct {
        name            string
        values          []string
        lengthAfterPop  int
    }{
        { "removes the last item from a list with length of 1", []string{"one"}, 0 },
        { "removes the last item from a list with a length of 3", []string{"one","two","three"}, 2 },
        { "throws an error if trying to pop from empty list", []string{}, 0 },
        { "popped node has a value of 'one'", []string{"one"}, 0 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, node := range test.values {
            l.Push(node)
        }

        popped, err := l.Pop()
        if err != nil {
            want := "list already has length of zero."

            if err.Error() == want {
                continue
            }
            
            t.Errorf("%s: expected empty list error. got %v, but wanted %v", test.name, want, emptyListErr)
            return
        }

        if l.Length != test.lengthAfterPop {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterPop)
        }

        want := test.values[len(test.values) - 1]

        if popped.Value != want {
            t.Errorf(incorrectNodeValueError, test.name, popped, want)
        }
    }
}

func Test_Shift(t *testing.T) {
    tests := []struct{
        name                string
        values              []string
        lengthAfterShift    int
    }{
        { "succesfully removes the first item", []string{"one"}, 0 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, _ := l.Shift()
        want := test.values[0]
        
        if got.Value != want {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, want)
        }

        if l.Length != test.lengthAfterShift {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterShift)
        }

        _, err := l.Shift()
        if err == nil {
            t.Errorf("%s: expected length error, but didn't get one", test.name)
        }

        if err != nil {
            want := "list already has length of zero."

            if err.Error() == want {
                continue
            }

            t.Errorf("%s: expected empty list error. got %v, but wanted %v", test.name, err.Error(), want)
        }
    }
}

func Test_Unshift(t *testing.T) {
    tests := []struct{
        name                string
        values              []string
        lengthAfterUnshift  int
    }{
        { "adds one node to beginning", []string{"one"}, 1 },
        { "adds three node to beginning", []string{"one","two","three"}, 3 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Unshift(value)
        }

        if l.Length != test.lengthAfterUnshift {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterUnshift)
        }
    }
}

func Test_Get(t *testing.T) {
    tests := []struct{
        name    string
        values  []string
        index   int
    }{
        { "gets node at index 2", []string{"one","two","three"}, 2 },
        { "returns out of bounds error", []string{"one","two","three"}, 5 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Get(test.index)
        if err != nil {
            want := "no node found at index %d."
            if err.Error() != fmt.Sprintf(want, test.index) {
                t.Errorf(incorrectErrError, test.name, err.Error(), fmt.Errorf(want, test.index))
            }
            continue
        }
        want := test.values[test.index]

        if got.Value != want {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, want)
        }
    }
}

func Test_Set(t *testing.T) {
    tests := []struct{
        name        string
        values      []string
        newValue    string
        index       int
    }{
        { "sets a new value at index 2", []string{"one","two","three"}, "test", 2 },
        { "sets a new value at index 1", []string{"one","two","three"}, "test", 1 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        _ = l.Set(test.index, test.newValue)
        got, _ := l.Get(test.index)

        if got.Value != test.newValue {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, test.newValue)
        }
    }
}

func Test_Insert(t *testing.T) {
    tests := []struct{
        name    string
        index   int
        value   string
    }{
        { "inserts at index 0", 0, "zero" },
        { "inserts at index 1", 1, "test" },
        { "throws out of bounds error", 4, "error" },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}
        l.Push("one")
        l.Push("two")
        err := l.Insert(test.index, test.value)
        if err != nil {
            if err.Error() != outOfBoundsError {
                t.Errorf(incorrectErrError, test.name, err.Error(), outOfBoundsError)
            }
            continue
        }

        node, _ := l.Get(test.index)

        if node.Value != test.value {
            t.Errorf(incorrectNodeValueError, test.name, node.Value, test.value)
        }
    }
}

func Test_Remove(t *testing.T) {
    tests := []struct{
        name                string
        values              []string
        index               int
        lengthAfterRemove   int
    }{
        { "removes node at index 2", []string{"one","two","three"}, 2, 2 },
        { "removes node at index 0", []string{"one"}, 0, 0 },
        { "throws out of bounds error", []string{"one","two","three"}, 5, 3 },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Remove(test.index)
        if err != nil {
            if err.Error() != outOfBoundsError {
                t.Errorf(incorrectErrError, test.name, err.Error(), outOfBoundsError)
            }

            continue
        }

        if got.Value != test.values[test.index] {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, test.values[test.index])
        }

        if l.Length != test.lengthAfterRemove {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterRemove)
        }
    }
}
