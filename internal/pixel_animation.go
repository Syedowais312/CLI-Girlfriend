package internal

import (
	"fmt"
	"time"
)

// RenderPixelAnimation takes a list of image file paths (frames) and repeatedly
// renders them as pixel art frames while *thinking is true. It clears the
// terminal between frames. delayMs controls frame duration in milliseconds.
func RenderPixelAnimation(frames []string, maxCols, maxRows int, thinking *bool, delayMs int) {
	if len(frames) == 0 {
		return
	}

	// Enter alternate screen buffer so animation doesn't clobber the user's shell.
	// Many terminals support the 1049 alternate buffer sequence.
	fmt.Print("\x1b[?1049h")
	// Ensure we restore the normal buffer when done.
	defer fmt.Print("\x1b[?1049l")

	i := 0
	for *thinking {
		// clear screen and move cursor to top-left before drawing each frame
		ConsoleLock.Lock()
		fmt.Print("\x1b[2J\x1b[H")

		// attempt to render the next frame (file path)
		framePath := frames[i%len(frames)]
		// best-effort: ignore errors and continue
		_ = RenderPixelArt(framePath, maxCols, maxRows)
		ConsoleLock.Unlock()

		i++
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}