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
	"fmt"
	"strings"

	"github.com/checkov-docs/checkov-docs/internal/logger"
)

// WriteTable returns a markdown table with arguments headers and rows
// revive:disable:unhandled-error ignore regex pattern in golangci-lint does not work
func WriteTable(headers []string, rows [][]string, logger *logger.Logger) (string, error) {
	logger.Info("create markdown table")
	var sb strings.Builder

	// Calculate the maximum length for each column
	columnLengths := getColumnLengths(headers, rows)
	logger.Debug("calculated maximum length for each column")

	// Create the header row
	sb.WriteString("|")
	for i, header := range headers {
		_, err := sb.WriteString(fmt.Sprintf(" %-*s |", columnLengths[i], header))
		if err != nil {
			logger.Error("failed to write header row", err.Error())
			return "", err
		}
	}
	sb.WriteString("\n")
	logger.Debug("added header row")

	// Create the separator row
	sb.WriteString("|")
	for _, length := range columnLengths {
		_, err := sb.WriteString(strings.Repeat("-", length+2))
		if err != nil {
			logger.Error("failed to write separator row", err.Error())
			return "", err
		}
		sb.WriteString("|")
	}
	if len(rows) > 0 {
		sb.WriteString("\n")
	}
	logger.Debug("added separator row")

	// Create the data rows
	for i, row := range rows {
		sb.WriteString("|")
		for i, value := range row {
			_, err := sb.WriteString(fmt.Sprintf(" %-*s |", columnLengths[i], value))
			if err != nil {
				logger.Error("failed to write data row", err.Error())
				return "", err
			}
		}
		// only add new line if it's not the last row, this allows the template string
		// to format space between opening tag, table and closing tag
		if i < len(rows)-1 {
			sb.WriteString("\n")
		}
	}
	logger.Debug("added data rows")

	return sb.String(), nil
}

// getColumnLengths returns the length of each column by comparing the maximum length
// in each element of `headers` and `rows`.
func getColumnLengths(headers []string, rows [][]string) []int {
	columnLengths := make([]int, len(headers))
	for i, header := range headers {
		length := len(header)
		if length > columnLengths[i] {
			columnLengths[i] = length
		}
	}
	for _, row := range rows {
		for i, value := range row {
			length := len(value)
			if length > columnLengths[i] {
				columnLengths[i] = length
			}
		}
	}

	return columnLengths
}
