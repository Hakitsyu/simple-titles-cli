package internal

import (
	"github.com/Hakitsyu/simple-titles-cli/internal/store"
	jsonstore "github.com/Hakitsyu/simple-titles-cli/internal/store/json"
)

func NewTitleStoreBySourceName(sourceName string) store.TitleStore {
	source := SourceStore.GetSource(sourceName)

	return jsonstore.NewJsonSourceStore(source.Path)
}
