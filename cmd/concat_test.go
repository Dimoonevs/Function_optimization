package main

import (
	"strings"
	"testing"
)

func main() {
}

// Func to generate an array of strings
func generateStringArray(n int) []string {
	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = "string"
	}
	return data
}

func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}
func BenchmarkConcatFunc(b *testing.B) {
	arrayStr := generateStringArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concat(arrayStr)
	}
}

// Optimization code with used strings.Builder
func concatOptimization(str []string) string {
	// Create strings.Builder, in order to avoid unnecessary memory allocation
	var builder strings.Builder
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}
func BenchmarkConcatOptimization(b *testing.B) {
	arrayStr := generateStringArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concatOptimization(arrayStr)
	}
}
