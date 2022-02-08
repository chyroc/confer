package tag_parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		want       *Tag
		errContain string
	}{
		{"1", "env", &Tag{Extractor: &Param{Name: "env"}}, ""},
		{"2", "env,key=val", &Tag{Extractor: &Param{Name: "env", Args: []KeyVal{{"key", "val"}}}}, ""},
		{"3", "env,a=b,c=d", &Tag{Extractor: &Param{Name: "env", Args: []KeyVal{{"a", "b"}, {"c", "d"}}}}, ""},
		{"4", "env,1=2", &Tag{Extractor: &Param{Name: "env", Args: []KeyVal{{"1", "2"}}}}, ""},

		{"2-1", "env,1=2;json", &Tag{
			Extractor:   &Param{Name: "env", Args: []KeyVal{{"1", "2"}}},
			Transformer: &Param{Name: "json"},
		}, ""},
		{"2-2", "env,1=2;json,k=v", &Tag{
			Extractor:   &Param{Name: "env", Args: []KeyVal{{"1", "2"}}},
			Transformer: &Param{Name: "json", Args: []KeyVal{{"k", "v"}}},
		}, ""},

		{"3-1", "env,1=2;json;required", &Tag{
			Extractor:   &Param{Name: "env", Args: []KeyVal{{"1", "2"}}},
			Transformer: &Param{Name: "json"},
			Required:    true,
		}, ""},
		{"3-2", "env,1=2;json;required,default=x", &Tag{
			Extractor:   &Param{Name: "env", Args: []KeyVal{{"1", "2"}}},
			Transformer: &Param{Name: "json"},
			Required:    true,
			Default:     "x",
		}, ""},
		{"3-3", "env,1=2;json;default=x,required", &Tag{
			Extractor:   &Param{Name: "env", Args: []KeyVal{{"1", "2"}}},
			Transformer: &Param{Name: "json"},
			Required:    true,
			Default:     "x",
		}, ""},
		{"3-4", "env;;required", &Tag{
			Extractor: &Param{Name: "env"},
			Required:  true,
		}, ""},

		{name: "err-1", args: ``, errContain: "tag conf can not be empty"},
		{name: "err-2", args: `;;`, errContain: "extractor can not be empty"},

		{name: "ok-1", args: `a`, want: &Tag{Extractor: &Param{Name: "a"}}, errContain: ""},
		{name: "ok-2", args: `a;b`, want: &Tag{Extractor: &Param{Name: "a"}, Transformer: &Param{Name: "b"}}, errContain: ""},
		{name: "ok-3", args: `a,k=1;b`, want: &Tag{Extractor: &Param{Name: "a", Args: []KeyVal{{"k", "1"}}}, Transformer: &Param{Name: "b"}}, errContain: ""},
		{name: "ok-4", args: `a,k=1,k2=2;b`, want: &Tag{Extractor: &Param{Name: "a", Args: []KeyVal{{"k", "1"}, {"k2", "2"}}}, Transformer: &Param{Name: "b"}}, errContain: ""},
		{name: "ok-5", args: `a , k = 1 , k2 = 2 ; b`, want: &Tag{Extractor: &Param{Name: "a", Args: []KeyVal{{"k", "1"}, {"k2", "2"}}}, Transformer: &Param{Name: "b"}}, errContain: ""},
		{name: "ok-6", args: `a , k =" 1 ", k2 = 2 ; b`, want: &Tag{Extractor: &Param{Name: "a", Args: []KeyVal{{"k", " 1 "}, {"k2", "2"}}}, Transformer: &Param{Name: "b"}}, errContain: ""},
		{name: "ok-7", args: `a , k = " 1 ", k2 = "2 3"; b`, want: &Tag{Extractor: &Param{Name: "a", Args: []KeyVal{{"k", " 1 "}, {"k2", "2 3"}}}, Transformer: &Param{Name: "b"}}, errContain: ""},

		{name: "keyword-required-1", args: `a;;required`, want: &Tag{Extractor: &Param{Name: "a"}, Required: true}},
		{name: "keyword-required-2", args: `a;;required;`, want: &Tag{Extractor: &Param{Name: "a"}, Required: true}},

		{name: "keyword-default-2", args: `a;;required,default=`, want: &Tag{Extractor: &Param{Name: "a"}, Required: true}},
		{name: "keyword-default-2", args: `a;;required,default=x`, want: &Tag{Extractor: &Param{Name: "a"}, Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = x `, want: &Tag{Extractor: &Param{Name: "a"}, Transformer: &Param{}, Required: true, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ; required , default = " x "`, want: &Tag{Extractor: &Param{Name: "a"}, Transformer: &Param{}, Required: true, Default: " x "}},
		{name: "keyword-default-2", args: `a ; ;   default = x `, want: &Tag{Extractor: &Param{Name: "a"}, Transformer: &Param{}, Required: false, Default: "x"}},
		{name: "keyword-default-2", args: `a ; ;  default = " x "`, want: &Tag{Extractor: &Param{Name: "a"}, Transformer: &Param{}, Required: false, Default: " x "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if tt.errContain != "" {
				if err == nil {
					t.Errorf("Parse() error = %v, wantErr %v", err, tt.errContain)
					return
				}
				if !strings.Contains(err.Error(), tt.errContain) {
					t.Errorf("Parse() error = %v, wantErr %v", err, tt.errContain)
					return
				}
			}
			if err := compareTag(got, tt.want); err != nil {
				t.Errorf("Parse() compare error = %v", err)
			}
		})
	}
}

func compareTag(a, b *Tag) error {
	if (a != nil && b == nil) || (a == nil && b != nil) {
		return fmt.Errorf("tag one nil, one not")
	}
	if a == nil && b == nil {
		return nil
	}
	checkParam := func(msg string, p1, p2 *Param) error {
		if (p1 != nil && p2 == nil) || (p1 == nil && p2 != nil) {
			return fmt.Errorf("%s one nil, one not", msg)
		}
		if p1 == nil && p2 == nil {
			return nil
		}
		if p1.Name != p2.Name {
			return fmt.Errorf("%s name not equal: %s %s", msg, p1.Name, p2.Name)
		}
		if len(p1.Args) != len(p2.Args) {
			return fmt.Errorf("%s args not equal: %d %d", msg, len(p1.Args), len(p2.Args))
		}
		for i := 0; i < len(p1.Args); i++ {
			if p1.Args[i].Key != p2.Args[i].Key {
				return fmt.Errorf("%s args key not equal: %s %s", msg, p1.Args[i].Key, p2.Args[i].Key)
			}
			if p1.Args[i].Val != p2.Args[i].Val {
				return fmt.Errorf("%s args val not equal: %s %s", msg, p1.Args[i].Val, p2.Args[i].Val)
			}
		}
		return nil
	}
	if err := checkParam("extractor", a.Extractor, b.Extractor); err != nil {
		return err
	}
	if err := checkParam("transformer", a.Transformer, b.Transformer); err != nil {
		return err
	}
	if a.Required != b.Required {
		return fmt.Errorf("required not equal: %t %t", a.Required, b.Required)
	}
	if a.Default != b.Default {
		return fmt.Errorf("default not equal: %s %s", a.Default, b.Default)
	}
	return nil
}
