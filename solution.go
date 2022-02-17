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
		return math.Floor(math.Pi*100) / 100 * sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * sideLen * sideLen
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else {
		return 0
	}
}
