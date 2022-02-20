package solution

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type cacltype int

const SidesTriangle cacltype = 3
const SidesSquare cacltype = 4
const SidesCircle cacltype = 0

func CalcSquare(sideLen float64, sidesNum cacltype) float64 {
	switch sidesNum {
	case SidesTriangle:
		return ((math.Pow(sideLen, 2) * math.Pow(3, 0.5)) / 4)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return (math.Pi * (math.Pow(sideLen, 2)))
	default:
		return 0
	}
}
