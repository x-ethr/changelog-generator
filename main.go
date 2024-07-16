package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"changelog-generator/internal/commands"
	"changelog-generator/internal/constants"
	"changelog-generator/internal/log"
	"changelog-generator/internal/types/level"
)

// logging is a variable that represents the current log level configuration.
var logging level.Type = level.Error

// --> https://goreleaser.com/cookbooks/using-main.version/

var (
	version = "development" // see go linking for compile-time variable overwrites
	commit  = "local"       // see go linking for compile-time variable overwrites
	date    = "unavailable" // see go linking for compile-time variable overwrites
)

func main() {
	var root = &cobra.Command{
		Use:                    constants.Name(),
		Short:                  fmt.Sprintf("%s - A Changelog Utilities CLI", constants.Name()),
		Long:                   fmt.Sprintf("%s is a tool that generates changelog(s) and reports git-related information.", constants.Name()),
		Example:                "",
		ValidArgs:              nil,
		ValidArgsFunction:      nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Annotations:            nil,
		Version:                version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var l level.Type
			if e := l.Set(logging.String()); e != nil {
				return e
			}

			log.Default(l.String())

			ctx = context.WithValue(ctx, "version", version)
			ctx = context.WithValue(ctx, "commit", commit)
			ctx = context.WithValue(ctx, "date", date)
			ctx = context.WithValue(ctx, "log-level", l)

			slog.DebugContext(ctx, "Metadata", slog.Group("build", slog.String("version", version), slog.String("commit", commit), slog.String("date", date), slog.String("log-level", l.String())))

			cmd.SetContext(ctx)

			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// @todo Logic to check if a newer version is available
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				if e := cmd.Help(); e != nil {
					panic(e)
				}
			}
		},
		PostRun:           nil,
		CompletionOptions: cobra.CompletionOptions{},
		TraverseChildren:  true,
		Hidden:            false,
		SilenceErrors:     false,
		SilenceUsage:      false,
	}

	root.PersistentFlags().VarP(&logging, "log-level", "l", "sets and configures logging verbosity")

	commands.Execute(root)
}

func init() {
	version = strings.TrimSpace(version)
	if e := os.Setenv("VERSION", version); e != nil {
		exception := fmt.Errorf("unable to set VERSION: %w", e)

		panic(exception)
	}
}
