package confer

import (
	"fmt"
	"reflect"

	"github.com/chyroc/confer/internal/loader"
)

// Load get data, and assign to source
//
// source must a pointer to struct
func Load(source interface{}, options ...ImplOption) error {
	r := &Impl{loader: map[string]Loader{}, transfer: map[string]Transfer{}}

	// default
	r.loader["env"] = loader.NewEnv()

	for _, v := range options {
		if err := v(r); err != nil {
			return err
		}
	}
	return r.load(source)
}

type Impl struct {
	loader   map[string]Loader
	transfer map[string]Transfer
}

func (r *Impl) load(source interface{}) error {
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
		tag, ok := ft.Tag.Lookup(tagName)
		if !ok {
			continue // TODO return error ?
		}
		tagConf, err := parseTagConf(tag)
		if err != nil {
			return err
		}
		var data string // load data by loader
		{
			if tagConf.loaderName == "" {
				return fmt.Errorf("expect get loader name")
			}
			loader, ok := r.loader[tagConf.loaderName]
			if !ok {
				return fmt.Errorf("%s loader not found", tagConf.loaderName)
			}
			data, err = loader.Load(tagConf.loaderArgs)
			if err != nil {
				return err
			}
		}
		dest := reflect.ValueOf(data)
		// fmt.Printf("load %q %q\n", tagConf.loaderName, data)

		// transfer data
		if tagConf.transferName != "" {
			transfer, ok := r.transfer[tagConf.transferName]
			if !ok {
				return fmt.Errorf("%s transfer not found", tagConf.transferName)
			}
			val, err := transfer.Transfer(data, tagConf.transferArgs, ft.Type)
			if err != nil {
				return err
			}
			dest = reflect.ValueOf(val)
		}

		// set
		if fv.CanSet() {
			fv.Set(dest)
		}
	}
	return nil
}

type ImplOption func(r *Impl) error

func WithLoader(loaders ...Loader) ImplOption {
	return func(r *Impl) error {
		for _, v := range loaders {
			if _, ok := r.loader[v.Name()]; ok {
				return fmt.Errorf("%s loader registed", v.Name())
			}
			r.loader[v.Name()] = v
		}
		return nil
	}
}

func WithTransfer(transfers ...Transfer) ImplOption {
	return func(r *Impl) error {
		for _, v := range transfers {
			if _, ok := r.transfer[v.Name()]; ok {
				return fmt.Errorf("%s transfer registed", v.Name())
			}
			r.transfer[v.Name()] = v
		}
		return nil
	}
}

type Loader interface {
	Name() string
	Load(args []string) (string, error)
}

type Transfer interface {
	Name() string
	Transfer(data string, args []string, typ reflect.Type) (interface{}, error)
}

const tagName = "conf"
