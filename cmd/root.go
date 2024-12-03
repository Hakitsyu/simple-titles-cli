package cmd

import (
	"fmt"
	"os"

	"github.com/Hakitsyu/simple-titles-cli/internal"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "simple-titles",
	Short: "A CLI to manage your titles",
	Long:  "A CLI to manage your titles",
	Run: func(cmd *cobra.Command, args []string) {
		print("hello world")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var Source string

func GetCurrentSource() string {
	source := Source

	if source == "" {
		source = internal.Store.GetDefaultSourceName()
	}

	if !internal.SourceStore.ExistsSource(source) {
		fmt.Printf(`
	Source '%s' not exists.
		`, source)

		os.Exit(1)
	}

	return source
}

func ConfigureSourceFlag(command *cobra.Command, description string) {
	command.PersistentFlags().StringVarP(&Source, "source", "s", "", description)
}
