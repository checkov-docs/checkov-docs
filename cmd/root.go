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

package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/checkov-docs/checkov-docs/internal/cli"
	"github.com/checkov-docs/checkov-docs/internal/logger"
	"github.com/checkov-docs/checkov-docs/internal/version"
)

var (
	cmdLogger  *logger.Logger
	cfgFile    string
	inputFile  string
	outputFile string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "checkov-docs",
	Short:         "Generate docs for checkov results",
	Long:          "Generate docs for checkov results",
	Annotations:   map[string]string{"command": "root"},
	Version:       version.GetFullVersion(),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cmdLogger.SetLogLevel(getLogLevel())
		err := initConfig()
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		in := viper.GetString("input-file")
		out := viper.GetString("output-file")
		dryrun := viper.GetBool("dry-run")
		cmdLogger.Info("run", "cmd", cmd.Aliases, "args", args, "input-file", in, "output-file", out, "dry-run", dryrun)
		if in != "" {
			err := cli.Generate(in, out, dryrun, cmdLogger)
			if err != nil {
				return err
			}
		} else {
			return errors.New("input file is required")
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	cmdLogger = logger.NewLogger("checkov-docs", "INFO")
	if err := rootCmd.Execute(); err != nil {
		cmdLogger.Error("command failed", err.Error())
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".checkov-docs.yaml", "config file")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input-file", "i", "", "input file, valid formats: json")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "o", "README.md", "output file")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show debug output")
	rootCmd.PersistentFlags().Bool("dry-run", false, "only print generated output")
	cobra.CheckErr(viper.BindPFlag("input-file", rootCmd.PersistentFlags().Lookup("input-file")))
	cobra.CheckErr(viper.BindPFlag("output-file", rootCmd.PersistentFlags().Lookup("output-file")))
	cobra.CheckErr(viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")))
	cobra.CheckErr(viper.BindPFlag("dry-run", rootCmd.PersistentFlags().Lookup("dry-run")))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			cmdLogger.Error("failed to find $HOME directory", err.Error())
			return err
		}

		// Search config in the following directories with name ".checkov-docs" (without extension)
		viper.AddConfigPath(home) // home directory
		viper.AddConfigPath(".")  // current directory
		viper.SetConfigType("yaml")
		viper.SetConfigName(".checkov-docs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()

	var pathError *os.PathError
	var notFoundError viper.ConfigFileNotFoundError
	switch {
	case err != nil && errors.As(err, &pathError):
		cmdLogger.Warn("no config file found", err.Error())
	case err != nil && !errors.As(err, &notFoundError):
		// pathError and notFoundError are produced when no config file is found
		// here we check and return an error produced when reading the config file
		cmdLogger.Error("failed to read config", err.Error())
		return err
	default:
		cmdLogger.Info("using config file", "path", viper.ConfigFileUsed())
	}

	return nil
}

// getViperLogLevel returns "DEBUG" if `debug` flag is used, else it returns "INFO".
func getLogLevel() string {
	switch viper.GetBool("verbose") {
	case true:
		return "DEBUG"
	default:
		return "INFO"
	}
}
