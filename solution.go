package square

import "math"

type Num int

const (
	SidesCircle   Num = 0
	SidesTriangle Num = 3
	SidesSquare   Num = 4
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	//return sideLen * sideLen
	// result := math.Sqrt(sideLen)
	// return result

	if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		return sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else {
		return 0
	}
}
