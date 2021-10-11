package helper

import (
	"fmt"
	"reflect"
)

func AssignValueToReflect(to, from reflect.Value) error {
	if to.Kind() == reflect.Ptr {
		return AssignValueToReflect(to.Elem(), from)
	}
	if !to.CanAddr() {
		return fmt.Errorf("cannnot set for unaddressable")
	}

	switch to.Kind() {
	case reflect.Int:
		val, err := ToInt(from.Interface())
		if err != nil {
			return err
		}
		to.SetInt(int64(val))
	case reflect.Int8:
		val, err := ToInt8(from.Interface())
		if err != nil {
			return err
		}
		to.SetInt(int64(val))
	case reflect.Int16:
		val, err := ToInt16(from.Interface())
		if err != nil {
			return err
		}
		to.SetInt(int64(val))
	case reflect.Int32:
		val, err := ToInt32(from.Interface())
		if err != nil {
			return err
		}
		to.SetInt(int64(val))
	case reflect.Int64:
		val, err := ToInt64(from.Interface())
		if err != nil {
			return err
		}
		to.SetInt(int64(val))
	case reflect.Uint:
		val, err := ToUint(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.Uint8:
		val, err := ToUint8(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.Uint16:
		val, err := ToUint16(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.Uint32:
		val, err := ToUint32(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.Uint64:
		val, err := ToUint64(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.Uintptr:
		val, err := ToUintptr(from.Interface())
		if err != nil {
			return err
		}
		to.SetUint(uint64(val))
	case reflect.String:
		val, err := ToString(from.Interface())
		if err != nil {
			return err
		}
		to.SetString(val)
	case reflect.Float32:
		val, err := ToFloat32(from.Interface())
		if err != nil {
			return err
		}
		to.SetFloat(float64(val))
	case reflect.Float64:
		val, err := ToFloat64(from.Interface())
		if err != nil {
			return err
		}
		to.SetFloat(float64(val))
	case reflect.Bool:
		val, err := ToBool(from.Interface())
		if err != nil {
			return err
		}
		to.SetBool(val)
	case reflect.Complex64:
		val, err := ToComplex64(from.Interface())
		if err != nil {
			return err
		}
		to.SetComplex(complex128(val))
	case reflect.Complex128:
		val, err := ToComplex128(from.Interface())
		if err != nil {
			return err
		}
		to.SetComplex(complex128(val))
	case reflect.Ptr:
		return AssignValueToReflect(to.Elem(), from)
	default:
		to.Set(from)
	}
	return nil
}
