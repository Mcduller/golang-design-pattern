package prototype_szl

type CloneAble interface {
	Clone() CloneAble
}

type PrototypeManager struct {
	BookMap map[string]CloneAble
}

func (m *PrototypeManager) Set(book *Book) {
	m.BookMap[book.Name] = book
}

func (m *PrototypeManager) Get(name string) CloneAble {
	return m.BookMap[name].Clone()
}

func NewPrototypeManager() PrototypeManager {
	return PrototypeManager{
		BookMap: make(map[string]CloneAble),
	}
}
