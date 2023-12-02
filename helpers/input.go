package helpers

import (
	"os"
	"path"
	"runtime"
	"strconv"
)

func ReadInput() string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input.txt")

	if err != nil {
		return ""
	}

	return string(input)
}

func ToInt(value string) int {
	result, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return result
}
