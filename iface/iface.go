package iface

import (
	"reflect"
)

type KeyVal struct {
	Key string
	Val string
}

type ExtractorReq struct {
	KeyVal []KeyVal
}

type TransformerReq struct {
	FieldType reflect.StructField
	KeyVal    []KeyVal
}

// Extractor define how to extract data from multi source
type Extractor interface {
	Name() string
	Extract(args *ExtractorReq) (string, error)
}

// Transformer define how to transform origin data to target data
type Transformer interface {
	Name() string
	Transform(data string, args *TransformerReq) (interface{}, error)
}

func (r *ExtractorReq) Get(key string) (string, bool) {
	for _, v := range r.KeyVal {
		if v.Key == key {
			return v.Val, true
		}
	}
	return "", false
}

func (r *TransformerReq) Get(key string) (string, bool) {
	for _, v := range r.KeyVal {
		if v.Key == key {
			return v.Val, true
		}
	}
	return "", false
}
