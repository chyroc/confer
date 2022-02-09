package load

import (
	"fmt"
	"reflect"

	"github.com/chyroc/go-loader/iface"
	"github.com/chyroc/go-loader/internal/helper"
)

type Option struct {
	TagName      string
	Extractors   map[string]iface.Extractor
	Transformers map[string]iface.Transformer
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
		tag, ok := ft.Tag.Lookup(opt.TagName) // default is loader
		if !ok {
			continue // TODO return error ?
		}
		tagConf, err := ParseTag(tag)
		if err != nil {
			return err
		}

		if !fv.CanSet() {
			return fmt.Errorf("field(%s) cannot set", ft.Name)
		}

		var data string // load data by extractors
		{
			loader, ok := opt.Extractors[tagConf.extractorName]
			if !ok {
				return fmt.Errorf("extractor(%q) not found", tagConf.extractorName)
			}
			data, err = loader.Extract(tagConf.extractorArgs)
			if err != nil {
				if tagConf.Default != "" {
					if err := helper.AssignValueToReflect(fv, reflect.ValueOf(tagConf.Default)); err != nil {
						return err
					}
					continue
				}
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
			val, err := transfer.Transform(data, &iface.TransformerReq{
				FieldType: ft,
				KeyVal:    tagConf.transformerArgs.KeyVal,
			})
			if err != nil {
				if tagConf.Default != "" {
					if err := helper.AssignValueToReflect(fv, reflect.ValueOf(tagConf.Default)); err != nil {
						return err
					}
					continue
				}
				return err
			}
			dest = reflect.ValueOf(val)
		} else {
			dest = reflect.ValueOf(data)
		}

		if dest.IsZero() {
			if tagConf.Default != "" {
				if err := helper.AssignValueToReflect(fv, reflect.ValueOf(tagConf.Default)); err != nil {
					return err
				}
				continue
			}
			if tagConf.Required {
				return fmt.Errorf("field(%q) required", ft.Name)
			}
		}

		if err := helper.AssignValueToReflect(fv, dest); err != nil {
			return err
		}
	}
	return nil
}
