package swagger

import (
	"reflect"
	"strings"
)

// GetSwaggerType returns the type of the swagger type that corresponds to the go type.
func GetSwaggerType(goType string) string {
	if strings.Contains(goType, "file") {
		return "file"
	}
	if strings.HasPrefix(goType, "[]") {
		return "array"
	}
	switch goType {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "time.Time":
		return "string"
	default:
		return "object"
	}
}

func getArrayElementType(goType string) string {
	return goType[2 : len(goType)-1]
}

// ConvertToSwaggerResponse converts a struct to a swagger response.
func ConvertToSwaggerResponse(data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["type"] = "object"
	response["properties"] = make(map[string]interface{})

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = field.Name
		}
		fieldValue := dataValue.Field(i).Interface()

		fieldType := field.Type.String()
		swaggerType := GetSwaggerType(fieldType)
		//if swaggerType == "object" && fieldValue == nil {
		//	fieldValue = reflect.New(field.Type).Elem().Interface()
		//}

		if swaggerType == "array" {
			itemTypeName := GetSwaggerType(field.Type.Elem().String())
			properties := make(map[string]interface{})
			if itemTypeName == "object" {
				objectVal := reflect.New(field.Type.Elem()).Elem().Interface()
				properties = ConvertToSwaggerResponse(objectVal)["properties"].(map[string]interface{})
			}

			itemsProperty := map[string]interface{}{
				"type":       GetSwaggerType(itemTypeName),
				"properties": properties,
			}

			response["properties"].(map[string]interface{})[fieldName] = map[string]interface{}{
				"type":  "array",
				"items": itemsProperty}
		} else {
			response["properties"].(map[string]interface{})[fieldName] = map[string]interface{}{
				"type":        swaggerType,
				"description": field.Tag.Get("doc"),
			}
		}

		if swaggerType == "object" {
			if fieldValue != nil {
				fieldMap := ConvertToSwaggerResponse(fieldValue)
				response["properties"].(map[string]interface{})[fieldName].(map[string]interface{})["properties"] = fieldMap["properties"]
			}
		}
	}

	return response
}
