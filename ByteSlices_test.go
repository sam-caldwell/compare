package compare

import (
	"fmt"
	"testing"
)

func TestByteSlices(t *testing.T) {
	tests := []struct {
		a, b   []byte
		expect bool
	}{
		{[]byte{1, 2, 3}, []byte{1, 2, 3}, true},
		{[]byte{1, 2, 3}, []byte{1, 2, 4}, false},
		{[]byte{1, 2, 3}, []byte{1, 2}, false},
		{[]byte{}, []byte{}, true},
		{nil, nil, true},
		{[]byte{}, nil, false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := ByteSlices(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("ByteSlices(%v, %v) = %v; expected %v", tt.a, tt.b, result, tt.expect)
			}
		})
	}
}
