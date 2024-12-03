package store

type SourceModel struct {
	Name string
	Path string
}

type TagModel struct {
	Name        string
	Symbol      string
	Description string
}

type TitleModel struct {
	Title string
	Tags  []string
}
