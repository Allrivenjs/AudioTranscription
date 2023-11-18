package Transform

import (
	"io/fs"
	"path/filepath"
)

func GetAllFilesPath(pathTemp string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(pathTemp, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
