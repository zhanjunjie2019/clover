package utils

import "reflect"

// MergeStructs 合并相同的类型的结构体指针对象
// 经过反射，性能较低，不建议用于用户级业务处理
// @skip
func MergeStructs[T any](src T, ts ...T) error {
	if len(ts) > 0 {
		for i := range ts {
			var err error
			err = MergeStruct[T](src, ts[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// MergeStruct 合并相同的类型的结构体指针对象
// 经过反射，性能较低，不建议用于用户级业务处理
// @skip
func MergeStruct[T any](src T, t T) error {
	srcVal := reflect.ValueOf(src).Elem()
	targetVal := reflect.ValueOf(t).Elem()
	mergeValue(srcVal, targetVal)
	return nil
}

// 值合并，反射
func mergeValue(src reflect.Value, t reflect.Value) {
	kind := src.Kind()
	if kind == reflect.Interface {
		// TODO 接口类型处理
	} else if kind == reflect.Struct {
		for i := 0; i < src.NumField(); i++ {
			srcField := src.Field(i)
			tField := t.Field(i)
			mergeValue(srcField, tField)
		}
	} else {
		if !t.IsZero() {
			src.Set(t)
		}
	}
}
