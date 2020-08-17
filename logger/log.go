package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type texFormat struct {
}

func (f *texFormat) Format(entry *logrus.Entry) ([]byte, error) {
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
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("] ")

	// timestamp
	b.WriteString(entry.Time.Format(time.RFC3339))

	// request_id
	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("] ")

	// msg
	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("] ")

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
