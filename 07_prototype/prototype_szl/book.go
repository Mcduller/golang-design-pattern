package prototype_szl

type Book struct {
	Name string
}

func (b *Book) Clone() CloneAble {
	t := *b
	return &t
}
