package kklogger_slack

import (
	"testing"
	"time"

	"github.com/yetiz-org/goth-kklogger"
)

func TestKKLoggerRollbarHook(t *testing.T) {
	hook := &KKLoggerSlackHook{
		ServiceHookUrl: "",
		Level:          kklogger.DebugLevel,
		Environment:    "test_env",
		ServerRoot:     "test_server_root",
		CodeVersion:    "test_code_version",
	}

	kklogger.AsyncWrite = false
	kklogger.SetLoggerHooks([]kklogger.LoggerHook{hook})
	kklogger.SetLogLevel("DEBUG")
	kklogger.DebugJ("djsType", "jsData")
	time.Sleep(time.Second * 2)
}

func TestExtendedLoggerHookImplementation(t *testing.T) {
	hook := &KKLoggerSlackHook{
		ServiceHookUrl: "https://hooks.slack.com/services/test",
		Level:          kklogger.InfoLevel,
		Environment:    "test",
		ServerRoot:     "/test",
		CodeVersion:    "v1.0.0",
	}

	var extHook kklogger.ExtendedLoggerHook = hook
	if extHook == nil {
		t.Fatal("Hook does not implement ExtendedLoggerHook interface")
	}

	testFunc := "test.function"
	testFile := "/path/to/file.go"
	testLine := 42

	hook.InfoWithCaller(testFunc, testFile, testLine, "test message")
	hook.ErrorWithCaller(testFunc, testFile, testLine, "error message")
}

func TestBasicHookBackwardCompatibility(t *testing.T) {
	hook := &KKLoggerSlackHook{
		ServiceHookUrl: "https://hooks.slack.com/services/test",
		Level:          kklogger.InfoLevel,
		Environment:    "test",
		ServerRoot:     "/test",
		CodeVersion:    "v1.0.0",
	}

	var basicHook kklogger.LoggerHook = hook
	if basicHook == nil {
		t.Fatal("Hook does not implement LoggerHook interface")
	}

	basicHook.Info("test message")
	basicHook.Error("error message")
}
