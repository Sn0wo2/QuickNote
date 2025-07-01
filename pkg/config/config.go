package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Sn0wo2/QuickNote/pkg/debug"
	"github.com/Sn0wo2/QuickNote/pkg/helper"
	"gopkg.in/yaml.v3"
)

var Instance Config

// Init contains validation config
func Init() error {
	cfgPath := "./data/config.yml"
	if debug.IsDebug() {
		cfgPath = "./data/config_test.yml"
	}

	cf, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(cf, &Instance); err != nil {
		return err
	}

	err = validate(&Instance)
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

func validate(cfg any) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := range t.NumField() {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}

		optional := field.Tag.Get("optional")
		if optional == "true" {
			continue
		}

		value := v.Field(i)

		var isEmpty bool

		//nolint:exhaustive
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
			if err := validate(value.Addr().Interface()); err != nil {
				return fmt.Errorf("%s.%s: %w", t.Name(), field.Name, err)
			}
		// reflect.Bool, reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Invalid
		default:
		}

		if isEmpty {
			return fmt.Errorf("config field [%s] is required but empty", field.Name)
		}
	}

	return nil
}
