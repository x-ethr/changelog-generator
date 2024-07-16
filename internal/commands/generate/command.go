package generate

import (
    "fmt"
    "strings"

    "github.com/spf13/cobra"

    "changelog-generator/internal/constants"
)

var Command = &cobra.Command{
    Use:        "generate",
    Aliases:    []string{},
    SuggestFor: nil,
    Short:      "The example's command short-description",
    Long:       "The example's command long-description -- value should be in full sentences, and can span multiple lines.",
    Example: strings.Join([]string{
        fmt.Sprintf("  %s", "# General command usage"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\"", constants.Name())),
        "",
        fmt.Sprintf("  %s", "# Extended usage demonstrating configuration of default(s)"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\" --output json", constants.Name())),
        "",
        fmt.Sprintf("  %s", "# Display help information and command usage"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s example --help", constants.Name())),
    }, "\n"),
    ValidArgs:              nil,
    ValidArgsFunction:      nil,
    Args:                   nil,
    ArgAliases:             nil,
    BashCompletionFunction: "",
    Deprecated:             "",
    Annotations:            nil,
    Version:                "",
    PreRun:                 func(cmd *cobra.Command, args []string) {},
    RunE: func(cmd *cobra.Command, args []string) error {
        ctx := cmd.Context()

        _ = ctx

        return nil
    },
    PostRun:           nil,
    CompletionOptions: cobra.CompletionOptions{},
    TraverseChildren:  false,
    Hidden:            false,
    SilenceErrors:     true,
    SilenceUsage:      false,
}

func init() {
    // flags := Command.Flags()

    // flags.StringVarP(&name, "name", "n", "", "a required example named-string-flag")
    // flags.VarP(&format, "output", "o", "structured data format")
    // if e := Command.MarkFlagRequired("name"); e != nil {
    // 	if exception := Command.Help(); exception != nil {
    // 		panic(exception)
    // 	}
    // }
}
