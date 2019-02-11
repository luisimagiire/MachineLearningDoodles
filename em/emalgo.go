package em

import (
	"fmt"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distmv"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
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

func PlotDataset(fileName string, dataset []float64){
	v := make(plotter.Values, len(dataset))
	for i, elem := range dataset{
		v[i] = elem
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	p.Add(h)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
		panic(err)
	}
}

func EMAlgo(dataset []float64, tol int) ([]float64, float64, float64,float64, float64,float64){

	// Initialize parameters
	sigmaZero := stat.Variance(dataset, nil)
	muZero := dataset[rand.Intn(len(dataset))]
	sigmaOne := sigmaZero
	muOne := dataset[rand.Intn(len(dataset))]
	piHat := rand.Float64()
	round := 0

	fmt.Printf("===================== \n")
	fmt.Printf("INITIALIZATION: \n")
	fmt.Printf("mu_zero: %v \n", muZero)
	fmt.Printf("mu_one: %v \n", muOne)
	fmt.Printf("sigma_zero: %v \n", sigmaZero)
	fmt.Printf("sigma_one: %v \n", sigmaOne)
	fmt.Printf("pi_hat: %v \n", piHat)

	// Auxiliary variables to the optimization step
	sigmaZeroRep := make([]float64, len(dataset))
	sigmaOneRep := make([]float64, len(dataset))
	muZeroRep := make([]float64, len(dataset))
	muOneRep := make([]float64, len(dataset))
	gammaRep :=  make([]float64, len(dataset))
	gamma := make([]float64, len(dataset))
	likelihoods := make([]float64, tol)

	for ; round < tol;{

		distOne := distuv.Normal{
			Mu:    muOne,
			Sigma: sigmaOne,
		}

		distZero := distuv.Normal{
			Mu:    muZero,
			Sigma: sigmaZero,
		}

		// Expectation Step - Compute Responsabilities
		for j:=0; j< len(dataset); j++{
			gamma[j] = (piHat*distZero.Prob(dataset[j]))/((1.0-piHat)*distOne.Prob(dataset[j]) + piHat*distZero.Prob(dataset[j]))
		}

		// Maximization Step
		for i:=0;i<len(gamma);i++{
			gammaRep[i] = 1.0 - gamma[i]
			muZeroRep[i] = gamma[i]*dataset[i]
			muOneRep[i] = (1.0 - gamma[i])*dataset[i]
		}

		gammaSum := floats.Sum(gamma)
		gammaRepSum := floats.Sum(gammaRep)

		// Update estimators
		muZero = floats.Sum(muZeroRep)/ gammaSum
		muOne = floats.Sum(muOneRep)/ gammaRepSum

		for i:=0;i<len(gamma);i++{
			sigmaOneRep[i] = (1.0 - gamma[i])*math.Pow(dataset[i] - muOne, 2)
			sigmaZeroRep[i] = gamma[i]*math.Pow(dataset[i] - muZero, 2)
		}

		sigmaZero = floats.Sum(sigmaZeroRep) / gammaSum
		sigmaOne = floats.Sum(sigmaOneRep) / gammaRepSum
		piHat = gammaSum / float64(len(gamma))

		likelihoods[round] = logLikelihood(piHat, dataset, distZero, distOne)
		round++
	}

	return likelihoods, muZero, muOne, sigmaZero, sigmaOne, piHat

}

func logLikelihood(pi float64, dataset []float64, normalZero distuv.Normal, normalOne distuv.Normal) float64{
	tmpRep := make([]float64, len(dataset))
	for i:= range tmpRep{
		tmpRep[i] = math.Log((1.0-pi)*normalOne.Prob(dataset[i]) + pi*normalZero.Prob(dataset[i]))
	}
	return floats.Sum(tmpRep)
}

