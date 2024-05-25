package utilities

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"
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
	dir = GoUpLevel(dir, 1)

	fmt.Println("Searching for files...")

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}
		if filepath.Ext(info.Name()) == ".locksmith" {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(matches, err)
		return
	}
	return matches, err
}

func GoUpLevel(dir string, level int) string {
	for range level {
		dir = filepath.Dir(dir)
	}
	return dir
}

func PrintOperationTime(ctx context.Context, Key CtxKey) {

	val := ctx.Value(Key).(Val)
	total := time.Since(val.Time)

	cpuTime := total - val.UserTime
	fmt.Fprintf(os.Stderr, "\nOperation Completed in\nTotal:%s UserWaitTime:%s CpuTime:%s \n", total, val.UserTime, cpuTime)
}

func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("\nRecovered from panic:", r)
	}
}

func ReadCtx(ctx context.Context, Key CtxKey) Val {
	val := ctx.Value(Key)
	value := val.(Val)
	return value
}

type CtxKey string
type Val struct {
	Time     time.Time
	UserTime time.Duration
}
