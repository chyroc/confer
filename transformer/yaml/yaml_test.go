package transformer_yaml

import (
	"testing"

	"github.com/chyroc/go-loader"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		type Bool struct {
			B bool `loader:"file,testdata/bool.yaml"`
		}
		res := new(Bool)

		err := loader.Load(res)
		as.Nil(err)
		as.Equal(true, res.B)
	})

	t.Run("", func(t *testing.T) {
		type Bool struct {
			B float32 `loader:"file,testdata/float.yaml"`
		}
		res := new(Bool)

		err := loader.Load(res)
		as.Nil(err)
		as.Equal(float32(1.1), res.B)
	})

	t.Run("", func(t *testing.T) {
		type Bool struct {
			B int32 `loader:"file,testdata/int.yaml"`
		}
		res := new(Bool)

		err := loader.Load(res)
		as.Nil(err)
		as.Equal(int32(123), res.B)
	})

	t.Run("", func(t *testing.T) {
		type Bool struct {
			B string `loader:"file,testdata/str.yaml"`
		}
		res := new(Bool)

		err := loader.Load(res)
		as.Nil(err)
		as.Equal("abc", res.B)
	})

	t.Run("", func(t *testing.T) {
		// - 1
		// - "x"
		// - false
		// - 1.0
		type Bool struct {
			A string `loader:"file,testdata/list-string.yaml;yaml,[0]"`
			B string `loader:"file,testdata/list-string.yaml;yaml,[1]"`
			C string `loader:"file,testdata/list-string.yaml;yaml,[2]"`
			D string `loader:"file,testdata/list-string.yaml;yaml,[3]"`
		}
		res := new(Bool)

		err := loader.Load(res, loader.WithTransform(New()))
		as.Nil(err)
		as.Equal("1", res.A)
		as.Equal("x", res.B)
		as.Equal("false", res.C)
		as.Equal("1", res.D)
	})

	t.Run("", func(t *testing.T) {
		// a: 1
		// b: "x"
		// c: true
		// d: 1.1
		// e: [1, 2, "x"]
		// f:
		//  - 1
		//  - "x"
		// g:
		//  g1: 1
		//  g2: "x"
		type Bool struct {
			A string `loader:"file,testdata/map.yaml;yaml,.a"`    // 1
			B string `loader:"file,testdata/map.yaml;yaml,.e[1]"` // 2
			C string `loader:"file,testdata/map.yaml;yaml,.f[1]"` // x
			D string `loader:"file,testdata/map.yaml;yaml,.g.g2"` // x
		}
		res := new(Bool)

		err := loader.Load(res, loader.WithTransform(New()))
		as.Nil(err)
		as.Equal("1", res.A)
		as.Equal("2", res.B)
		as.Equal("x", res.C)
		as.Equal("x", res.D)
	})

	t.Run("", func(t *testing.T) {
		// g:
		//  g1: 1
		//  g2: "x"
		type Bool struct {
			D struct {
				G1 uint32 `loader:"inherit,as:yaml;yaml,.g1"`
				G2 string `loader:"inherit,as:yaml;yaml,.g2"`
			} `loader:"file,testdata/map.yaml;yaml,.g,to:string"` // x
		}
		res := new(Bool)

		err := loader.Load(res, loader.WithTransform(New()))
		as.Nil(err)
		as.Equal("x", res.D.G2)
	})
}
