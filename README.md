# mwx-go-thumbnail

A lightweight asynchronous thumbnail generator for Go.

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
