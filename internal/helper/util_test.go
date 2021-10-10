package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_min(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "1", a: 1, b: 2, want: 1},
		{name: "2", a: 0, b: 0, want: 0},
		{name: "3", a: 0, b: -1, want: -1},
		{name: "4", a: -1, b: 0, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as.Equal(tt.want, Min(tt.a, tt.b))
		})
	}
}
