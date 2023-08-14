package main

import (
	"reflect"
	"testing"
)

func BenchmarkDelSliceElementOld(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DelSliceElementOld[int]([]int{1, 2, 3}, 2)
	}
}

func BenchmarkDelSliceElement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DelSliceElement[int]([]int{1, 2, 3}, 2)
	}
}

func TestDelSliceElement(t *testing.T) {
	type args[T AllowT] struct {
		src   []T
		index int
	}
	type testCase[T AllowT] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}

	testInts := []testCase[int]{
		{
			name: "正常范围index",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: 1,
			},
			want:    []int{1, 3},
			wantErr: false,
		},
		{
			name: "index<0",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: -1,
			},
			want:    []int{1, 2, 3},
			wantErr: true,
		},
		{
			name: "index>length",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: 3,
			},
			want:    []int{1, 2, 3},
			wantErr: true,
		},
	}
	testStrs := []testCase[string]{
		{
			name: "正常范围index",
			args: args[string]{
				src:   []string{"a", "b", "c"},
				index: 1,
			},
			want:    []string{"a", "c"},
			wantErr: false,
		},
		{
			name: "index<0",
			args: args[string]{
				src:   []string{"a", "b", "c"},
				index: -1,
			},
			want:    []string{"a", "b", "c"},
			wantErr: true,
		},
		{
			name: "index>length",
			args: args[string]{
				src:   []string{"a", "b", "c"},
				index: 3,
			},
			want:    []string{"a", "b", "c"},
			wantErr: true,
		},
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DelSliceElement(tt.args.src, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelSliceElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelSliceElement() got = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range testStrs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DelSliceElement(tt.args.src, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelSliceElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelSliceElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}
