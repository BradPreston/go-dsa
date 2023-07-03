package dataStructures

import (
	"testing"
)

var incorrectLengthError = "%s: incorrect list length. got %d, but wanted %d"
var lengthOfZeroError = "list already has length of zero."
var incorrectNodeValueError = "%s: node contains incorrect value. got %s, but wanted %s"
var incorrectErrError = "%s: incorrect error. got %s, but wanted %s"
var unexpectedErrError = "%s: got an error, but didn't expect one"
var outOfBoundsError = "index must be greater than zero and less than the length of the list."
var expectedErrError = "%s: expected an error, but didn't get one"

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
        expectedErr     bool
    }{
        { "removes the last item from a list with length of 1", []string{"one"}, 0, false },
        { "removes the last item from a list with a length of 3", []string{"one","two","three"}, 2, false },
        { "throws an error if trying to pop from empty list", []string{}, 0, true },
        { "popped node has a value of 'one'", []string{"one"}, 0, false },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Pop()
        if err != nil {
            if test.expectedErr {
                if err.Error() != lengthOfZeroError {
                    t.Errorf(incorrectErrError, test.name, err.Error(), lengthOfZeroError)
                }
            } else {
                t.Errorf(unexpectedErrError, test.name)
            }

            continue
        }


        if l.Length != test.lengthAfterPop {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterPop)
        }

        want := test.values[len(test.values) - 1]

        if got.Value != want {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, want)
        }
    }
}

func Test_Shift(t *testing.T) {
    tests := []struct{
        name                string
        values              []string
        lengthAfterShift    int
        expectedErr         bool
    }{
        { "succesfully removes the first item", []string{"one"}, 0, false },
        { "throws empty list error", []string{}, 0, true },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Shift()
        if err != nil {
            // if an error IS expected, but it isn't the corrrect error
            if test.expectedErr {
                if err.Error() != lengthOfZeroError {
                    t.Errorf(incorrectErrError, test.name, err.Error(), lengthOfZeroError)
                }
            } else {
                t.Errorf(unexpectedErrError, test.name)
            }

            continue
        }

        if err == nil && test.expectedErr {
            t.Errorf(expectedErrError, test.name)
        }

        want := test.values[0]
        
        if got.Value != want {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, want)
        }

        if l.Length != test.lengthAfterShift {
            t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterShift)
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
        name        string
        values      []string
        index       int
        expectedErr bool
    }{
        { "gets node at index 2", []string{"one","two","three"}, 2, false },
        { "returns out of bounds error", []string{"one","two","three"}, 5, true },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Get(test.index)
        if err != nil {
            if test.expectedErr {
                if err.Error() != outOfBoundsError {
                    t.Errorf(incorrectErrError, test.name, err.Error(), outOfBoundsError)
                }
            } else {
                t.Errorf(unexpectedErrError, test.name)
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
        name        string
        index       int
        value       string
        expectedErr bool
    }{
        { "inserts at index 0", 0, "zero", false },
        { "inserts at index 1", 1, "test", false },
        { "throws out of bounds error", 4, "error", true },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}
        l.Push("one")
        l.Push("two")
        err := l.Insert(test.index, test.value)
        if err != nil {
            if test.expectedErr {
                if err.Error() != outOfBoundsError {
                    t.Errorf(incorrectErrError, test.name, err.Error(), outOfBoundsError)
                }
            } else {
                t.Errorf(unexpectedErrError, test.name)
            }

            continue
        }

        got, _ := l.Get(test.index)

        if got.Value != test.value {
            t.Errorf(incorrectNodeValueError, test.name, got.Value, test.value)
        }
    }
}

func Test_Remove(t *testing.T) {
    tests := []struct{
        name                string
        values              []string
        index               int
        lengthAfterRemove   int
        expectedErr         bool
    }{
        { "removes node at index 2", []string{"one","two","three"}, 2, 2, false },
        { "removes node at index 0", []string{"one"}, 0, 0, false },
        { "throws out of bounds error", []string{"one","two","three"}, 5, 3, true },
    }

    for _, test := range tests {
        l := SinglyLinkedList{}

        for _, value := range test.values {
            l.Push(value)
        }

        got, err := l.Remove(test.index)
        if err != nil {
            if test.expectedErr {
                if err.Error() != outOfBoundsError {
                    t.Errorf(incorrectErrError, test.name, err.Error(), outOfBoundsError)
                }
            } else {
                t.Errorf(unexpectedErrError, test.name)
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

func Test_Reverse(t *testing.T) {
    l := SinglyLinkedList{}
    l.Push("one")
    l.Push("two")
    l.Push("three")

    l.Reverse()

    head, _ := l.Get(0)
    tail, _ := l.Get(2)
    if head.Value != "three" && tail.Value != "one" {
        t.Error("list did not reverse")
    }
}
