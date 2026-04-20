package math

// Maps a value from an float64 range into another float64 range
func MapRange(value float64, xmin float64, xmax float64, ymin float64, ymax float64) float64 {
	return ymin + (value-xmin)*(ymax-ymin)/(xmax-xmin)
}
