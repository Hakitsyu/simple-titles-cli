package internal

import (
	"github.com/Hakitsyu/simple-titles-cli/internal/store"
)

const AppName string = "Simple Titles CLI"

var (
	Store       store.AppStore
	SourceStore store.SourceStore
	TagStore    store.TagStore
)

func init() {
	LoadPaths()

	currentStore := store.NewJsonAppStore(AppConfigPath)

	Store = currentStore
	SourceStore = currentStore
	TagStore = currentStore
}
