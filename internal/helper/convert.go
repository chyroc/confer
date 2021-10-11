package helper

import (
	"fmt"
	"math"
	"strconv"
)

func ToInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint64:
		if uint64(v) <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int", v, v)
	case uint:
		if uint64(v) <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int", v, v)
	case int64:
		if uint64(v) <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int", v, v)
	case uint32:
		if uint64(v) <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int", v, v)
	case uintptr:
		if uint64(v) <= math.MaxInt {
			return int(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToInt(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to int", v, v)
}

func ToInt8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case int8:
		return v, nil

	case int:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case int16:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case int32:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case int64:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uint:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uint8:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uint16:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uint32:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uint64:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case uintptr:
		if v <= math.MaxInt8 {
			return int8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int8", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToInt8(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to int8", v, v)
}

func ToInt16(v interface{}) (int16, error) {
	switch v := v.(type) {
	case int8:
		return int16(v), nil
	case int16:
		return v, nil

	case uint8:
		return int16(v), nil
	case int:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case int32:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case int64:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case uint:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case uint16:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case uint32:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case uint64:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case uintptr:
		if v <= math.MaxInt16 {
			return int16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int16", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToInt16(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to int16", v, v)
}

func ToInt32(v interface{}) (int32, error) {
	switch v := v.(type) {
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil

	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case int:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case int64:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case uint:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case uint32:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case uint64:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case uintptr:
		if v <= math.MaxInt32 {
			return int32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int32", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToInt32(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to int32", v, v)
}

func ToInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil

	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int64", v, v)
	case uint64:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int64", v, v)
	case uintptr:
		if v <= math.MaxInt64 {
			return int64(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max int64", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToInt64(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to int64", v, v)
}

func ToUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case int:
		return uint(v), nil
	case int8:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case int32:
		return uint(v), nil
	case uint:
		return v, nil

	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case int64:
		if uint64(v) <= math.MaxUint {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint", v, v)
	case uint64:
		if uint64(v) <= math.MaxUint {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint", v, v)
	case uintptr:
		if uint64(v) <= math.MaxUint {
			return uint(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUint(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uint", v, v)
}

func ToUint8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case int8:
		return uint8(v), nil
	case uint8:
		return v, nil

	case int:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case int16:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case int32:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case int64:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case uint:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case uint16:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case uint32:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case uint64:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case uintptr:
		if v <= math.MaxUint8 {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint8", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUint8(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uint8", v, v)
}

func ToUint16(v interface{}) (uint16, error) {
	switch v := v.(type) {
	case int8:
		return uint16(v), nil
	case int16:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil

	case int:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case int32:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case int64:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case uint:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case uint32:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case uint64:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case uintptr:
		if v <= math.MaxUint16 {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint16", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUint16(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uint16", v, v)
}

func ToUint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case int8:
		return uint32(v), nil
	case int16:
		return uint32(v), nil
	case int32:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil

	case int:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint32", v, v)
	case int64:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint32", v, v)
	case uint:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint32", v, v)
	case uint64:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint32", v, v)
	case uintptr:
		if v <= math.MaxUint32 {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max uint32", v, v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUint32(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uint32", v, v)
}

func ToUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case int:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil

	case uintptr:
		return uint64(v), nil
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUint64(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uint64", v, v)
}

func ToUintptr(v interface{}) (uintptr, error) {
	switch v := v.(type) {
	case int:
		return uintptr(v), nil
	case int8:
		return uintptr(v), nil
	case int16:
		return uintptr(v), nil
	case int32:
		return uintptr(v), nil
	case int64:
		return uintptr(v), nil
	case uint:
		return uintptr(v), nil
	case uint8:
		return uintptr(v), nil
	case uint16:
		return uintptr(v), nil
	case uint32:
		return uintptr(v), nil
	case uint64:
		return uintptr(v), nil
	case uintptr:
		return v, nil
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return ToUintptr(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to uintptr", v, v)
}

func ToFloat32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case float32:
		return v, nil
	case float64:
		if v <= math.MaxFloat32 {
			return float32(v), nil
		}
		return 0, fmt.Errorf("%v(%T) overflow max float32", v, v)
	case string:
		if i, err := strconv.ParseFloat(v, 32); err == nil {
			return ToFloat32(i)
		}
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			return ToFloat32(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to float32", v, v)
}

func ToFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		if i, err := strconv.ParseFloat(v, 32); err == nil {
			return ToFloat64(i)
		}
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			return ToFloat64(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to float64", v, v)
}

func ToBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	case string:
		if i, err := strconv.ParseBool(v); err == nil {
			return ToBool(i)
		}
	}
	return false, fmt.Errorf("%v(%T) can't convert to bool", v, v)
}

func ToComplex64(v interface{}) (complex64, error) {
	switch v := v.(type) {
	case complex64:
		return v, nil
	case string:
		if i, err := strconv.ParseComplex(v, 64); err == nil {
			return ToComplex64(i)
		}
		if i, err := strconv.ParseComplex(v, 128); err == nil {
			return ToComplex64(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to complex64", v, v)
}

func ToComplex128(v interface{}) (complex128, error) {
	switch v := v.(type) {
	case complex64:
		return complex128(v), nil
	case complex128:
		return v, nil
	case string:
		if i, err := strconv.ParseComplex(v, 64); err == nil {
			return ToComplex128(i)
		}
		if i, err := strconv.ParseComplex(v, 128); err == nil {
			return ToComplex128(i)
		}
	}
	return 0, fmt.Errorf("%v(%T) can't convert to complex128", v, v)
}

func ToString(v interface{}) (string, error) {
	switch v := v.(type) {
	case []rune:
		return string(v), nil
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(v), 10), nil
	case uintptr:
		return strconv.FormatUint(uint64(v), 10), nil
	case bool:
		return strconv.FormatBool(v), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32), nil
	case float64:
		return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
	case complex64:
		return strconv.FormatComplex(complex128(v), 'g', -1, 64), nil
	case complex128:
		return strconv.FormatComplex(complex128(v), 'g', -1, 128), nil
	default:
		return fmt.Sprintf("%x", v), nil
	}
}
