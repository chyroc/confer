package extractor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loaderEnv(t *testing.T) {
	as := assert.New(t)

	as.Equal("env", NewEnv().Name())

	os.Setenv("a", "val-a")

	tests := []struct {
		name       string
		args       []string
		want       string
		errContain string
	}{
		{name: "ok-1", args: []string{"a"}, want: "val-a"},
		{name: "ok-2", args: []string{"b"}, want: ""},

		{name: "err-1", args: []string{}, errContain: "env extractor expect one args"},
		{name: "err-2", args: []string{"a", "b"}, errContain: "env extractor expect one args"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewEnv()
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
