package main

import (
	"reflect"
	"testing"
)

func TestRotateArraySpace(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				input: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateArraySpace(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateArraySpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkRotateArraySpace(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		RotateArraySpace(arr)
	}

}

func BenchmarkRotateArray(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		RotateArray(arr)
	}
}

func BenchmarkIsPalindramNumber(b *testing.B) {
	arr := []int64{1221221166693838383, 3322223234444555231}
	for i := 0; i < b.N; i++ {
		IsPalindramNumber(arr[b.N%2])
	}
}

func BenchmarkIsPalindramNumber1(b *testing.B) {
	arr := []int64{1221221166693838383, 3322223234444555231}
	for i := 0; i < b.N; i++ {
		IsPalindramNumber1(arr[b.N%2])
	}
}
