package main

import (
	"fmt"
	"testing"
)

func Test_calcLine(t *testing.T) {
	tests := []struct {
		line string
		want int64
	}{
		{
			line: "@~don't()mul(683,461)",
			want: 314863,
		},
		{
			line: "mul(2,3)mul(5,20)",
			want: 106,
		},
		{
			line: "mul(2,3)aaaaaaaaamul(500,2)",
			want: 1006,
		},
		{
			line: "mul(2,3)*****ul(2,4)mul(5,2)fgdfgdfgmul(8,2sdtermul(5.4)fgdemul(7,1)fgdfgdfgfdgdfg",
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v => %v", tt.line, tt.want), func(t *testing.T) {
			if got := calcLine(tt.line); got != tt.want {
				t.Errorf("calcLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
