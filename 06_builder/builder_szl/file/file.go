package file

type FileBuilder interface {
	BuildHeader()
	BuildBody()
	BuildTail()
}
