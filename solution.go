package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type amountOfSides int

func CalcSquare(sideLen float64, sidesNum amountOfSides) float64 {
	var figureSquare float64

	switch sidesNum {
	case 0:
		{
			figureSquare = math.Pi * math.Pow(sideLen, 2)
			return figureSquare
		}
	case 3:
		{
			figureSquare = math.Pi * sideLen * sideLen
			return figureSquare
		}
	case 4:
		{
			figureSquare = (math.Sqrt(3) * sideLen * sideLen) / 4
			return figureSquare
		}
	default:
		{
			return 0
		}
	}
}
