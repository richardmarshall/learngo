package main

import (
	"reflect"
	"runtime"
)

func Swap(input []int, i, j int) {
	temp := input[j]
	input[j] = input[i]
	input[i] = temp
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
