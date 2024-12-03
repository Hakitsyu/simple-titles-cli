package internal

import (
	"os"
	"path"
	"sync"
)

var (
	AppDirPath        string
	AppSourcesDirPath string
	AppConfigPath     string

	once sync.Once
)

func LoadPaths() {
	once.Do(func() {
		AppDirPath = getAppDirPath()
		AppSourcesDirPath = getAppSourcesDirPath()
		AppConfigPath = getAppConfigPath()
	})
}

func getAppDirPath() string {
	dir, err := os.UserConfigDir()

	if err != nil {
		panic(err)
	}

	return path.Join(dir, AppName)
}

func getAppSourcesDirPath() string {
	return path.Join(getAppDirPath(), "sources")
}

func getAppConfigPath() string {
	return path.Join(getAppDirPath(), "app.json")
}
