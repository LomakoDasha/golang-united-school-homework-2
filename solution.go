package square

import (
	"math"
)

type amountOfSides int

const (
	SidesTriangle amountOfSides = 3
	SidesSquare   amountOfSides = 4
	SidesCircle   amountOfSides = 0
)

func CalcSquare(sideLen float64, sidesNum amountOfSides) float64 {
	var figureSquare float64

	switch sidesNum {
	case SidesTriangle:
		{
			figureSquare = math.Pi * math.Pow(sideLen, 2)
		}
	case SidesSquare:
		{
			figureSquare = math.Pi * sideLen * sideLen
		}
	case SidesCircle:
		{
			figureSquare = (math.Sqrt(3) * sideLen * sideLen) / 4
		}
	default:
		{
			figureSquare = 0
		}
	}
	return figureSquare
}
