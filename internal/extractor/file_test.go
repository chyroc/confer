package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_file(t *testing.T) {
	as := assert.New(t)

	as.Equal("file", NewFile().Name())

	file1 := newTmpFile("val-1")

	tests := []struct {
		name       string
		args       []string
		want       string
		errContain string
	}{
		{name: "ok-1", args: []string{file1}, want: "val-1"},
		// {name: "ok-2", args: []string{"b"}, want: ""},
		// {name: "ok-3", args: []string{"b", "default=x"}, want: "x"},
		// {name: "ok-4", args: []string{"b", "default = x"}, want: "x"},
		// {name: "ok-5", args: []string{"b", "default = ' x '"}, want: " x "},

		{name: "err-1", args: []string{}, errContain: "file extractor expect one args"},
		{name: "err-2", args: []string{"a"}, errContain: "open a: no such file or directory"},
		{name: "err-3", args: []string{"a", "b"}, errContain: "file extractor expect one args"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewFile()
			got, err := r.Extract(tt.args)
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
