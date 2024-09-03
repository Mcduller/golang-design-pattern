package observer_szl

import "testing"

func TestObserver(t *testing.T) {
	paper := NewPaper()

	reader1 := NewReader("read1")
	reader2 := NewReader("read2")
	reader3 := NewReader("read3")

	paper.Attach(reader1)
	paper.Attach(reader2)
	paper.Attach(reader3)

	paper.Notify()
}
