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
	"bytes"
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
)

// MockLogger provides a mocked implementation of the `hclog.Logger` interface used for testing.
type mockLogger struct {
	level  hclog.Level
	output *bytes.Buffer
}

// NewMockLogger returns an instance of mockLogger
func NewMockLogger(output *bytes.Buffer) *Logger {
	return &Logger{logger: &mockLogger{output: output}}
}

// GetLogLevel returns the current log level
func (m *mockLogger) GetLogLevel() hclog.Level {
	return m.level
}

// revive:disable:unused-parameter safe to ignore in mock logger
// revive:disable:unhandled-error safe to ignore in mock logger

// Trace logs a mocked trace message
func (m *mockLogger) Trace(msg string, args ...interface{}) {
	m.output.WriteString("TRACE: " + msg + "\n")
}

// Debug logs a mocked debug message
func (m *mockLogger) Debug(msg string, args ...interface{}) {
	m.output.WriteString("DEBUG: " + msg + "\n")
}

// Info logs a mocked info message
func (m *mockLogger) Info(msg string, args ...interface{}) {
	m.output.WriteString("INFO: " + msg + "\n")
}

// Warn logs a mocked warning message
func (m *mockLogger) Warn(msg string, args ...interface{}) {
	m.output.WriteString("WARN: " + msg + "\n")
}

// Trace logs a mocked error message
func (m *mockLogger) Error(msg string, args ...interface{}) {
	m.output.WriteString("ERROR: " + msg + "\n")
}

// revive:disable:exported safe to ignore in mock logger

func (m *mockLogger) IsTrace() bool {
	return true
}

func (m *mockLogger) IsDebug() bool {
	return true
}

func (m *mockLogger) IsInfo() bool {
	return true
}

func (m *mockLogger) IsWarn() bool {
	return true
}

func (m *mockLogger) IsError() bool {
	return true
}

func (m *mockLogger) Log(level hclog.Level, msg string, args ...interface{}) {}

func (m *mockLogger) ImpliedArgs() []interface{} {
	return nil
}

func (m *mockLogger) With(args ...interface{}) hclog.Logger {
	return m
}

func (m *mockLogger) Name() string {
	return ""
}

func (m *mockLogger) Named(name string) hclog.Logger {
	return m
}

func (m *mockLogger) ResetNamed(name string) hclog.Logger {
	return m
}

func (m *mockLogger) SetLevel(level hclog.Level) {
	m.level = level
}

func (m *mockLogger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return nil
}

func (m *mockLogger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return nil
}

// revive:enable:unused-parameter safe to ignore in mock logger
// revive:enable:unhandled-error safe to ignore in mock logger
// revive:enable:exported safe to ignore in mock logger
