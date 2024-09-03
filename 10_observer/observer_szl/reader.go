package observer_szl

import "fmt"

type Reader struct {
	Name string
}

func (r *Reader) Update(paper *Paper) {
	fmt.Printf("reader:%s reciver the paper\n", r.Name)
}

func NewReader(name string) *Reader {
	return &Reader{Name: name}
}
