package thumbnail

const (
	DefaultWidth   = 200
	DefaultHeight  = 200
	DefaultQuality = 80
	DefaultWorkers = 4
)

var rootDir string

type Config struct {
	Width   int
	Height  int
	Quality int
	Workers int

	RootDir string
}

// Result represents the result of generating a thumbnail.
type Result struct {
	SourcePath    string
	ThumbnailPath string // output path
	Err           error
}

// DefaultConfig returns the default thumbnail generation settings.
func DefaultConfig() Config {
	return Config{
		Width:   DefaultWidth,
		Height:  DefaultHeight,
		Quality: DefaultQuality,
		Workers: DefaultWorkers,
		RootDir: ".",
	}
}

// Setup sets the root directory containing source images.
func SetRoot(root string) {
	rootDir = root
}
