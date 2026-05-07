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
	return &fake.Address{
		Rng: km.rng,
	}
}

func (km KMock) Boolean() *fake.Boolean {
	return &fake.Boolean{
		Rng: km.rng,
	}
}

func (km KMock) Car() *fake.Car {
	return &fake.Car{
		Rng: km.rng,
	}
}

func (km KMock) Company() *fake.Company {
	return &fake.Company{
		Rng: km.rng,
	}
}

func (km KMock) Currency() *fake.Currency {
	return &fake.Currency{
		Rng: km.rng,
	}
}

func (km KMock) Date() *fake.Date {
	return &fake.Date{
		Rng: km.rng,
	}
}

func (km KMock) File() *fake.File {
	return &fake.File{
		Rng: km.rng,
	}
}

func (km KMock) Finance() *fake.Finance {
	return &fake.Finance{
		Rng: km.rng,
	}
}

func (km KMock) ID() *fake.ID {
	return &fake.ID{}
}

func (km KMock) Internet() *fake.Internet {
	return &fake.Internet{
		Rng: km.rng,
	}
}

func (km KMock) Lorem() *fake.Lorem {
	return &fake.Lorem{
		Rng: km.rng,
	}
}

func (km KMock) Number() *fake.Number {
	return &fake.Number{
		Rng: km.rng,
	}
}

func (km KMock) Person() *fake.Person {
	return &fake.Person{
		Rng: km.rng,
	}
}

func (km KMock) Regex() *fake.Regex {
	return &fake.Regex{
		Rng: km.rng,
	}
}
