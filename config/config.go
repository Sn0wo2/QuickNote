package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Sn0wo2/QuickNote/helper"
	"gopkg.in/yaml.v3"
)

var Instance Config

// Init or Reload config
func Init() error {
	cf, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(cf, &Instance); err != nil {
		return err
	}

	err = ValidateConfig(&Instance)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) String() string {
	cb, err := yaml.Marshal(*c)
	if err != nil {
		return ""
	}

	return helper.BytesToString(cb)
}

func ValidateConfig(cfg any) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		optional := field.Tag.Get("optional")
		if optional == "true" {
			continue
		}
		value := v.Field(i)
		isEmpty := false
		switch value.Kind() {
		case reflect.String:
			isEmpty = strings.TrimSpace(value.String()) == ""
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			isEmpty = value.Int() == 0
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			isEmpty = value.Uint() == 0
		case reflect.Float32, reflect.Float64:
			isEmpty = value.Float() == 0
		case reflect.Ptr, reflect.Interface:
			isEmpty = value.IsNil()
		case reflect.Slice, reflect.Map, reflect.Array:
			isEmpty = value.Len() == 0
		case reflect.Struct:
			/*if err := ValidateConfig(value.Addr().Interface()); err != nil {
				return fmt.Errorf("%s.%s: %w", t.Name(), field.Name, err)
			}*/
		case reflect.Bool, reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Invalid:
		default:
		}
		if isEmpty {
			return fmt.Errorf("config field [%s] is required but empty", field.Name)
		}
	}
	return nil
}
