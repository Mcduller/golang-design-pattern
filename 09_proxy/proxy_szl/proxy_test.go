package proxy_szl

import "testing"

func TestEmailProxy_SendMsg(t *testing.T) {
	var yahoo = &Yahoo{}
	var Proxy = EmailProxy{Yahoo: yahoo}

	Proxy.SendMsg()
}
