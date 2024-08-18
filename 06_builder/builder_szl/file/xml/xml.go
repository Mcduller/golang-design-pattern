package xml

type XmlBuilder struct {
	content string
}

func (x *XmlBuilder) BuildHeader() {
	x.content += "XmlHead"
}

func (x *XmlBuilder) BuildBody() {
	x.content += "XmlBody"
}

func (x *XmlBuilder) BuildTail() {
	x.content += "XmlTail"
}

func (x *XmlBuilder) GetRes() string {
	return x.content
}
