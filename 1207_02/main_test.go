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
				{"+", "||"},
				{"*", "+"},
				{"*", "*"},
				{"*", "||"},
				{"||", "+"},
				{"||", "*"},
				{"||", "||"},
			},
		},
		{
			input: make([]int64, 4),
			want: [][]string{
				{"+", "+", "+"},
				{"+", "+", "*"},
				{"+", "+", "||"},
				{"+", "*", "+"},
				{"+", "*", "*"},
				{"+", "*", "||"},
				{"+", "||", "+"},
				{"+", "||", "*"},
				{"+", "||", "||"},
				{"*", "+", "+"},
				{"*", "+", "*"},
				{"*", "+", "||"},
				{"*", "*", "+"},
				{"*", "*", "*"},
				{"*", "*", "||"},
				{"*", "||", "+"},
				{"*", "||", "*"},
				{"*", "||", "||"},
				{"||", "+", "+"},
				{"||", "+", "*"},
				{"||", "+", "||"},
				{"||", "*", "+"},
				{"||", "*", "*"},
				{"||", "*", "||"},
				{"||", "||", "+"},
				{"||", "||", "*"},
				{"||", "||", "||"},
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
				TestValue: 159, // 10 + 5 || 6 + 3
				Numbers:   []int64{10, 5, 6, 3},
			},
			want:    true,
			wantErr: false,
		},
		{
			fields: fields{
				TestValue: 1518, // 10 * 5 || 6 * 3
				Numbers:   []int64{10, 5, 6, 3},
			},
			want:    true,
			wantErr: false,
		},
		{
			fields: fields{
				TestValue: 1, // invalid
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

func Test_calc(t *testing.T) {
	type args struct {
		numbers   []int64
		operators []string
	}
	tests := []struct {
		args    args
		want    int64
		wantErr bool
	}{
		{
			args: args{
				numbers:   []int64{10, 5},
				operators: []string{"+"},
			},
			want:    15,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5},
				operators: []string{"*"},
			},
			want:    50,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5},
				operators: []string{"||"},
			},
			want:    105,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"+", "+"},
			},
			want:    18,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"+", "*"},
			},
			want:    45,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"+", "||"},
			},
			want:    153,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"*", "+"},
			},
			want:    53,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"*", "*"},
			},
			want:    150,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"*", "||"},
			},
			want:    503,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"||", "+"},
			},
			want:    108,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"||", "*"},
			},
			want:    315,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 3},
				operators: []string{"||", "||"},
			},
			want:    1053,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 6, 3},
				operators: []string{"+", "+", "+"},
			},
			want:    24,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 6, 3},
				operators: []string{"+", "||", "+"},
			},
			want:    159,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{10, 5, 6, 3},
				operators: []string{"||", "+", "+"},
			},
			want:    114,
			wantErr: false,
		},
		{
			args: args{
				numbers:   []int64{6, 8, 6, 15},
				operators: []string{"*", "||", "*"},
			},
			want:    7290,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v", tt.args, tt.want), func(t *testing.T) {
			got, err := calc(tt.args.numbers, tt.args.operators)
			if (err != nil) != tt.wantErr {
				t.Errorf("calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}