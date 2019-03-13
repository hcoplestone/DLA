package main

import (
	"math/rand"
)

// RandomGenerator is a wrapper for a random number generator
type RandomGenerator struct {
	seed int64
	rand *rand.Rand
}

// NewRandomGenerator creates a random generator
func NewRandomGenerator(seed int64) *RandomGenerator {
	rg := new(RandomGenerator)
	rg.seed = seed
	rg.rand = rand.New(rand.NewSource(seed))
	return rg
}

// RandomInt generates a random integer between 0 and maxInt
func (rg RandomGenerator) RandomInt(maxInt int) int {
	return rg.rand.Intn(maxInt)
}
