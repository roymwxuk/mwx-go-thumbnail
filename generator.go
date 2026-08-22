package thumbnail

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func GenerateThumbnail(
	imgPath string,
	config *Config,
) (string, error) {
	if !IsSupportedImage(imgPath) {
		return "", fmt.Errorf("unsupported image: %s", imgPath)
	}

	width := config.Width
	height := config.Height
	quality := config.Quality
	if width <= 0 {
		width = DefaultWidth
	}
	if height <= 0 {
		height = DefaultHeight
	}
	if quality <= 0 {
		quality = DefaultQuality
	}
	outputPath, err := buildThumbnailPath(imgPath, width, height, config.RootDir)
	if err != nil {
		return "", fmt.Errorf("build thumbnail path: %w", err)
	}
	fmt.Printf("outputPath=%s\n", outputPath)

	// Create the thumbnail output directory if needed.
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return "", fmt.Errorf("create thumbnail directory: %w", err)
	}

	img, err := imaging.Open(imgPath)
	if err != nil {
		return "", fmt.Errorf("open image: %w", err)
	}

	thumbnail := imaging.Fit(
		img,
		width,
		height,
		imaging.Lanczos,
	)

	f, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("create thumbnail file: %w", err)
	}
	defer f.Close()

	err = webp.Encode(
		f,
		thumbnail,
		&webp.Options{
			Quality: float32(quality),
		},
	)
	if err != nil {
		return "", fmt.Errorf("encode webp: %w", err)
	}

	return outputPath, nil
}
