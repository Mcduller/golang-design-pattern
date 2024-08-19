package prototype_szl

import "testing"

func TestPrototype(t *testing.T) {
	prototypeManager := NewPrototypeManager()
	book := &Book{Name: "szl"}
	prototypeManager.Set(book)
	book2 := prototypeManager.Get(book.Name)
	if book2 == book {
		t.Errorf("should be clone")
	}
}
