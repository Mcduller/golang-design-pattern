package json

type JsonBuilder struct {
	content string
}

func (j *JsonBuilder) BuildHeader() {
	j.content += "JsonHead"
}

func (j *JsonBuilder) BuildBody() {
	j.content += "JsonBody"
}

func (j *JsonBuilder) BuildTail() {
	j.content += "JsonTail"
}

func (j *JsonBuilder) GetRes() string {
	return j.content
}
