package utilities

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func PrintStruct(s interface{}) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(s).Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func DetectFile() (matches []string, err error) {
	dir, _ := os.Getwd()
	dir = GoUpLevel(dir, 2)
	pattern := filepath.Join(dir, "*", "*", "*.locksmith")
	matches, err = filepath.Glob(pattern)
	return matches, err
}

func GoUpLevel(dir string, level int) string {
	for range level {
		dir = filepath.Dir(dir)
	}
	return dir
}
