package main

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// handle pointer x
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	// loop invariant
	case reflect.String:
		fn(val.String())
	}
}
