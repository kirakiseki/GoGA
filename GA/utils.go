package GA

import (
	"fmt"
	"math"
)

type Args struct {
	Generation           int
	Size                 int
	Range                [2]float64
	Precision            float64
	ChromosomeLength     int
	MutationProbability  float64
	CrossoverProbability float64
	FitFunc              FitnessFunc
	SelectFunc           SelectionFunc
}

func CalcChromosomeLength(min, max, precision float64) int {
	fmt.Println("min:", min, "max:", max, "precision:", precision)
	return int(math.Ceil(math.Log2((max - min) / precision)))
}

// Deep copy of individuals
func copyIndividuals(individuals []*Individual) []*Individual {
	newIndividuals := make([]*Individual, len(individuals))
	for i := range individuals {
		newIndividuals[i] = individuals[i].copy()
	}
	return newIndividuals
}
