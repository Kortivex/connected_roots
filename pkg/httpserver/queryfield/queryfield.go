// Package queryfield helps to decode query string fields based on
// the conventions of the public API.
package queryfield

import (
	"strconv"
	"strings"
)

const (
	opSeparator    = ":"
	fieldSeparator = ":"
)

// ops is the translation table for the operators.
var ops = map[string]string{
	"eq":  "==",
	"neq": "!=",
	"gt":  ">",
	"gte": ">=",
	"lt":  "<",
	"lte": "<=",
}

// Decode performs all the operations over a key-value filter and returns the property name,
// the operator decoded and the type decoded.
func Decode(name, value, prefix string) (string, string, string) {
	propertyName := DecodeEmbeddedName(name, prefix)
	operator, valueCleaned := DecodeOperator(value)
	valueType := DecodeType(valueCleaned)

	return propertyName, operator, valueType
}

// DecodeEmbeddedName returns the name of an embed field without its prefix and
// separated by dots.
func DecodeEmbeddedName(name, prefix string) string {
	return strings.ReplaceAll(strings.TrimPrefix(name, prefix), fieldSeparator, ".")
}

// DecodeType returns a string quoted or unquoted based on its "underlying" type.
func DecodeType(value string) string {
	_, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return value
	}

	return `"` + value + `"`
}

// DecodeOperator returns the operator based on the ops table and the value cleaned, without
// the operator prefixed.
func DecodeOperator(value string) (string, string) {
	for op, opDecoded := range ops {
		opPrefix := op + opSeparator
		if strings.HasPrefix(value, opPrefix) {
			return opDecoded, strings.TrimPrefix(value, opPrefix)
		}
	}

	return "==", value
}
