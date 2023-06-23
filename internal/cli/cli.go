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
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/checkov-docs/checkov-docs/internal/config"
	"github.com/checkov-docs/checkov-docs/internal/filewriter"
	"github.com/checkov-docs/checkov-docs/internal/logger"
	"github.com/checkov-docs/checkov-docs/internal/markdown"
	"github.com/checkov-docs/checkov-docs/internal/models"
)

// Generate markdown table from checkov results in `inputFile`
// and write generated content to `outputFile` which defaults to a 'README.md' file in the current directory.
func Generate(inputFile, outputFile string, dryRun bool, logger *logger.Logger) error {

	// Read file with checkov results
	jsonData, err := os.ReadFile(filepath.Clean(inputFile))
	if err != nil {
		logger.Error("failed to read checkov json file", err.Error())
		return err
	}
	logger.Info("read checkov results json data")

	// Parse file with checkov results
	findings := &models.CheckovResults{}
	err = json.Unmarshal(jsonData, findings)
	if err != nil {
		logger.Error("failed to parse checkov json file", err.Error())
		return err
	}
	logger.Info("parsed checkov results json data")

	// Create markdown header and data rows
	headers := config.OutputFileHeader
	rows := make([][]string, len(findings.Results.SkippedChecks))
	for i, finding := range findings.Results.SkippedChecks {
		rows[i] = []string{
			finding.FilePath,
			finding.CheckID,
			finding.Resource,
			finding.CheckResult.SuppressComment,
		}
	}
	logger.Debug("created header and data rows", "headers", headers, "rows", rows)

	// Create markdown table
	table, err := markdown.WriteTable(headers, rows, logger)
	if err != nil {
		logger.Error("failed to generate markdown table", err.Error())
		return err
	}

	if dryRun {
		// write generated content to stdout
		_, err = os.Stdout.WriteString(table)
		if err != nil {
			return err
		}
	} else {
		// write generated content to output file
		w := &filewriter.FileWriter{
			Filepath:   outputFile,
			Template:   config.OutputTemplate,
			OpeningTag: config.TemplateBeginTag,
			ClosingTag: config.TemplateEndTag,
			Logger:     logger,
		}
		_, err = io.WriteString(w, table)
		if err != nil {
			return err
		}
		logger.Info("output file updated successfully")
	}

	return nil
}
