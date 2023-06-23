/*
Copyright © 2023 The checkov-docs Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package logger

import (
	"github.com/hashicorp/go-hclog"
)

// Logger provides logging functionality
type Logger struct {
	logger hclog.Logger
}

// NewLogger returns a new instance of Logger
func NewLogger(name, level string) *Logger {
	newLogger := hclog.New(&hclog.LoggerOptions{
		Name:                     name,
		Level:                    getLogLevel(level),
		Output:                   nil,
		Mutex:                    nil,
		JSONFormat:               false,
		IncludeLocation:          false,
		AdditionalLocationOffset: 0,
		TimeFormat:               "",
		DisableTime:              false,
		Color:                    hclog.AutoColor,
		ColorHeaderOnly:          false,
		IndependentLevels:        false,
	})

	return &Logger{logger: newLogger}
}

// Debug logs a message at the DEBUG level
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

// Info logs a message at the INFO level
func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

// Warn logs a message at the WARN level
func (l *Logger) Warn(msg, warning string) {
	if warning != "" {
		l.logger.Warn(msg, "warning", warning)
	} else {
		l.logger.Warn(msg)
	}
}

// Error logs a message at the ERROR level
func (l *Logger) Error(msg, err string) {
	if err != "" {
		l.logger.Error(msg, "error", err)
	} else {
		l.logger.Error(msg)
	}
}

// SetLogLevel updates the log level
func (l *Logger) SetLogLevel(level string) {
	l.logger.SetLevel(getLogLevel(level))
}

// getLogLevel returns a hclog.Level using its string representation from `level`
func getLogLevel(level string) hclog.Level {
	switch {
	case level == "DEBUG":
		return hclog.Debug
	default:
		return hclog.Info
	}
}
