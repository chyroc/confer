package extractor

import (
	"fmt"
	"io/ioutil"
	"os"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (r *File) Name() string {
	return "file"
}

// Extract impl file Extract
//
// file Extractor accept one arg as filepath
// filepath support use $ENV_NAME as part of filepath
func (r *File) Extract(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("file extractor expect one args")
	}
	path := args[0]
	path = os.ExpandEnv(path)

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
