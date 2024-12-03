package main

import (
	"path"

	"github.com/Hakitsyu/simple-titles-cli/internal"
	jsonstore "github.com/Hakitsyu/simple-titles-cli/internal/store/json"
)

func main() {
	internal.CreateInitialResources()

	println(internal.Store.GetDefaultSourceName())

	internal.Store.SetDefaultSource("Hello World")
	println(internal.Store.GetDefaultSourceName())

	internal.SourceStore.AddSource("OPA", "teste")

	internal.TagStore.AddTag("Anime", "a", "Tag relacionada a titulos de anime")

	sourceStore := jsonstore.NewJsonSourceStore(path.Join(internal.AppSourcesDirPath, "store.json"))

	sourceStore.AddTitle("Naruto", []string{"a"})
}
