package GA

import (
	"math"
	"math/rand"
)

type Individual struct {
	chromosome []byte
	Fitness    float64
	X          float64
}

func (ind *Individual) decode(xmin, xmax float64) float64 {
	var x float64
	for i := 0; i < len(ind.chromosome); i++ {
		x += float64(ind.chromosome[i]) * math.Exp2(float64(i))
	}
	ind.X = xmin + (xmax-xmin)*x/(math.Exp2(float64(len(ind.chromosome)))-1)
	return ind.X
}

func (ind *Individual) calcFitness(fitFunc func(float64) float64) float64 {
	ind.Fitness = fitFunc(ind.X)
	return ind.Fitness
}

func (ind *Individual) mutate(mutationProbability float64) {
	for i := range ind.chromosome {
		if rand.Float64() < mutationProbability {
			if ind.chromosome[i] == 0 {
				ind.chromosome[i] = 1
			} else if ind.chromosome[i] == 1 {
				ind.chromosome[i] = 0
			}
		}
	}
}

// Use one point crossover and change the rear section of the two individuals
func (ind *Individual) crossover(target *Individual) {
	pos := rand.Intn(len(ind.chromosome))
	for i := pos; i < len(ind.chromosome); i++ {
		ind.chromosome[i], target.chromosome[i] = target.chromosome[i], ind.chromosome[i]
	}
}

// Deep copy of an individual and return a pointer to the copy of this individual
func (ind *Individual) copy() *Individual {
	chromosomeCopy := make([]byte, len(ind.chromosome))
	for j := range ind.chromosome {
		chromosomeCopy[j] = ind.chromosome[j]
	}
	indNew := &Individual{
		chromosome: chromosomeCopy,
		Fitness:    ind.Fitness,
		X:          ind.X,
	}
	return indNew
}
