package store

import "github.com/google/uuid"

type AppStore interface {
	GetDefaultSource() *SourceModel
	GetDefaultSourceName() string
	SetDefaultSource(sourceName string)
}

type SourceStore interface {
	AddSource(name string, path string, description string)
	RemoveSource(name string)
	GetSources() []SourceModel
	GetSource(sourceName string) *SourceModel
	ExistsSource(sourceName string) bool
}

type TagStore interface {
	AddTag(name string, symbol string, description string)
}

type TitleStore interface {
	AddTitle(title string, tags []string) uuid.UUID
	RemoveTitle(id uuid.UUID)
	GetTitles() []TitleModel
}
