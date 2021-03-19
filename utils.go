package simpdf

var (
	// PIXEL_INCH is the number of pixels in an Imperial inch
	PIXEL_INCH = float64(96)
	// POINTS_INCH is the number of points in an Imperial inch
	POINTS_INCH = float64(72)
)

// PixelsToPoints function returns the number of points for the given pixels
func PixelsToPoints(p float64) float64 {
	return p * (POINTS_INCH / PIXEL_INCH)
}

// PointsToPixels function returns the number of pixels for the given points
func PointsToPixels(p float64) float64 {
	return p * (PIXEL_INCH / POINTS_INCH)
}
