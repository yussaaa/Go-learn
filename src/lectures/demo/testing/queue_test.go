package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < q.capacity; i += 1 {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: expected %d, got %d", i, len(q.items))
		}
		if !q.Append(i) {
			t.Errorf("Failed to Append")
		}
	}
	if q.Append(6) {
		t.Errorf("Queue should be full")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	q.Append(1)
	q.Append(2)
	q.Append(3)

	for i := 1; i <= 3; i += 1 {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Next() failed")
		}
		if item != i {
			t.Errorf("Incorrect item returned: expected %d, got %d", i, item)
		}
	}
	_, ok := q.Next()
	if ok {
		t.Errorf("Queue shuold be empty already!")
	}
}
