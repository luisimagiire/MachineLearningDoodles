package em

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"

	"testing"
)

// Test K-means on gaussian generated values
func TestGenerateMVGaussian(t *testing.T){
	gaussianMeans := []float64{1.0, -1.0}
	sigma := mat.NewDiagDense(2, []float64{1,1})
	numObs := 10

	var dataSet = GenerateMVGaussian(gaussianMeans, sigma, numObs,uint64(99))
	fmt.Printf("2x2 - MV Gaussian:")
	for _, vec := range dataSet{
		fmt.Printf("%v \n", vec)
	}
}

func TestEMAlgo(t *testing.T) {
	means := []int{-3,3}
	numObs := 1000
	sigma := 1.5
	responsability := 0.7
	maxRounds := 200
	rand.Seed(int64(1))

	var gaussians = make([][]float64, len(means))

	for i := 0; i<len(means); i++{
		var tmpDist = distuv.Normal{Mu:float64(means[i]), Sigma: sigma}
		tmpvec := make([]float64, numObs)
		for j := 0; j<numObs; j++{
			tmpvec[j] = tmpDist.Rand()
		}
		gaussians[i] = tmpvec
	}

	dataset := GenerateMixtureGaussians(responsability, gaussians)
	PlotDataset("hist.png", dataset)

	run:=0
	maxRun := 20

	for ;run<maxRun;{
		rand.Seed(rand.Int63())
		logs, muZero, muOne, sigmaZero, sigmaOne, piHat:= EMAlgo(dataset, maxRounds)

		fmt.Printf("===================== \n")
		fmt.Printf("RUN: %v \n", run)
		fmt.Printf("mu_zero: %v \n", muZero)
		fmt.Printf("mu_one: %v \n", muOne)
		fmt.Printf("sigma_zero: %v \n", sigmaZero)
		fmt.Printf("sigma_one: %v \n", sigmaOne)
		fmt.Printf("pi_hat: %v \n", piHat)
		fmt.Printf("last_log: %v \n", logs[len(logs)-1])
		run++
	}

}