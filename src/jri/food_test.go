package jri

import (
	"math"
	"testing"
)

const samplesPerItem = 10000
const errorThreshold = 1.0
const userId = 2

// Tests if weighted distribution works as expected
func TestDistribution(t *testing.T) {
	preset := BasedPreset
	totalSamples := samplesPerItem * len(preset)
	distr := map[string]int{}
	for i := 0; i < totalSamples; i++ {
		food, err := Jri(userId)
		if err != nil {
			t.Fatalf("%v", err)
		}

		distr[food]++
	}

	weightsSum := 0
	for _, food := range preset {
		weightsSum += food.Weight
	}

	for _, food := range preset {
		foodDistr, ok := distr[food.Name]
		if !ok && food.Weight > 0 {
			t.Errorf("Food %s not found in distr", food.Name)
			continue
		}

		expected := float64(food.Weight)
		actual := float64(foodDistr) / float64(totalSamples) * float64(weightsSum)

		if math.Abs(expected-actual) >= errorThreshold {
			t.Errorf("Food %s distribution is %.2f, expected %.2f", food.Name, actual, expected)
		}
	}
}
