package logger

import (
	"testing"
)

func TestInitLog(t *testing.T) {
	config := NewDefaultLogConfig()
	InitLog(config)
	Info("log info")

}
