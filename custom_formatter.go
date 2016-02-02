package logrus

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type CustomFormatter struct {
	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string
}

func (f *CustomFormatter) Format(entry *Entry) ([]byte, error) {
	var keys []string = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	b := &bytes.Buffer{}

	prefixFieldClashes(entry.Data)

	levelText := strings.ToUpper(entry.Level.String())[0:4]
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampNano
	}
	fmt.Fprintf(b, "%s[%s] %-44s ", levelText, entry.Time.Format(timestampFormat), entry.Message)

	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " %s=%+v", k, v)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
