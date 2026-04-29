// Package funcs
package funcs

import (
	"fmt"
	"reflect"
	"server/internal/err/panics"
)

func Ternary(a any, b any, result bool) any {
	if result {
		return a
	}
	return b
}

func SPrintStruct[T any](data T) string {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() != reflect.Struct || v.Kind() != reflect.Struct {
		panics.PanicMisuse("SPrintStruct", "wrong type")
	}
	ret := t.Name() + " {\n"
	for i := range t.NumField() {
		ret += "  " + t.Field(i).Name + ": " + fmt.Sprintf("%v\n", v.Field(i))
	}
	ret += "}\n"
	return ret
}

func TypeVal(data any) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() == reflect.Pointer || v.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct || v.Kind() != reflect.Struct {
		panics.PanicMisuse("Type Val", "wrong type")
	}
	return t, v
}
