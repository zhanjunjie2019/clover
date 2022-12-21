package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// LoadEnvToStruct 加载环境变量到指针对象中
// 经过反射，性能较低，不建议用于用户级业务处理
func LoadEnvToStruct(conf any) error {
	v := reflect.ValueOf(conf).Elem()
	t := reflect.TypeOf(conf).Elem()
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		err := loadEnvToValue(v.Field(i), structField)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadEnvToValue(v reflect.Value, t reflect.StructField) error {
	kind := v.Kind()
	if kind == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			structField := t.Type.Field(i)
			err := loadEnvToValue(v.Field(i), structField)
			if err != nil {
				return err
			}
		}
	} else {
		envKey := t.Tag.Get("env")
		if len(envKey) > 0 {
			envVal := os.Getenv(envKey)
			if len(envVal) > 0 {
				if kind == reflect.String {
					v.SetString(envVal)
				} else if kind == reflect.Bool {
					parseBool, err := strconv.ParseBool(envVal)
					if err != nil {
						return err
					}
					v.SetBool(parseBool)
				} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
					atoi, err := strconv.Atoi(envVal)
					if err != nil {
						return err
					}
					v.SetInt(int64(atoi))
				} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
					atoi, err := strconv.Atoi(envVal)
					if err != nil {
						return err
					}
					v.SetUint(uint64(atoi))
				} else {
					return fmt.Errorf("file kind type error")
				}
			}
		}
	}
	return nil
}
