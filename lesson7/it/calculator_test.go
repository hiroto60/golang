package it

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	var calc *Calculator

	// テストデータの準備
	setup := func() {
		calc = NewCalculator()
	}

	// テストデータのクリーンアップ
	teardown := func() {
		calc = nil
	}

	// 加算のテスト
	t.Run("Addition", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(2)
		calc.Add(3)
		assertEqual(t, calc.Result(), 5)
	})

	// 減算のテスト
	t.Run("Subtraction", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(10)
		calc.Subtract(3)
		assertEqual(t, calc.Result(), 7)
	})

	// 乗算のテスト
	t.Run("Multiplication", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(2)
		calc.Multiply(3)
		assertEqual(t, calc.Result(), 6)
	})

	// 除算のテスト
	t.Run("Division", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(10)
		err := calc.Divide(2)
		assertNoError(t, err)
		assertEqual(t, calc.Result(), 5)
	})

	// ゼロ除算のテスト
	t.Run("DivisionByZero", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(10)
		err := calc.Divide(0)
		assertError(t, err)
	})

	// 複雑な計算のテスト
	t.Run("ComplexCalculation", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(2)
		calc.Multiply(3)
		calc.Subtract(1)
		err := calc.Divide(2)
		assertNoError(t, err)
		assertEqual(t, calc.Result(), 2)
	})

}

// 期待値と実際の値が等しいかどうかを検証
func assertEqual(t *testing.T, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

// エラーがnilであることを検証
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// エラーが存在することを検証
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error, got nil")
	}
}
