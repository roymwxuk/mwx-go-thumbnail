# mwx-go-thumbnail

A lightweight asynchronous thumbnail generator for Go.

## Usage

`Generate` accepts either relative or absolute file paths.

Generate a thumbnail for a single image:

```go
gen.Generate("car/a.png")
```

To generate thumbnails for all supported images in a directory recursively, use `GenerateDir()`:

```go
gen.GenerateDir("car")
```
