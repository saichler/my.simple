package common

import "reflect"

func TypeOf(any interface{}) reflect.Type {
	v := reflect.ValueOf(any)
	if !v.IsValid() {
		return nil
	}
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	return v.Type()
}

func TypeName(any interface{}) string {
	t := TypeOf(any)
	if t == nil {
		return ""
	}
	return t.Name()
}
