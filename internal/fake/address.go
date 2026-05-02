package fake

import (
	"kmock/internal/randkit"
	"math/rand/v2"
)

var (
	country = []string{
		"Afghanistan", "Albania", "Algeria", "Angola", "Argentina",
		"Australia", "Austria", "Bangladesh", "Belgium", "Bolivia",
		"Brazil", "Bulgaria", "Cambodia", "Canada", "Chile",
		"China", "Colombia", "Costa Rica", "Croatia", "Cuba",
		"Czech Republic", "Denmark", "Dominican Republic", "Ecuador", "Egypt",
		"El Salvador", "Ethiopia", "Finland", "France", "Germany",
		"Ghana", "Greece", "Guatemala", "Honduras", "Hungary",
		"India", "Indonesia", "Iran", "Iraq", "Ireland",
		"Israel", "Italy", "Jamaica", "Japan", "Jordan",
		"Kazakhstan", "Kenya", "Kuwait", "Lebanon", "Libya",
		"Malaysia", "Mexico", "Morocco", "Mozambique", "Myanmar",
		"Nepal", "Netherlands", "New Zealand", "Nicaragua", "Nigeria",
		"Norway", "Pakistan", "Panama", "Paraguay", "Peru",
		"Philippines", "Poland", "Portugal", "Romania", "Russia",
		"Saudi Arabia", "Senegal", "Serbia", "Singapore", "Slovakia",
		"South Africa", "South Korea", "Spain", "Sri Lanka", "Sudan",
		"Sweden", "Switzerland", "Syria", "Taiwan", "Tanzania",
		"Thailand", "Tunisia", "Turkey", "Uganda", "Ukraine",
		"United Arab Emirates", "United Kingdom", "United States", "Uruguay", "Uzbekistan",
		"Venezuela", "Vietnam", "Yemen", "Zambia", "Zimbabwe",
	}

	countryCode = []string{
		"AF", "AL", "DZ", "AO", "AR",
		"AU", "AT", "BD", "BE", "BO",
		"BR", "BG", "KH", "CA", "CL",
		"CN", "CO", "CR", "HR", "CU",
		"CZ", "DK", "DO", "EC", "EG",
		"SV", "ET", "FI", "FR", "DE",
		"GH", "GR", "GT", "HN", "HU",
		"IN", "ID", "IR", "IQ", "IE",
		"IL", "IT", "JM", "JP", "JO",
		"KZ", "KE", "KW", "LB", "LY",
		"MY", "MX", "MA", "MZ", "MM",
		"NP", "NL", "NZ", "NI", "NG",
		"NO", "PK", "PA", "PY", "PE",
		"PH", "PL", "PT", "RO", "RU",
		"SA", "SN", "RS", "SG", "SK",
		"ZA", "KR", "ES", "LK", "SD",
		"SE", "CH", "SY", "TW", "TZ",
		"TH", "TN", "TR", "UG", "UA",
		"AE", "GB", "US", "UY", "UZ",
		"VE", "VN", "YE", "ZM", "ZW",
	}

	state = map[string][]string{
		"en-US": {
			"Alabama", "Alaska", "Arizona", "Arkansas", "California",
			"Colorado", "Connecticut", "Delaware", "Florida", "Georgia",
			"Hawaii", "Idaho", "Illinois", "Indiana", "Iowa",
			"Kansas", "Kentucky", "Louisiana", "Maine", "Maryland",
		},
		"pt-BR": {
			"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
			"Ceará", "Distrito Federal", "Espírito Santo", "Goiás", "Maranhão",
			"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
			"Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		},
	}

	stateCode = map[string][]string{
		"en-US": {
			"AL", "AK", "AZ", "AR", "CA",
			"CO", "CT", "DE", "FL", "GA",
			"HI", "ID", "IL", "IN", "IA",
			"KS", "KY", "LA", "ME", "MD",
		},
		"pt-BR": {
			"AC", "AL", "AP", "AM", "BA",
			"CE", "DF", "ES", "GO", "MA",
			"MT", "MS", "MG", "PA", "PB",
			"PR", "PE", "PI", "RJ", "RN",
		},
	}

	city = map[string]map[string][]string{
		"en-US": {
			"Alabama":     {"Birmingham", "Montgomery", "Huntsville", "Mobile", "Tuscaloosa", "Hoover", "Dothan", "Auburn", "Decatur", "Madison"},
			"Alaska":      {"Anchorage", "Fairbanks", "Juneau", "Sitka", "Ketchikan", "Wasilla", "Kenai", "Kodiak", "Bethel", "Palmer"},
			"Arizona":     {"Phoenix", "Tucson", "Mesa", "Chandler", "Scottsdale", "Glendale", "Gilbert", "Tempe", "Peoria", "Surprise"},
			"Arkansas":    {"Little Rock", "Fort Smith", "Fayetteville", "Springdale", "Jonesboro", "North Little Rock", "Conway", "Rogers", "Bentonville", "Pine Bluff"},
			"California":  {"Los Angeles", "San Diego", "San Jose", "San Francisco", "Fresno", "Sacramento", "Long Beach", "Oakland", "Bakersfield", "Anaheim"},
			"Colorado":    {"Denver", "Colorado Springs", "Aurora", "Fort Collins", "Lakewood", "Thornton", "Arvada", "Westminster", "Pueblo", "Boulder"},
			"Connecticut": {"Bridgeport", "New Haven", "Hartford", "Stamford", "Waterbury", "Norwalk", "Danbury", "New Britain", "West Hartford", "Greenwich"},
			"Delaware":    {"Wilmington", "Dover", "Newark", "Middletown", "Smyrna", "Milford", "Seaford", "Georgetown", "Elsmere", "New Castle"},
			"Florida":     {"Jacksonville", "Miami", "Tampa", "Orlando", "St. Petersburg", "Hialeah", "Port St. Lucie", "Cape Coral", "Tallahassee", "Fort Lauderdale"},
			"Georgia":     {"Atlanta", "Augusta", "Columbus", "Macon", "Savannah", "Athens", "Sandy Springs", "South Fulton", "Roswell", "Albany"},
			"Hawaii":      {"Honolulu", "East Honolulu", "Pearl City", "Hilo", "Kailua", "Waipahu", "Kaneohe", "Mililani Town", "Kahului", "Ewa Gentry"},
			"Idaho":       {"Boise", "Nampa", "Meridian", "Idaho Falls", "Pocatello", "Caldwell", "Coeur d'Alene", "Twin Falls", "Lewiston", "Post Falls"},
			"Illinois":    {"Chicago", "Aurora", "Naperville", "Joliet", "Rockford", "Springfield", "Elgin", "Peoria", "Champaign", "Waukegan"},
			"Indiana":     {"Indianapolis", "Fort Wayne", "Evansville", "South Bend", "Carmel", "Fishers", "Bloomington", "Hammond", "Gary", "Muncie"},
			"Iowa":        {"Des Moines", "Cedar Rapids", "Davenport", "Sioux City", "Iowa City", "Waterloo", "Ames", "West Des Moines", "Ankeny", "Dubuque"},
			"Kansas":      {"Wichita", "Overland Park", "Kansas City", "Olathe", "Topeka", "Lawrence", "Shawnee", "Manhattan", "Lenexa", "Salina"},
			"Kentucky":    {"Louisville", "Lexington", "Bowling Green", "Owensboro", "Covington", "Hopkinsville", "Richmond", "Florence", "Georgetown", "Henderson"},
			"Louisiana":   {"New Orleans", "Baton Rouge", "Shreveport", "Metairie", "Lafayette", "Lake Charles", "Kenner", "Bossier City", "Monroe", "Alexandria"},
			"Maine":       {"Portland", "Lewiston", "Bangor", "South Portland", "Auburn", "Biddeford", "Sanford", "Augusta", "Saco", "Westbrook"},
			"Maryland":    {"Baltimore", "Frederick", "Rockville", "Gaithersburg", "Bowie", "Hagerstown", "Annapolis", "College Park", "Salisbury", "Laurel"},
		},
		"pt-BR": {
			"Acre":                {"Rio Branco", "Cruzeiro do Sul", "Sena Madureira", "Tarauacá", "Feijó", "Brasiléia", "Epitaciolândia", "Juruá", "Mâncio Lima", "Plácido de Castro"},
			"Alagoas":             {"Maceió", "Arapiraca", "Palmeira dos Índios", "Rio Largo", "Penedo", "União dos Palmares", "São Miguel dos Campos", "Coruripe", "Delmiro Gouveia", "Marechal Deodoro"},
			"Amapá":               {"Macapá", "Santana", "Laranjal do Jari", "Oiapoque", "Mazagão", "Porto Grande", "Tartarugalzinho", "Pedra Branca do Amapari", "Calçoene", "Ferreira Gomes"},
			"Amazonas":            {"Manaus", "Parintins", "Itacoatiara", "Manacapuru", "Coari", "Tefé", "Tabatinga", "Maués", "Humaitá", "São Gabriel da Cachoeira"},
			"Bahia":               {"Salvador", "Feira de Santana", "Vitória da Conquista", "Camaçari", "Juazeiro", "Itabuna", "Lauro de Freitas", "Ilhéus", "Jequié", "Alagoinhas"},
			"Ceará":               {"Fortaleza", "Caucaia", "Juazeiro do Norte", "Maracanaú", "Sobral", "Crato", "Itapipoca", "Maranguape", "Iguatu", "Quixadá"},
			"Distrito Federal":    {"Brasília", "Ceilândia", "Taguatinga", "Samambaia", "Planaltina", "Águas Claras", "Recanto das Emas", "Gama", "Guará", "Sobradinho"},
			"Espírito Santo":      {"Vitória", "Vila Velha", "Cariacica", "Serra", "Linhares", "Cachoeiro de Itapemirim", "Colatina", "Guarapari", "São Mateus", "Aracruz"},
			"Goiás":               {"Goiânia", "Aparecida de Goiânia", "Anápolis", "Rio Verde", "Luziânia", "Águas Lindas de Goiás", "Valparaíso de Goiás", "Trindade", "Formosa", "Novo Gama"},
			"Maranhão":            {"São Luís", "Imperatriz", "São José de Ribamar", "Timon", "Caxias", "Codó", "Paço do Lumiar", "Açailândia", "Bacabal", "Balsas"},
			"Mato Grosso":         {"Cuiabá", "Várzea Grande", "Rondonópolis", "Sinop", "Tangará da Serra", "Cáceres", "Sorriso", "Lucas do Rio Verde", "Primavera do Leste", "Barra do Garças"},
			"Mato Grosso do Sul":  {"Campo Grande", "Dourados", "Três Lagoas", "Corumbá", "Ponta Porã", "Naviraí", "Nova Andradina", "Aquidauana", "Sidrolândia", "Maracaju"},
			"Minas Gerais":        {"Belo Horizonte", "Uberlândia", "Contagem", "Juiz de Fora", "Betim", "Montes Claros", "Ribeirão das Neves", "Uberaba", "Governador Valadares", "Ipatinga"},
			"Pará":                {"Belém", "Ananindeua", "Santarém", "Marabá", "Castanhal", "Parauapebas", "Altamira", "Itaituba", "Abaetetuba", "Tucuruí"},
			"Paraíba":             {"João Pessoa", "Campina Grande", "Santa Rita", "Patos", "Bayeux", "Sousa", "Cajazeiras", "Cabedelo", "Guarabira", "Sapé"},
			"Paraná":              {"Curitiba", "Londrina", "Maringá", "Ponta Grossa", "Cascavel", "São José dos Pinhais", "Foz do Iguaçu", "Colombo", "Guarapuava", "Paranaguá"},
			"Pernambuco":          {"Recife", "Caruaru", "Petrolina", "Olinda", "Paulista", "Jaboatão dos Guararapes", "Garanhuns", "Cabo de Santo Agostinho", "Camaragibe", "Vitória de Santo Antão"},
			"Piauí":               {"Teresina", "Parnaíba", "Picos", "Piripiri", "Floriano", "Campo Maior", "Barras", "União", "Altos", "José de Freitas"},
			"Rio de Janeiro":      {"Rio de Janeiro", "São Gonçalo", "Duque de Caxias", "Nova Iguaçu", "Niterói", "Belford Roxo", "São João de Meriti", "Campos dos Goytacazes", "Petrópolis", "Volta Redonda"},
			"Rio Grande do Norte": {"Natal", "Mossoró", "Parnamirim", "São Gonçalo do Amarante", "Ceará-Mirim", "Caicó", "Açu", "Currais Novos", "Santa Cruz", "Macaíba"},
		},
	}

	neighborhood = map[string][]string{
		"en-US": {
			"Downtown", "Midtown", "Uptown", "East Side", "West Side",
			"North End", "South End", "Chinatown", "Little Italy", "Financial District",
			"Arts District", "Historic District", "Old Town", "Riverside", "Lakeside",
			"Hillcrest", "Sunset", "Parkview", "Garden District", "University District",
			"Capitol Hill", "Beacon Hill", "Back Bay", "Greenville", "Maplewood",
			"Cedarwood", "Oakwood", "Elmwood", "Birchwood", "Pinehurst",
			"Rosewood", "Brookside", "Creekside", "Meadowbrook", "Willowbrook",
			"Springdale", "Fairview", "Grandview", "Highview", "Clearview",
			"Westview", "Eastview", "Northview", "Southgate", "Northgate",
			"Westgate", "Eastgate", "Millbrook", "Stonegate", "Harborview",
		},
		"pt-BR": {
			"Centro", "Bairro Alto", "Vila Nova", "Jardim América", "Boa Vista",
			"Pinheiros", "Lapa", "Consolação", "Liberdade", "Bela Vista",
			"Campos Elíseos", "Santa Cecília", "Bom Retiro", "Brás", "Mooca",
			"Belenzinho", "Tatuapé", "Vila Prudente", "Ipiranga", "Saúde",
			"Vila Mariana", "Moema", "Ibirapuera", "Brooklin", "Santo Amaro",
			"Jabaquara", "Cursino", "Vila Olímpia", "Itaim Bibi", "Butantã",
			"Perdizes", "Pompeia", "Sumaré", "Higienópolis", "Pacaembu",
			"Barra Funda", "Santana", "Tucuruvi", "Vila Guilherme", "Penha",
			"Ermelino Matarazzo", "São Miguel", "Itaquera", "Guaianazes", "Cidade Tiradentes",
			"Parelheiros", "Grajaú", "Campo Limpo", "Capão Redondo", "Jardim Ângela",
		},
	}

	streetName = map[string][]string{
		"en-US": {
			"Main Street", "Oak Avenue", "Maple Street", "Cedar Lane", "Pine Street",
			"Elm Street", "Washington Boulevard", "Park Avenue", "Sunset Drive", "Riverside Road",
			"Highland Avenue", "Lakeview Drive", "Forest Road", "Hill Street", "Valley Road",
			"Spring Street", "Church Street", "Mill Road", "River Road", "Lake Drive",
			"Broadway", "Jefferson Avenue", "Lincoln Way", "Adams Street", "Monroe Drive",
			"Harrison Road", "Jackson Boulevard", "Madison Lane", "Franklin Street", "Grant Avenue",
			"Willow Court", "Birchwood Lane", "Rosewood Drive", "Stonegate Road", "Harborview Way",
			"Meadow Brook Lane", "Creekside Drive", "Hillcrest Avenue", "Fairview Court", "Grandview Terrace",
			"Northgate Boulevard", "Southgate Drive", "Westview Lane", "Eastview Road", "Clearwater Drive",
			"Pinehurst Circle", "Oakwood Trail", "Elmwood Place", "Maplewood Court", "Cedarwood Way",
		},
		"pt-BR": {
			"Rua das Flores", "Avenida Brasil", "Rua São João", "Avenida Paulista", "Rua da Paz",
			"Rua XV de Novembro", "Avenida Rio Branco", "Rua Tiradentes", "Rua da Liberdade", "Avenida Atlântica",
			"Rua das Palmeiras", "Avenida das Américas", "Rua do Comércio", "Avenida Central", "Rua Nova",
			"Rua Santa Clara", "Avenida Independência", "Rua dos Pinheiros", "Rua Sete de Setembro", "Avenida Norte",
			"Rua Marechal Deodoro", "Avenida Getúlio Vargas", "Rua Visconde de Mauá", "Rua Benjamin Constant", "Avenida Dom Pedro II",
			"Rua Duque de Caxias", "Avenida Presidente Vargas", "Rua General Osório", "Rua Barão do Rio Branco", "Avenida Ipiranga",
			"Rua das Acácias", "Rua dos Jacarandás", "Rua das Mangueiras", "Rua das Goiabeiras", "Avenida das Pitangueiras",
			"Rua dos Ipês", "Rua das Aroeiras", "Rua das Perobas", "Rua dos Cedros", "Avenida dos Eucaliptos",
			"Travessa das Rosas", "Alameda Santos", "Alameda Campinas", "Alameda Jaú", "Rua Augusta",
			"Rua Oscar Freire", "Rua Haddock Lobo", "Avenida Rebouças", "Rua da Consolação", "Avenida Brigadeiro Faria Lima",
		},
	}
)

type Address struct{}

// Country generates a random country name using the provided random number generator.
func (a Address) Country(rng *rand.Rand) string {
	return randkit.PickFromList(rng, country)
}

// CountryCode generates a random country code using the provided random number generator.
func (a Address) CountryCode(rng *rand.Rand) string {
	return randkit.PickFromList(rng, countryCode)
}

// State generates a random state name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) State(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, state[locale])
}

// StateCode generates a random state code for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StateCode(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, stateCode[locale])
}

// City generates a random city name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) City(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	randomState := randkit.PickFromList(rng, state[locale])
	return randkit.PickFromList(rng, city[locale][randomState])
}

// CityFromState generates a random city name for the specified locale and state using the provided random number generator.
// It panics if the locale or state is not supported.
func (a Address) CityFromState(rng *rand.Rand, locale, stateName string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	if _, ok := city[locale][stateName]; !ok {
		panic("state not supported for locale: " + stateName)
	}
	return randkit.PickFromList(rng, city[locale][stateName])
}

// Neighborhood generates a random neighborhood name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) Neighborhood(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, neighborhood[locale])
}

// StreetName generates a random street name for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetName(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	return randkit.PickFromList(rng, streetName[locale])
}

// StreetNumber generates a random street number using the provided random number generator, following common formatting patterns.
func (a Address) StreetNumber(rng *rand.Rand) string {
	templates := []string{"#####", "####", "###", "##"}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates))
}

// StreetComplement generates a random street complement for the specified locale using the provided random number generator.
// It panics if the locale is not supported.
func (a Address) StreetComplement(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"Apt. ###", "Suite ###", "Floor #", "Unit ###", "Building #"},
		"pt-BR": {"Apto. ###", "Sala ###", "Andar #", "Unidade ###", "Bloco #"},
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates[locale]))
}

// ZipCode generates a random zip code for the specified locale using the provided random number generator, following common formatting patterns.
// It panics if the locale is not supported.
func (a Address) ZipCode(rng *rand.Rand, locale string) string {
	if _, ok := availableLocales[locale]; !ok {
		panic("locale not supported: " + locale)
	}
	templates := map[string][]string{
		"en-US": {"#####", "#####-####"},
		"pt-BR": {"#####-###"},
	}
	return randkit.RandomStringTemplate(rng, randkit.PickFromList(rng, templates[locale]))
}

// Latitude generates a random latitude value between -90 and 90 degrees using the provided random number generator.
func (a Address) Latitude(rng *rand.Rand) float64 {
	return rng.Float64()*180 - 90
}

// Longitude generates a random longitude value between -180 and 180 degrees using the provided random number generator.
func (a Address) Longitude(rng *rand.Rand) float64 {
	return rng.Float64()*360 - 180
}

// RuntimeDocs provides runtime documentation for the Address struct and its methods
func (a Address) RuntimeDocs() *RunTimeDocs {
	return &RunTimeDocs{
		Struct: "Address",
		Methods: map[string]RunTimeDocsMethod{
			"Country": {
				Name:        "Country",
				Description: "Generates a random country name",
				Params:      []string{},
			},
			"CountryCode": {
				Name:        "CountryCode",
				Description: "Generates a random country code",
				Params:      []string{},
			},
			"State": {
				Name:        "State",
				Description: "Generates a random state name for the specified locale",
				Params:      []string{"locale"},
			},
			"StateCode": {
				Name:        "StateCode",
				Description: "Generates a random state code for the specified locale",
				Params:      []string{"locale"},
			},
			"City": {
				Name:        "City",
				Description: "Generates a random city name for the specified locale",
				Params:      []string{"locale"},
			},
			"CityFromState": {
				Name:        "CityFromState",
				Description: "Generates a random city name for the specified locale and state",
				Params:      []string{"locale", "state"},
			},
			"Neighborhood": {
				Name:        "Neighborhood",
				Description: "Generates a random neighborhood name for the specified locale",
				Params:      []string{"locale"},
			},
			"StreetName": {
				Name:        "StreetName",
				Description: "Generates a random street name for the specified locale",
				Params:      []string{"locale"},
			},
			"StreetNumber": {
				Name:        "StreetNumber",
				Description: "Generates a random street number",
				Params:      []string{},
			},
			"StreetComplement": {
				Name:        "StreetComplement",
				Description: "Generates a random street complement for the specified locale",
				Params:      []string{"locale"},
			},
			"ZipCode": {
				Name:        "ZipCode",
				Description: "Generates a random zip code for the specified locale",
				Params:      []string{"locale"},
			},
			"Latitude": {
				Name:        "Latitude",
				Description: "Generates a random latitude value between -90 and 90 degrees",
				Params:      []string{},
			},
			"Longitude": {
				Name:        "Longitude",
				Description: "Generates a random longitude value between -180 and 180 degrees",
				Params:      []string{},
			},
		},
	}
}
