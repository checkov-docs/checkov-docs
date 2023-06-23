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

package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkov-docs/checkov-docs/internal/logger"
)

func TestGenerate_WithSkips(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	expectedOutputFile := "testdata/with-skips.md"
	inputFile := "testdata/with-skips.json"
	tmpOutputFile := createTempFile(t, nil)
	defer os.Remove(tmpOutputFile)
	logger := logger.NewMockLogger(&bytes.Buffer{})

	// Test generating output file
	err := Generate(inputFile, tmpOutputFile, false, logger)
	assert.Nil(err, "unexpected error returned by function", err)

	// Assert that the output file exists
	assert.FileExistsf(tmpOutputFile, "output file was not created")

	// Assert the content of the output file
	expected, err := os.ReadFile(expectedOutputFile)
	assert.Nil(err, "unexpected error reading expected output file", err)
	output, err := os.ReadFile(tmpOutputFile)
	assert.Nil(err, "unexpected error reading output file", err)
	assert.Equal(string(expected), string(output))
}

func TestGenerate_NoSkips(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	expectedOutputFile := "testdata/no-skips.md"
	inputFile := "testdata/no-skips.json"
	tmpOutputFile := createTempFile(t, nil)
	defer os.Remove(tmpOutputFile)
	logger := logger.NewMockLogger(&bytes.Buffer{})

	// Test generating output file
	err := Generate(inputFile, tmpOutputFile, false, logger)
	assert.Nil(err, "unexpected error returned by function", err)

	// Assert that the output file exists
	assert.FileExistsf(tmpOutputFile, "output file was not created")

	// Assert the content of the output file
	expected, err := os.ReadFile(expectedOutputFile)
	assert.Nil(err, "unexpected error reading expected output file", err)
	output, err := os.ReadFile(tmpOutputFile)
	assert.Nil(err, "unexpected error reading output file", err)
	assert.Equal(string(expected), string(output))
}

// Helper function to create a temporary file and write content to it
func createTempFile(t *testing.T, content []byte) string {
	tmpFile, err := os.CreateTemp(".", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err.Error())
	}
	defer tmpFile.Close()

	if content != nil {
		_, err = tmpFile.Write(content)
		if err != nil {
			t.Fatalf("Failed to write content to temporary file: %s", err.Error())
		}
	}

	return tmpFile.Name()
}
