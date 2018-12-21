package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func EncodingJson(value interface{}) []byte {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func DecodingJson(value []byte, arg interface{}) {
	if err := json.Unmarshal(value, &arg); err != nil {
		panic(err)
	}
}

func SetCookie(key string, value string, httpOnly bool, c *gin.Context) {
	c.SetCookie(key, value, 0, "", "", false, httpOnly)
}
