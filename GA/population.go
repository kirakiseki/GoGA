package GA

import (
	"math/rand"
	"sort"
)

type Population struct {
	Args                Args
	individuals         []*Individual
	fitnessSum          float64
	fitnessRunningTotal []float64
}

// Individuals Implement sort.Interface
type Individuals []*Individual

func (individuals Individuals) Len() int {
	return len(individuals)
}

func (individuals Individuals) Less(i, j int) bool {
	return individuals[i].Fitness > individuals[j].Fitness
}

func (individuals Individuals) Swap(i, j int) {
	individuals[i], individuals[j] = individuals[j], individuals[i]
}

// Calculate fitness of each individual and return the sum of all fitness, running total (prefix sum of fitness) and a map containing fitness and x value for each individual
func (pop *Population) calcFitness() (float64, []float64, map[*Individual][]float64) {
	var fitMap = make(map[*Individual][]float64, pop.Args.Size)
	var sum float64
	runningTotal := make([]float64, pop.Args.Size)
	for i := range pop.individuals {
		sum += pop.individuals[i].calcFitness(pop.Args.FitFunc)
		fitMap[pop.individuals[i]] = []float64{pop.individuals[i].X, pop.individuals[i].Fitness}
		runningTotal[i] = sum
	}
	for i := range runningTotal {
		runningTotal[i] /= sum
	}
	pop.fitnessSum, pop.fitnessRunningTotal = sum, runningTotal
	return sum, runningTotal, fitMap
}

// Crossover (one point crossover) between two individuals
// Shuffle all the individuals to confirm that the process is random
func (pop *Population) crossover() {
	crossoverCount := 0
	for range pop.individuals {
		if rand.Float64() < pop.Args.CrossoverProbability {
			crossoverCount++
		}
	}
	if crossoverCount%2 != 0 {
		crossoverCount--
	}
	rand.Shuffle(pop.Args.Size, func(i, j int) {
		pop.individuals[i], pop.individuals[j] = pop.individuals[j], pop.individuals[i]
	})
	for i := 0; i < crossoverCount; i += 2 {
		pop.individuals[i].crossover(pop.individuals[i+1])
	}
}

// Evolve Implement of the genetic algorithm
// return next generation, optimized value at this generation and all the fitness and x value of all individuals in this generation
func (pop *Population) Evolve() (*Population, float64, map[*Individual][]float64) {
	for i := range pop.individuals {
		pop.individuals[i].decode(pop.Args.Range[0], pop.Args.Range[1])
	}
	var fitMap map[*Individual][]float64
	pop.fitnessSum, pop.fitnessRunningTotal, fitMap = pop.calcFitness()

	nextGeneration := &Population{
		Args:        pop.Args,
		individuals: make([]*Individual, pop.Args.Size),
	}

	for i := range nextGeneration.individuals {
		nextGeneration.individuals[i] = pop.Args.SelectFunc(pop).copy()
	}

	nextGeneration.crossover()
	for i := range nextGeneration.individuals {
		nextGeneration.individuals[i].mutate(pop.Args.MutationProbability)
	}
	nextGeneration.Args.Generation++

	// evaluate the fitness of selected individuals after crossover and mutation
	for i := range pop.individuals {
		nextGeneration.individuals[i].decode(pop.Args.Range[0], pop.Args.Range[1])
	}

	nextGeneration.fitnessSum, nextGeneration.fitnessRunningTotal, _ = nextGeneration.calcFitness()

	// combine this generation and individuals after crossover and mutation
	allIndividuals := append(copyIndividuals(pop.individuals), copyIndividuals(nextGeneration.individuals)...)
	sort.Sort(Individuals(allIndividuals))
	// select the best individuals (copy process in GA algorithm)(next generation contains part of this generation and part of individuals after crossover and mutation)
	nextGeneration.individuals = allIndividuals[:pop.Args.Size]
	optimizedValue := nextGeneration.individuals[0].Fitness
	// get optimized value of this generation
	for i := range nextGeneration.individuals {
		if nextGeneration.individuals[i].Fitness > optimizedValue {
			optimizedValue = nextGeneration.individuals[i].Fitness
		}
	}
	return nextGeneration, optimizedValue, fitMap
}

func (pop *Population) InitRandomPopulation() {
	pop.Args.ChromosomeLength = CalcChromosomeLength(pop.Args.Range[0], pop.Args.Range[1], pop.Args.Precision)
	pop.individuals = make([]*Individual, pop.Args.Size)
	for i := range pop.individuals {
		pop.individuals[i] = &Individual{
			chromosome: make([]byte, pop.Args.ChromosomeLength),
		}
		// initialize the chromosome of each individual randomly
		for j := range pop.individuals[i].chromosome {
			// for each bit of chromosome, randomly assign 0 or 1
			pop.individuals[i].chromosome[j] = byte(rand.Intn(2))
		}
	}
}
