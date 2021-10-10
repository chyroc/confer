package helper

import (
	"fmt"
	"reflect"
)

func ReflectSet(v reflect.Value, t reflect.Type, s string) error {
	v.SetString(s)
	return nil
	switch v.Kind() {
	case reflect.Struct:
		v.SetString(s)
	default:
		return fmt.Errorf("cannot set %q to %s", s, v.Kind())
	}
	return nil
}
