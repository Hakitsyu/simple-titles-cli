package store

import (
	"encoding/json"
	"os"
)

type JsonAppStore struct {
	FilePath string
	Content  *AppJson
}

func NewJsonAppStore(filePath string) *JsonAppStore {
	content := readAppJson(filePath)

	return &JsonAppStore{
		FilePath: filePath,
		Content:  content,
	}
}

func readAppJson(filePath string) *AppJson {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	var content AppJson
	err = json.Unmarshal(data, &content)
	if err != nil {
		return nil
	}

	return &content
}

func (s JsonAppStore) GetDefaultSource() *SourceModel {
	sourceName := s.GetDefaultSourceName()
	for _, source := range s.Content.Sources {
		if source.Name == sourceName {
			return source.ToSourceModel()
		}
	}
	return nil
}

func (s JsonAppStore) GetDefaultSourceName() string {
	return s.Content.DefaultSource
}

func (s JsonAppStore) SetDefaultSource(sourceName string) {
	s.Content.DefaultSource = sourceName
	s.SaveContent()
}

func (s JsonAppStore) AddSource(name string, path string) {
	s.Content.Sources = append(s.Content.Sources, SourceJson{
		Name: name,
		Path: path,
	})

	s.SaveContent()
}

func (s JsonAppStore) RemoveSource(name string) {
	for i, source := range s.Content.Sources {
		if source.Name == name {
			s.Content.Sources = append(s.Content.Sources[:i], s.Content.Sources[i+1:]...)
			s.SaveContent()
			break
		}
	}
}

func (s JsonAppStore) AddTag(name string, symbol string, description string) {
	s.Content.Tags = append(s.Content.Tags, TagJson{
		Name:        name,
		Symbol:      symbol,
		Description: description,
	})

	s.SaveContent()
}

func (s JsonAppStore) RemoveTag(name string) {
	for i, tag := range s.Content.Tags {
		if tag.Name == name {
			s.Content.Tags = append(s.Content.Tags[:i], s.Content.Tags[i+1:]...)
			s.SaveContent()
			break
		}
	}
}

func (s JsonAppStore) ReloadContent() {
	s.Content = readAppJson(s.FilePath)
}

func (s JsonAppStore) SaveContent() {
	data, err := json.Marshal(s.Content)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(s.FilePath, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

type AppJson struct {
	DefaultSource string       `json:"defaultSource"`
	Sources       []SourceJson `json:"sources"`
	Tags          []TagJson    `json:"tags"`
}

type SourceJson struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (s SourceJson) ToSourceModel() *SourceModel {
	return &SourceModel{
		Name: s.Name,
		Path: s.Path,
	}
}

type TagJson struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

func (t TagJson) ToTagModel() *TagModel {
	return &TagModel{
		Name:        t.Name,
		Symbol:      t.Symbol,
		Description: t.Description,
	}
}
