package knn

import (
	"fmt"
	"math/rand"
	"testing"
)


// Test K-means on gaussian generated values
func TestKmeans(t *testing.T){
	var gaussianMeans =  []int{50,30}

	var numObs = 100
	var sigmaRange = float64(5.0)
	var numClusters = len(gaussianMeans)
	var randomSeed = rand.Int63() //30
	var dataSet = GenerateGaussian(gaussianMeans, numObs, sigmaRange)

	var _, finalMeans  = Kmeans(&dataSet, numClusters, randomSeed, 50)

	for _,e := range finalMeans{
		fmt.Printf("means : %v \n", e)
	}

}
