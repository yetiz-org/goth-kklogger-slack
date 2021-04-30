package kklogger_slack

import (
	"testing"
	"time"

	"github.com/kklab-com/goth-kklogger"
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
