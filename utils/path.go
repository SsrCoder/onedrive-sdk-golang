package utils

import "strings"

type path struct {
	parent *path
	name   string
}

func NewPath(pathStr string) *path {
	parts := strings.Split(pathStr, "/")
	res := &path{}
	curr := res
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if part != "" {
			curr.parent = &path{
				name: part,
			}
			curr = curr.parent
		}
	}
	return res.parent
}

func CleanPath(path string) string {
	if path == "" || path == "/" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	return path
}
