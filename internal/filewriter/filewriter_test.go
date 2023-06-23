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

package filewriter

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkov-docs/checkov-docs/internal/config"
	"github.com/checkov-docs/checkov-docs/internal/logger"
)

var testLogger = []logger.Logger{*logger.NewLogger("test-checkov-docs", "INFO")}

func TestFileWriter_Write(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	filePath := "output.md"

	fw := &FileWriter{
		Filepath:   filePath,
		Template:   config.OutputTemplate,
		OpeningTag: config.TemplateBeginTag,
		ClosingTag: config.TemplateEndTag,
		Logger:     &testLogger[0],
	}

	// Test writing to a non-existing file
	input := "foo"
	expected := getExpected(input)
	_, err := io.WriteString(fw, input)
	assert.Nil(err, "unexpected error while writing to non-existing file", err)

	// Verify file content
	actual, err := os.ReadFile(filePath)
	assert.Nil(err, "unexpected error while reading file")
	assert.Equal(expected, string(actual))

	// Clear existing file contents
	err = os.WriteFile(filePath, []byte{}, 0644)
	assert.Nil(err, "unexpected error while clearing content in existing file", err)

	// Test writing to an existing file with empty content
	input = "foo bar"
	expected = getExpected(input)
	_, err = io.WriteString(fw, input)
	assert.Nil(err, "unexpected error while writing to file with empty content", err)

	// Verify file content
	actual, err = os.ReadFile(filePath)
	assert.Nil(err, "unexpected error while reading file", err)
	assert.Equal(expected, string(actual))

	// Test writing to an existing file with opening and closing tags
	input = "foo bar baz"
	expected = getExpected(input)
	_, err = io.WriteString(fw, input)
	assert.Nil(err, "unexpected error while writing to file with opening and closing tags")

	// Verify file content
	actual, err = os.ReadFile(filePath)
	assert.Nil(err, "unexpected error while reading file", err)
	assert.Equal(expected, string(actual))

	// Clean up test files
	err = os.Remove(filePath)
	assert.Nil(err, "unexpected error while cleaning up test files", err)
}

func getExpected(input string) string {
	return config.TemplateBeginTag + "\n\n" + input + "\n\n" + config.TemplateEndTag
}

func TestFileWriter_Write_Error(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	fw := &FileWriter{
		Filepath:   "nonexistentpath/output.txt",
		Template:   config.OutputTemplate,
		OpeningTag: config.TemplateBeginTag,
		ClosingTag: config.TemplateEndTag,
		Logger:     &testLogger[0],
	}

	// Test error when writing to a non-existent directory
	content := "non-existing-path"
	_, err := io.WriteString(fw, content)
	assert.NotNil(err, "expected an error when writing to a non-existent directory, but got no error")
}

func TestFileWriter_render(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := []byte("Hello, world!")
	fw := &FileWriter{
		Template: config.OutputTemplate,
		Logger:   &testLogger[0],
	}

	// Test rendering the template
	buf, err := fw.render(input)
	assert.Nil(err, "unexpected error while applying template", err)

	expected := getExpected(string(input))
	assert.Equal(expected, buf.String())
}

func TestFileWriter_render_Error(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	input := []byte("Hello, world!")
	fw := FileWriter{
		Template: "{{.UndefinedVar}}",
		Logger:   &testLogger[0],
	}

	// Test applying the template with an undefined variable
	_, err := fw.render(input)
	assert.NotNil(err, "expected an error while applying template with undefined variable, but got no error")
}
