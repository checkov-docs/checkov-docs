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

package markdown

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkov-docs/checkov-docs/internal/logger"
)

func TestWriteTable(t *testing.T) {
	assert := assert.New(t)
	logger := logger.NewMockLogger(&bytes.Buffer{})

	// Prepare test data
	headers := []string{"Name", "Age", "Location"}
	rows := [][]string{
		{"John Doe", "30", "New York"},
		{"Jane Smith", "25", "London"},
	}
	expected := `
| Name       | Age | Location |
|------------|-----|----------|
| John Doe   | 30  | New York |
| Jane Smith | 25  | London   |
`

	// Test writing table
	output, err := WriteTable(headers, rows, logger)
	assert.Nil(err, "unexpected error writing table", err)
	assert.Equal(strings.TrimSpace(expected), strings.TrimSpace(output))
}

func TestGetColumnLengths(t *testing.T) {
	assert := assert.New(t)

	// Prepare test data
	headers := []string{"Name", "Age", "Location"}
	rows := [][]string{
		{"John Doe", "30", "New York"},
		{"Jane Smith", "25", "London"},
	}

	// Test computing column lengths
	columnLengths := getColumnLengths(headers, rows)

	expectedColumnLengths := []int{10, 3, 8}
	assert.ElementsMatchf(expectedColumnLengths, columnLengths, fmt.Sprintf("Unexpected number of column lengths. Got %d, expected %d.", len(columnLengths), len(expectedColumnLengths)))

	for i, length := range columnLengths {
		assert.Equalf(length, expectedColumnLengths[i], fmt.Sprintf("unexpected column length at index %d. Got %d, expected %d.", i, length, expectedColumnLengths[i]))
	}
}
