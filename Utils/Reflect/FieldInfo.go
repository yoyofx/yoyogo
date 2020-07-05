package Reflect

import (
	"errors"
	"reflect"
)

type FieldInfo struct {
	Name  string
	Type  reflect.Type
	Kind  reflect.Kind
	Tags  reflect.StructTag
	Value reflect.Value
}

func (field FieldInfo) SetValue(v interface{}) {
	field.Value.Set(reflect.ValueOf(v))
}

func (field FieldInfo) GetValue() interface{} {
	return field.Value.Interface()
}

func (field FieldInfo) AsTypeInfo() (TypeInfo, error) {
	if field.Kind == reflect.Struct || field.Kind == reflect.Ptr {
		return GetTypeInfoWithValueType(field.Value, field.Type)
	}
	return TypeInfo{}, errors.New("must be struct")
}
