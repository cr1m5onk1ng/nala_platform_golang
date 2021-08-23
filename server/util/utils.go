package util

import (
	"fmt"
	"strings"
)

func CheckStringInSlice(slice []string, searchedString string) bool {
	for _, s := range slice {
		if s == searchedString {
			return true
		}
	}
	return false
}

func ParseRequestParam(param string) (string, error) {
	splitted := strings.Split(param, "=")
	if len(splitted) != 2 {
		return "", fmt.Errorf("unable to parse request parameter %s", param)
	}
	return splitted[1], nil
}
