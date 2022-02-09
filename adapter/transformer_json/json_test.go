package transformer_json

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/chyroc/go-loader/iface"
	"github.com/stretchr/testify/assert"
)

func Test_Name(t *testing.T) {
	as := assert.New(t)
	as.Equal("json", New().Name())
}

func Test_All(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		data       string
		path       string
		dataType   interface{}
		want       interface{}
		errContain string
	}{
		{"str-1", `{"x":""}`, "x", string(""), "", ""},
		{"str-2", `{"x":"123"}`, "x", string(""), "123", ""},

		{"int-1", `{"x":0}`, "x", int(0), int(0), ""},
		{"int-2", `{"x":1}`, "x", int(0), int(1), ""},
		{"int-3", `{"x":-1}`, "x", int(0), int(-1), ""},

		{"int8-1", `{"x":0}`, "x", int8(0), int8(0), ""},
		{"int8-2", `{"x":1}`, "x", int8(0), int8(1), ""},
		{"int8-3", `{"x":-1}`, "x", int8(0), int8(-1), ""},

		{"int16-1", `{"x":0}`, "x", int16(0), int16(0), ""},
		{"int16-2", `{"x":1}`, "x", int16(0), int16(1), ""},
		{"int16-3", `{"x":-1}`, "x", int16(0), int16(-1), ""},

		{"int32-1", `{"x":0}`, "x", int32(0), int32(0), ""},
		{"int32-2", `{"x":1}`, "x", int32(0), int32(1), ""},
		{"int32-3", `{"x":-1}`, "x", int32(0), int32(-1), ""},

		{"int64-1", `{"x":0}`, "x", int64(0), int64(0), ""},
		{"int64-2", `{"x":1}`, "x", int64(0), int64(1), ""},
		{"int64-3", `{"x":-1}`, "x", int64(0), int64(-1), ""},

		{"uint-1", `{"x":0}`, "x", uint(0), uint(0), ""},
		{"uint-2", `{"x":1}`, "x", uint(0), uint(1), ""},

		{"uint8-1", `{"x":0}`, "x", uint8(0), uint8(0), ""},
		{"uint8-2", `{"x":1}`, "x", uint8(0), uint8(1), ""},

		{"uint16-1", `{"x":0}`, "x", uint16(0), uint16(0), ""},
		{"uint16-2", `{"x":1}`, "x", uint16(0), uint16(1), ""},

		{"uint32-1", `{"x":0}`, "x", uint32(0), uint32(0), ""},
		{"uint32-2", `{"x":1}`, "x", uint32(0), uint32(1), ""},

		{"uint64-1", `{"x":0}`, "x", uint64(0), uint64(0), ""},
		{"uint64-2", `{"x":1}`, "x", uint64(0), uint64(1), ""},

		{"bool-1", `{"x":false}`, "x", false, false, ""},
		{"bool-2", `{"x":true}`, "x", true, true, ""},

		// {"list-1", `{"x":[1,2,3]}`, "x", []int{}, []int{1, 2, 3}, ""},

		//
		{"str-1", `{"x":{"y":""}}`, "x.y", string(""), "", ""},
		{"str-2", `{"x":{"y":"123"}}`, "x.y", string(""), "123", ""},

		{"int-1", `{"x":{"y":0}}`, "x.y", int(0), int(0), ""},
		{"int-2", `{"x":{"y":1}}`, "x.y", int(0), int(1), ""},
		{"int-3", `{"x":{"y":-1}}`, "x.y", int(0), int(-1), ""},

		{"int8-1", `{"x":{"y":0}}`, "x.y", int8(0), int8(0), ""},
		{"int8-2", `{"x":{"y":1}}`, "x.y", int8(0), int8(1), ""},
		{"int8-3", `{"x":{"y":-1}}`, "x.y", int8(0), int8(-1), ""},

		{"int16-1", `{"x":{"y":0}}`, "x.y", int16(0), int16(0), ""},
		{"int16-2", `{"x":{"y":1}}`, "x.y", int16(0), int16(1), ""},
		{"int16-3", `{"x":{"y":-1}}`, "x.y", int16(0), int16(-1), ""},

		{"int32-1", `{"x":{"y":0}}`, "x.y", int32(0), int32(0), ""},
		{"int32-2", `{"x":{"y":1}}`, "x.y", int32(0), int32(1), ""},
		{"int32-3", `{"x":{"y":-1}}`, "x.y", int32(0), int32(-1), ""},

		{"int64-1", `{"x":{"y":0}}`, "x.y", int64(0), int64(0), ""},
		{"int64-2", `{"x":{"y":1}}`, "x.y", int64(0), int64(1), ""},
		{"int64-3", `{"x":{"y":-1}}`, "x.y", int64(0), int64(-1), ""},

		{"uint-1", `{"x":{"y":0}}`, "x.y", uint(0), uint(0), ""},
		{"uint-2", `{"x":{"y":1}}`, "x.y", uint(0), uint(1), ""},

		{"uint8-1", `{"x":{"y":0}}`, "x.y", uint8(0), uint8(0), ""},
		{"uint8-2", `{"x":{"y":1}}`, "x.y", uint8(0), uint8(1), ""},

		{"uint16-1", `{"x":{"y":0}}`, "x.y", uint16(0), uint16(0), ""},
		{"uint16-2", `{"x":{"y":1}}`, "x.y", uint16(0), uint16(1), ""},

		{"uint32-1", `{"x":{"y":0}}`, "x.y", uint32(0), uint32(0), ""},
		{"uint32-2", `{"x":{"y":1}}`, "x.y", uint32(0), uint32(1), ""},

		{"uint64-1", `{"x":{"y":0}}`, "x.y", uint64(0), uint64(0), ""},
		{"uint64-2", `{"x":{"y":1}}`, "x.y", uint64(0), uint64(1), ""},

		{"bool-1", `{"x":{"y":false}}`, "x.y", false, false, ""},
		{"bool-2", `{"x":{"y":true}}`, "x.y", true, true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := iface.TransformerReq{
				FieldType: reflect.StructField{Type: reflect.TypeOf(tt.dataType)},
				KeyVal:    []iface.KeyVal{{"path", tt.path}},
			}
			got, err := New().Transform(tt.data, &req)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, fmt.Sprintf("%s: %v", tt.name, err))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %T:%v, want %T:%v", got, got, tt.want, tt.want)
			}
		})
	}
}
