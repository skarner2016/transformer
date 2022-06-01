package log

import (
	"fmt"
	"go.uber.org/zap"
	"transformer/library/config"
	"testing"
)

func TestNewLogger(t *testing.T) {
	config.InitConfig()
	logger := NewLogger()

	fmt.Println(fmt.Sprintf("%p", logger))

	logger.Info("abc", zap.String("ahaha", "sjdkfmks"))
}

func TestLog_GetFile(t *testing.T) {
	fmt.Println(getFile())
}

func TestLog_GetLogLevel(t *testing.T) {
	level := "debug"
	fmt.Println(getLogLevel(level))
}
