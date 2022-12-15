package GA

import (
	"math/rand"
	"sort"
)

type SelectionFunc func(*Population) *Individual

func RouletteSelection(pop *Population) *Individual {
	N := rand.Float64()
	for i := range pop.individuals {
		if pop.fitnessRunningTotal[i] >= N {
			return pop.individuals[i]
		}
	}
	return nil
}

func TournamentSelection(pop *Population) *Individual {
	M := pop.Args.Size / 2
	individuals := make([]*Individual, M)
	for i := range individuals {
		individuals[i] = pop.individuals[rand.Intn(pop.Args.Size)]
	}
	sort.Sort(Individuals(individuals))
	return individuals[0]
}
