package GA

import "math"

type FitnessFunc func(float64) float64

func TargetFunc(x float64) float64 {
	return x*math.Cos(x) + 2
}
