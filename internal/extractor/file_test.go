package extractor

import (
	"testing"

	"github.com/chyroc/go-loader/internal"
	"github.com/stretchr/testify/assert"
)

func Test_file(t *testing.T) {
	as := assert.New(t)

	as.Equal("file", NewFile().Name())

	file1 := newTmpFile("val-1")

	tests := []struct {
		name       string
		args       []internal.KeyVal
		want       string
		errContain string
	}{
		{name: "ok-1", args: []internal.KeyVal{{"path", file1}}, want: "val-1"},
		// {name: "ok-2", args: []string{"b"}, want: ""},
		// {name: "ok-3", args: []string{"b", "default=x"}, want: "x"},
		// {name: "ok-4", args: []string{"b", "default = x"}, want: "x"},
		// {name: "ok-5", args: []string{"b", "default = ' x '"}, want: " x "},

		{name: "err-1", args: []internal.KeyVal{}, errContain: "file extractor expect `path` args"},
		{name: "err-2", args: []internal.KeyVal{{"path", "a"}}, errContain: "open a: no such file or directory"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewFile()
			got, err := r.Extract(&internal.ExtractorReq{KeyVal: tt.args})
			if tt.errContain != "" {
				as.NotNil(err)
				as.Empty(got)
				as.Contains(err.Error(), tt.errContain)
			} else {
				as.Nil(err)
				as.Equal(tt.want, got)
			}
		})
	}
}
