package internal

import (
	"github.com/Hakitsyu/simple-titles-cli/internal/store"
	jsonstore "github.com/Hakitsyu/simple-titles-cli/internal/store/json"
)

const AppName string = "Simple Titles CLI"

var (
	Store       store.AppStore
	SourceStore store.SourceStore
	TagStore    store.TagStore
)

func init() {
	LoadPaths()
	CreateInitialResources()

	currentStore := jsonstore.NewJsonAppStore(AppConfigPath)

	Store = currentStore
	SourceStore = currentStore
	TagStore = currentStore
}
