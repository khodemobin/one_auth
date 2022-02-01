package zap

import (
	l "log"
	"os"

	"github.com/khodemobin/pilo/auth/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	zap *zap.SugaredLogger
}

func New() logger.Logger {
	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l.Fatalln(err)
	}

	ws := zapcore.AddSync(f)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(enc, ws, zapcore.ErrorLevel)

	z := zap.New(core)
	sugarLogger := z.Sugar()

	return &log{sugarLogger}
}

func (l *log) Error(err error) {
	l.zap.Error(err.Error())
}

func (l *log) Fatal(err error) {
	l.zap.Fatal(err.Error())
}

func (l *log) Warn(msg string) {
	l.zap.Warn(msg)
}

func (l *log) Info(msg string) {
	l.zap.Info(msg)
}
