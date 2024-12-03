package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hakitsyu/simple-titles-cli/internal"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	titleCommand.AddCommand(titleListCommand)
	titleCommand.AddCommand(titleAddCommand)
	titleCommand.AddCommand(titleRemoveCommand)
	titleCommand.AddCommand(titleImportCommand)
	ConfigureSourceFlag(titleCommand, "Source used to handle your titles")

	rootCommand.AddCommand(titleCommand)
}

var titleCommand = &cobra.Command{
	Use:   "title",
	Short: "Manage your titles",
}

var titleListCommand = &cobra.Command{
	Use:   "list",
	Short: "List your titles",
	Run: func(cmd *cobra.Command, args []string) {
		source := GetCurrentSource()

		store := internal.NewTitleStoreBySourceName(source)

		titles := store.GetTitles()

		if len(titles) == 0 {
			fmt.Println(`
	You don't have any title yet.
			`)

			return
		}

		fmt.Println("")

		for _, title := range store.GetTitles() {
			var tags string
			if len(title.Tags) > 0 {
				tags = " [" + strings.Join(title.Tags, ", ") + "]"
			}

			fmt.Printf("	- %s%s (%s)\n", title.Name, tags, title.Id.String())
		}

		fmt.Println("")
	},
}

var titleAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a title",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := internal.Store.GetDefaultSource()

		store := internal.NewTitleStoreBySourceName(source.Name)

		title := args[0]

		var tags []string
		if len(args) > 1 {
			tags = strings.Split(args[1], ",")
		}

		store.AddTitle(title, tags)

		fmt.Println(`
	Title added successfully.
		`)
	},
}

// Desenvolva o comando de remover titles, recebendo o ID do title
var titleRemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a title",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := uuid.Parse(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		source := GetCurrentSource()

		store := internal.NewTitleStoreBySourceName(source)

		store.RemoveTitle(id)

		fmt.Println(`
	Title removed successfully.
		`)
	},
}

var titleImportCommand = &cobra.Command{
	Use:   "import",
	Short: "Import titles from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		source := GetCurrentSource()
		store := internal.NewTitleStoreBySourceName(source)

		titlesQty := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			content := scanner.Text()

			var (
				title string
				tags  []string
			)

			if strings.Contains(content, "(") && strings.Contains(content, ")") {
				tagsContent := strings.Split(content, "(")[1]
				tagsContent = strings.Split(tagsContent, ")")[0]
				tagsContent = strings.ReplaceAll(tagsContent, " ", "")

				tags = strings.Split(tagsContent, "-")
				title = strings.Split(content, ")")[1]
			} else {
				title = content
			}

			titlesQty++
			store.AddTitle(title, tags)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(`
	Error reading file:
			`, err)
		} else {
			fmt.Printf(`
	%d Titles imported successfully.
			`, titlesQty)
		}
	},
}
