package packer

import (
	"os"
	"path/filepath"
	"regexp"
)

// NormalizePath will clean the dir to be an nodejs pure dir
func NormalizePath(directoryPath string) []string {
	dir := []string{}
	basePath := filepath.Base(directoryPath)

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		nodeModulesFilesOrPackageLock, err := regexp.MatchString("node_modules|package-lock", path)
		if err != nil {
			return err
		}

		isConfigFile := info.Name()[:1] == "."

		if nodeModulesFilesOrPackageLock || isConfigFile || info.IsDir() {
			return nil
		}

		if info.Name() == basePath && info.IsDir() {
			return nil
		}

		dir = append(dir, path)

		return nil
	})

	if err != nil {
		return []string{}
	}

	return dir
}
