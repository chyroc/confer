package extractor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chyroc/go-loader/internal"
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
func (r *File) Extract(args *internal.ExtractorReq) (string, error) {
	path, _ := args.Get("path")
	if path == "" {
		return "", fmt.Errorf("file extractor expect `path` args")
	}
	path = os.ExpandEnv(path)

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
