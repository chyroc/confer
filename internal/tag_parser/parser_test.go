package tag_parser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Tag
		wantErr bool
	}{
		{"1", "env", &Tag{Extractor: &Param{Name: "env"}}, false},
		{"2", "env,key=val", &Tag{Extractor: &Param{Name: "env", Args: []KeyVal{{"key", "val"}}}}, false},
		{"3", "env,1=2,3=4", &Tag{Extractor: &Param{Name: "env", Args: []KeyVal{{"1", "2"}, {"3", "4"}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
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
