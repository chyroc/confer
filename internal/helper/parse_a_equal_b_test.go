package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// parse a=b
func Test_parseAEqualB(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       string
		want       []string
		errContain string
	}{
		{name: "err-1", args: `1`, errContain: "reach end of data, `=` cannot found"},
		{name: "err-1", args: `a='`, errContain: "expect end with quota(')"},
		{name: "err-1", args: `default='`, errContain: "expect end with quota(')"},

		{name: "ok-1", args: ``, want: []string{"", ""}},
		{name: "ok-2", args: `1=`, want: []string{"1", ""}},
		{name: "ok-3", args: ` 1 = 2 `, want: []string{"1", "2"}},
		{name: "ok-4", args: `1='2'`, want: []string{"1", "2"}},
		{name: "ok-5", args: `1 = ' 2 ' `, want: []string{"1", " 2 "}},
		{name: "ok-5", args: ` ' 1 ' = ' 2 ' `, want: []string{" 1 ", " 2 "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := fmt.Sprintf("%s - in: %q, out: %#v", tt.name, tt.args, tt.want)
			a, b, err := ParseAEqualB(tt.args)
			if tt.errContain == "" {
				as.Nil(err, msg)
				as.Equal(tt.want, []string{a, b})
			} else {
				as.NotNil(err, msg)
				as.Contains(err.Error(), tt.errContain, msg)
			}
		})
	}
}
