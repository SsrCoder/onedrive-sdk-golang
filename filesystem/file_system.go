package filesystem

import (
	"fmt"
	"github.com/SsrCoder/onedrive-sdk-golang/client"
	"github.com/SsrCoder/onedrive-sdk-golang/microsoft"
	"github.com/SsrCoder/onedrive-sdk-golang/utils"
	"net/url"
)

type FileSystem struct {
	client *client.Client
	root   *File
}

func New(client *client.Client) *FileSystem {
	return &FileSystem{
		client: client,
	}
}

func (fs *FileSystem) OpenDir(dir string) (file *File, err error) {
	dir = utils.CleanPath(dir)
	path := dir
	if dir != "/" {
		dir = url.QueryEscape(dir)
		dir = fmt.Sprintf(":%s:", dir)
	} else {
		dir = ""
	}

	var res microsoft.FileInfoResp
	uri := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root%s", dir)
	err = fs.client.Get(uri).BindJSON(&res).Do()
	file = newFileFromMsResp(fs, path, &res.FileInfo)
	return
}

func (fs *FileSystem) Root() (root *File, err error) {
	if fs.root == nil {
		fs.root, err = fs.OpenDir("/")
	}
	return fs.root, err
}

func (fs *FileSystem) GetChildren(path string) (files []*File, err error) {
	path = utils.CleanPath(path)
	dir := path
	if dir != "/" {
		dir = url.QueryEscape(dir)
		dir = fmt.Sprintf(":%s:", dir)
	} else {
		dir = ""
	}
	var res microsoft.DirChildrenInfoResp
	uri := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root%s/children", dir)
	err = fs.client.Get(uri).BindJSON(&res).Do()
	for _, file := range res.Value {
		file := file
		files = append(files, newFileFromMsResp(fs, path+"/"+file.Name, &file))
	}
	return
}
