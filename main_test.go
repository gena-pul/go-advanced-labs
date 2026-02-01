package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"0 factorial", 0, 1, false},
		{"1 factorial", 1, 1, false},
		{"5 factorial", 5, 120, false},
		{"3 factorial", 3, 6, false},
		{"negative", -1, 0, true},
		{"7 factorial", 7, 5040, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch")
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"2 is prime", 2, true, false},
		{"3 is prime", 3, true, false},
		{"4 not prime", 4, false, false},
		{"9 not prime", 9, false, false},
		{"1 error", 1, false, true},
		{"17 factorial", 17, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch")
				return
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name    string
		base    int
		exp     int
		want    int
		wantErr bool
	}{
		{"2 ^ 3", 2, 3, 8, false},
		{"5 ^ 0", 5, 0, 1, false},
		{"0 ^ 3", 0, 3, 0, false},
		{"1 ^ 10", 1, 10, 1, false},
		{"negative exp", 2, -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("error mismatch")
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
