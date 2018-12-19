package utils

import (
	"encoding/json"
)

func EncodingJson(value interface{}) []byte {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func DecodingJson(value []byte, arg *interface{}) {
	if err := json.Unmarshal(value, &arg); err != nil {
		panic(err)
	}
}
