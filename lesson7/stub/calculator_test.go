package stub

import (
	"testing"
)

type StubLogger struct {
	messages []string
}

func (l *StubLogger) Log(message string) {
	l.messages = append(l.messages, message)
}

func TestCalculator(t *testing.T) {
	var calc *Calculator
	var logger *StubLogger

	// テストデータの準備
	setup := func() {
		logger = &StubLogger{}
		calc = NewCalculator(logger)
	}

	// テストデータのクリーンアップ
	teardown := func() {
		calc = nil
		logger = nil
	}

	// 加算のテスト
	t.Run("Addition", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(2)
		calc.Add(3)
		assertEqual(t, calc.Result(), 5)
		assertLogMessage(t, logger.messages, "Add(2): result = 2")
		assertLogMessage(t, logger.messages, "Add(3): result = 5")
	}) // 減算のテスト t.Run("Subtraction", func(t *testing.T) { setup() defer teardown() calc.Add(10) calc.Subtract(3) assertEqual(t, calc.Result(), 7) assertLogMessage(t, logger.messages, "Add(10): result = 10") assertLogMessage(t, logger.messages, "Subtract(3): result = 7") }) // 乗算のテスト
	t.Run("Multiplication", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(2)
		calc.Multiply(3)
		assertEqual(t, calc.Result(), 6)
		assertLogMessage(t, logger.messages, "Add(2): result = 2")
		assertLogMessage(t, logger.messages, "Multiply(3): result = 6")
	})

	// 除算のテスト
	t.Run("Division", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(10)
		err := calc.Divide(2)
		assertNoError(t, err)
		assertEqual(t, calc.Result(), 5)
		assertLogMessage(t, logger.messages, "Add(10): result = 10")
		assertLogMessage(t, logger.messages, "Divide(2): result = 5")
	})

	// ゼロ除算のテスト
	t.Run("DivisionByZero", func(t *testing.T) {
		setup()
		defer teardown()
		calc.Add(10)
		err := calc.Divide(0)
		assertError(t, err)
		assertLogMessage(t, logger.messages, "Add(10): result = 10")
		assertLogMessage(t, logger.messages, "Divide(0): error = division by zero")
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
		assertLogMessage(t, logger.messages, "Add(2): result = 2")
		assertLogMessage(t, logger.messages, "Multiply(3): result = 6")
		assertLogMessage(t, logger.messages, "Subtract(1): result = 5")
		assertLogMessage(t, logger.messages, "Divide(2): result = 2")
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

// ログメッセージが存在することを検証
func assertLogMessage(t *testing.T, messages []string, expectedMessage string) {
	t.Helper()
	found := false
	for _, message := range messages {
		if message == expectedMessage {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected log message '%s' not found", expectedMessage)
	}
}
