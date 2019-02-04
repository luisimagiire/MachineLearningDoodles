package em

import (
	"fmt"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)


func GenerateMVGaussian(means []float64, sigma *mat.DiagDense, numObs int, seed uint64) [][]float64{
	normalMv, err := distmv.NewNormal(means, sigma, rand.NewSource(seed))
	if err != true{
		panic(fmt.Sprintf("Something went wrong generating mvn random variables !"))
	}

	var data = make([][]float64, numObs)
	for j := 0; j< numObs; j++{
		data[j] = normalMv.Rand(nil)
		}

	return data
}


func EMAlgo(){

}