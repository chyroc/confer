package helper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignValueToReflect(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v reflect.Value
		s reflect.Value
	}
	tests := []struct {
		name      string
		to        reflect.Value
		in        reflect.Value
		want      interface{}
		erContain string
	}{
		{name: "str -> str", to: reflect.ValueOf(newEmptyStringPtr()), in: reflect.ValueOf("x"), want: "x"},
		{name: "str -> int", to: reflect.ValueOf(newEmptyIntPtr()), in: reflect.ValueOf("1"), want: int(1)},
		{name: "int -> int", to: reflect.ValueOf(newEmptyIntPtr()), in: reflect.ValueOf(int(1)), want: int(1)},
		{name: "int64 -> int", to: reflect.ValueOf(newEmptyIntPtr()), in: reflect.ValueOf(int64(1)), want: int(1)},
		{name: "str -> bool", to: reflect.ValueOf(newEmptyBoolPtr()), in: reflect.ValueOf("true"), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AssignValueToReflect(tt.to, tt.in)
			if tt.erContain != "" {
				as.NotNil(err)
				as.Contains(err.Error(), tt.erContain)
			} else {
				as.Nil(err)
				as.Equal(tt.want, tt.to.Elem().Interface())
			}
		})
	}
}

func newEmptyStringPtr() *string {
	s := ""
	return &s
}

func newEmptyIntPtr() *int {
	s := 1
	return &s
}

func newEmptyBoolPtr() *bool {
	s := false
	return &s
}
