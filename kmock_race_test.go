package kmock

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MockKmockRaceTestSuite struct {
	suite.Suite
}

func TestMockKmockRaceTestSuite(t *testing.T) {
	suite.Run(t, new(MockKmockRaceTestSuite))
}

func (suite *MockKmockRaceTestSuite) TestKMockConcurrentAccess() {
	const numGoroutines = 8

	calls := []struct {
		fn     string
		params []any
	}{
		{fn: "Address.City", params: []any{EN_US}},
		{fn: "Car.Brand", params: []any{}},
		{fn: "Company.Name", params: []any{EN_US}},
		{fn: "Currency.Code", params: []any{}},
		{fn: "Date.Date", params: []any{"2000-01-01", "2020-12-31", "YYYY/MM/DD"}},
		{fn: "Internet.MacAddress", params: []any{}},
		{fn: "Lorem.Sentence", params: []any{0}},
		{fn: "Number.IntBetween", params: []any{1, 100}},
		{fn: "Person.Name", params: []any{EN_US}},
		{fn: "Regex.Generate", params: []any{"[A-Z]{5}\\d{4}"}},
	}

	var wg sync.WaitGroup
	for range numGoroutines {
		wg.Go(func() {
			km := New()
			for _, call := range calls {
				switch call.fn {
				case "Address.City":
					_, _ = km.Address.City(call.params[0].(string))
				case "Car.Brand":
					_ = km.Car.Brand()
				case "Company.Name":
					_, _ = km.Company.Name(call.params[0].(string))
				case "Currency.Code":
					_ = km.Currency.Code()
				case "Date.Date":
					_ = km.Date.Date(call.params[0].(string), call.params[1].(string), call.params[2].(string))
				case "Internet.MacAddress":
					_ = km.Internet.MacAddress()
				case "Lorem.Sentence":
					_ = km.Lorem.Sentence(call.params[0].(int))
				case "Number.IntBetween":
					_ = km.Number.IntBetween(call.params[0].(int), call.params[1].(int))
				case "Person.Name":
					_, _ = km.Person.Name(call.params[0].(string))
				case "Regex.Generate":
					_, _ = km.Regex.Generate(call.params[0].(string))
				}
			}
		})
	}
	wg.Wait()
}
