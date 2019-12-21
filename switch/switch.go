package control

import "fmt"

func mySwitch(contents interface{}) (typeName string) {
	switch contents.(type) {
	case string:
		typeName = "string"
	case int:
		typeName = "int"
	case int64:
		typeName = "int64"
	case float32:
		typeName = "float32"
	case []string:
		typeName = "slice"
	default:
		typeName = fmt.Sprintf("%v", contents)
	}
	return typeName
}
