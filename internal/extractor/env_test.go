package extractor

import (
	"os"
	"testing"

	"github.com/chyroc/go-loader/internal"
	"github.com/stretchr/testify/assert"
)

func Test_loaderEnv(t *testing.T) {
	as := assert.New(t)

	as.Equal("env", NewEnv().Name())

	os.Setenv("a", "val-a")

	tests := []struct {
		name       string
		args       []internal.KeyVal
		want       string
		errContain string
	}{
		{name: "ok-1", args: []internal.KeyVal{{"key", "a"}}, want: "val-a"},
		{name: "ok-2", args: []internal.KeyVal{{"key", "b"}}, want: ""},

		{name: "err-1", args: []internal.KeyVal{}, errContain: "env extractor expect `key` args"},
		{name: "err-2", args: []internal.KeyVal{{"not-key", "a"}}, errContain: "env extractor expect `key` args"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewEnv()
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
