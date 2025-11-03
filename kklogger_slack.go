package kklogger_slack

import (
	"fmt"
	"net/http"
	"net/url"

	kklogger "github.com/yetiz-org/goth-kklogger"
	"github.com/yetiz-org/goth-util/value"
)

type KKLoggerSlackHook struct {
	ServiceHookUrl string
	Environment    string
	CodeVersion    string
	ServerRoot     string
	Level          kklogger.Level
}

func (h *KKLoggerSlackHook) LogString(args ...interface{}) string {
	if args == nil || len(args) == 0 {
		return ""
	}

	if len(args) == 1 {
		if slice, ok := args[0].([]interface{}); ok {
			args = slice
		}
	}

	argl := len(args)

	if argl == 1 {
		switch tp := args[0].(type) {
		case string:
			return tp
		}
	} else if argl > 1 {
		switch tp := args[0].(type) {
		case string:
			pargs := args[1:]
			return fmt.Sprintf(tp, pargs...)
		}
	}

	return fmt.Sprint(args...)
}

func (h *KKLoggerSlackHook) Trace(args ...interface{}) {
	if h.Level < kklogger.TraceLevel {
		return
	}

	h.Send(kklogger.TraceLevel, "", "", 0, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Debug(args ...interface{}) {
	if h.Level < kklogger.DebugLevel {
		return
	}

	h.Send(kklogger.DebugLevel, "", "", 0, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Info(args ...interface{}) {
	if h.Level < kklogger.InfoLevel {
		return
	}

	h.Send(kklogger.InfoLevel, "", "", 0, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Warn(args ...interface{}) {
	if h.Level < kklogger.WarnLevel {
		return
	}

	h.Send(kklogger.WarnLevel, "", "", 0, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Error(args ...interface{}) {
	if h.Level < kklogger.ErrorLevel {
		return
	}

	h.Send(kklogger.ErrorLevel, "", "", 0, h.LogString(args...))
}

func (h *KKLoggerSlackHook) TraceWithCaller(funcName, file string, line int, args ...interface{}) {
	if h.Level < kklogger.TraceLevel {
		return
	}

	h.Send(kklogger.TraceLevel, funcName, file, line, h.LogString(args...))
}

func (h *KKLoggerSlackHook) DebugWithCaller(funcName, file string, line int, args ...interface{}) {
	if h.Level < kklogger.DebugLevel {
		return
	}

	h.Send(kklogger.DebugLevel, funcName, file, line, h.LogString(args...))
}

func (h *KKLoggerSlackHook) InfoWithCaller(funcName, file string, line int, args ...interface{}) {
	if h.Level < kklogger.InfoLevel {
		return
	}

	h.Send(kklogger.InfoLevel, funcName, file, line, h.LogString(args...))
}

func (h *KKLoggerSlackHook) WarnWithCaller(funcName, file string, line int, args ...interface{}) {
	if h.Level < kklogger.WarnLevel {
		return
	}

	h.Send(kklogger.WarnLevel, funcName, file, line, h.LogString(args...))
}

func (h *KKLoggerSlackHook) ErrorWithCaller(funcName, file string, line int, args ...interface{}) {
	if h.Level < kklogger.ErrorLevel {
		return
	}

	h.Send(kklogger.ErrorLevel, funcName, file, line, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Send(level kklogger.Level, funcName, file string, line int, msg string) {
	fields := []KKLoggerSlackBlockField{
		{
			Type: "mrkdwn",
			Text: "*Level*",
		},
		{
			Type: "plain_text",
			Text: h.Level.String(),
		},
		{
			Type: "mrkdwn",
			Text: "*Environment*",
		},
		{
			Type: "plain_text",
			Text: h.Environment,
		},
		{
			Type: "mrkdwn",
			Text: "*ServerRoot*",
		},
		{
			Type: "plain_text",
			Text: h.ServerRoot,
		},
		{
			Type: "mrkdwn",
			Text: "*CodeVersion*",
		},
		{
			Type: "plain_text",
			Text: h.CodeVersion,
		},
	}

	if funcName != "" {
		fields = append(fields, KKLoggerSlackBlockField{
			Type: "mrkdwn",
			Text: "*Function*",
		}, KKLoggerSlackBlockField{
			Type: "plain_text",
			Text: funcName,
		})
	}

	if file != "" {
		fileInfo := file
		if line > 0 {
			fileInfo = fmt.Sprintf("%s:%d", file, line)
		}
		fields = append(fields, KKLoggerSlackBlockField{
			Type: "mrkdwn",
			Text: "*File*",
		}, KKLoggerSlackBlockField{
			Type: "plain_text",
			Text: fileInfo,
		})
	}

	d := url.Values{}
	d.Set("payload", fmt.Sprintf(value.JsonMarshal(map[string]interface{}{
		"blocks": []KKLoggerSlackBlock{
			{
				Type: "section",
				Text: KKLoggerSlackBlockField{
					Type: "plain_text",
					Text: msg,
				},
				Fields: fields,
			},
		},
	})))

	http.DefaultClient.PostForm(h.ServiceHookUrl, d)
}

type KKLoggerSlackBlock struct {
	Type   string                    `json:"type"`
	Text   KKLoggerSlackBlockField   `json:"text"`
	Fields []KKLoggerSlackBlockField `json:"fields"`
}

type KKLoggerSlackBlockField struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
