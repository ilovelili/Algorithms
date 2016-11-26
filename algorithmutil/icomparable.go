// Package algorithmutil provides basic algorithm and data structure implementation and defination
package algorithmutil

import (
	"bytes"
)

// IComparable define a generalized type-specific comparison method that a value type or class implements to order or sort its instances.
type IComparable interface {
	CompareTo(IComparable) int
}

// Integer define alias of int for CompareTo method
type Integer int

// MakeIntegerSlice make an interger slice
// this method is used to keep the compability of IComparable
func MakeIntegerSlice(a []int) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = Integer(a[i])
	}
	return b
}

// CompareTo compare int with int
func (input1 Integer) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(Integer); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	}

	panic("Must compare an integer type\n")
}

// Long define alias of int64 for CompareTo method
type Long int64

// MakeLongSlice make an int64 slice
// this method is used to keep the compability of IComparable
func MakeLongSlice(a []int64) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = Long(a[i])
	}
	return b
}

// CompareTo compare long with long
func (input1 Long) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(Long); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	}

	panic("Must compare a long type\n")
}

// Float define alias of float32
type Float float32

// MakeFloatSlice make an float32 slice
// this method is used to keep the compability of IComparable
func MakeFloatSlice(a []int64) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = Float(a[i])
	}
	return b
}

// CompareTo compare float with float
func (input1 Float) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(Float); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	}

	panic("Must compare a float type\n")
}

// Double define alias of float64
type Double float64

// MakeDoubleSlice make an float64 slice
// this method is used to keep the compability of IComparable
func MakeDoubleSlice(a []int64) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = Double(a[i])
	}
	return b
}

// CompareTo compare double with double
func (input1 Double) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(Double); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	}

	panic("Must compare a double type\n")
}

// Byte define alias of byte
type Byte byte

// MakeByteSlice make an byte slice
// this method is used to keep the compability of IComparable
func MakeByteSlice(a []int64) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = Byte(a[i])
	}
	return b
}

// CompareTo compare string with string
func (input1 Byte) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(Byte); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	}

	panic("Must compare a byte type\n")
}

// String define alias of string
type String string

// MakeStringSlice make an string slice
// this method is used to keep the compability of IComparable
func MakeStringSlice(a []int64) []IComparable {
	b := make([]IComparable, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = String(a[i])
	}
	return b
}

// CompareTo compare string with string
func (input1 String) CompareTo(input2 IComparable) int {
	this := input1
	if that, ok := input2.(String); ok {
		thisBytes := []byte(this)
		thatBytes := []byte(that)

		return bytes.Compare(thisBytes, thatBytes)
	}

	panic("Must compare a string type\n")
}
