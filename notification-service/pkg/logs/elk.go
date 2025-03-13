package elk

import (
	"net"
	"sync"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var Log *Logger

type Logger struct {
	mu     *sync.Mutex
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) Initialization() error {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		return err
	}

	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type": "notification-service",
	}))

	l.logger.Hooks.Add(hook)

	return nil
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	l.mu.Lock()
	l.logger.Level = logrus.ErrorLevel
	l.logger.WithFields(fields).Error(msg)
	l.mu.Unlock()
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.mu.Lock()
	l.logger.Level = logrus.InfoLevel
	l.logger.WithFields(fields).Info(msg)
	l.mu.Unlock()
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	l.mu.Lock()
	l.logger.Level = logrus.DebugLevel
	l.logger.WithFields(fields).Debug(msg)
	l.mu.Unlock()
}

func init() {
	Log = NewLogger()
	if err := Log.Initialization(); err != nil {
		panic(err)
	}
}
