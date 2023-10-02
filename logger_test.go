package go_logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	go New("TEST: ")
}
