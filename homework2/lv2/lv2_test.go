package main

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "add",
			args: args{num1: 10, num2: 20},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_control(t *testing.T) {
	type args struct {
		mark string
	}
	tests := []struct {
		name string
		args args
		want operate
	}{
		// TODO: Add test cases.
		{name: "control", args: args{mark: "*"}, want: mul},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := control(tt.args.mark); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("control() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "div", args: args{num1: 20, num2: 20}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jisuan(t *testing.T) {
	type args struct {
		control operate
		num1    int
		num2    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "jisuan", args: args{control: mul, num1: 10, num2: 20}, want: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jisuan(tt.args.control, tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("jisuan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "mul", args: args{num1: 10, num2: 20}, want: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "sub", args: args{num1: 20, num2: 20}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
