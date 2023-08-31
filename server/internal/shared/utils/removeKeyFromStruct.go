package utils

import "reflect"

func RemoveKeyFromStruct(p interface{}, keyName string) {
	r := reflect.ValueOf(p).Elem()
	t := r.Type()

	newStruct := reflect.New(t).Elem()

	for i := 0; i < r.NumField(); i++ {
		fieldName := t.Field(i).Name
		if fieldName != keyName {
			newStruct.Field(i).Set(r.Field(i))
		}
	}

	r.Set(newStruct)
}