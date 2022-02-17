package square

import (
	"math"
)

type Num int

const (
	SidesCircle   Num = 0
	SidesTriangle Num = 3
	SidesSquare   Num = 4
)

func CalcSquare(sideLen float64, sidesNum Num) float64 {
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
