package kmock

import (
	"kmock/internal/fake"
	"kmock/internal/randkit"
	"math/rand/v2"
)

type KMock struct {
	rng *rand.Rand
}

// New creates a new KMock instance with a unique random seed to ensure different random sequences across multiple instances and runs.
func New() *KMock {
	seeds := []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	return &KMock{
		rng: rand.New(rand.NewPCG(seeds[0], seeds[1])),
	}
}

func (km KMock) Address() *fake.Address {
	return &fake.Address{}
}

func (km KMock) Boolean() *fake.Boolean {
	return &fake.Boolean{}
}

func (km KMock) Car() *fake.Car {
	return &fake.Car{}
}
