package thumbnail

import (
	"fmt"
	"path/filepath"
	"strings"
)

const DefaultFolder = ".thumbnail"
const DefaultExt = ".webp"

func buildThumbnailPath(imgPath string, width int, height int) string {
	thumbnailDir := fmt.Sprintf("%s/%dx%d", DefaultFolder, width, height)
	dir := filepath.Dir(imgPath)
	filename := strings.TrimSuffix(filepath.Base(imgPath), filepath.Ext(imgPath)) + DefaultExt
	return filepath.Join(thumbnailDir, dir, filename)
}
