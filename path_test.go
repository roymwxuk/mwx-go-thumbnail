package thumbnail

import "testing"

type testCase struct {
	path   string
	width  int
	height int
	expect string
}

func TestBuildThumbnailPath(t *testing.T) {
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
