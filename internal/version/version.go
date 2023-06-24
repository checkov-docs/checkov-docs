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

package version

import (
	_ "embed"

	goversion "github.com/caarlos0/go-version"
)

var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

//go:embed art.txt
var asciiArt string

const projectURL = "https://github.com/checkov-docs/checkov-docs"

// GetVersion returns the latest version including runtime GOOS and GOARCH and more.
func GetVersion() string {
	v := goversion.GetVersionInfo(
		goversion.WithAppDetails("checkov-docs", "Generate docs for checkov results", projectURL),
		goversion.WithASCIIName(asciiArt),
		goversion.WithBuiltBy(builtBy),
		func(i *goversion.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			if version != "" {
				i.GitVersion = version
			}
			if date != "" {
				i.BuildDate = date
			}
			if builtBy != "" {
				i.BuiltBy = builtBy
			}
		},
	)

	return v.String()
}
