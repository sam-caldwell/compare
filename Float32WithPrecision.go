package compare

// Float32WithPrecision - compares two floats precisely up to a specified number of decimal places.
func Float32WithPrecision(a, b float32, numberDecimalPlaces int) (result bool) {
	return TruncFloat32(a, numberDecimalPlaces) == TruncFloat32(b, numberDecimalPlaces)
}
