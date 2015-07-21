package boxer

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	waiters := []point{point{1, 1}, point{1, 2}, point{1, 3}, point{1, 4}, point{1, 5}}
	q := newQueue()

	// queue 4 waiter
	for i := 0; i < 4; i++ {
		q.q(waiters[i])
		if q.length != i+1 {
			t.Fatal("Unexpected length of q")
		}
	}
	log.Printf("The queue is: %v\n", q)

	// Dequeue 3
	for i := 0; i < 3; i++ {
		a := q.dq()
		if *a != waiters[i] {
			t.Fatalf("q.dq() did not return the expected waiter. %v != %v", *a, waiters[i])
		}
	}

	// Queue 1 more
	q.q(waiters[4])

	// Dequeue last 2
	if *(q.dq()) != waiters[3] {
		t.Fatal("q.next() did not return waiters[3]")
	}

	if *(q.dq()) != waiters[4] {
		t.Fatal("q.next() did not return waiters[4]")
	}

	if q.length != 0 {
		t.Fatal("q.length is not 0 after all waiters are dequeued")
	}
}
