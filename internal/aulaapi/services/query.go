package services

import (
	"fmt"
	"net/url"
)

// EncodeValue percent-encodes a single query parameter value for safe inclusion in a URL.
// Uses application/x-www-form-urlencoded encoding.
func EncodeValue(value string) string {
	return url.QueryEscape(value)
}

// Param formats a single key=encoded_value query parameter pair.
func Param(key, value string) string {
	return fmt.Sprintf("%s=%s", key, EncodeValue(value))
}

// ParamNum formats a single key=value query parameter pair for numeric values (no encoding needed).
func ParamNum[N interface{ ~int | ~int32 | ~int64 | ~float64 }](key string, value N) string {
	return fmt.Sprintf("%s=%v", key, value)
}
