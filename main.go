package main

import (
	"errors"
	"fmt"
)

type AllowT interface {
	int | int8 | int16 | int32 | int64 | string
}

func DelSliceElement[T AllowT](s []T, index int) ([]T, error) {
	if index < 0 || index > len(s)-1 {
		return s, errors.New(fmt.Sprintf("index:%d out of memory", index))
	}
	ret := []T{}
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)
	return ret, nil
}

func main() {
	r1, err1 := DelSliceElement([]int{1, 2, 3}, 1)
	r2, err2 := DelSliceElement([]int{1, 2, 3}, -1)
	r3, err3 := DelSliceElement([]int{1, 2, 3}, 3)
	r4, err4 := DelSliceElement([]string{"a", "b", "c"}, 0)
	r5, err5 := DelSliceElement([]string{"a", "b", "c"}, -1)
	r6, err6 := DelSliceElement([]string{"a", "b", "c"}, 2)

	println("err1", err1, fmt.Sprintf("===%v,%v", r1, err1))
	println("err2", err2, fmt.Sprintf("===%v,%v", r2, err2))
	println("err3", err3, fmt.Sprintf("===%v,%v", r3, err3))
	println(" err4", err4, fmt.Sprintf("===%v,%v", r4, err4))
	println("err5", err5, fmt.Sprintf("===%v,%v", r5, err5))
	println("err6", err6, fmt.Sprintf("===%v,%v", r6, err6))
}
