package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Hakitsyu/simple-titles-cli/configs"
	"github.com/Hakitsyu/simple-titles-cli/internal"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	sourceCommand.AddCommand(sourceListCommand)
	sourceCommand.AddCommand(sourceAddCommand)
	sourceCommand.AddCommand(sourceRemoveCommand)
	sourceCommand.AddCommand(sourceSetDefaultCommand)

	rootCommand.AddCommand(sourceCommand)
}

var sourceCommand = &cobra.Command{
	Use:   "source",
	Short: "Manage your sources",
}

var sourceListCommand = &cobra.Command{
	Use:   "list",
	Short: "List your sources",
	Run: func(cmd *cobra.Command, args []string) {
		defaultSource := internal.Store.GetDefaultSourceName()

		fmt.Println("")

		for _, source := range internal.SourceStore.GetSources() {
			var description string
			if source.Description != "" {
				description = "(" + source.Description + ")"
			}

			if source.Name == defaultSource {
				fmt.Printf("	* %s %s\n", source.Name, description)
			} else {
				fmt.Printf("	- %s %s\n", source.Name, description)
			}
		}

		fmt.Println("")
	},
}

var sourceAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a source",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var description string

		if len(args) > 1 {
			description = args[1]
		}

		if internal.SourceStore.ExistsSource(name) {
			fmt.Printf(`
	Source '%s' already exists.
		`, name)
		}

		defaultSourceContent, err := configs.GetEmbeddedDefaultSourceAsString()
		if err != nil {
			panic(err)
		}

		fileName := uuid.NewString() + ".json"
		filePath := path.Join(internal.AppSourcesDirPath, fileName)

		err = os.WriteFile(filePath, []byte(defaultSourceContent), os.ModePerm)
		if err != nil {
			panic(err)
		}

		internal.SourceStore.AddSource(name, filePath, description)

		fmt.Printf(`
	Source '%s' added successfully.
		`, name)
	},
}

var sourceRemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceName := args[0]

		if !internal.SourceStore.ExistsSource(sourceName) {
			fmt.Printf(`
	Source '%s' not exists.
		`, sourceName)

			return
		}

		internal.SourceStore.RemoveSource(sourceName)

		fmt.Printf(`
	Source '%s' removed successfully.
		`, sourceName)
	},
}

var sourceSetDefaultCommand = &cobra.Command{
	Use:   "set-default",
	Short: "Set the default source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceName := args[0]

		if !internal.SourceStore.ExistsSource(sourceName) {
			fmt.Printf(`
	Source '%s' does not exist.\n
			`, sourceName)
			return
		}

		internal.Store.SetDefaultSource(sourceName)

		fmt.Println(`
	Default source set successfully.
		`)
	},
}
