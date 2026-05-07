# The project

Kmock is a simple Go library to facilitate the creation of fake data. Heavly inspired in [faker/v2](https://github.com/jaswdr/faker).

It provides several functions to generate data for:

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
| `Company` | `IE` | |
| `Company` | `CNAE` | |
| `Currency` | `Name` | |
| `Currency` | `Code` | |
| `Currency` | `Symbol` | |
| `Currency` | `Full` | |
| `Date` | `Date` | `[from, to, format]` |
| `Date` | `Time` | `[from, to, format]` |
| `Date` | `DateTime` | `[from, to, format]` |
| `Date` | `Now` | `[format]` |
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
| `Regex` | `Generate` | `[pattern]` |
| `Regex` | `ClearCache` | |

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

Use `kmock.New()` to instanciate the faker generator, and access the domain generator types to call their methods.

Each `kmock` instance is created with a unique set of random seeds to ensure different random sequences across multiple instances.

```go
import "github.com/lfsc09/kmock"

func main() {
  k := kmock.New() // Returns *KMock instance

  k.Address.Country()

  k.Person.Name("en-US")
}
```

| Available locales |
| :-: |
| `en-US` |
| `pt-BR` |

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

### Maintain

Updating dependencies

```bash
# Download updates for all dependencies to their latest minor/patch versions
go get -u ./...

# Tidy: remove unused deps and add any missing ones
go mod tidy
```
