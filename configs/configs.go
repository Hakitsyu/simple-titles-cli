package configs

import (
	"embed"
)

//go:embed store.json
var configFile embed.FS

func GetEmbeddedStoreAsString() (string, error) {
	data, err := configFile.ReadFile("store.json")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

//go:embed default-source.json
var defaultSourceFile embed.FS

func GetEmbeddedDefaultSourceAsString() (string, error) {
	data, err := defaultSourceFile.ReadFile("default-source.json")
	if err != nil {
		return "", err
	}

	return string(data), nil
}