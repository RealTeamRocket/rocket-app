package logger_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"rocket-backend/pkg/logger"
)

// mockLogger implements Logger and writes to a buffer for assertions.
type mockLogger struct {
	buf *bytes.Buffer
}

func newMockLogger() *mockLogger {
	return &mockLogger{buf: &bytes.Buffer{}}
}

func (m *mockLogger) Debug(v ...interface{}) { m.buf.WriteString("[DEBUG] " + sprintArgs(v...) + "\n") }
func (m *mockLogger) Info(v ...interface{})  { m.buf.WriteString("[INFO] " + sprintArgs(v...) + "\n") }
func (m *mockLogger) Warn(v ...interface{})  { m.buf.WriteString("[WARN] " + sprintArgs(v...) + "\n") }
func (m *mockLogger) Error(v ...interface{}) { m.buf.WriteString("[ERROR] " + sprintArgs(v...) + "\n") }
func (m *mockLogger) Fatal(v ...interface{}) { m.buf.WriteString("[FATAL] " + sprintArgs(v...) + "\n") }

func sprintArgs(v ...interface{}) string {
	var b strings.Builder
	for i, arg := range v {
		if i > 0 {
			b.WriteRune(' ')
		}
		b.WriteString(strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(fmt.Sprint(arg), "\n"), "\n")))
	}
	return b.String()
}

func TestLoggerSetAndGet(t *testing.T) {
	mock := newMockLogger()
	logger.Set(mock)
	if logger.Get() != mock {
		t.Error("Get() should return the logger set by Set()")
	}
}

func TestLoggerDebugInfoWarnError(t *testing.T) {
	mock := newMockLogger()
	logger.Set(mock)

	logger.Debug("foo", 123)
	logger.Info("bar", "baz")
	logger.Warn("warn", "msg")
	logger.Error("err", "msg")

	out := mock.buf.String()
	if !strings.Contains(out, "[DEBUG] foo 123") {
		t.Error("Debug output missing or incorrect:", out)
	}
	if !strings.Contains(out, "[INFO] bar baz") {
		t.Error("Info output missing or incorrect:", out)
	}
	if !strings.Contains(out, "[WARN] warn msg") {
		t.Error("Warn output missing or incorrect:", out)
	}
	if !strings.Contains(out, "[ERROR] err msg") {
		t.Error("Error output missing or incorrect:", out)
	}
}

func TestLoggerFatal(t *testing.T) {
	mock := newMockLogger()
	logger.Set(mock)
	logger.Fatal("fatal", "error")
	out := mock.buf.String()
	if !strings.Contains(out, "[FATAL] fatal error") {
		t.Error("Fatal output missing or incorrect:", out)
	}
}
