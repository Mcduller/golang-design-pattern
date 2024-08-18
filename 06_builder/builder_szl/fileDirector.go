package builder_szl

import "builder_szl/file"

type FileDirector struct {
	fb file.FileBuilder
}

func NewFileDirector(fb file.FileBuilder) *FileDirector {
	return &FileDirector{fb: fb}
}

func (f *FileDirector) Construct() {
	f.fb.BuildHeader()
	f.fb.BuildBody()
	f.fb.BuildTail()
}
