package calc

import "testing"

// Add 関数のベンチマークテスト
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

// Subtract 関数のベンチマークテスト
func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(30, 20)
	}
}

// Multiply 関数のベンチマークテスト
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(10, 20)
	}
}

// Divide 関数のベンチマークテスト
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 20)
	}
}
