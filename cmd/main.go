package main

import (
	"fmt"

	thumbnail "github.com/roymwxuk/mwx-go-thumbnail"
)

func main() {
	config := thumbnail.DefaultConfig()
	config.RootDir = "testdata"

	results, err := thumbnail.GenerateDir(&config)
	if err != nil {
		fmt.Printf("failed to generate thumbnails: %v\n", err)
		return
	}

	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("failed: %s: %v\n", result.SourcePath, result.Err)
			continue
		}

		fmt.Printf(
			"generated: %s -> %s\n",
			result.SourcePath,
			result.ThumbnailPath,
		)
	}
}
