package compare

// ByteSlices - compare two byte slices
func ByteSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
