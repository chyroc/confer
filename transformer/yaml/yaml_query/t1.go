package yaml_query

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/chyroc/go-loaders/transformer/transformer_yaml/query_key"
	"github.com/goccy/go-yaml"
)

// string
// uint64
// bool
// float64
// (interface {}) <nil>
// map[string]interface
// []interface {}

// func main() {
// 	bs := readfile(os.Args[1])
// 	fmt.Println(string(bs))
// 	key := os.Args[2]
//
// 	resp, err := AccessYaml(bs, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	spew.Dump(resp)
// }

func QueryYaml(bs []byte, key string) (interface{}, error) {
	queryKey, err := query_key.Parse(key)
	if err != nil {
		return nil, err
	}

	var obj interface{}
	if err = yaml.Unmarshal(bs, &obj); err != nil {
		return nil, err
	}

	return accessYaml(obj, queryKey)
}

func accessYaml(obj interface{}, queryKey *query_key.QueryKey) (interface{}, error) {
	if queryKey == nil || queryKey.Key == "." {
		return obj, nil
	}
	switch queryKey.Type {
	case "key":
		if obj == nil {
			return nil, nil
		}
		objm, ok := obj.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("%s dont support dot access", reflect.TypeOf(objm).Name())
		}
		return accessYaml(objm[queryKey.Key], queryKey.Next)
	case "index":
		if obj == nil {
			return nil, nil
		}
		objl, ok := obj.([]interface{})
		if !ok {
			return nil, fmt.Errorf("%s dont support index access", reflect.TypeOf(objl).Name())
		}
		return accessYaml(objl[queryKey.Index], queryKey.Next)
	default:
		panic("unreachable")
	}
}

func readfile(f string) []byte {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return bs
}
