package compare

// Float64WithPrecision - compares two floats precisely up to a specified number of decimal places.
func Float64WithPrecision(a, b float64, numberDecimalPlaces int) (result bool) {
	return TruncFloat64(a, numberDecimalPlaces) == TruncFloat64(b, numberDecimalPlaces)
}
