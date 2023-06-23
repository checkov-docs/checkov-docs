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
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/checkov-docs/checkov-docs/internal/logger"
)

// FileWriter implements the io.Writer interface to write content to a file.
//
// Step 1. Validate filepath
// Step 2. Render template
// Step 3. Write generated content to output file, i.e. append or inject
type FileWriter struct {
	Filepath   string
	Template   string
	OpeningTag string
	ClosingTag string
	Logger     *logger.Logger
}

// Write content to file
func (fw *FileWriter) Write(p []byte) (int, error) {
	fw.Logger.Info("write content to output file")

	buf, err := fw.render(p)
	if err != nil {
		return 0, err
	}

	existingContent, err := os.ReadFile(filepath.Clean(fw.Filepath))
	if err != nil {
		// if file doesn't exist, create it and write generated output
		return 0, os.WriteFile(fw.Filepath, buf.Bytes(), 0644)
	}

	if len(existingContent) == 0 {
		// if file exists but it's empty, write generated output
		return 0, os.WriteFile(fw.Filepath, buf.Bytes(), 0644)
	}

	return fw.write(string(existingContent), buf.String())
}

// write appends or injects generated output to a file
func (fw *FileWriter) write(content, generated string) (int, error) {
	// Find the position of the opening and closing tags
	openingIndex := strings.Index(content, fw.OpeningTag)
	closingIndex := strings.Index(content, fw.ClosingTag)

	// if no tags found, simply append generated output to existing content
	if openingIndex == -1 && closingIndex == -1 {
		return 0, os.WriteFile(fw.Filepath, []byte(content+"\n"+generated), 0644)
	}

	if openingIndex == -1 {
		return 0, errors.New("opening comment tag is not found")
	}

	if closingIndex == -1 {
		return 0, errors.New("closing comment tag is not found")
	}

	// if both tags found, merge generated output with existing content
	// note that both tags in the existing content are omitted
	// because the generated output already includes them
	mergedContent := fmt.Sprintf("%s%s%s", content[:openingIndex], generated, content[closingIndex+len(fw.ClosingTag):])

	return 0, os.WriteFile(fw.Filepath, []byte(mergedContent), 0644)
}

// render parses and applies the template to generate output content
func (fw *FileWriter) render(p []byte) (bytes.Buffer, error) {
	fw.Logger.Info("render output template")

	type templateData struct {
		Content string
	}

	tpl, err := template.New("generated").Parse(fw.Template)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "generated", templateData{Content: string(p)})
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf, err
}
