package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil

}

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	total := initial
	add := func(x int) {
		total += x
	}
	subtract := func(x int) {
		total -= x
	}
	get := func() int {
		return total
	}
	return add, subtract, get
}
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
func Reduce(nums []int, initial int, operation func(int, int) int) int {
	acc := initial
	for _, v := range nums {
		acc = operation(acc, v)
	}
	return acc
}

func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Println("===== Process Information =====")

	// A process ID (PID) is a unique number assigned by the operating system
	// to identify each running program.
	fmt.Println("Current Process PID: ", os.Getpid())

	// The Parent process ID is the PID of the program that launched this one
	fmt.Println("Parent Process PID: ", os.Getppid())

	// Create sample data in a slice
	data := []int{1, 2, 3, 4, 5}

	// The slice header is a small structure in GO that contains:
	// 1) a pointer to the underlying array
	// 2) the length of the slice
	// 3) the capacity of the slice
	// Printing &data shows the address of this slice header, NOT the array itself
	fmt.Printf("Memory address of slice: %p\n", &data)

	// This is the actual memory address of the fist element
	// inside the underlying array that the slice refers to.
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	// Process isolation mean each running program has its own  protected memory space.
	// Other processes cannot read or modify these memory addresses directly,
	// which prevents bugs, crashes, and security violations.
	fmt.Println("Note: Other processes cannot access these addresses due to process isolation.")

}
func DoubleValue(x int) {
	x = x * 2
	// Does not change original because GO passes values, not variables.
}

func DoublePointer(x *int) {
	*x = *x * 2
	// DOES change original because we modify memory directly
}

func CreateOnStack() int {
	v := 10
	return v // stays on stack
}

func CreateOnHeap() *int {
	v := 20
	return &v //escapes to heap
}

func SwapValues(a, b int) (int, int) {
	return b, a
}
func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}
