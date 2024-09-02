package proxy_szl

import "fmt"

type Yahoo struct {
}

func (y Yahoo) SendMsg() {
	fmt.Printf("yahoo send message\n")
}
