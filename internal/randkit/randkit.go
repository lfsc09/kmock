package randkit

import (
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type weigthedChoice[T any] struct {
	value  T
	weight int
}

const (
	letterLowerCaseStart = 97  // ASCII code for 'a'
	letterLowerCaseEnd   = 122 // ASCII code for 'z'
	letterUpperCaseStart = 65  // ASCII code for 'A'
	letterUpperCaseEnd   = 90  // ASCII code for 'Z'
	asciiPrintableStart  = 32  // ASCII code for space character
	asciiPrintableEnd    = 126 // ASCII code for '~' character

	stringLength = 1000 // default length for generated strings
)

const (
	// fiboGoldenRatio is a large constant derived from the golden ratio, used to help ensure good distribution of seeds.
	fiboGoldenRatio = 0x9E3779B97F4A7C15

	// piMixingConstant is a large constant derived from the digits of pi, used to further mix the seed values for better randomness.
	piMixingConstant = 0x517cc1b727220a95
)

var seedCounter atomic.Uint64

// Pool of string builders to reduce allocations when generating random strings
var stringBuilderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

func GetStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

func PutStringBuilder(sb *strings.Builder) {
	sb.Reset()
	stringBuilderPool.Put(sb)
}

// RandomSeed creates a unique seed by combining the current time, process ID, and an atomic counter, to minimize the chances of collision.
func RandomSeed() uint64 {
	return uint64(time.Now().UnixNano()) ^ (uint64(os.Getpid()) * piMixingConstant) ^ (seedCounter.Add(1) * fiboGoldenRatio)
}

// RandomBool generates a random boolean value based on the provided probability of being true using the given random number generator.
// probTrue should be a value between 0.0 and 1.0, where 0.0 means always false and 1.0 means always true.
func RandomBool(rng *rand.Rand, probTrue float64) bool {
	if probTrue <= 0 {
		return false
	}
	if probTrue >= 1 {
		return true
	}
	return rng.Float64() < probTrue
}

// RandomDigit generates a random digit (0-9) using the provided random number generator.
func RandomDigit(rng *rand.Rand) int {
	return rng.Int() % 10
}

// RandomIntegerBetween generates a random integer between min and max (inclusive) using the provided random number generator.
// If min is greater than max, they will be swapped to ensure a valid range.
// If max is 0, it will be treated as the maximum value for the type T (e.g., math.MaxInt for int).
func RandomIntegerBetween[T int | int32 | int64 | uint | uint32 | uint64](rng *rand.Rand, min, max T) T {
	if min > max {
		min, max = max, min
	}

	switch any(T(0)).(type) {
	case int:
		if max == T(0) {
			max = T(any(math.MaxInt).(int))
		}
		return T(rng.Int()%(int(max)-int(min)+1) + int(min))
	case int32:
		if max == T(0) {
			max = T(any(int32(math.MaxInt32)).(int32))
		}
		return T(rng.Int32()%(int32(max)-int32(min)+1) + int32(min))
	case int64:
		if max == T(0) {
			max = T(any(int64(math.MaxInt64)).(int64))
		}
		return T(rng.Int64()%(int64(max)-int64(min)+1) + int64(min))
	case uint:
		if max == T(0) {
			max = T(any(uint(math.MaxUint)).(uint))
		}
		return T(rng.Uint()%(uint(max)-uint(min)+1) + uint(min))
	case uint32:
		if max == T(0) {
			max = T(any(uint32(math.MaxUint32)).(uint32))
		}
		return T(rng.Uint32()%(uint32(max)-uint32(min)+1) + uint32(min))
	case uint64:
		if max == T(0) {
			max = T(any(uint64(math.MaxUint64)).(uint64))
		}
		return T(rng.Uint64()%(uint64(max)-uint64(min)+1) + uint64(min))
	}
	return T(0)
}

// RandomInteger generates a random integer of the specified type using the provided random number generator.
func RandomInteger[T int | int32 | int64 | uint | uint32 | uint64](rng *rand.Rand) T {
	switch any(T(0)).(type) {
	case int:
		return T(rng.Int())
	case int32:
		return T(rng.Int32())
	case int64:
		return T(rng.Int64())
	case uint:
		return T(rng.Uint())
	case uint32:
		return T(rng.Uint32())
	case uint64:
		return T(rng.Uint64())
	}
	return T(0)
}

// RandomFloatBetween generates a random float between min and max (inclusive) with the specified number of decimal places using the provided random number generator.
// If min is greater than max, they will be swapped to ensure a valid range.
// If max is 0, it will be treated as math.MaxFloat64.
// If decimalsExact is true, the result will be rounded to the specified 'decimals' decimal places.
// If false, the result may have up to 'decimals' decimal places.
func RandomFloatBetween(rng *rand.Rand, decimals int, decimalsExact bool, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}

	if max == 0 {
		max = math.MaxFloat64
	}

	if decimals < 0 {
		decimals = 0
	}

	scale := math.Pow10(decimals)
	randomValue := min + rng.Float64()*(max-min)

	// Round to specific number of decimal places
	if decimalsExact {
		return math.Round(randomValue*scale) / scale
	}

	// Round up to the specified number of decimal places
	return math.Ceil(randomValue*scale) / scale
}

func RandomFloat(rng *rand.Rand) float64 {
	return rng.Float64()
}

// RandomLetter generates a random letter (A-Z or a-z) based on the upperCase parameter using the provided random number generator.
func RandomLetter(rng *rand.Rand, upperCase bool) string {
	if upperCase {
		return string(byte(rng.Int()%(letterUpperCaseEnd-letterUpperCaseStart+1) + letterUpperCaseStart))
	}
	return string(byte(rng.Int()%(letterLowerCaseEnd-letterLowerCaseStart+1) + letterLowerCaseStart))
}

// RandomAlphanumeric generates a random alphanumeric character (0-9, a-z, A-Z) based on the upperCase parameter using the provided random number generator.
func RandomAlphanumeric(rng *rand.Rand, upperCase bool) string {
	choices := []int{0, 1}
	picked := PickFromList(rng, choices)
	if picked == 0 {
		return strconv.Itoa(RandomDigit(rng))
	}
	return RandomLetter(rng, upperCase)
}

// RandomASCIICharacter generates a random printable ASCII character (from 32 to 126) using the provided random number generator.
func RandomASCIICharacter(rng *rand.Rand) string {
	return string(byte(rng.Int()%(asciiPrintableEnd-asciiPrintableStart+1) + asciiPrintableStart))
}

// RandomString generates a random string of the specified length and case using the provided random number generator.
// If upperCase is true, the string will contain only uppercase letters.
// The length is capped at 'stringLength' to prevent excessive memory usage.
func RandomString(rng *rand.Rand, length int, upperCase bool) string {
	if length <= 0 {
		return ""
	}

	// Cap length to prevent excessive memory usage
	if length > stringLength {
		length = stringLength
	}

	sb := GetStringBuilder()
	defer PutStringBuilder(sb)

	sb.Grow(length)
	for range length {
		sb.WriteString(RandomLetter(rng, upperCase))
	}

	return sb.String()
}

// RandomStringTemplate generates a random string filling in the provided template, where:
// - '\d' is replaced with a random digit (0-9)
// - '\a' is replaced with a random alphanumeric character (0-9, a-z)
// - '\A' is replaced with a random alphanumeric character (0-9, A-Z)
// - '\l' is replaced with a random lowercase letter (a-z)
// - '\L' is replaced with a random uppercase letter (A-Z)
// - '\.' is replaced with a random printable ASCII 32-126 character
// Any other characters in the template are left unchanged.
func RandomStringTemplate(rng *rand.Rand, template string) string {
	sb := GetStringBuilder()
	defer PutStringBuilder(sb)

	i := 0
	for i < len(template) {
		if template[i] == '\\' && i+1 < len(template) {
			switch template[i+1] {
			case 'd':
				sb.WriteString(strconv.Itoa(RandomDigit(rng)))
				i += 2
				continue
			case 'a':
				sb.WriteString(RandomAlphanumeric(rng, false))
				i += 2
				continue
			case 'A':
				sb.WriteString(RandomAlphanumeric(rng, true))
				i += 2
				continue
			case 'l':
				sb.WriteString(RandomLetter(rng, false))
				i += 2
				continue
			case 'L':
				sb.WriteString(RandomLetter(rng, true))
				i += 2
				continue
			case '.':
				sb.WriteString(RandomASCIICharacter(rng))
				i += 2
				continue
			}
		}
		sb.WriteByte(template[i])
		i++
	}

	return sb.String()
}

// PickFromList selects a random element from the provided list using the given random number generator. If the list is empty, it returns the zero value for the type T.
func PickFromList[T any](rng *rand.Rand, list []T) T {
	// Return zero value for the list type
	if len(list) == 0 {
		var zero T
		return zero
	}

	return list[rng.Int()%len(list)]
}

// PickFromWeightedList selects a random element from the provided list of weighted choices using the given random number generator. If the list is empty, it returns the zero value for the type T.
func PickFromWeightedList[T any](rng *rand.Rand, list []weigthedChoice[T]) T {
	// Return zero value for the list type
	if len(list) == 0 {
		var zero T
		return zero
	}

	totalWeight := 0
	for _, item := range list {
		totalWeight += item.weight
	}

	randomWeight := rng.Int() % totalWeight
	for _, item := range list {
		if randomWeight < item.weight {
			return item.value
		}
		randomWeight -= item.weight
	}

	// Fallback in case of an unexpected issue (should not happen)
	var zero T
	return zero
}

// PickFromMapKeys selects a random key from the provided map using the given random number generator. If the map is empty, it returns the zero value for the key type K.
func PickFromMapKeys[K comparable, V any](rng *rand.Rand, m map[K]V) K {
	// Return zero value for the key type
	if len(m) == 0 {
		var zero K
		return zero
	}

	n := rng.Int() % len(m)
	for k := range m {
		if n == 0 {
			return k
		}
		n--
	}

	// Fallback in case of an unexpected issue (should not happen)
	var zero K
	return zero
}

// PickFromMapValues selects a random value from the provided map using the given random number generator. If the map is empty, it returns the zero value for the value type V.
func PickFromMapValues[K comparable, V any](rng *rand.Rand, m map[K]V) V {
	// Return zero value for the value type
	if len(m) == 0 {
		var zero V
		return zero
	}

	n := rng.Int() % len(m)
	for _, v := range m {
		if n == 0 {
			return v
		}
		n--
	}

	// Fallback in case of an unexpected issue (should not happen)
	var zero V
	return zero
}

// RandomDateTime generates a random time between the provided 'from' and 'to' times using the provided random number generator.
func RandomDateTime(rng *rand.Rand, from time.Time, to time.Time) time.Time {
	if to.Before(from) {
		from, to = to, from
	}
	diff := to.Sub(from)
	randomDuration := time.Duration(RandomIntegerBetween(rng, 0, diff.Nanoseconds()))
	return from.Add(randomDuration)
}

// ShuffleString takes a string and returns a new string with the characters shuffled in random order using the provided random number generator.
func ShuffleString(rng *rand.Rand, s string) string {
	runes := []rune(s)

	// Fisher-Yates shuffle algorithm
	for i := len(runes) - 1; i > 0; i-- {
		j := rng.IntN(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
