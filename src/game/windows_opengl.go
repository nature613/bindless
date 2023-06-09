//go:build windows

package game

import "os"

// currently directX performance is quite bad and can cause lag when
// skipping text screens which load many glyphs at once in a single frame
func init() {
	// allow directX if passed as program flag
	for _, arg := range os.Args {
		if arg == "--directX" || arg == "--directx" { return }
	}

	// set openGL as the graphics backend otherwise
	err := os.Setenv("EBITENGINE_GRAPHICS_LIBRARY", "opengl")
	if err != nil { panic(err) }
}
