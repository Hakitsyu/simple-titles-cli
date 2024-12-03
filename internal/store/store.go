package store

type AppStore interface {
	GetDefaultSource() *SourceModel
	GetDefaultSourceName() string
	SetDefaultSource(sourceName string)
}

type SourceStore interface {
	AddSource(name string, path string)
	RemoveSource(name string)
}

type TagStore interface {
	AddTag(name string, symbol string, description string)
}
