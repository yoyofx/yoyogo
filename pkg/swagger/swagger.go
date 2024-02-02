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
	switch goType {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
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
			response["properties"].(map[string]interface{})[fieldName] = map[string]interface{}{
				"type":  "array",
				"items": map[string]interface{}{"type": GetSwaggerType(getArrayElementType(fieldType))},
			}
		} else {
			response["properties"].(map[string]interface{})[fieldName] = map[string]interface{}{
				"type": swaggerType,
			}
		}

		if swaggerType == "object" {
			if fieldValue != nil {
				response["properties"].(map[string]interface{})[fieldName].(map[string]interface{})["properties"] = ConvertToSwaggerResponse(fieldValue)
			}
		}
	}

	return response
}
