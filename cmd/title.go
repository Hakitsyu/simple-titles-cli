package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Hakitsyu/simple-titles-cli/internal"
	"github.com/google/uuid"
	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
)

func init() {
	titleCommand.AddCommand(titleSearchCommand)
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

			titleName := strings.TrimSpace(title.Name)
			fmt.Printf("	- %s%s (%s)\n", titleName, tags, title.Id.String())
		}

		fmt.Println("")
	},
}

var titleSearchCommand = &cobra.Command{
	Use:   "search [query]",
	Short: "Search titles using JMESPath expressions",
	Long: `Search titles using JMESPath query expressions. Examples:

	# Search by name containing "test"
	title search "[?contains(Name, 'test')]"

	# Search by specific tag
	title search "[?contains(Tags, 'important')]" 

	# Search by ID
	title search "[?Id=='123e4567-e89b-12d3-a456-426614174000']"

	# Combine conditions
	title search "[?contains(Name,'test') && contains(Tags, 'important')]"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := GetCurrentSource()

		store := internal.NewTitleStoreBySourceName(source)

		titles := store.GetTitles()
		titlesJson := make([]map[string]interface{}, len(titles))
		for i, title := range titles {
			titlesJson[i] = map[string]interface{}{
				"Id":   title.Id.String(),
				"Name": title.Name,
				"Tags": title.Tags,
			}
		}

		jsonData, err := json.Marshal(titlesJson)
		if err != nil {
			fmt.Printf("Error converting data: %v\n", err)
			return
		}

		var data interface{}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}

		query := args[0]

		result, err := jmespath.Search(query, data)
		if err != nil {
			fmt.Printf("Error executing query: %v\n", err)
			return
		}

		resultJson, _ := json.MarshalIndent(result, "", "    ")
		fmt.Println(string(resultJson))
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
