package builder_szl

import (
	"builder_szl/file/json"
	"builder_szl/file/xml"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		j := &json.JsonBuilder{}
		fileDirector := NewFileDirector(j)
		fileDirector.Construct()
		want := "JsonHeadJsonBodyJsonTail"
		if got := j.GetRes(); got != want {
			t.Errorf("want:%s,got:%s", want, got)
		}
	})

	t.Run("xml", func(t *testing.T) {
		xmlBuilder := &xml.XmlBuilder{}
		director := NewFileDirector(xmlBuilder)
		director.Construct()
		want := "XmlHeadXmlBodyXmlTail"
		if got := xmlBuilder.GetRes(); got != want {
			t.Errorf("want:%s，got:%s", want, got)
		}
	})
}
