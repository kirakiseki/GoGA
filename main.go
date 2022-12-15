package main

import (
	"GA/GA"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Set up the command line flags
	min := flag.Float64("min", -5.0, "minimum x value")
	max := flag.Float64("max", 5.0, "maximum x value")
	maxGeneration := flag.Int("g", 30, "maximum generation")
	populationSize := flag.Int("s", 100, "population size")
	mutationProbability := flag.Float64("m", 0.1, "MutationProbability")
	crossoverProbability := flag.Float64("c", 0.6, "CrossoverProbability")
	precision := flag.Float64("p", 0.01, "Precision")
	function := flag.String("f", "r", "function to select individuals (r for Roulette, t for Tournament)")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	// OptimizedValue in each generation
	optimizedValue := make([]float64, *maxGeneration)

	// Set init args for GA
	args := GA.Args{
		Generation:           1,
		Size:                 *populationSize,
		Range:                [2]float64{*min, *max},
		Precision:            *precision,
		MutationProbability:  *mutationProbability,
		CrossoverProbability: *crossoverProbability,
		FitFunc:              GA.TargetFunc,
		SelectFunc:           GA.RouletteSelection,
	}
	if *function == "t" {
		args.SelectFunc = GA.TournamentSelection
	}
	population := &GA.Population{
		Args: args,
	}
	population.InitRandomPopulation()

	fmt.Println("ChromosomeLength", population.Args.ChromosomeLength)

	// Evolve the population and store the fitness value of all individuals in each generation
	fitMap := make([]map[*GA.Individual][]float64, *maxGeneration)
	for i := 0; i < *maxGeneration; i++ {
		population, optimizedValue[i], fitMap[i] = population.Evolve()
		fmt.Println("Gen", i, optimizedValue[i])
	}

	// Visualize the fitness value of all individuals in each generation and save it as a gif file
	buffer := &bytes.Buffer{}
	generateGIF(buffer, fitMap, args.Range, GA.TargetFunc)
	_ = os.WriteFile("GA.gif", buffer.Bytes(), 0644)
}

func generateGIF(out io.Writer, fitMap []map[*GA.Individual][]float64, rangeX [2]float64, target GA.FitnessFunc) {
	palette := []color.Color{
		color.White,
		color.Black,
		color.NRGBA{R: 0xff, G: 0xae, B: 0x00, A: 0xff},
		color.NRGBA{R: 0x66, G: 0xcc, B: 0xff, A: 0xff},
	}
	anim := gif.GIF{LoopCount: len(fitMap)}
	const (
		scale       = 70
		blackIndex  = 1
		orangeIndex = 2
		blueIndex   = 3
	)
	delay := 16

	// Find Y range of target function to set the height of image
	var max, min float64
	for i := rangeX[0]; i <= rangeX[1]; i += 0.01 {
		if target(i) > max {
			max = target(i)
		}
		if target(i) < min {
			min = target(i)
		}
	}

	xSize := int((rangeX[1] - rangeX[0]) * scale)
	ySize := int((math.Max(math.Abs(max), math.Abs(min))*2 + 2) * scale)
	if ySize%2 == 1 {
		ySize += 1
	}
	if xSize%2 == 1 {
		xSize += 1
	}

	// Draw each generation in one frame
	for i := range fitMap {
		rect := image.Rect(0, 0, xSize, ySize)
		img := image.NewPaletted(rect, palette)
		// Draw X axis and mark
		for j := 0; j < xSize; j += 1 {
			img.SetColorIndex(j, ySize/2, blackIndex)
		}
		for j := xSize / 2; j < xSize; j += scale {
			img.SetColorIndex(j, ySize/2, blackIndex)
			img.SetColorIndex(j, ySize/2-1, blackIndex)
			img.SetColorIndex(j, ySize/2-2, blackIndex)
			img.SetColorIndex(j, ySize/2-3, blackIndex)
		}
		for j := xSize / 2; j > 0; j -= scale {
			img.SetColorIndex(j, ySize/2, blackIndex)
			img.SetColorIndex(j, ySize/2-1, blackIndex)
			img.SetColorIndex(j, ySize/2-2, blackIndex)
			img.SetColorIndex(j, ySize/2-3, blackIndex)
		}
		// Draw Y axis and mark
		for j := 0; j < ySize; j += 1 {
			img.SetColorIndex(xSize/2, j, blackIndex)
		}
		for j := ySize / 2; j < ySize; j += scale {
			img.SetColorIndex(xSize/2, j, blackIndex)
			img.SetColorIndex(xSize/2-1, j, blackIndex)
			img.SetColorIndex(xSize/2-2, j, blackIndex)
			img.SetColorIndex(xSize/2-3, j, blackIndex)
		}
		for j := ySize / 2; j > 0; j -= scale {
			img.SetColorIndex(xSize/2, j, blackIndex)
			img.SetColorIndex(xSize/2-1, j, blackIndex)
			img.SetColorIndex(xSize/2-2, j, blackIndex)
			img.SetColorIndex(xSize/2-3, j, blackIndex)
		}
		// Draw target function
		for j := rangeX[0]; j < rangeX[1]; j += 0.0001 {
			img.SetColorIndex(int((j-rangeX[0])*scale), int((-target(j))*scale)+ySize/2, orangeIndex)
		}
		// Draw each individual as a blue cross
		for j := range fitMap[i] {
			img.SetColorIndex(int((j.X-rangeX[0])*scale), int((-j.Fitness)*scale)+ySize/2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale), int((-j.Fitness)*scale)+ySize/2+1, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale), int((-j.Fitness)*scale)+ySize/2+2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale), int((-j.Fitness)*scale)+ySize/2-1, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale), int((-j.Fitness)*scale)+ySize/2-2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale)-1, int((-j.Fitness)*scale)+ySize/2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale)-2, int((-j.Fitness)*scale)+ySize/2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale)+1, int((-j.Fitness)*scale)+ySize/2, blueIndex)
			img.SetColorIndex(int((j.X-rangeX[0])*scale)+2, int((-j.Fitness)*scale)+ySize/2, blueIndex)
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
