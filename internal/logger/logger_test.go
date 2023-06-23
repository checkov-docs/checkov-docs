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
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestLogger_Debug(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := "Debug message"
	output := &bytes.Buffer{}
	logger := NewMockLogger(output)

	// Test logging a debug message
	logger.Debug("Debug message")
	expected := getExpected("DEBUG", input)
	assert.Equal(expected, output.String())
}

func TestLogger_Info(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := "Info message"
	output := &bytes.Buffer{}
	logger := NewMockLogger(output)

	// Test logging an info message
	logger.Info(input)
	expected := getExpected("INFO", input)
	assert.Equal(expected, output.String())
}

func TestLogger_Warn(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := "Warning message"
	output := &bytes.Buffer{}
	logger := NewMockLogger(output)

	// Test logging a warning message
	logger.Warn(input, "foo")
	expected := "WARN: Warning message\n"
	assert.Equal(expected, output.String())
}

func TestLogger_Error(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := "Error message"
	output := &bytes.Buffer{}
	logger := NewMockLogger(output)

	// Test logging a warning message
	logger.Error(input, "foo")
	expected := "ERROR: Error message\n"
	assert.Equal(expected, output.String())
}

func TestLogger_SetLogLevel(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	mockLogger := &mockLogger{} // Mock logger to track log level changes
	logger := Logger{
		logger: mockLogger,
	}

	// Test setting log level
	logger.SetLogLevel("DEBUG")
	expected := hclog.Debug
	assert.Equal(expected, mockLogger.level)

	// Test updating log level
	logger.SetLogLevel("INFO")
	expected = hclog.Info
	assert.Equal(expected, mockLogger.level)

	// Test setting unknown log level returns default level
	logger.SetLogLevel("FOO")
	expected = hclog.Info
	assert.Equal(expected, mockLogger.level)
}

func getExpected(level, input string) string {
	return fmt.Sprintf("%s: %s\n", level, input)
}
