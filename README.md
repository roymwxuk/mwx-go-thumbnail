# mwx-go-thumbnail

A lightweight concurrent thumbnail generator for Go.

## Features

- Concurrent thumbnail generation using a worker pool
- Recursive directory scanning
- Configurable worker count and thumbnail size
- WebP output

## Usage

`Generate()` accepts a path relative to the project root.

Generate a thumbnail for a single image:

```go
gen.Generate("car/a.png")
```

To generate thumbnails for all supported images in a directory recursively, use `GenerateDir()`:

```go
gen.GenerateDir("car")
```

