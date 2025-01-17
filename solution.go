package square

import "math"

type SideAmount int

const SidesCircle SideAmount = 0
const SidesTriangle SideAmount = 3
const SidesSquare SideAmount = 4

func CalcSquare(sideLen float64, sidesNum SideAmount) float64 {
	switch sidesNum {
	case 0:
		return math.Pow(sideLen, 2.0) * math.Pi
	case 3:
		return (math.Sqrt(3.0) / 4.0) * math.Pow(sideLen, 2.0)
	case 4:
		return math.Pow(sideLen, 2.0)
	default:
		return 0.0
	}
}
