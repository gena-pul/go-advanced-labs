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

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name      string
		start     int
		calls     int
		wantAfter int
	}{
		{"start at 0, call once", 0, 1, 1},
		{"start at 0, call twice", 0, 2, 2},
		{"start at 10, call once", 10, 1, 11},
		{"independent counters", 5, 3, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			var got int
			for i := 0; i < tt.calls; i++ {
				got = counter()
			}

			if got != tt.wantAfter {
				t.Errorf("got %d, want %d", got, tt.wantAfter)
			}
		})
	}
}
func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{"double 5", 2, 5, 10},
		{"3 is prime", 3, 5, 15},
		{"zero factor", 0, 7, 0},
		{"negative factor", -2, 4, -8},
		{"factor 1", 1, 9, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mult := MakeMultiplier(tt.factor)
			got := mult(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name      string
		initial   int
		addVal    int
		subVal    int
		wantFinal int
	}{
		{"basic case", 100, 50, 30, 120},
		{"start at zero", 0, 10, 5, 5},
		{"all negative", -10, -5, -5, -10},
		{"no change", 20, 0, 0, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, sub, get := MakeAccumulator(tt.initial)
			add(tt.addVal)
			sub(tt.subVal)
			if get() != tt.wantFinal {
				t.Errorf("got %d, want %d", get(), tt.wantFinal)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		operation func(int) int
		want      []int
	}{
		{
			name:      "square numbers",
			input:     []int{1, 2, 3, 4},
			operation: func(x int) int { return x * x },
			want:      []int{1, 4, 9, 16},
		},
		{
			name:      "double numbers",
			input:     []int{1, 2, 3},
			operation: func(x int) int { return x * 2 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "negate numbers",
			input:     []int{1, -2, 3},
			operation: func(x int) int { return -x },
			want:      []int{-1, 2, -3},
		},
		{
			name:      "empty slice",
			input:     []int{},
			operation: func(x int) int { return x + 1 },
			want:      []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.input, tt.operation)
			if len(got) != len(tt.want) {
				t.Fatalf("length mismatch: got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("index %d: got %d want %d", i, got[i], tt.want[i])
				}
			}
			if &got == &tt.input {
				t.Errorf("Apply modified the original slice")
			}
		})
	}
}
func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "only evens",
			input:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "only positives",
			input:     []int{-2, -1, 0, 1, 2},
			predicate: func(x int) bool { return x > 0 },
			want:      []int{1, 2},
		},
		{
			name:      "greater than 10",
			input:     []int{5, 10, 15, 20},
			predicate: func(x int) bool { return x > 10 },
			want:      []int{15, 20},
		},
		{
			name:      "nothing passes",
			input:     []int{1, 3, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.input, tt.predicate)
			if len(got) != len(tt.want) {
				t.Fatalf("length mismatch: got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("index %d: got %d want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{
			name:      "sum",
			input:     []int{1, 2, 3, 4},
			initial:   0,
			operation: func(acc, x int) int { return acc + x },
			want:      10,
		},
		{
			name:      "product",
			input:     []int{1, 2, 3, 4},
			initial:   1,
			operation: func(acc, x int) int { return acc * x },
			want:      24,
		},
		{
			name:    "max",
			input:   []int{5, 2, 9, 1},
			initial: 0,
			operation: func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			want: 9,
		},
		{
			name:    "min",
			input:   []int{5, 2, 9, 1},
			initial: 100,
			operation: func(acc, x int) int {
				if x < acc {
					return x
				}
				return acc
			},
			want: 1,
		},
		{
			name:      "empty slice",
			input:     []int{},
			initial:   42,
			operation: func(acc, x int) int { return acc + x },
			want:      42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.input, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
func TestCompose(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{
			name:  "double then add two",
			g:     func(x int) int { return x * 2 },
			f:     func(x int) int { return x + 2 },
			input: 5,
			want:  12,
		},
		{
			name:  "add five then square",
			g:     func(x int) int { return x + 5 },
			f:     func(x int) int { return x * x },
			input: 3,
			want:  64,
		},
		{
			name:  "negate then double",
			g:     func(x int) int { return -x },
			f:     func(x int) int { return x * 2 },
			input: 4,
			want:  -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.input)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"basic swap", 1, 2, 2, 1},
		{"same values", 5, 5, 5, 5},
		{"negatives", -1, 3, 3, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			SwapPointers(&a, &b)
			if a != tt.wantA || b != tt.wantB {
				t.Errorf("got (%d, %d), want (%d,%d)",
					a, b, tt.wantA, tt.wantB)
			}
		})
	}
}
