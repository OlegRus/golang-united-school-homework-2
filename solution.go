package square

import "math"

type SideAmount int

func CalcSquare(sideLen float64, sidesNum SideAmount) float64 {
	switch sidesNum {
	case 0:
		r := (sideLen / math.Pi) / 2.0
		return math.Pow(r, 2.0) * math.Pi
	case 3:
		return (math.Sqrt(3.0) / 4.0) * math.Pow(sideLen, 2.0)
	case 4:
		return math.Pow(sideLen, 2.0)
	default:
		return 0.0
	}
}
