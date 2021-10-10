package load

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLoad(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       string
		want       *Tag
		errContain string
	}{
		{name: "err-1", args: ``, errContain: "tag conf can not be empty"},
		{name: "err-2", args: `;;`, errContain: "expect get extractor name"},
		{name: "err-4", args: `a,"`, errContain: "expect end with quota(\")"},
		{name: "err-5", args: `a,"x`, errContain: "expect end with quota(\")"},
		{name: "err-6", args: `a,"x'`, errContain: "expect end with quota(\")"},

		{name: "ok-1", args: `a`, want: &Tag{extractorName: "a"}},
		{name: "ok-2", args: `a;b`, want: &Tag{extractorName: "a", transformerName: "b"}},
		{name: "ok-3", args: `a;b`, want: &Tag{extractorName: "a", transformerName: "b"}},
		{name: "ok-3", args: `a,1;b`, want: &Tag{extractorName: "a", extractorArgs: []string{"1"}, transformerName: "b"}},
		{name: "ok-4", args: `a,1,1;b`, want: &Tag{extractorName: "a", extractorArgs: []string{"1", "1"}, transformerName: "b"}},
		{name: "ok-5", args: `a , 1 , 2 , 3 ; b`, want: &Tag{extractorName: "a", extractorArgs: []string{"1", "2", "3"}, transformerName: "b"}},
		{name: "ok-6", args: `a , 1 , 2 , 3 ; b,,`, want: &Tag{extractorName: "a", extractorArgs: []string{"1", "2", "3"}, transformerName: "b", transformerArgs: []string{"", ""}}},
		{name: "ok-7", args: `a," 1 ", 2 , " 3 "`, want: &Tag{extractorName: "a", extractorArgs: []string{" 1 ", "2", " 3 "}}},
		{name: "ok-8", args: `a," 1 ", 2 , ' 3 '`, want: &Tag{extractorName: "a", extractorArgs: []string{" 1 ", "2", " 3 "}}},

		{name: "keyword-required-1", args: `a;;required`, want: &Tag{extractorName: "a", Required: true}},
		{name: "keyword-required-2", args: `a;;required;`, want: &Tag{extractorName: "a", Required: true}},

		{name: "keyword-default-1", args: `a;;required,default=`, want: &Tag{extractorName: "a", Required: true, Default: ""}},
		{name: "keyword-default-2", args: `a;;required,default= 'x'`, want: &Tag{extractorName: "a", Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = x `, want: &Tag{extractorName: "a", Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = ' x ' `, want: &Tag{extractorName: "a", Required: true, Default: " x "}},
		{name: "keyword-default-1", args: `a;;default=`, want: &Tag{extractorName: "a", Required: false, Default: ""}},
		{name: "keyword-default-2", args: `a;;default= 'x'`, want: &Tag{extractorName: "a", Required: false, Default: "x"}},
		{name: "keyword-default-2", args: `a ;  ; default = x `, want: &Tag{extractorName: "a", Required: false, Default: "x"}},
		{name: "keyword-default-2", args: `a ;  ; default = ' x ' `, want: &Tag{extractorName: "a", Required: false, Default: " x "}},
		{name: "keyword-default-2", args: `a ;  ; default = '  `, errContain: "expect end with quota(')"},

		{name: "keyword-default-1", args: `a;;required,default=;`, want: &Tag{extractorName: "a", Required: true, Default: ""}},
		{name: "keyword-default-2", args: `a;;required,default= 'x';`, want: &Tag{extractorName: "a", Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = x ;`, want: &Tag{extractorName: "a", Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = ' x ' ;`, want: &Tag{extractorName: "a", Required: true, Default: " x "}},
		{name: "keyword-default-1", args: `a;;default=;`, want: &Tag{extractorName: "a", Required: false, Default: ""}},
		{name: "keyword-default-2", args: `a;;default= 'x';`, want: &Tag{extractorName: "a", Required: false, Default: "x"}},
		{name: "keyword-default-2", args: `a ;  ; default = x ;`, want: &Tag{extractorName: "a", Required: false, Default: "x"}},
		{name: "keyword-default-2", args: `a ;  ; default = ' x ' ;`, want: &Tag{extractorName: "a", Required: false, Default: " x "}},
		{name: "keyword-default-2", args: `a ;  ; default = '  ;`, errContain: "expect end with quota(')"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := fmt.Sprintf("%s - in: %q, out: %#v", tt.name, tt.args, tt.want)
			got, err := ParseTag(tt.args)
			if tt.errContain == "" {
				as.Nil(err, msg)
				as.NotNil(got, msg)
				as.Equal(*tt.want, *got)
			} else {
				as.NotNil(err, msg)
				as.Nil(got, msg)
				as.Contains(err.Error(), tt.errContain, msg)
			}
		})
	}
}
