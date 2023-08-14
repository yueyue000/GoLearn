package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = "index:%d out of range"

type AllowT interface {
	int | int8 | int16 | int32 | int64 | string
}

func DelSliceElementOld[T AllowT](s []T, index int) ([]T, error) {
	if index < 0 || index > len(s)-1 {
		return s, errors.New(fmt.Sprintf("index:%d out of memory", index))
	}
	ret := []T{}
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)
	return ret, nil
}

func DelSliceElement[T AllowT](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return src, errors.New(fmt.Sprintf(ErrIndexOutOfRange, index))
	}
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	return Shrink(src)[:length-1], nil
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}
