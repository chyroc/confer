package query_key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	as := assert.New(t)
	tests := []struct {
		name       string
		args       string
		want       *QueryKey
		errContain string
	}{
		{name: "1", args: "", errContain: "empty query key"},
		{name: "1", args: "a", errContain: "invalid query key: a"},
		{name: "2", args: ".", want: &QueryKey{Type: "key", Key: ".", Index: 0, Next: (*QueryKey)(nil)}},
		{name: "3", args: ".a", want: &QueryKey{Type: "key", Key: "a", Index: 0, Next: (*QueryKey)(nil)}},
		{name: "4", args: "[0]", want: &QueryKey{Type: "index", Key: "", Index: 0, Next: (*QueryKey)(nil)}},
		{name: "4", args: "[0][1]", want: &QueryKey{Type: "index", Key: "", Index: 0, Next: &QueryKey{Type: "index", Key: "", Index: 1, Next: (*QueryKey)(nil)}}},
		{name: "4", args: "[0].a[1]", want: &QueryKey{Type: "index", Key: "", Index: 0, Next: &QueryKey{Type: "key", Key: "a", Next: &QueryKey{Type: "index", Key: "", Index: 1, Next: (*QueryKey)(nil)}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if tt.errContain != "" {
				as.NotNil(err)
				as.Nil(got)
				as.Contains(err.Error(), tt.errContain)
			} else {
				as.Nil(err)
				as.NotNil(got)
				as.Equal(tt.want, got)
			}
		})
	}
}
