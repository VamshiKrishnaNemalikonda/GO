package functions

import (
        log "github.com/sirupsen/logrus"
        lumberjack "gopkg.in/natefinch/lumberjack.v2"
        "os"
        "config"
)

func LumberJackLogger(filePath string, maxSize int, maxBackups int, maxAge int) *lumberjack.Logger {

        return &lumberjack.Logger{

                Filename: filePath,

                MaxSize: maxSize, // megabytes

                MaxBackups: maxBackups,

                MaxAge: maxAge, //days

        }

}

func init() {

        // Log as JSON instead of the default ASCII formatter.
        log.SetFormatter(&log.JSONFormatter{})

        // Output to stdout instead of the default stderr
        // Can be any io.Writer, see below for File example
        log.SetOutput(os.Stdout)

        // Only log the warning severity or above.
        //log.SetLevel(WarnLevel)

        out := LumberJackLogger(config.ErrorLogFilePath+config.ErrorLogFileExtension, config.ErrorLogMaxSize, config.ErrorLogMaxBackups, config.ErrorLogMaxAge)

        // defer f.Close()

        // logrus.SetOutput(os.Stderr)

        log.SetOutput(out)

        log.SetLevel(log.DebugLevel)
}

// Debug logs a message with debug log level.
func Debug(msg string) {
        log.Debug(msg)
}

// Debugf logs a formatted message with debug log level.
func Debugf(msg string, args ...interface{}) {
        log.Debugf(msg, args...)
}

// Info logs a message with info log level.
func Info(msg string) {
        log.Info(msg)
}

// Infof logs a formatted message with info log level.
func Infof(msg string, args ...interface{}) {
        log.Infof(msg, args...)
}

func WithFields(msg string, fkey string, fvalue string) {

        log.WithFields(log.Fields{fkey: fvalue}).Error(msg)
}
