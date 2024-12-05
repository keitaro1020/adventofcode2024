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
			name: "pattern1",
			args: args{
				i: 1,
				j: 1,
				table: [][]string{
					{"M", ".", "S"},
					{".", "A", "."},
					{"M", ".", "S"},
				},
			},
			want: 1,
		},
		{
			name: "pattern2",
			args: args{
				i: 1,
				j: 1,
				table: [][]string{
					{"M", ".", "M"},
					{".", "A", "."},
					{"S", ".", "S"},
				},
			},
			want: 1,
		},
		{
			name: "pattern3",
			args: args{
				i: 1,
				j: 1,
				table: [][]string{
					{"S", ".", "M"},
					{".", "A", "."},
					{"S", ".", "M"},
				},
			},
			want: 1,
		},
		{
			name: "pattern4",
			args: args{
				i: 1,
				j: 1,
				table: [][]string{
					{"S", ".", "S"},
					{".", "A", "."},
					{"M", ".", "M"},
				},
			},
			want: 1,
		},
		{
			name: "invalid1",
			args: args{
				i: 0,
				j: 0,
				table: [][]string{
					{"A", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid2",
			args: args{
				i: 0,
				j: 1,
				table: [][]string{
					{".", "A", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid3",
			args: args{
				i: 0,
				j: 2,
				table: [][]string{
					{".", ".", "A"},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid4",
			args: args{
				i: 1,
				j: 2,
				table: [][]string{
					{".", ".", "."},
					{".", ".", "A"},
					{".", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid5",
			args: args{
				i: 2,
				j: 2,
				table: [][]string{
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "A"},
				},
			},
			want: 0,
		},
		{
			name: "invalid6",
			args: args{
				i: 2,
				j: 1,
				table: [][]string{
					{".", ".", "."},
					{".", ".", "."},
					{".", "A", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid7",
			args: args{
				i: 2,
				j: 0,
				table: [][]string{
					{".", ".", "."},
					{".", ".", "."},
					{"A", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid8",
			args: args{
				i: 1,
				j: 0,
				table: [][]string{
					{".", ".", "."},
					{"A", ".", "."},
					{".", ".", "."},
				},
			},
			want: 0,
		},
		{
			name: "invalid9",
			args: args{
				i: 1,
				j: 1,
				table: [][]string{
					{".", ".", "."},
					{".", "A", "."},
					{".", ".", "."},
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
