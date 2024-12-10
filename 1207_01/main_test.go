package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLine_getOperatorCombination(t *testing.T) {
	tests := []struct {
		input   []int64
		want    [][]string
		wantErr bool
	}{
		{
			input: make([]int64, 3),
			want: [][]string{
				{"+", "+"},
				{"+", "*"},
				{"*", "+"},
				{"*", "*"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v", tt.input, tt.want), func(t *testing.T) {
			l := &Line{
				Numbers: tt.input,
			}
			got, err := l.getOperatorCombination()
			if (err != nil) != tt.wantErr {
				t.Errorf("getOperatorCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOperatorCombination() got = %v, want %v", got, tt.want)
			}
		})
	}
}
