package filesystem

import (
	"github.com/SsrCoder/onedrive-sdk-golang/microsoft"
	"github.com/SsrCoder/onedrive-sdk-golang/utils"
)

type File struct {
	fs     *FileSystem
	path   string
	msFile *microsoft.FileInfo
}

func newFileFromMsResp(fs *FileSystem, path string, resp *microsoft.FileInfo) *File {
	path = utils.CleanPath(path)
	return &File{
		fs:     fs,
		path:   path,
		msFile: resp,
	}
}

func (f *File) IsFile() bool {
	return f.msFile.File != nil
}

func (f *File) IsFolder() bool {
	return f.msFile.Folder != nil
}

func (f *File) Children() (children []*File, err error) {
	return f.fs.GetChildren(f.path)
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Walk(fn func(file *File)) {
	fn(f)
	if f.IsFolder() {
		children, err := f.Children()
		if err != nil {
			panic(err)
		}
		for _, child := range children {
			child.Walk(fn)
		}
	}
}
