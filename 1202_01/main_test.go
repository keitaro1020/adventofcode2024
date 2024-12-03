package main

import (
	"fmt"
	"testing"
)

func Test_checkSafe(t *testing.T) {
	tests := []struct {
		args []int64
		want bool
	}{
		{
			args: []int64{1, 2},
			want: true,
		},
		{
			args: []int64{2, 1},
			want: true,
		},
		{
			args: []int64{4, 2, 1},
			want: true,
		},
		{
			args: []int64{1, 2, 4},
			want: true,
		},
		{
			args: []int64{5, 2, 1},
			want: true,
		},
		{
			args: []int64{1, 2, 5},
			want: true,
		},
		{
			args: []int64{2, 2},
			want: false,
		},
		{
			args: []int64{4, 2, 2},
			want: false,
		},
		{
			args: []int64{2, 2, 4},
			want: false,
		},
		{
			args: []int64{6, 2, 1},
			want: false,
		},
		{
			args: []int64{1, 2, 6},
			want: false,
		},
		{
			args: []int64{1, 2, 4, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v => %v", tt.args, tt.want), func(t *testing.T) {
			if got := checkSafe(tt.args); got != tt.want {
				t.Errorf("checkSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
