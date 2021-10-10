package loader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Conf1 struct {
	A          string `loader:"env,KEY_A"`
	B_Default1 string `loader:"env,KEY_B;;default=x"`
	B_Default2 string `loader:"env,KEY_B;;default = x "`
	B_Default3 string `loader:"env,KEY_B;;default = ' x ' "`
}

func Test_Load(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		setup      func()
		distroy    func()
		want       *Conf1
		errContain string
	}{
		{name: "err-1", args: 1, errContain: "source need to be a pointer to a struct"},
		{name: "err-2", args: Conf1{}, errContain: "source need to be a pointer to a struct"},

		{name: "ok-1", args: &Conf1{}, want: &Conf1{
			A:          "",
			B_Default1: "x",
			B_Default2: "x",
			B_Default3: " x ",
		}},
		{name: "ok-2", setup: func() {
			os.Setenv("KEY_A", "a")
		}, distroy: func() {
			os.Setenv("KEY_A", "")
		}, args: &Conf1{}, want: &Conf1{
			A:          "a",
			B_Default1: "x",
			B_Default2: "x",
			B_Default3: " x ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.distroy != nil {
				defer tt.distroy()
			}

			var req interface{}
			if tt.args == nil {
				req = new(Conf1)
			} else {
				req = tt.args
			}
			err := Load(req)
			if tt.errContain != "" {
				as.NotNil(err)
				as.Contains(err.Error(), tt.errContain)
			} else {
				as.Nil(err)
				as.Equal(tt.want, req)
			}
		})
	}
}

type Conf2 struct {
	A string `loader:"env,KEY_A;;required"`
}

func Test_LoadRequired(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		setup      func()
		distroy    func()
		want       *Conf2
		errContain string
	}{
		{name: "ok-1", setup: func() {
			os.Setenv("KEY_A", "a")
		}, distroy: func() {
			os.Setenv("KEY_A", "")
		}, args: &Conf2{}, want: &Conf2{
			A: "a",
		}},

		{name: "fail-1", args: &Conf2{}, errContain: "field(\"A\") required"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.distroy != nil {
				defer tt.distroy()
			}

			var req interface{}
			if tt.args == nil {
				req = new(Conf1)
			} else {
				req = tt.args
			}
			err := Load(req)
			if tt.errContain != "" {
				as.NotNil(err)
				as.Contains(err.Error(), tt.errContain)
			} else {
				as.Nil(err)
				as.Equal(tt.want, req)
			}
		})
	}
}
