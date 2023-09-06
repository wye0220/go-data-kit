package slice

import (
	"github.com/wye0220/GoDataKit/util"
)

// Sum calculates the sum of a slice of numbers of a generic type T.
func Sum[T util.Number](ts []T) T {
	if len(ts) == 0 {
		return T(0) // Return zero value if slice is empty
	}

	var res T
	for _, n := range ts {
		res += n
	}
	return res
}
