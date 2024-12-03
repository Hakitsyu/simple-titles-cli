package main

import (
	"github.com/Hakitsyu/simple-titles-cli/internal"
)

func main() {
	internal.CreateInitialResources()

	println(internal.Store.GetDefaultSourceName())

	internal.Store.SetDefaultSource("Hello World")
	println(internal.Store.GetDefaultSourceName())

	internal.SourceStore.AddSource("OPA", "teste")

	internal.TagStore.AddTag("Anime", "a", "Tag relacionada a titulos de anime")
}
