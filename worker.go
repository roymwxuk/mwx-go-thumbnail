package thumbnail

import (
	"fmt"
	"sync"
)

// GenerateAll generates thumbnails concurrently using a worker pool.
func GenerateAll(config *Config) ([]Result, error) {
	paths, err := Scan(config.RootDir)
	if err != nil {
		return nil, err
	}

	return generateAll(paths, config), nil
}

func generateAll(paths []string, config *Config) []Result {
	if config.Workers <= 0 {
		config.Workers = DefaultWorkers
	}

	jobs := make(chan string)
	results := make(chan Result, len(paths))

	var wg sync.WaitGroup

	// Start workers.
	for i := 0; i < config.Workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for path := range jobs {
				fmt.Printf("worker: processing %s\n", path)
				result := Result{SourcePath: path}
				outputPath, err := GenerateThumbnail(path, config)
				if err != nil {
					result.Err = err
				}
				result.ThumbnailPath = outputPath

				results <- result
			}
		}()
	}

	// Send jobs to workers.
	go func() {
		for _, path := range paths {
			jobs <- path
		}
		close(jobs)
	}()

	// Close results after all workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	output := make([]Result, 0, len(paths))

	for result := range results {
		output = append(output, result)
	}

	return output
}
