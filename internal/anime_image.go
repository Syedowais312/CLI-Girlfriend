package internal

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"

	xdraw "golang.org/x/image/draw"
)

// RenderPixelArt renders an image specifically optimized for pixel art.
// It removes black backgrounds automatically and keeps edges sharp.
// Suggest maxCols: 80, maxRows: 40 for standard terminals.
func RenderPixelArt(path string, maxCols, maxRows int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// --- 1. Scaling (Preserve Aspect Ratio) ---
	bounds := img.Bounds()
	origW := float64(bounds.Dx())
	origH := float64(bounds.Dy())

	// Terminal characters are ~1:2 ratio. We use 2 vertical pixels per char.
	targetW := float64(maxCols)
	targetH := float64(maxRows * 2)

	scale := math.Min(targetW/origW, targetH/origH)

	// Calculate new size
	newW := int(math.Max(1, math.Round(origW*scale)))
	newH := int(math.Max(1, math.Round(origH*scale)))

	// Force height to be even for half-block rendering
	if newH%2 != 0 {
		newH--
	}

	// --- 2. Resize using NearestNeighbor (CRITICAL for Pixel Art) ---
	// This keeps the pixels "blocky" and sharp instead of blurry.
	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
	xdraw.NearestNeighbor.Scale(dst, dst.Bounds(), img, bounds, xdraw.Over, nil)

	// --- 3. Render Loop ---
	for y := 0; y < newH; y += 2 {
		for x := 0; x < newW; x++ {
			// Get colors
			c1 := dst.At(x, y)
			c2 := dst.At(x, y+1)

			// Convert to NRGBA
			p1 := color.NRGBAModel.Convert(c1).(color.NRGBA)
			p2 := color.NRGBAModel.Convert(c2).(color.NRGBA)

			// --- 4. Magic Background Removal ---
			// If the image is a JPG, transparency is lost and becomes black.
			// We manually check if the pixel is "Dark enough" to be background.
			// Threshold: (R+G+B) < 30 means it's almost pure black.
			isTransparent1 := p1.A < 10 || (int(p1.R)+int(p1.G)+int(p1.B) < 30)
			isTransparent2 := p2.A < 10 || (int(p2.R)+int(p2.G)+int(p2.B) < 30)

			// --- 5. Draw Blocks ---
			if isTransparent1 && isTransparent2 {
				fmt.Print(" ") // Transparent
			} else if !isTransparent1 && !isTransparent2 {
				// Both Opaque: Bottom half (FG) + Top half (BG)
				fmt.Printf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm▄\x1b[0m",
					p2.R, p2.G, p2.B, // FG
					p1.R, p1.G, p1.B, // BG
				)
			} else if !isTransparent1 && isTransparent2 {
				// Top Opaque Only: Upper Block (▀)
				fmt.Printf("\x1b[38;2;%d;%d;%dm▀\x1b[0m", p1.R, p1.G, p1.B)
			} else if isTransparent1 && !isTransparent2 {
				// Bottom Opaque Only: Lower Block (▄)
				fmt.Printf("\x1b[38;2;%d;%d;%dm▄\x1b[0m", p2.R, p2.G, p2.B)
			}
		}
		fmt.Println()
	}

	return nil
}