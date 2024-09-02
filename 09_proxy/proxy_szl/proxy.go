package proxy_szl

import "fmt"

type EmailProxy struct {
	Yahoo *Yahoo
}

func (e *EmailProxy) SendMsg() {
	fmt.Printf("Send Meaasge Before\n")
	e.Yahoo.SendMsg()
	fmt.Printf("Send Message After\n")
}
