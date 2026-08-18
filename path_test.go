package thumbnail

import "testing"

func TestBuildThumbnailPath(t *testing.T) {
	type testCase struct {
		path   string
		width  int
		height int
		expect string
	}
	testCases := []testCase{
		{"car.png", 64, 64, ".thumbnail/64x64/car.webp"},
		{"car.jpg", 64, 64, ".thumbnail/64x64/car.webp"},
		{"foo/car.png", 64, 64, ".thumbnail/64x64/foo/car.webp"},
		{"foo/car.png", 300, 300, ".thumbnail/300x300/foo/car.webp"},
	}

	for i, item := range testCases {
		got := buildThumbnailPath(item.path, item.width, item.height)
		if got != item.expect {
			t.Fatalf("Case-%d failed: got %q, want %q", i, got, item.expect)
		}
	}
}

func TestIsSupportedImage(t *testing.T) {
	type testCase struct {
		path   string
		expect bool
	}
	testCases := []testCase{
		{"car.png", true},
		{"car.jpg", true},
		{"car.jpeg", true},
		{"car.webp", true},
		{"car.html", false},
		{"car", false},
		{"foo/car.png", true},
		{"foo/bar/car.png", true},
	}

	for i, item := range testCases {
		got := IsSupportedImage(item.path)
		if got != item.expect {
			t.Fatalf("Case-%d failed: got %v for path: %s\n", i, got, item.path)
		}
	}
}
