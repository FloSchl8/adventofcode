package main

import (
	"testing"
)

func Test_day01(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Input",
			args: args{
				input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day01(tt.args.input); got != tt.want {
				t.Errorf("day01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day02(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Input",
			args: args{
				input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day02(tt.args.input); got != tt.want {
				t.Errorf("day02() = %v, want %v", got, tt.want)
			}
		})
	}
}
