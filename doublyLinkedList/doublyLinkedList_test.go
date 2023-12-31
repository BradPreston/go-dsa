package doublyLinkedList

import "testing"

var (
	incorrectLengthError    = "%s: incorrect list length. got %d, but wanted %d"
	lengthOfZeroError       = "list already has length of zero."
	incorrectNodeValueError = "%s: node contains incorrect value. got %s, but wanted %s"
	incorrectErrError       = "%s: incorrect error. got %s, but wanted %s"
	unexpectedErrError      = "%s: got an error, but didn't expect one"
	outOfBoundsError        = "index must be greater than zero and less than the length of the list."
	expectedErrError        = "%s: expected an error, but didn't get one"
)

func Test_Push(t *testing.T) {
	tests := []struct {
		name            string
		values          []string
		lengthAfterPush int
	}{
		{"adds one node to end of list", []string{"one"}, 1},
		{"adds three nodes to end of list", []string{"one", "two", "three"}, 3},
	}

	for _, test := range tests {
		l := DoublyLinkedList{}

		for _, value := range test.values {
			l.Push(value)
		}

		if l.Length != test.lengthAfterPush {
			t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterPush)
		}
	}
}

func Test_Pop(t *testing.T) {
	tests := []struct {
		name           string
		values         []string
		lengthAfterPop int
		expectedErr    bool
	}{
		{"removes the last item from a list with length of 1", []string{"one"}, 0, false},
		{"removes the last item from a list with a length of 3", []string{"one", "two", "three"}, 2, false},
		{"throws an error if trying to pop from empty list", []string{}, 0, true},
		{"popped node has a value of 'one'", []string{"one"}, 0, false},
	}

	for _, test := range tests {
		l := DoublyLinkedList{}

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

		want := test.values[len(test.values)-1]

		if got.Value != want {
			t.Errorf(incorrectNodeValueError, test.name, got.Value, want)
		}
	}
}

func Test_Shift(t *testing.T) {
	tests := []struct {
		name             string
		values           []string
		lengthAfterShift int
		expectedErr      bool
	}{
		{"succesfully removes the first item", []string{"one"}, 0, false},
		{"throws empty list error", []string{}, 0, true},
	}

	for _, test := range tests {
		l := DoublyLinkedList{}

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
	tests := []struct {
		name               string
		values             []string
		lengthAfterUnshift int
	}{
		{"adds one node to beginning", []string{"one"}, 1},
		{"adds three node to beginning", []string{"one", "two", "three"}, 3},
	}

	for _, test := range tests {
		l := DoublyLinkedList{}

		for _, value := range test.values {
			l.Unshift(value)
		}

		if l.Length != test.lengthAfterUnshift {
			t.Errorf(incorrectLengthError, test.name, l.Length, test.lengthAfterUnshift)
		}
	}
}
