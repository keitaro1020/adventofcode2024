package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sumCenterPageNumber(t *testing.T) {
	type args struct {
		pages [][]int64
	}
	tests := []struct {
		args    args
		want    int64
		wantErr bool
	}{
		{
			args: args{
				pages: [][]int64{
					{1, 2, 3},
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			args: args{
				pages: [][]int64{
					{1, 2, 3},
					{4, 5, 6, 7, 8},
				},
			},
			want:    8,
			wantErr: false,
		},
		{
			args: args{
				pages: [][]int64{
					{1, 2, 3},
					{4, 5, 6, 7},
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v,wantErr:%v", tt.args, tt.want, tt.wantErr), func(t *testing.T) {
			got, err := sumCenterPageNumber(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumCenterPageNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sumCenterPageNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPageRule(t *testing.T) {
	type args struct {
		pageRule     []int64
		inputedPages []int64
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				pageRule:     []int64{2, 3},
				inputedPages: []int64{1},
			},
			want: true,
		},
		{
			args: args{
				pageRule:     []int64{},
				inputedPages: []int64{1, 2, 3},
			},
			want: true,
		},
		{
			args: args{
				pageRule:     []int64{2, 3},
				inputedPages: []int64{},
			},
			want: true,
		},
		{
			args: args{
				pageRule:     []int64{2, 3, 5},
				inputedPages: []int64{1, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v", tt.args, tt.want), func(t *testing.T) {
			if got := checkPageRule(tt.args.pageRule, tt.args.inputedPages); got != tt.want {
				t.Errorf("checkPageRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePageRule(t *testing.T) {
	type args struct {
		line     string
		pageRule map[int64][]int64
	}
	tests := []struct {
		args    args
		want    map[int64][]int64
		wantErr bool
	}{
		{
			args: args{
				line:     "1|2",
				pageRule: map[int64][]int64{},
			},
			want: map[int64][]int64{
				1: {2},
			},
			wantErr: false,
		},
		{
			args: args{
				line: "1|3",
				pageRule: map[int64][]int64{
					1: {2},
				},
			},
			want: map[int64][]int64{
				1: {2, 3},
			},
			wantErr: false,
		},
		{
			args: args{
				line: "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				line: "1|a",
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				line: "a|1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v,wantErr:%v", tt.args, tt.want, tt.wantErr), func(t *testing.T) {
			got, err := parsePageRule(tt.args.line, tt.args.pageRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePageRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePageRule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePages(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args    args
		want    []int64
		wantErr bool
	}{
		{
			args: args{
				line: "1,2,3",
			},
			want:    []int64{1, 2, 3},
			wantErr: false,
		},
		{
			args: args{
				line: "1,2,3,4",
			},
			want:    []int64{1, 2, 3, 4},
			wantErr: false,
		},
		{
			args: args{
				line: "1,2,a",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v,wantErr:%v", tt.args, tt.want, tt.wantErr), func(t *testing.T) {
			got, err := parsePages(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterPages(t *testing.T) {
	type args struct {
		pageRule map[int64][]int64
		pages    [][]int64
	}
	tests := []struct {
		args args
		want [][]int64
	}{
		{
			args: args{
				pageRule: map[int64][]int64{
					75: {29, 53, 47, 61, 13},
					47: {53},
					61: {13, 53, 29},
					53: {29, 13},
					97: {13, 61, 47, 29, 53, 75},
					29: {13},
				},
				pages: [][]int64{
					{75, 47, 61, 53, 29},
					{97, 61, 53, 29, 13},
					{75, 29, 13},
					{75, 97, 47, 61, 53},
					{61, 13, 29},
					{97, 13, 75, 29, 47},
				},
			},
			want: [][]int64{
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v", tt.args, tt.want), func(t *testing.T) {
			if got := filterPages(tt.args.pageRule, tt.args.pages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPageLine(t *testing.T) {
	pageRule := map[int64][]int64{
		75: {29, 53, 47, 61, 13},
		47: {53},
		61: {13, 53, 29},
		53: {29, 13},
		97: {13, 61, 47, 29, 53, 75},
		29: {13},
	}
	type args struct {
		pageRule map[int64][]int64
		pageLine []int64
	}
	tests := []struct {
		args args
		want []int64
	}{
		{
			args: args{
				pageRule: pageRule,
				pageLine: []int64{75, 97, 47, 61, 53},
			},
			want: []int64{97, 75, 47, 61, 53},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("args:%v,want:%v", tt.args, tt.want), func(t *testing.T) {
			if got := sortPageLine(tt.args.pageRule, tt.args.pageLine); !matchInt64Array(t, got, tt.want) {
				t.Errorf("sortPageLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func matchInt64Array(t *testing.T, a, b []int64) bool {
	fmt.Printf("a: %v, b: %v\n", a, b)
	if len(a) != len(b) {
		fmt.Println("len(a) != len(b)")
		return false
	}
	for i, v := range a {
		if v != b[i] {
			fmt.Printf("v != b[i] => v: %d, b[i]: %d\n", v, b[i])
			return false
		}
	}
	fmt.Println("return true")
	return true
}
