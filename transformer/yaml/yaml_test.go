package transformer_yaml

import (
	"testing"

	"github.com/chyroc/go-loader"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	as := assert.New(t)

	type Bool struct {
		B bool `loader:"file,testdata/bool.yaml"`
	}
	res := new(Bool)

	err := loader.Load(res)
	as.Nil(err)
}
