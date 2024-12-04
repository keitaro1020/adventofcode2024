package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		i     int
		j     int
		table [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "下方向",
			args: args{
				i: 0,
				j: 1,
				table: [][]string{
					{".", "X", ".", "."},
					{".", "M", ".", "."},
					{".", "A", ".", "."},
					{".", "S", ".", "."},
				},
			},
			want: 1,
		},
		{
			name: "上方向",
			args: args{
				i: 3,
				j: 1,
				table: [][]string{
					{".", "S", ".", "."},
					{".", "A", ".", "."},
					{".", "M", ".", "."},
					{".", "X", ".", "."},
				},
			},
			want: 1,
		},
		{
			name: "右斜め下方向",
			args: args{
				i: 0,
				j: 0,
				table: [][]string{
					{"X", ".", ".", "."},
					{".", "M", ".", "."},
					{".", ".", "A", "."},
					{".", ".", ".", "S"},
				},
			},
			want: 1,
		},
		{
			name: "右斜め上方向",
			args: args{
				i: 3,
				j: 3,
				table: [][]string{
					{"S", ".", ".", "."},
					{".", "A", ".", "."},
					{".", ".", "M", "."},
					{".", ".", ".", "X"},
				},
			},
			want: 1,
		},
		{
			name: "左斜め下方向",
			args: args{
				i: 0,
				j: 0,
				table: [][]string{
					{"X", ".", ".", "X"},
					{".", "M", ".", "."},
					{".", ".", "A", "."},
					{".", ".", ".", "S"},
				},
			},
			want: 1,
		},
		{
			name: "左斜め上方向",
			args: args{
				i: 3,
				j: 3,
				table: [][]string{
					{"S", ".", ".", "."},
					{".", "A", ".", "."},
					{".", ".", "M", "."},
					{".", ".", ".", "X"},
				},
			},
			want: 1,
		},
		{
			name: "右方向",
			args: args{
				i: 0,
				j: 3,
				table: [][]string{
					{"S", "A", "M", "X"},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
				},
			},
			want: 1,
		},
		{
			name: "全方向",
			args: args{
				i: 3,
				j: 3,
				table: [][]string{
					{"S", ".", ".", "S", ".", ".", "S"},
					{".", "A", ".", "A", ".", "A", "."},
					{".", ".", "M", "M", "M", ".", "."},
					{"S", "A", "M", "X", "M", "A", "S"},
					{".", ".", "M", "M", "M", ".", "."},
					{".", "A", ".", "A", ".", "A", "."},
					{"S", ".", ".", "S", ".", ".", "S"},
				},
			},
			want: 8,
		},
		{
			name: "右方向無し",
			args: args{
				i: 0,
				j: 1,
				table: [][]string{
					{".", "X", "M", "A"},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.i, tt.args.j, tt.args.table); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
