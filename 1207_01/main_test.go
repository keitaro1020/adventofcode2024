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

func TestLine_determineTrueLine(t *testing.T) {
	type fields struct {
		TestValue int64
		Numbers   []int64
	}
	tests := []struct {
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			fields: fields{
				TestValue: 24, // 10 + 5 + 6 + 3
				Numbers:   []int64{10, 5, 6, 3},
			},
			want:    true,
			wantErr: false,
		},
		{
			fields: fields{
				TestValue: 69, // 10 + 5 || 6 + 3
				Numbers:   []int64{10, 5, 6, 3},
			},
			want:    false,
			wantErr: false,
		},
		{
			fields: fields{
				TestValue: 1680, // 10 * 5 || 6 * 3
				Numbers:   []int64{10, 5, 6, 3},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("fields:%v,want:%v", tt.fields, tt.want), func(t *testing.T) {
			l := &Line{
				TestValue: tt.fields.TestValue,
				Numbers:   tt.fields.Numbers,
			}
			got, err := l.determineTrueLine()
			if (err != nil) != tt.wantErr {
				t.Errorf("determineTrueLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("determineTrueLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
