package structs

import (
	"sort"
	"strconv"
)

// BigInteger is a structure holding a int number
// of great length in string format.
type BigInteger struct {

	// Value is the string format of a great number
	Value string
}

// BigIntegers sorts an array of BigInteger objects
func BigIntegers(a []BigInteger) { sort.Sort(BigIntegerSlice(a)) }

type BigIntegerSlice []BigInteger

func (p BigIntegerSlice) Sort()              { sort.Sort(p) }
func (p BigIntegerSlice) Len() int           { return len(p) }
func (p BigIntegerSlice) Less(i, j int) bool { return compare(p[i], p[j]) == -1 }
func (p BigIntegerSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func compare(a, b BigInteger) int {
	vA := a.Value
	vB := b.Value
	sizeA := len(vA)
	sizeB := len(vB)
	if sizeA > sizeB {
		return 1
	}
	if sizeA < sizeB {
		return -1
	}
	if sizeA == sizeB {
		if sizeA > 19 {

			// we break big integer into two parts
			// and we compare the first part.
			wA1, wA2 := breakBigIntegerValue(a)
			wB1, wB2 := breakBigIntegerValue(b)

			if wA1 > wB1 {
				return 1
			}
			if wA1 < wB1 {
				return -1
			}

			// if are equal we execute the compare function with the rest
			return compare(wA2, wB2)
		}
	}

	// size of both string is equal and they are less than 19
	iA, _ := strconv.Atoi(vA)
	iB, _ := strconv.Atoi(vB)
	if iA > iB {
		return 1
	}
	if iA == iB {
		return 0
	}
	return -1
}

func breakBigIntegerValue(a BigInteger) (intPart int, bigPart BigInteger) {
	value := a.Value
	size := len(value)
	if size > 19 {
		intPart, _ = strconv.Atoi(value[0:19])
		bigPart = BigInteger{
			Value: value[19:size],
		}
	} else {
		intPart, _ = strconv.Atoi(value[0:size])
		bigPart = BigInteger{
			Value: "0",
		}
	}
	return intPart, bigPart
}
