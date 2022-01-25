package tool

import "strings"

func GetFileSuffix(fileName string) string {
	position := strings.Index(fileName, ".")
	if position == -1 {
		return ""
	}
	position++
	return fileName[position:]
}
