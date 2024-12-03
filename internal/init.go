package internal

import (
	"errors"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/Hakitsyu/simple-titles-cli/configs"
)

const defaultStoreFileName = "default.json"

func CreateInitialResources() {
	if existsAppFolder() {
		return
	}

	createAppFolder()
	createStoreFile()
	createSourcesFolder()
	createDefaultSourceFile()
}

func existsAppFolder() bool {
	_, err := os.Stat(AppDirPath)
	return !(err != nil && errors.Is(err, os.ErrNotExist))
}

func createAppFolder() {
	os.Mkdir(AppDirPath, os.ModePerm)
}

func createStoreFile() {
	configContent, err := configs.GetEmbeddedStoreAsString()
	if err != nil {
		panic(err)
	}

	escapedAppDirPath := strings.ReplaceAll(strconv.Quote(AppDirPath), "\"", "")
	configContent = strings.ReplaceAll(configContent, "%APPDIR%", escapedAppDirPath)

	err = os.WriteFile(AppConfigPath, []byte(configContent), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func createSourcesFolder() {
	os.Mkdir(AppSourcesDirPath, os.ModePerm)
}

func createDefaultSourceFile() {
	defaultSourceContent, err := configs.GetEmbeddedDefaultSourceAsString()
	if err != nil {
		panic(err)
	}

	defaultSourcePath := path.Join(AppSourcesDirPath, defaultStoreFileName)

	err = os.WriteFile(defaultSourcePath, []byte(defaultSourceContent), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
