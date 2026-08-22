package thumbnail

import (
	"io/fs"
	"path/filepath"
)

// Scan returns all supported images under rootDir.
func Scan(rootDir string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip generated thumbnails.
		if d.IsDir() && d.Name() == DefaultFolder {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		if IsSupportedImage(path) {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}
