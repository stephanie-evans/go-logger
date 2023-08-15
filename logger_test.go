package go_logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	l := New("TEST: ")
	l.Info("This is an info message")
	l.Error("This is an error message")
}
