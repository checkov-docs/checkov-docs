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

package config

import (
	"fmt"
	"strings"
)

// Values used to generate template string
const (
	TemplateBeginTag      = "<!-- BEGIN_CHECKOV_DOCS -->"
	templateDataStructure = "{{ .Content }}"
	TemplateEndTag        = "<!-- END_CHECKOV_DOCS -->"
)

// OutputTemplate stores the template used to generate content
var OutputTemplate = fmt.Sprintf("%s\n\n%s\n\n%s", TemplateBeginTag, templateDataStructure, TemplateEndTag)

// OutputFileHeader stores the fields used to generate header in markdown table
var OutputFileHeader = []string{"File", "Check ID", "Resource ID", "Reason"}

// GetMarkdownHeader returns the markdown-formatted header
// revive:disable:unhandled-error ignore error in `WriteString`
func GetMarkdownHeader() string {
	var sb strings.Builder

	sb.WriteString("|")
	for _, val := range OutputFileHeader {
		sb.WriteString(fmt.Sprintf(" %s |", val))
	}

	return sb.String()
}

// revive:enable:unhandled-error ignore error in `sb.WriteString`
