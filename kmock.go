package kmock

import (
	"math/rand/v2"

	"github.com/lfsc09/kmock/internal/fake"
	"github.com/lfsc09/kmock/internal/randkit"
)

const (
	EN_US      = fake.EN_US
	PT_BR      = fake.PT_BR
	PWD_WEAK   = fake.PWD_WEAK
	PWD_MEDIUM = fake.PWD_MEDIUM
	PWD_STRONG = fake.PWD_STRONG
)

type KMock struct {
	seeds    []uint64
	rng      *rand.Rand
	Address  *fake.Address
	Boolean  *fake.Boolean
	Car      *fake.Car
	Company  *fake.Company
	Currency *fake.Currency
	Date     *fake.Date
	File     *fake.File
	Finance  *fake.Finance
	ID       *fake.ID
	Internet *fake.Internet
	Lorem    *fake.Lorem
	Number   *fake.Number
	Person   *fake.Person
	Regex    *fake.Regex
}

// New creates a new KMock instance with a unique random seed to ensure different random sequences across multiple instances and runs.
func New() *KMock {
	seeds := []uint64{randkit.RandomSeed(), randkit.RandomSeed()}
	rng := rand.New(rand.NewPCG(seeds[0], seeds[1]))
	return &KMock{
		seeds:    seeds,
		rng:      rng,
		Address:  &fake.Address{Rng: rng},
		Boolean:  &fake.Boolean{Rng: rng},
		Car:      &fake.Car{Rng: rng},
		Company:  &fake.Company{Rng: rng},
		Currency: &fake.Currency{Rng: rng},
		Date:     &fake.Date{Rng: rng},
		File:     &fake.File{Rng: rng},
		Finance:  &fake.Finance{Rng: rng},
		ID:       &fake.ID{},
		Internet: &fake.Internet{Rng: rng},
		Lorem:    &fake.Lorem{Rng: rng},
		Number:   &fake.Number{Rng: rng},
		Person:   &fake.Person{Rng: rng},
		Regex:    &fake.Regex{Rng: rng},
	}
}
