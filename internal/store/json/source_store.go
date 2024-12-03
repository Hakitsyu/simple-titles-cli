package jsonstore

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type JsonSourceStore struct {
	FilePath string
	Content  *SourceJson
}

func NewJsonSourceStore(filePath string) *JsonSourceStore {
	content := readSourceJson(filePath)

	return &JsonSourceStore{
		FilePath: filePath,
		Content:  content,
	}
}

func readSourceJson(filePath string) *SourceJson {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	var content SourceJson
	err = json.Unmarshal(data, &content)
	if err != nil {
		return nil
	}

	return &content
}

func (s JsonSourceStore) AddTitle(title string, tags []string) uuid.UUID {
	id := uuid.New()

	s.Content.Titles = append(s.Content.Titles, SourceTitleJson{
		Id:   id.String(),
		Name: title,
		Tags: tags,
	})

	s.SaveContent()

	return id
}

func (s JsonSourceStore) RemoveTitle(id uuid.UUID) {
	for i, title := range s.Content.Titles {
		if title.Id == id.String() {
			s.Content.Titles = append(s.Content.Titles[:i], s.Content.Titles[i+1:]...)
			s.SaveContent()
			break
		}
	}
}

func (s JsonSourceStore) ReloadContent() {
	s.Content = readSourceJson(s.FilePath)
}

func (s JsonSourceStore) SaveContent() {
	data, err := json.Marshal(s.Content)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(s.FilePath, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

type SourceJson struct {
	Titles []SourceTitleJson `json:"titles"`
}

type SourceTitleJson struct {
	Id   string   `json:"guid"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
