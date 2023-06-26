package dataStructures

import (
	"testing"
)

// TestEnqueue_LengthOfZero verifies that a new queue has a length of zero.
func TestEnqueue_LengthOfZero(t *testing.T) {
    q := Queue{}
    if q.Length != 0 {
        t.Errorf("enqueue - length of zero. expected 0, but got %d", q.Length)
    }
}

// TestEnqueue_LengthOfThree verifies that a test has a length of 3 after enqueueing 3 nodes.
func TestEnqueue_LengthOfThree(t *testing.T) {
    q := Queue{}
    q.Enqueue("one")
    q.Enqueue("two")
    q.Enqueue("three")
    if q.Length != 3 {
        t.Errorf("enqueue - length of three. expected 3, but got %d", q.Length)
    }
}

// TestEnqueue_NextIsNil verifies that the last node (or first) are nil on a new queue.
func TestEnqueue_NextIsNil(t *testing.T) {
    q := Queue{}

    if q.Last != nil {
        t.Errorf("enqueue - next is nil. expected next to be nil, but got :%v", q.Last)
    }
}

// TestDequeue_ErrorIfQueueIsEmpty verifies that an error is returned if you try to dequeue an empty queue.
func TestDequeue_ErrorIfQueueIsEmpty(t *testing.T) {
    q := Queue{}
    _, err := q.Dequeue()

    if err == nil {
        t.Error("dequeue - error if queue is empty. expected an error, but didn't get one")
    }
}

// TestDequeue_NotExpectingError verifies that no error should be thrown if you dequeue a queue that is not empty.
func TestDequeue_NotExpectingError(t *testing.T) {
    q := Queue{}
    q.Enqueue("test")
    _, err := q.Dequeue()

    if err != nil {
        t.Error("dequeue - queue is not empty. got an error, but shouldn't have.")
    }
}

// TestDequeue_ReturnsCorrectValue verifies that the correct value is returned when you dequeue a node.
func TestDequeue_ReturnsCorrectValue(t *testing.T) {
    q := Queue{}
    want := "test"
    q.Enqueue(want)
    got, _ := q.Dequeue()

    if got != want {
        t.Errorf("dequeue - returns correct value. wanted %v, but got %v", want, got)
    }
}
