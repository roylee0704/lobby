package lobby

import "testing"

// for the sake of unit testing
func TestPop(t *testing.T) {

	var tests = []struct {
		p Pool
	}{
		{p: []*Worker{&Worker{}, &Worker{}, &Worker{}}},
		{p: []*Worker{&Worker{}}},
		//{p: []*Worker{}}, pop empty queue, does not happen in heap.interface
	}

	for _, test := range tests {
		n := len(test.p)
		w := test.p[n-1]
		if w != test.p.Pop() {
			t.Error("priority queue poping failed.")
		}
	}
}
