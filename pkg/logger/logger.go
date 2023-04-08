package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"

	"github.com/krasun/logrus2telegram"
)

func NewLoggerTG(logToken string, chatid int64) *log.Logger {
	l := log.New()
	hook, err := logrus2telegram.NewHook(
		logToken,
		[]int64{chatid},

		logrus2telegram.Levels(log.AllLevels),
		// default: []log.Level{log.ErrorLevel, log.FatalLevel, log.PanicLevel, log.WarnLevel, log.InfoLevel}
		logrus2telegram.NotifyOn([]log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel, log.InfoLevel}),
		// default: 3 * time.Second
		logrus2telegram.RequestTimeout(10*time.Second),
		// default: entry.String() time="2021-12-22T14:48:56+02:00" level=debug msg="example"
		logrus2telegram.Format(func(e *log.Entry) (string, error) {
			return fmt.Sprintf("%s %s", strings.ToUpper(e.Level.String()), e.Message), nil
		}),
	)
	if err != nil {
		panic(err)
	}
	l.AddHook(hook)
	return l
}
