package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseMuls(t *testing.T) {
	tests := []struct {
		lines []string
		want  []string
	}{
		{
			lines: []string{"mul(683,461)"},
			want:  []string{"mul(683,461)"},
		},
		{
			lines: []string{"mul(2,3)", "mul(5,20)"},
			want:  []string{"mul(2,3)", "mul(5,20)"},
		},
		{
			lines: []string{"mul(2,3)aaaaaaaaamul(500,2)"},
			want:  []string{"mul(2,3)", "mul(500,2)"},
		},
		{
			lines: []string{"mul(2,3)*****ul(2,4)mul(5,2)fgdfgdfgmul(8,2sdtermul(5.4)fgdemul(7,1)fgdfgdfgfdgdfg"},
			want:  []string{"mul(2,3)", "mul(5,2)", "mul(7,1)"},
		},
		{
			lines: []string{"@~don't()mul(683,461)", "mul(2,3)mul(5,20)"},
			want:  nil,
		},
		{
			lines: []string{"@~don't()mul(683,461)", "mul(2,3)do()mul(5,20)"},
			want:  []string{"mul(5,20)"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v => %v", tt.lines, tt.want), func(t *testing.T) {
			if got := parseMuls(tt.lines); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("calcLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcMuls(t *testing.T) {

	tests := []struct {
		muls []string
		want int64
	}{
		{
			muls: []string{"mul(2,3)", "mul(5,20)"},
			want: 106,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v => %v", tt.muls, tt.want), func(t *testing.T) {
			assert.Equalf(t, tt.want, calcMuls(tt.muls), "calcMuls(%v)", tt.muls)
		})
	}
}
