package logmsg

import (
	"log/slog"
)

type LogMsg struct {
	URL    string
	Method string
	Text   string
	Status int
}

func NewLogMsg(url, method string) *LogMsg {
	return &LogMsg{
		URL:    url,
		Method: method,
	}
}

func (msg *LogMsg) WithText(text string) *LogMsg {
	return &LogMsg{
		Text:   text,
		Status: msg.Status,
		URL:    msg.URL,
		Method: msg.Method,
	}
}

func (msg *LogMsg) WithStatus(status int) *LogMsg {
	return &LogMsg{
		Text:   msg.Text,
		Status: status,
		URL:    msg.URL,
		Method: msg.Method,
	}
}

func (msg *LogMsg) Info() {
	slog.Info(msg.Text, getArgs(msg)...)
}

func (msg *LogMsg) Error() {
	slog.Error(msg.Text, getArgs(msg)...)
}

func getArgs(msg *LogMsg) []any {
	return []any{
		"status", msg.Status,
		"url", msg.URL,
		"method", msg.Method,
	}
}
