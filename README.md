# The project

Kmock is a Go library that generates random, structurally valid fake data on demand. Each call produces a different value, and each kmock.New() instance is seeded independently, so parallel test runs and multiple instances never produce the same sequence.

Inspired by [faker/v2](https://github.com/jaswdr/faker).

It covers a wide range of domains, with locale support for region-specific data.

### Advantages

- **Locale-aware generation** — the same API produces data for en-US or pt-BR without any extra setup.
- **Reproducible randomness** — seeds are exposed on each instance, so a failing test can record its seed and be replayed exactly.
- **Single dependency** — no external data files, no network calls, no configuration; just import and call.
- **Structurally valid output** — credit card numbers pass Luhn, CPF passes checksum, UUIDs conform to v4/v6/v7 spec, etc.
- Simple, flat API with no global state.
- Each instance is independently seeded, safe to use in parallel tests.
- [Regex-based generation](#randexp-package) via `k.Regex.Generate(pattern)` for custom formats.

### Disadvantages

- Data variety is limited to the built-in datasets — adding new locales or domains requires a code change.
- Not suitable for cryptographic or security-sensitive use cases — randomness comes from `math/rand/v2`, not `crypto/rand`.

### Domains

| Domain | Function | Params |
| --- | --: | --: |
| `Address` | `Country` | |
| `Address` | `CountryCode` | |
| `Address` | `State` | `[locale]` |
| `Address` | `StateCode` | `[locale]` |
| `Address` | `City` | `[locale]` |
| `Address` | `CityFromState` | `[locale, stateCode]` |
| `Address` | `Neighborhood` | `[locale]` |
| `Address` | `StreetName` | `[locale]` |
| `Address` | `StreetNumber` | |
| `Address` | `StreetComplement` | `[locale]` |
| `Address` | `ZipCode` | `[locale]` |
| `Address` | `Latitude` | |
| `Address` | `Longitude` | |
| `Boolean` | `Random` | |
| `Boolean` | `RandomWithProbability` | `[probability]` |
| `Car` | `Brand` | |
| `Car` | `Model` | |
| `Car` | `LicensePlate` | `[locale]` |
| `Car` | `Color` | |
| `Company` | `Name` | `[locale]` |
| `Company` | `Dba` | `[locale]` |
| `Company` | `Industry` | `[locale]` |
| `Company` | `Suffix` | `[locale]` |
| `Company` | `EIN` | |
| `Company` | `CNPJLegacyValid` | |
| `Company` | `CNPJLegacyInvalid` | |
| `Company` | `CNPJAlphanumericValid` | |
| `Company` | `CNPJAlphanumericInvalid` | |
| `Company` | `IE` | |
| `Company` | `CNAE` | |
| `Currency` | `Name` | |
| `Currency` | `Code` | |
| `Currency` | `Symbol` | |
| `Currency` | `Full` | |
| [`Date`](#date-generator) | `Date` | `[from, to, format]` |
| [`Date`](#date-generator) | `Time` | `[from, to, format]` |
| [`Date`](#date-generator) | `DateTime` | `[from, to, format]` |
| [`Date`](#date-generator) | `Now` | `[format]` |
| `File` | `Name` | |
| `File` | `Extension` | |
| `File` | `NameWithExtension` | |
| `File` | `MimeType` | |
| `Finance` | `CreditCardVendor` | |
| `Finance` | `CreditCardNumber` | |
| `Finance` | `CreditCardCVV` | |
| `Finance` | `CreditCardExpirationDate` | |
| `Finance` | `CreditCardHolder` | `[locale]` |
| `ID` | `UUIDv4` | |
| `ID` | `UUIDv6` | |
| `ID` | `UUIDv7` | |
| `ID` | `SequentialID` | `[startFrom]` |
| `Internet` | `Domain` | |
| `Internet` | `DomainWithSubdomain` | |
| `Internet` | `Ipv4` | |
| `Internet` | `Ipv6` | |
| `Internet` | `MacAddress` | |
| `Internet` | `Url` | |
| `Internet` | `UserAgent` | |
| `Lorem` | `Word` | |
| `Lorem` | `Sentence` | `[wordCount]` |
| `Lorem` | `Paragraph` | `[sentenceCount]` |
| `Number` | `Int` | |
| `Number` | `IntBetween` | `[min, max]` |
| `Number` | `Float` | |
| `Number` | `FloatBetween` | `[decimals, min, max]` |
| `Person` | `Name` | `[locale]` |
| `Person` | `FirstName` | `[locale]` |
| `Person` | `MiddleName` | `[locale]` |
| `Person` | `LastName` | `[locale]` |
| `Person` | `Phone` | |
| `Person` | `Email` | |
| `Person` | `Username` | |
| `Person` | `Password` | `[strength]` |
| `Person` | `JobTitle` | `[locale]` |
| `Person` | `CPFValid` | |
| `Person` | `CPFInvalid` | |
| [`Regex`](#randexp-package) | [`Generate`](#regex-generator) | `[pattern]` |
| [`Regex`](#randexp-package) | `ClearCache` | |

### Date generator

The `Date` domain generates random dates, times, and datetimes within a given range. All methods accept optional `from`, `to`, and `format` parameters — pass an empty string `""` to use the defaults.

#### `from` to `to` expected formats

| Method | Param | Format | Default |
| -- | -- | -- | -- |
| `Date` | `from` | `YYYY-MM-DD` | `1970-01-01` |
| `Date` | `to` | `YYYY-MM-DD` | current date |
| `Time` | `from` | `hh:mm:ss` | `00:00:00` |
| `Time` | `to` | `hh:mm:ss` | `23:59:59` |
| `DateTime` | `from` | `YYYY-MM-DD hh:mm:ss` | `1970-01-01 00:00:00` |
| `DateTime` | `to` | `YYYY-MM-DD hh:mm:ss` | current date |

#### `format` tokens

| Token | Meaning | Example |
| --- | --: | --: |
| `YYYY` | 4-digit year | `2024` |
| `MM` | 2-digit month | `01` |
| `DD` | 2-digit day | `31` |
| `hh` | Hour (24h) | `15` |
| `mm` | Minute | `04` |
| `ss` | Second | `05` |
| `sss` | Miliseconds | `000` |

#### Examples

```go
// Date
k.Date.Date("", "", "")                     // "2013-07-22"
k.Date.Date("2000-01-01", "2020-12-31", "") // "2008-03-15"
k.Date.Date("", "", "DD/MM/YYYY")           // "22/07/2013"

// Time
k.Date.Time("", "", "")                 // "14:32:07"
k.Date.Time("08:00:00", "18:00:00", "") // "11:47:52"
k.Date.Time("", "", "hh:mm")            // "14:32"

// DateTime
k.Date.DateTime("", "", "")                                       // "2013-07-22 14:32:07"
k.Date.DateTime("2000-01-01 00:00:00", "2020-12-31 23:59:59", "") // "2008-03-15 11:47:52"
k.Date.DateTime("", "", "DD/MM/YYYY hh:mm")                       // "22/07/2013 14:32"

// Now
k.Date.Now("")           // "2026-05-08 13:45:00"
k.Date.Now("DD/MM/YYYY") // "08/05/2026"
k.Date.Now("hh:mm")      // "13:45"
```

### Regex generator

The `Regex` domain generates random strings that match a given regular expression pattern. Internally it delegates to the `randexp` package (bellow) for pattern parsing and generation.

**Caching** — compiled pattern generators are cached on the `Regex` instance, on `Regex.Generator(pattern)` calls, so subsequent calls with the same pattern reuse the cached generator at no extra cost.

The cache is scoped to the `Regex` instance, so different `kmock.New()` instances have independent caches.

`Regex.ClearCache()` empties the cache, releasing all compiled generators for GC.

```go
k.Regex.Generate(`\d{5}`)    // compiled and cached
k.Regex.Generate(`[a-z]{8}`) // compiled and cached

k.Regex.ClearCache()         // both entries released

k.Regex.Generate(`\d{5}`)    // compiled again from scratch
```

### randexp package

The `randexp` package powers `k.Regex.Generate(pattern)`. It parses a regular expression using Go's [`regexp/syntax`](https://pkg.go.dev/regexp/syntax) package and walks the resulting AST to produce a random string that satisfies the pattern.

```go
k.Regex.Generate(`[A-Z]{2}\d{4}`)   // "BK7392"
k.Regex.Generate(`\d{3}-\d{2}-\d{4}`) // "482-90-1637"
k.Regex.Generate(`(foo|bar)-\d{2}`)  // "bar-17"
```

Supported constructs:

| Construct | Example | Notes |
| -- | -- | -- |
| Literals | `abc` | Case-folding (`(?i)`) supported |
| Character classes | `[a-z]`, `\d`, `\w`, `\s` | Negated classes (`[^0-9]`) supported |
| Quantifiers | `x?`, `x*`, `x+`, `x{n}`, `x{n,m}` | See limitations bellow |
| Alternation | `foo\|bar` | One branch picked uniformly at random |
| Grouping | `(abc)` | Capturing groups are transparent |
| Dot `.` | `.` | Restricted to printable ASCII (see limitations) |
| Anchors | `^`, `$`, `\b` | Silently ignored — produce no output |

#### Limitations

- **Unbounded quantifiers** (`*`, `+`, `{n,}`) are capped at 10 repetitions to keep output finite.
- **Dot** (`.`) is restricted to printable ASCII [0x20, 0x7E] — full Unicode semantics are not emulated.
- **Lookaheads and lookbehinds** (`(?=...)`, `(?!...)`, `(?<=...)`, `(?<!...)`) are not supported and will return an error.
- **Back-references** (`\1`, `\k<name>`) are not supported and will return an error.
- **Not cryptographically random** — uses `math/rand/v2`, not `crypto/rand`.

</br>

# Installation

```bash
go get github.com/lfsc09/kmock
```

```go
import "github.com/lfsc09/kmock"
```

</br>

# Basic Usage

Use `kmock.New()` to instantiate the faker generator. Each instance is created with a unique random seed, so multiple instances produce independent random sequences.

```go
import "github.com/lfsc09/kmock"

func main() {
  k := kmock.New() // *KMock instance

  fmt.Println(k.Address.Country())            // "Norway"
  fmt.Println(k.Car.Brand())                  // "Toyota"
  fmt.Println(k.Currency.Code())              // "USD"
  fmt.Println(k.Internet.Ipv4())              // "192.168.1.42"
  fmt.Println(k.Lorem.Sentence(0))            // "Lorem ipsum dolor sit amet."
  fmt.Println(k.Number.IntBetween(1, 100))    // 73
  fmt.Println(k.Regex.Generate("\\L-\\d{5}")) // ADR-16548
}
```

Some methods accept a `locale` to generate region-specific data:

| Locale | Constant |
| --- | --: |
| `en-US` | `kmock.EN_US` |
| `pt-BR` | `kmock.PT_BR` |

```go
name, err := k.Person.Name(kmock.EN_US)
if err != nil {
  // handle unsupported locale
}
fmt.Println(name) // "Adam Carter"

city, err := k.Address.City(kmock.PT_BR)
if err != nil {
  // handle unsupported locale
}
fmt.Println(city) // "São Paulo"
```

Multiple independent instances can be used concurrently, each has its own random state:

```go
k1 := kmock.New()
k2 := kmock.New()

// k1 and k2 produce different sequences
fmt.Println(k1.Person.Phone()) // "+1 555-0192"
fmt.Println(k2.Person.Phone()) // "+55 11 91234-5678"
```

</br>

# Development Details

### Installation

Clone the repository.

```bash
git clone git@github.com:lfsc09/kmock.git
cd kmock
```

Install dependencies.

```bash
go mod download
```

Configure git hooks for auto-bump version on commits.

```bash
make install-hooks
```

Run tests.

```bash
go test -race ./...
```
