package store

import "github.com/google/uuid"

type SourceModel struct {
	Name        string
	Path        string
	Description string
}

type TagModel struct {
	Name        string
	Symbol      string
	Description string
}

type TitleModel struct {
	Id   uuid.UUID
	Name string
	Tags []string
}
