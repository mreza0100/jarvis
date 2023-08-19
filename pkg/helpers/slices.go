package helpers

import (
	"strconv"
	"strings"
)

func MapKeysToInterfaceArray(in map[int64]struct{}) []interface{} {
	result := make([]interface{}, 0, len(in))

	for key := range in {
		result = append(result, key)
	}

	return result
}

func StringSliceContainsString(slice []string, value string) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

func Int64SliceContainsInt64(slice []int64, value int64) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

func StringSliceToCommaSeparatedString(slice []string) string {
	var result strings.Builder
	for _, element := range slice {
		result.WriteString(element)
		result.WriteRune(',')
	}
	return result.String()
}

func CommaSeparatedStringToStringSlice(inputString string) []string {
	elements := strings.Split(inputString, ",")
	elements = elements[0 : len(elements)-1]
	return elements
}

func Int64SliceToCommaSeparatedString(slice []int64) string {
	var result strings.Builder
	for _, element := range slice {
		result.WriteString(strconv.FormatInt(element, 10))
		result.WriteRune(',')
	}
	return result.String()
}

func CommaSeparatedStringToInt64Slice(inputString string) []int64 {
	elements := strings.Split(inputString, ",")
	elements = elements[0 : len(elements)-1]
	result := make([]int64, 0, len(elements))
	for _, element := range elements {
		value, _ := strconv.ParseInt(element, 10, 64)
		result = append(result, value)
	}
	return result
}
