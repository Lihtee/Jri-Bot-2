package jri

import (
	"math"
	"testing"

	wr "github.com/mroth/weightedrand"
)

const samplesPerItem = 10000
const errorThreshold = 1.0
const userId = 2

// Tests if weighted distribution works as expected
func TestDistribution(t *testing.T) {
	storage := NewTestStorage()
	che := NewChe(storage)

	preset := make([]wr.Choice, len(BasedPreset))
	for i, food := range BasedPreset {
		preset[i] = food.toChoice()
	}
	totalSamples := samplesPerItem * len(preset)
	distr := map[string]int{}
	for i := 0; i < totalSamples; i++ {
		food, err := che.Sojrat(userId)
		if err != nil {
			t.Fatalf("%v", err)
		}

		distr[food]++
	}

	weightsSum := uint(0)
	for _, food := range preset {
		weightsSum += food.Weight
	}

	for _, food := range preset {
		foodName := food.Item.(string)
		foodDistr, ok := distr[foodName]
		if !ok && food.Weight > 0 {
			t.Errorf("Food %s not found in distr", foodName)
			continue
		}

		expected := float64(food.Weight)
		actual := float64(foodDistr) / float64(totalSamples) * float64(weightsSum)

		if math.Abs(expected-actual) >= errorThreshold {
			t.Errorf("Food %s distribution is %.2f, expected %.2f", foodName, actual, expected)
		}
	}
}
