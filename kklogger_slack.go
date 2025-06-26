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
	if args == nil {
		return ""
	}

	args = args[0].([]interface{})
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

	h.Send(kklogger.TraceLevel, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Debug(args ...interface{}) {
	if h.Level < kklogger.DebugLevel {
		return
	}

	h.Send(kklogger.DebugLevel, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Info(args ...interface{}) {
	if h.Level < kklogger.InfoLevel {
		return
	}

	h.Send(kklogger.InfoLevel, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Warn(args ...interface{}) {
	if h.Level < kklogger.WarnLevel {
		return
	}

	h.Send(kklogger.WarnLevel, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Error(args ...interface{}) {
	if h.Level < kklogger.ErrorLevel {
		return
	}

	h.Send(kklogger.ErrorLevel, h.LogString(args...))
}

func (h *KKLoggerSlackHook) Send(level kklogger.Level, msg string) {
	d := url.Values{}
	d.Set("payload", fmt.Sprintf(value.JsonMarshal(map[string]interface{}{
		"blocks": []KKLoggerSlackBlock{
			{
				Type: "section",
				Text: KKLoggerSlackBlockField{
					Type: "plain_text",
					Text: msg,
				},
				Fields: []KKLoggerSlackBlockField{
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
				},
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
