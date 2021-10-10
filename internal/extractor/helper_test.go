package extractor

import (
	"io/ioutil"
)

func newTmpFile(s string) string {
	f, err := ioutil.TempFile("", "go-loader-*")
	if err != nil {
		panic(err)
	}
	f.WriteString(s)
	f.Close()
	return f.Name()
}
