package leetcode

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1",
			args{2},
			[]int{0, 1, 3, 2},
		},
		{"2",
			args{6},
			[]int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8, 24, 25, 27, 26, 30, 31, 29, 28, 20, 21, 23, 22, 18, 19, 17, 16, 48, 49, 51, 50, 54, 55, 53, 52, 60, 61, 63, 62, 58, 59, 57, 56, 40, 41, 43, 42, 46, 47, 45, 44, 36, 37, 39, 38, 34, 35, 33, 32},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
