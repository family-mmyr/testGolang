package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const LoggerVer = "v1"

type TextFormatter struct {
	RequestID string
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// level
	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("] ")

	// version
	b.WriteByte('[')
	b.WriteString(LoggerVer)
	b.WriteString("] ")

	// timestamp
	b.WriteString(entry.Time.Format(time.RFC3339))

	// request_id
	b.WriteString(" [")
	b.WriteString(f.RequestID)
	b.WriteString("] ")

	// msg
	if entry.Message != "" {
		b.WriteString(entry.Message)
	}

	// data
	for key, value := range entry.Data {
		b.WriteString(" ,")
		b.WriteString(key)
		b.WriteByte('=')
		fmt.Fprint(b, value)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
