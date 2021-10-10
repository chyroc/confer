package load

import (
	"fmt"
	"reflect"
)

type Extractor interface {
	Name() string
	Extract(args []string) (string, error)
}

type Transformer interface {
	Name() string
	Transform(data string, args []string, typ reflect.Type) (interface{}, error)
}

type Option struct {
	TagName      string
	Extractors   map[string]Extractor
	Transformers map[string]Transformer
}

func Load(source interface{}, opt *Option) error {
	vv := reflect.ValueOf(source)
	vt := reflect.TypeOf(source)
	if vv.Kind() != reflect.Ptr {
		return fmt.Errorf("source need to be a pointer to a struct")
	}
	vv = vv.Elem()
	vt = vt.Elem()
	if vv.Kind() != reflect.Struct {
		return fmt.Errorf("source need to be a pointer to a struct")
	}
	for i := 0; i < vv.NumField(); i++ {
		fv := vv.Field(i)
		ft := vt.Field(i)
		tag, ok := ft.Tag.Lookup(opt.TagName)
		if !ok {
			continue // TODO return error ?
		}
		tagConf, err := ParseTag(tag)
		if err != nil {
			return err
		}
		var data string // load data by extractors
		{
			loader, ok := opt.Extractors[tagConf.extractorName]
			if !ok {
				return fmt.Errorf("extractor(%q) not found", tagConf.extractorName)
			}
			data, err = loader.Extract(tagConf.extractorArgs)
			if err != nil {
				return err
			}
		}
		// fmt.Printf("load %q %q\n", Tag.extractorName, data)

		// transformers data
		var dest reflect.Value
		if tagConf.transformerName != "" {
			transfer, ok := opt.Transformers[tagConf.transformerName]
			if !ok {
				return fmt.Errorf("%s transformers not found", tagConf.transformerName)
			}
			val, err := transfer.Transform(data, tagConf.transformerArgs, ft.Type)
			if err != nil {
				return err
			}
			dest = reflect.ValueOf(val)
		} else {
			dest = reflect.ValueOf(data)
		}

		if tagConf.Required && dest.IsZero() {
			return fmt.Errorf("field(%q) required", ft.Name)
		}

		// set
		if fv.CanSet() {
			fv.Set(dest)
		}
	}
	return nil
}
