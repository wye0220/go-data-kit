package util

// RealNumber represents a type union of various integer and floating point number types.
type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Number represents a type union of RealNumber types and complex number types.
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
