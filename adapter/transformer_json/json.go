package transformer_json

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/chyroc/go-loader/iface"
)

type JSON struct{}

func New() *JSON {
	return &JSON{}
}

func (r *JSON) Name() string {
	return "json"
}

func (r *JSON) Transform(data string, args *iface.TransformerReq) (interface{}, error) {
	path, _ := args.Get("path")
	if path == "" {
		return "", fmt.Errorf("json transformer expect `path` args")
	}
	tmp := strings.Split(path, ".")
	paths := make([]interface{}, len(tmp))
	for i := range tmp {
		paths[i] = tmp[i]
	}

	node, err := sonic.GetFromString(data, paths...)
	if err != nil {
		return nil, err
	}

	switch args.FieldType.Type.Kind() {
	case reflect.String:
		return node.String()
	case reflect.Int:
		s, err := node.Int64()
		return int(s), err
	case reflect.Int8:
		s, err := node.Int64()
		return int8(s), err
	case reflect.Int16:
		s, err := node.Int64()
		return int16(s), err
	case reflect.Int32:
		s, err := node.Int64()
		return int32(s), err
	case reflect.Int64:
		s, err := node.Int64()
		return int64(s), err
	case reflect.Uint:
		s, err := node.Int64()
		return uint(s), err
	case reflect.Uint8:
		s, err := node.Int64()
		return uint8(s), err
	case reflect.Uint16:
		s, err := node.Int64()
		return uint16(s), err
	case reflect.Uint32:
		s, err := node.Int64()
		return uint32(s), err
	case reflect.Uint64:
		s, err := node.Int64()
		return uint64(s), err
	case reflect.Uintptr:
		s, err := node.Int64()
		return uint64(s), err
	case reflect.Float32:
		s, err := node.Float64()
		return float32(s), err
	case reflect.Float64:
		return node.Float64()
	case reflect.Bool:
		return node.Bool()
	case reflect.Array, reflect.Slice:
		args.FieldType.Type.Elem()
		node.ArrayUseNode()
	}

	// arr, map, slice

	return node.Interface()
}
