package Reflect

import (
	"errors"
	"reflect"
)

// FieldInfo : field info
type FieldInfo struct {
	Name  string
	Type  reflect.Type
	Kind  reflect.Kind
	Tags  reflect.StructTag
	Value reflect.Value
}

// SetValue : set value to field, field must be kind of reflect.Ptr
func (field FieldInfo) SetValue(v interface{}) {
	if field.Value.CanSet() {
		field.Value.Set(reflect.ValueOf(v))
	}
}

// GetValue : get value of field
func (field FieldInfo) GetValue() interface{} {
	return field.Value.Interface()
}

// AsTypeInfo : convert field to TypeInfo
func (field FieldInfo) AsTypeInfo() (TypeInfo, error) {
	if field.Kind == reflect.Struct || field.Kind == reflect.Ptr {
		return GetTypeInfoWithValueType(field.Value, field.Type)
	}
	return TypeInfo{}, errors.New("must be struct")
}
