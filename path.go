package thumbnail

import (
	"fmt"
	"path/filepath"
	"strings"
)

const DefaultFolder = ".thumbnail"
const DefaultExt = ".webp"

func buildThumbnailPath(imgPath string, width int, height int, rootDir string) (string, error) {
	thumbnailDir := fmt.Sprintf("%s/%dx%d", DefaultFolder, width, height)
	relPath, err := filepath.Rel(rootDir, imgPath)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(relPath)
	filename := strings.TrimSuffix(filepath.Base(imgPath), filepath.Ext(imgPath)) + DefaultExt
	return filepath.Join(rootDir, thumbnailDir, dir, filename), nil
}

var extMap = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".webp": true,
}

func IsSupportedImage(imgPath string) bool {
	return extMap[filepath.Ext(imgPath)]
}
