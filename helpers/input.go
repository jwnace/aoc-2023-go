package helpers

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput() string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input.txt")

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(input))
}

func ReadInputLines() []string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input.txt")

	if err != nil {
		return []string{}
	}

	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

func ToInt(value string) int {
	result, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return result
}
