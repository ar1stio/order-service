package util

import "encoding/json"

func StrStruct(structs interface{}) string {
	str, _ := json.Marshal(structs)
	return string(str)
}

func JsonStruct(structs interface{}) []byte {
	json, _ := json.Marshal(structs)
	return json
}
