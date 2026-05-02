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

const (
	sliceMinSize = 1  // minimum size for generated slices/arrays/maps
	sliceMaxSize = 10 // maximum size for generated slices/arrays/maps

	lowerCaseA = 97  // ASCII code for 'a'
	lowerCaseZ = 122 // ASCII code for 'z'
	upperCaseA = 65  // ASCII code for 'A'
	upperCaseZ = 90  // ASCII code for 'Z'
	asciiStart = 97  // start of printable ASCII characters
	asciiEnd   = 126 // end of printable ASCII characters

	stringLength = 10 // default length for generated strings
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

func getStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

func putStringBuilder(sb *strings.Builder) {
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

// RandomDigitNotZero generates a random digit (1-9) using the provided random number generator.
func RandomDigitNotZero(rng *rand.Rand) int {
	return (rng.Int() % 9) + 1
}

// RandomIntegerBetween generates a random integer between min and max (inclusive) using the provided random number generator.
func RandomIntegerBetween[T int | int32 | int64 | uint | uint32 | uint64](rng *rand.Rand, min, max T) T {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	switch any(min).(type) {
	case int:
		return T(rng.Int()%(int(max)-int(min)+1) + int(min))
	case int32:
		return T(rng.Int32()%(int32(max)-int32(min)+1) + int32(min))
	case int64:
		return T(rng.Int64()%(int64(max)-int64(min)+1) + int64(min))
	case uint:
		return T(rng.Uint()%(uint(max)-uint(min)+1) + uint(min))
	case uint32:
		return T(rng.Uint32()%(uint32(max)-uint32(min)+1) + uint32(min))
	case uint64:
		return T(rng.Uint64()%(uint64(max)-uint64(min)+1) + uint64(min))
	default:
		panic("unsupported type")
	}
}

// RandomIntegerByDecimals generates a random integer with the specified number of decimal digits using the provided random number generator.
func RandomIntegerByDecimals[T int | int32 | int64 | uint | uint32 | uint64](rng *rand.Rand, decimals int) T {
	if decimals <= 1 {
		return T(RandomDigit(rng))
	}

	minN := T(math.Pow10(decimals - 1))
	maxN := T(math.Pow10(decimals)) - 1

	return RandomIntegerBetween(rng, minN, maxN)
}

// RandomFloatBetween generates a random float between min and max with the specified number of decimal places using the provided random number generator.
// If decimalsExact is true, the result will be rounded to the specified 'decimals' decimal places. If false, the result may have up to 'decimals' decimal places.
func RandomFloatBetween(rng *rand.Rand, decimals int, decimalsExact bool, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
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

// RandomASCIICharacter generates a random printable ASCII character (from 97 to 126) using the provided random number generator.
func RandomASCIICharacter(rng *rand.Rand) string {
	return string(byte(rng.Int()%(asciiEnd-asciiStart+1) + asciiStart))
}

// RandomLetter generates a random lowercase letter (a-z) using the provided random number generator.
func RandomLowerCaseLetter(rng *rand.Rand) string {
	return string(byte(rng.Int()%(lowerCaseZ-lowerCaseA+1) + lowerCaseA))
}

// RandomUpperCaseLetter generates a random uppercase letter (A-Z) using the provided random number generator.
func RandomUpperCaseLetter(rng *rand.Rand) string {
	return string(byte(rng.Int()%(upperCaseZ-upperCaseA+1) + upperCaseA))
}

// RandomLowerCaseString generates a random string of the specified length consisting of lowercase letters (a-z) using the provided random number generator.
func RandomLowerCaseString(rng *rand.Rand, length int) string {
	if length <= 0 {
		return ""
	}

	// Cap length to prevent excessive memory usage
	if length > 1000 {
		length = 1000
	}

	sb := getStringBuilder()
	defer putStringBuilder(sb)

	sb.Grow(length)
	for range length {
		sb.WriteString(RandomLowerCaseLetter(rng))
	}

	return sb.String()
}

// RandomUpperCaseString generates a random string of the specified length consisting of uppercase letters (A-Z) using the provided random number generator.
func RandomUpperCaseString(rng *rand.Rand, length int) string {
	if length <= 0 {
		return ""
	}

	// Cap length to prevent excessive memory usage
	if length > 1000 {
		length = 1000
	}

	sb := getStringBuilder()
	defer putStringBuilder(sb)

	sb.Grow(length)
	for range length {
		sb.WriteString(RandomUpperCaseLetter(rng))
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

// RandomStringTemplate generates a random string filling in the provided template, where:
// - '#' is replaced with a random digit (0-9)
// - '%' is replaced with a random alphanumeric character (0-9, a-z)
// - '&' is replaced with a random alphanumeric character (0-9, A-Z)
// - '?' is replaced with a random lowercase letter (a-z)
// - '!' is replaced with a random uppercase letter (A-Z)
// - '*' is replaced with a random printable ASCII 97-126 character
// Any other characters in the template are left unchanged.
func RandomStringTemplate(rng *rand.Rand, template string) string {
	sb := getStringBuilder()
	defer putStringBuilder(sb)

	for _, ch := range template {
		switch ch {
		case '#':
			sb.WriteString(strconv.Itoa(RandomDigit(rng)))
		case '?':
			sb.WriteString(RandomLowerCaseLetter(rng))
		case '!':
			sb.WriteString(RandomUpperCaseLetter(rng))
		case '*':
			sb.WriteString(RandomASCIICharacter(rng))
		default:
			sb.WriteRune(ch)
		}
	}

	return sb.String()
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
