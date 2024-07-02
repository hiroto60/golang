package calc

import (
	"testing"
)

// 加算関数 Add のテストを行う
func TestAdd(t *testing.T) {
	// テストケースのスライスを定義
	// 各ケースは加算の入力値 (a, b) と期待する結果 (expected)
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-5, -3, -8},
		{7, -2, 5},
		{0, 5, 5},
		{3, 0, 3},
	}

	// 各テストケースに対してループを実行
	for _, tc := range testCases {
		result := Add(tc.a, tc.b) // Add 関数を呼出
		// 加算の結果が期待値と異なる場合、エラーメッセージを出力
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

// 減算関数 Subtract のテストを行う
func TestSubtract(t *testing.T) {
	// 減算のテストケースを定義
	testCases := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{-3, -5, 2},
		{7, -2, 9},
		{5, 0, 5},
		{0, 3, -3},
	}

	// 各テストケースで Subtract 関数をテスト
	for _, tc := range testCases {
		result := Subtract(tc.a, tc.b)
		// 減算結果が期待値と異なる場合、エラーメッセージを出力
		if result != tc.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

// 乗算関数 Multiply のテストを行う
func TestMultiply(t *testing.T) {
	// 乗算のテストケースを定義
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, -3, 6},
		{2, -3, -6},
		{5, 0, 0},
		{0, 3, 0},
	}

	// 各テストケースで Multiply 関数をテスト
	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		// 乗算結果が期待値と異なる場合、エラーメッセージを出力
		if result != tc.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

// 除算関数 Divide のテストを行う
func TestDivide(t *testing.T) {
	// 除算のテストケースを定義します。テストケースには除算可能かどうかのエラー状態も含む
	testCases := []struct {
		a, b          int
		expected      int
		expectedError bool
	}{
		{6, 3, 2, false},
		{-6, -3, 2, false},
		{6, -3, -2, false},
		{0, 5, 0, false},
		{5, 0, 0, true}, // 0 で割る場合、エラーが期待
		{0, 0, 0, true}, // 0 で割る場合、エラーが期待
	}

	// 各テストケースで Divide 関数をテスト
	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectedError {
			// エラーが期待されている場合、エラーが返されなかったらテストを失敗させる
			if err == nil {
				t.Errorf("Divide(%d, %d) did not return an error", tc.a, tc.b)
			}
		} else {
			// エラーが期待されていない場合、エラーが返されたらテストを失敗させる
			if err != nil {
				t.Errorf("Divide(%d, %d) returned an error: %v", tc.a, tc.b, err)
			}
			// 除算結果が期待値と異なる場合、エラーメッセージを出力
			if result != tc.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		}
	}
}
