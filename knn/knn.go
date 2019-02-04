package knn

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"math/rand"
)

func GenerateGaussian(means []int, numObs int, sigma float64) []float64{
	totalObs := len(means)*numObs
	var data = make([]float64, totalObs, totalObs)

	for idx, elem := range means{
		var tmpDist = distuv.Normal{Mu:float64(elem), Sigma: sigma}
		for j := 0; j< numObs; j++{
			data[idx*numObs + j] = tmpDist.Rand()
		}
	}

	return data
}

func Kmeans(dataSet *[]float64, k int, seed int64, tol int) ([]int, []float32){

	// Initialization
	var targets = randomInitialize(len(*dataSet), k, seed)
	fmt.Printf("Random Init %v \n", targets)
	var means []float32
	var keepRunning = true
	var maxTry = tol
	var tries = 0

	for ; keepRunning && tries < maxTry;{
		// Compute classes means
		means = computeMeans(k, targets, *dataSet)
		fmt.Printf("tr means: %v \n", means)

		// Re-classify based on new means
		var newTargets = make([]int, len(*dataSet))

		for idx, elem := range *dataSet{
			var minK int
			var minDist float32 = 0

			for jdx, mean := range means{
				if minDist > distance(mean, float32(elem)) || jdx == 0{
					minK = jdx
					minDist = distance(mean, float32(elem))
				}
			}

			newTargets[idx] = minK
			//fmt.Printf("K = %v, Min distance = %v \n", minK, minDist)
			//fmt.Printf("Initial class %v,  New Class %v \n", targets[idx], newTargets[idx])
		}

		// Check if there is a difference
		keepRunning = false
		tries++
		fmt.Printf("OLD Init %v \n", targets)
		fmt.Printf("NEW Init %v \n", newTargets)
		for i, elem:= range newTargets{
			if elem != targets[i]{
				keepRunning = true
				targets = newTargets
				break
			}
		}
	}

	return targets, means
}

func computeMeans(k int, targets []int, dataSet []float64) []float32{
	var means = make([]float32, k)

	for class :=0; class < k; class++{
		var sum float64 = 0
		var classCounter = 0

		for j, elem:= range targets{
			if elem == class {
				sum += dataSet[j]
				classCounter ++
			}
		}
		means[class] = float32(sum / float64(classCounter))
	}

	return means
}

func distance(x1 float32, x2 float32) float32{
	return float32(math.Abs(float64(x1 - x2)))
}

func randomInitialize(dataSetLength int, numClasses int, seed int64) []int{
	// Set seed
	r := rand.New(rand.NewSource(seed))

	binsLen := float64(1.0 / float64(numClasses))
	var bins = make([]float64, dataSetLength)
	lastBin := 0.0

	for i:=0; i < numClasses; i++{
		lastBin += binsLen
		bins[i] = lastBin
	}

	// Random set classes
	var targets = make([]int, dataSetLength)

	for i:=0;i<dataSetLength;i++{
		var draw = r.Float64()
		for j, elem := range bins{
			if draw < elem{
				targets[i] = j
				break
			}
		}
	}

	return targets
}