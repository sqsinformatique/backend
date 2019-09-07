package utils

import (
	"io"
	"os"

	logger "github.com/sirupsen/logrus"
)

// Level type
type Level uint32

func LevelByName(name string) Level {
	switch name {
	case "Panic":
		return PanicL
	case "Fatal":
		return FatalL
	case "Error":
		return ErrorL
	case "Warn":
		return WarnL
	case "Info":
		return InfoL
	case "Debug":
		return DebugL
	case "Trace":
		return TraceL
	default:
		return 0
	}
}

const (
	PanicL Level = iota
	FatalL
	ErrorL
	WarnL
	InfoL
	DebugL
	TraceL
)

var log = logger.NewEntry(logger.New())

// Init log subsystem
func InitLogger(LogName string) {

	if LogName != "" {
		file, err := os.OpenFile(LogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, file)
			logger.SetOutput(mw)
		} else {
			logger.Info("Failed to log to file, using default stderr")
		}
	} else {
		mw := io.MultiWriter(os.Stdout)
		logger.SetOutput(mw)
	}

	// Default InfoLevel
	logger.SetLevel(logger.InfoLevel)
}

// Set log level
func SetLogLevel(lvl Level) {
	if lvl == TraceL {
		logger.SetLevel(logger.Level(TraceL))
		logger.SetReportCaller(true)
	}
	logger.SetLevel(logger.Level(lvl))
	logger.SetReportCaller(false)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Trace(args ...interface{}) {
	log.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	log.Traceln(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
