package fake

var availableLocales = map[string]struct{}{
	"en-US": {},
	"pt-BR": {},
}

type countryData struct {
	name string
	code string
}

type stateData struct {
	name string
	code string
}

var (
	country = []countryData{
		{"Afghanistan", "AF"}, {"Albania", "AL"}, {"Algeria", "DZ"}, {"Angola", "AO"}, {"Argentina", "AR"},
		{"Australia", "AU"}, {"Austria", "AT"}, {"Bangladesh", "BD"}, {"Belgium", "BE"}, {"Bolivia", "BO"},
		{"Brazil", "BR"}, {"Bulgaria", "BG"}, {"Cambodia", "KH"}, {"Canada", "CA"}, {"Chile", "CL"},
		{"China", "CN"}, {"Colombia", "CO"}, {"Costa Rica", "CR"}, {"Croatia", "HR"}, {"Cuba", "CU"},
		{"Czech Republic", "CZ"}, {"Denmark", "DK"}, {"Dominican Republic", "DO"}, {"Ecuador", "EC"}, {"Egypt", "EG"},
		{"El Salvador", "SV"}, {"Ethiopia", "ET"}, {"Finland", "FI"}, {"France", "FR"}, {"Germany", "DE"},
		{"Ghana", "GH"}, {"Greece", "GR"}, {"Guatemala", "GT"}, {"Honduras", "HN"}, {"Hungary", "HU"},
		{"India", "IN"}, {"Indonesia", "ID"}, {"Iran", "IR"}, {"Iraq", "IQ"}, {"Ireland", "IE"},
		{"Israel", "IL"}, {"Italy", "IT"}, {"Jamaica", "JM"}, {"Japan", "JP"}, {"Jordan", "JO"},
		{"Kazakhstan", "KZ"}, {"Kenya", "KE"}, {"Kuwait", "KW"}, {"Lebanon", "LB"}, {"Libya", "LY"},
		{"Malaysia", "MY"}, {"Mexico", "MX"}, {"Morocco", "MA"}, {"Mozambique", "MZ"}, {"Myanmar", "MM"},
		{"Nepal", "NP"}, {"Netherlands", "NL"}, {"New Zealand", "NZ"}, {"Nicaragua", "NI"}, {"Nigeria", "NG"},
		{"Norway", "NO"}, {"Pakistan", "PK"}, {"Panama", "PA"}, {"Paraguay", "PY"}, {"Peru", "PE"},
		{"Philippines", "PH"}, {"Poland", "PL"}, {"Portugal", "PT"}, {"Romania", "RO"}, {"Russia", "RU"},
		{"Saudi Arabia", "SA"}, {"Senegal", "SN"}, {"Serbia", "RS"}, {"Singapore", "SG"}, {"Slovakia", "SK"},
		{"South Africa", "ZA"}, {"South Korea", "KR"}, {"Spain", "ES"}, {"Sri Lanka", "LK"}, {"Sudan", "SD"},
		{"Sweden", "SE"}, {"Switzerland", "CH"}, {"Syria", "SY"}, {"Taiwan", "TW"}, {"Tanzania", "TZ"},
		{"Thailand", "TH"}, {"Tunisia", "TN"}, {"Turkey", "TR"}, {"Uganda", "UG"}, {"Ukraine", "UA"},
		{"United Arab Emirates", "AE"}, {"United Kingdom", "GB"}, {"United States", "US"}, {"Uruguay", "UY"}, {"Uzbekistan", "UZ"},
		{"Venezuela", "VE"}, {"Vietnam", "VN"}, {"Yemen", "YE"}, {"Zambia", "ZM"}, {"Zimbabwe", "ZW"},
	}

	state = map[string][]stateData{
		"en-US": {
			{"Alabama", "AL"}, {"Alaska", "AK"}, {"Arizona", "AZ"}, {"Arkansas", "AR"}, {"California", "CA"},
			{"Colorado", "CO"}, {"Connecticut", "CT"}, {"Delaware", "DE"}, {"Florida", "FL"}, {"Georgia", "GA"},
			{"Hawaii", "HI"}, {"Idaho", "ID"}, {"Illinois", "IL"}, {"Indiana", "IN"}, {"Iowa", "IA"},
			{"Kansas", "KS"}, {"Kentucky", "KY"}, {"Louisiana", "LA"}, {"Maine", "ME"}, {"Maryland", "MD"},
		},
		"pt-BR": {
			{"Acre", "AC"}, {"Alagoas", "AL"}, {"Amapá", "AP"}, {"Amazonas", "AM"}, {"Bahia", "BA"},
			{"Ceará", "CE"}, {"Distrito Federal", "DF"}, {"Espírito Santo", "ES"}, {"Goiás", "GO"}, {"Maranhão", "MA"},
			{"Mato Grosso", "MT"}, {"Mato Grosso do Sul", "MS"}, {"Minas Gerais", "MG"}, {"Pará", "PA"}, {"Paraíba", "PB"},
			{"Paraná", "PR"}, {"Pernambuco", "PE"}, {"Piauí", "PI"}, {"Rio de Janeiro", "RJ"}, {"Rio Grande do Norte", "RN"},
		},
	}

	city = map[string]map[string][]string{
		"en-US": {
			"AL": {"Birmingham", "Montgomery", "Huntsville", "Mobile", "Tuscaloosa", "Hoover", "Dothan", "Auburn", "Decatur", "Madison"},
			"AK": {"Anchorage", "Fairbanks", "Juneau", "Sitka", "Ketchikan", "Wasilla", "Kenai", "Kodiak", "Bethel", "Palmer"},
			"AZ": {"Phoenix", "Tucson", "Mesa", "Chandler", "Scottsdale", "Glendale", "Gilbert", "Tempe", "Peoria", "Surprise"},
			"AR": {"Little Rock", "Fort Smith", "Fayetteville", "Springdale", "Jonesboro", "North Little Rock", "Conway", "Rogers", "Bentonville", "Pine Bluff"},
			"CA": {"Los Angeles", "San Diego", "San Jose", "San Francisco", "Fresno", "Sacramento", "Long Beach", "Oakland", "Bakersfield", "Anaheim"},
			"CO": {"Denver", "Colorado Springs", "Aurora", "Fort Collins", "Lakewood", "Thornton", "Arvada", "Westminster", "Pueblo", "Boulder"},
			"CT": {"Bridgeport", "New Haven", "Hartford", "Stamford", "Waterbury", "Norwalk", "Danbury", "New Britain", "West Hartford", "Greenwich"},
			"DE": {"Wilmington", "Dover", "Newark", "Middletown", "Smyrna", "Milford", "Seaford", "Georgetown", "Elsmere", "New Castle"},
			"FL": {"Jacksonville", "Miami", "Tampa", "Orlando", "St. Petersburg", "Hialeah", "Port St. Lucie", "Cape Coral", "Tallahassee", "Fort Lauderdale"},
			"GA": {"Atlanta", "Augusta", "Columbus", "Macon", "Savannah", "Athens", "Sandy Springs", "South Fulton", "Roswell", "Albany"},
			"HI": {"Honolulu", "East Honolulu", "Pearl City", "Hilo", "Kailua", "Waipahu", "Kaneohe", "Mililani Town", "Kahului", "Ewa Gentry"},
			"ID": {"Boise", "Nampa", "Meridian", "Idaho Falls", "Pocatello", "Caldwell", "Coeur d'Alene", "Twin Falls", "Lewiston", "Post Falls"},
			"IL": {"Chicago", "Aurora", "Naperville", "Joliet", "Rockford", "Springfield", "Elgin", "Peoria", "Champaign", "Waukegan"},
			"IN": {"Indianapolis", "Fort Wayne", "Evansville", "South Bend", "Carmel", "Fishers", "Bloomington", "Hammond", "Gary", "Muncie"},
			"IA": {"Des Moines", "Cedar Rapids", "Davenport", "Sioux City", "Iowa City", "Waterloo", "Ames", "West Des Moines", "Ankeny", "Dubuque"},
			"KS": {"Wichita", "Overland Park", "Kansas City", "Olathe", "Topeka", "Lawrence", "Shawnee", "Manhattan", "Lenexa", "Salina"},
			"KY": {"Louisville", "Lexington", "Bowling Green", "Owensboro", "Covington", "Hopkinsville", "Richmond", "Florence", "Georgetown", "Henderson"},
			"LA": {"New Orleans", "Baton Rouge", "Shreveport", "Metairie", "Lafayette", "Lake Charles", "Kenner", "Bossier City", "Monroe", "Alexandria"},
			"ME": {"Portland", "Lewiston", "Bangor", "South Portland", "Auburn", "Biddeford", "Sanford", "Augusta", "Saco", "Westbrook"},
			"MD": {"Baltimore", "Frederick", "Rockville", "Gaithersburg", "Bowie", "Hagerstown", "Annapolis", "College Park", "Salisbury", "Laurel"},
		},
		"pt-BR": {
			"AC": {"Rio Branco", "Cruzeiro do Sul", "Sena Madureira", "Tarauacá", "Feijó", "Brasiléia", "Epitaciolândia", "Juruá", "Mâncio Lima", "Plácido de Castro"},
			"AL": {"Maceió", "Arapiraca", "Palmeira dos Índios", "Rio Largo", "Penedo", "União dos Palmares", "São Miguel dos Campos", "Coruripe", "Delmiro Gouveia", "Marechal Deodoro"},
			"AP": {"Macapá", "Santana", "Laranjal do Jari", "Oiapoque", "Mazagão", "Porto Grande", "Tartarugalzinho", "Pedra Branca do Amapari", "Calçoene", "Ferreira Gomes"},
			"AM": {"Manaus", "Parintins", "Itacoatiara", "Manacapuru", "Coari", "Tefé", "Tabatinga", "Maués", "Humaitá", "São Gabriel da Cachoeira"},
			"BA": {"Salvador", "Feira de Santana", "Vitória da Conquista", "Camaçari", "Juazeiro", "Itabuna", "Lauro de Freitas", "Ilhéus", "Jequié", "Alagoinhas"},
			"CE": {"Fortaleza", "Caucaia", "Juazeiro do Norte", "Maracanaú", "Sobral", "Crato", "Itapipoca", "Maranguape", "Iguatu", "Quixadá"},
			"DF": {"Brasília", "Ceilândia", "Taguatinga", "Samambaia", "Planaltina", "Águas Claras", "Recanto das Emas", "Gama", "Guará", "Sobradinho"},
			"ES": {"Vitória", "Vila Velha", "Cariacica", "Serra", "Linhares", "Cachoeiro de Itapemirim", "Colatina", "Guarapari", "São Mateus", "Aracruz"},
			"GO": {"Goiânia", "Aparecida de Goiânia", "Anápolis", "Rio Verde", "Luziânia", "Águas Lindas de Goiás", "Valparaíso de Goiás", "Trindade", "Formosa", "Novo Gama"},
			"MA": {"São Luís", "Imperatriz", "São José de Ribamar", "Timon", "Caxias", "Codó", "Paço do Lumiar", "Açailândia", "Bacabal", "Balsas"},
			"MT": {"Cuiabá", "Várzea Grande", "Rondonópolis", "Sinop", "Tangará da Serra", "Cáceres", "Sorriso", "Lucas do Rio Verde", "Primavera do Leste", "Barra do Garças"},
			"MS": {"Campo Grande", "Dourados", "Três Lagoas", "Corumbá", "Ponta Porã", "Naviraí", "Nova Andradina", "Aquidauana", "Sidrolândia", "Maracaju"},
			"MG": {"Belo Horizonte", "Uberlândia", "Contagem", "Juiz de Fora", "Betim", "Montes Claros", "Ribeirão das Neves", "Uberaba", "Governador Valadares", "Ipatinga"},
			"PA": {"Belém", "Ananindeua", "Santarém", "Marabá", "Castanhal", "Parauapebas", "Altamira", "Itaituba", "Abaetetuba", "Tucuruí"},
			"PB": {"João Pessoa", "Campina Grande", "Santa Rita", "Patos", "Bayeux", "Sousa", "Cajazeiras", "Cabedelo", "Guarabira", "Sapé"},
			"PR": {"Curitiba", "Londrina", "Maringá", "Ponta Grossa", "Cascavel", "São José dos Pinhais", "Foz do Iguaçu", "Colombo", "Guarapuava", "Paranaguá"},
			"PE": {"Recife", "Caruaru", "Petrolina", "Olinda", "Paulista", "Jaboatão dos Guararapes", "Garanhuns", "Cabo de Santo Agostinho", "Camaragibe", "Vitória de Santo Antão"},
			"PI": {"Teresina", "Parnaíba", "Picos", "Piripiri", "Floriano", "Campo Maior", "Barras", "União", "Altos", "José de Freitas"},
			"RJ": {"Rio de Janeiro", "São Gonçalo", "Duque de Caxias", "Nova Iguaçu", "Niterói", "Belford Roxo", "São João de Meriti", "Campos dos Goytacazes", "Petrópolis", "Volta Redonda"},
			"RN": {"Natal", "Mossoró", "Parnamirim", "São Gonçalo do Amarante", "Ceará-Mirim", "Caicó", "Açu", "Currais Novos", "Santa Cruz", "Macaíba"},
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

var (
	carBrand = []string{
		"Audi", "BMW", "Chevrolet", "Chrysler", "Citroën",
		"Dodge", "Ferrari", "Fiat", "Ford", "Honda",
		"Hyundai", "Jaguar", "Jeep", "Kia", "Lamborghini",
		"Land Rover", "Lexus", "Maserati", "Mazda", "Mercedes-Benz",
		"Mitsubishi", "Nissan", "Peugeot", "Porsche", "Renault",
		"Subaru", "Tesla", "Toyota", "Volkswagen", "Volvo",
	}

	carModel = map[string][]string{
		"Audi":          {"A3", "A4", "A6", "Q5", "Q7"},
		"BMW":           {"3 Series", "5 Series", "X3", "X5", "M3"},
		"Chevrolet":     {"Silverado", "Equinox", "Malibu", "Camaro", "Tahoe"},
		"Chrysler":      {"300", "Pacifica", "Voyager", "Aspen", "Sebring"},
		"Citroën":       {"C3", "C4", "C5 Aircross", "Berlingo", "Jumper"},
		"Dodge":         {"Charger", "Challenger", "Durango", "Ram 1500", "Journey"},
		"Ferrari":       {"488 GTB", "F8 Tributo", "Roma", "SF90 Stradale", "Portofino"},
		"Fiat":          {"500", "Pulse", "Fastback", "Doblo", "Toro"},
		"Ford":          {"F-150", "Mustang", "Explorer", "Escape", "Bronco"},
		"Honda":         {"Civic", "Accord", "CR-V", "Pilot", "HR-V"},
		"Hyundai":       {"Elantra", "Sonata", "Tucson", "Santa Fe", "Kona"},
		"Jaguar":        {"XE", "XF", "F-Pace", "E-Pace", "I-Pace"},
		"Jeep":          {"Wrangler", "Cherokee", "Grand Cherokee", "Compass", "Renegade"},
		"Kia":           {"Forte", "Optima", "Sportage", "Sorento", "Telluride"},
		"Lamborghini":   {"Huracán", "Urus", "Aventador", "Gallardo", "Murciélago"},
		"Land Rover":    {"Defender", "Discovery", "Range Rover", "Freelander", "Evoque"},
		"Lexus":         {"IS", "ES", "RX", "NX", "GX"},
		"Maserati":      {"Ghibli", "Quattroporte", "Levante", "GranTurismo", "GranCabrio"},
		"Mazda":         {"Mazda3", "Mazda6", "CX-5", "CX-9", "MX-5 Miata"},
		"Mercedes-Benz": {"C-Class", "E-Class", "S-Class", "GLC", "GLE"},
		"Mitsubishi":    {"Outlander", "Eclipse Cross", "Pajero", "L200", "ASX"},
		"Nissan":        {"Altima", "Sentra", "Rogue", "Murano", "Frontier"},
		"Peugeot":       {"208", "308", "3008", "5008", "508"},
		"Porsche":       {"911", "Cayenne", "Macan", "Panamera", "Taycan"},
		"Renault":       {"Clio", "Megane", "Duster", "Sandero", "Captur"},
		"Subaru":        {"Impreza", "Legacy", "Outback", "Forester", "Crosstrek"},
		"Tesla":         {"Model 3", "Model S", "Model X", "Model Y", "Cybertruck"},
		"Toyota":        {"Camry", "Corolla", "RAV4", "Highlander", "Tacoma"},
		"Volkswagen":    {"Golf", "Jetta", "Passat", "Tiguan", "Touareg"},
		"Volvo":         {"S60", "S90", "XC40", "XC60", "XC90"},
	}
)

var (
	companyName = map[string][]string{
		"en-US": {
			"Apex Solutions", "Blue Ridge Technologies", "Cascade Ventures", "Delta Dynamics", "Evergreen Enterprises",
			"Falcon Industries", "Gateway Group", "Harbor Holdings", "Ironclad Systems", "Juniper Partners",
			"Keystone Capital", "Lighthouse Labs", "Meridian Corp", "Nexus Networks", "Orion Consulting",
			"Pinnacle Services", "Quantum Strategies", "Redwood Innovations", "Summit Technologies", "Titan Works",
			"Unified Solutions", "Vanguard Ventures", "Westbrook Associates", "Xenon Digital", "Yellowstone Group",
			"Zenith Global", "Crestview Holdings", "Dawnridge Partners", "Edgewood Systems", "Fairfield Capital",
			"Granite Peak Corp", "Highpoint Solutions", "Ironwood Enterprises", "Jasper Technologies", "Kestrel Group",
			"Lakeview Partners", "Maple Grove Inc", "Northstar Ventures", "Oakdale Industries", "Pacific Ridge Corp",
		},
		"pt-BR": {
			"Ápice Soluções", "Bravo Tecnologia", "Caminho Digital", "Delta Sistemas", "Estrela Ventures",
			"Falcão Indústrias", "Gávea Participações", "Horizonte Holdings", "Inova Sistemas", "Jatobá Parceiros",
			"Kinect Capital", "Luminar Labs", "Meridiano Corp", "Nexus Redes", "Orion Consultoria",
			"Pináculo Serviços", "Quantum Estratégias", "Redwood Inovações", "Sumitec Tecnologia", "Titan Obras",
			"Unida Soluções", "Vanguarda Ventures", "Westbrook Associados", "Xênon Digital", "Yellowstone Grupo",
			"Zênite Global", "Crestview Holdings", "Alvorada Parceiros", "Edgewood Sistemas", "Fairfield Capital",
			"Granito Pico Corp", "Altoponto Soluções", "Ironwood Empreendimentos", "Jasper Tecnologias", "Kestrel Grupo",
			"Lakeview Parceiros", "Maplewood Ltda", "Estrela do Norte Ventures", "Oakdale Indústrias", "Crista do Pacífico Corp",
		},
	}

	companyDba = map[string][]string{
		"en-US": {
			"Apex Pro", "Blue Tech", "CascadeX", "DeltaNet", "EverGreen Co",
			"FalconTech", "GateGroup", "HarborHub", "IronSys", "JuniperPro",
			"KeyCap", "LightLabs", "MeriCorp", "NexNet", "OrionCo",
			"PinnServ", "QuantumX", "RedwoodTech", "SummitCo", "TitanWorks",
			"UniSol", "VanVentures", "WestAssoc", "XenoDigital", "YellowGroup",
			"ZenithCo", "CrestHold", "DawnPart", "EdgeSys", "FairCap",
			"GranitePeak", "HighSol", "IronEnter", "JasperTech", "KestrelGrp",
			"LakePartners", "MapleGrove", "NorthVentures", "OakIndustries", "PacificRidge",
		},
		"pt-BR": {
			"Ápice Pro", "BlueTech BR", "CaminhoX", "DeltaNet BR", "EstrelaCo",
			"FalcãoTech", "GáveaGrupo", "HorizonteHub", "InovaSys", "JatobáPro",
			"KinectCap", "LuminarLabs", "MeriCorp BR", "NexRedes", "OrionBR",
			"PinácServ", "QuantumBR", "RedwoodTech BR", "SumitecCo", "TitanObras",
			"UnidaSol", "VanguardaBR", "WestAssoc BR", "XênoBR", "YeloBR",
			"ZêniteBR", "CrestHold BR", "AlvoradaPart", "EdgeSys BR", "FairCap BR",
			"GranitoPico", "AltopSol", "IronEmpreend", "JasperTech BR", "KestrelGrp BR",
			"LakePart BR", "MapleWood BR", "EstrelaNorte", "OakIndustriais", "CristaPacífico",
		},
	}

	companyIndustry = map[string][]string{
		"en-US": {
			"Technology", "Healthcare", "Finance", "Education", "Retail",
			"Manufacturing", "Transportation", "Energy", "Real Estate", "Agriculture",
			"Entertainment", "Hospitality", "Construction", "Telecommunications", "Pharmaceuticals",
			"Automotive", "Aerospace", "Food & Beverage", "Media", "Logistics",
		},
		"pt-BR": {
			"Tecnologia", "Saúde", "Finanças", "Educação", "Varejo",
			"Manufatura", "Transporte", "Energia", "Imóveis", "Agronegócio",
			"Entretenimento", "Hotelaria", "Construção Civil", "Telecomunicações", "Farmacêutica",
			"Automotivo", "Aeroespacial", "Alimentos e Bebidas", "Mídia", "Logística",
		},
	}

	companySuffix = map[string][]string{
		"en-US": {
			"Inc.", "LLC", "Ltd.", "Corp.", "Co.",
			"Group", "Associates", "Partners", "Holdings", "Enterprises",
		},
		"pt-BR": {
			"Ltda.", "S.A.", "EIRELI", "MEI", "S/S",
			"EPP", "ME", "S.C.", "COOP",
		},
	}
)

type currencyData struct {
	name   string
	code   string
	symbol string
}

var currency = []currencyData{
	{"Afghan Afghani", "AFN", "؋"},
	{"Albanian Lek", "ALL", "L"},
	{"Algerian Dinar", "DZD", "د.ج"},
	{"Angolan Kwanza", "AOA", "Kz"},
	{"Argentine Peso", "ARS", "$"},
	{"Australian Dollar", "AUD", "A$"},
	{"Euro", "EUR", "€"},
	{"Bangladeshi Taka", "BDT", "৳"},
	{"Euro", "EUR", "€"},
	{"Boliviano", "BOB", "Bs."},
	{"Brazilian Real", "BRL", "R$"},
	{"Bulgarian Lev", "BGN", "лв"},
	{"Cambodian Riel", "KHR", "៛"},
	{"Canadian Dollar", "CAD", "CA$"},
	{"Chilean Peso", "CLP", "$"},
	{"Chinese Yuan", "CNY", "¥"},
	{"Colombian Peso", "COP", "$"},
	{"Costa Rican Colón", "CRC", "₡"},
	{"Euro", "EUR", "€"},
	{"Cuban Peso", "CUP", "$"},
	{"Czech Koruna", "CZK", "Kč"},
	{"Danish Krone", "DKK", "kr"},
	{"Dominican Peso", "DOP", "RD$"},
	{"US Dollar", "USD", "$"},
	{"Egyptian Pound", "EGP", "£"},
	{"US Dollar", "USD", "$"},
	{"Ethiopian Birr", "ETB", "Br"},
	{"Euro", "EUR", "€"},
	{"Euro", "EUR", "€"},
	{"Euro", "EUR", "€"},
	{"Ghanaian Cedi", "GHS", "₵"},
	{"Euro", "EUR", "€"},
	{"Guatemalan Quetzal", "GTQ", "Q"},
	{"Honduran Lempira", "HNL", "L"},
	{"Hungarian Forint", "HUF", "Ft"},
	{"Indian Rupee", "INR", "₹"},
	{"Indonesian Rupiah", "IDR", "Rp"},
	{"Iranian Rial", "IRR", "﷼"},
	{"Iraqi Dinar", "IQD", "ع.د"},
	{"Euro", "EUR", "€"},
	{"Israeli New Shekel", "ILS", "₪"},
	{"Euro", "EUR", "€"},
	{"Jamaican Dollar", "JMD", "J$"},
	{"Japanese Yen", "JPY", "¥"},
	{"Jordanian Dinar", "JOD", "د.ا"},
	{"Kazakhstani Tenge", "KZT", "₸"},
	{"Kenyan Shilling", "KES", "KSh"},
	{"Kuwaiti Dinar", "KWD", "د.ك"},
	{"Lebanese Pound", "LBP", "ل.ل"},
	{"Libyan Dinar", "LYD", "ل.د"},
	{"Malaysian Ringgit", "MYR", "RM"},
	{"Mexican Peso", "MXN", "$"},
	{"Moroccan Dirham", "MAD", "د.م."},
	{"Mozambican Metical", "MZN", "MT"},
	{"Burmese Kyat", "MMK", "K"},
	{"Nepalese Rupee", "NPR", "₨"},
	{"Euro", "EUR", "€"},
	{"New Zealand Dollar", "NZD", "NZ$"},
	{"Nicaraguan Córdoba", "NIO", "C$"},
	{"Nigerian Naira", "NGN", "₦"},
	{"Norwegian Krone", "NOK", "kr"},
	{"Pakistani Rupee", "PKR", "₨"},
	{"Panamanian Balboa", "PAB", "B/."},
	{"Paraguayan Guaraní", "PYG", "₲"},
	{"Peruvian Sol", "PEN", "S/"},
	{"Philippine Peso", "PHP", "₱"},
	{"Polish Złoty", "PLN", "zł"},
	{"Euro", "EUR", "€"},
	{"Romanian Leu", "RON", "lei"},
	{"Russian Ruble", "RUB", "₽"},
	{"Saudi Riyal", "SAR", "﷼"},
	{"West African CFA Franc", "XOF", "Fr"},
	{"Serbian Dinar", "RSD", "din"},
	{"Singapore Dollar", "SGD", "S$"},
	{"Euro", "EUR", "€"},
	{"South African Rand", "ZAR", "R"},
	{"South Korean Won", "KRW", "₩"},
	{"Euro", "EUR", "€"},
	{"Sri Lankan Rupee", "LKR", "₨"},
	{"Sudanese Pound", "SDG", "ج.س."},
	{"Swedish Krona", "SEK", "kr"},
	{"Swiss Franc", "CHF", "Fr"},
	{"Syrian Pound", "SYP", "£"},
	{"New Taiwan Dollar", "TWD", "NT$"},
	{"Tanzanian Shilling", "TZS", "TSh"},
	{"Thai Baht", "THB", "฿"},
	{"Tunisian Dinar", "TND", "د.ت"},
	{"Turkish Lira", "TRY", "₺"},
	{"Ugandan Shilling", "UGX", "USh"},
	{"Ukrainian Hryvnia", "UAH", "₴"},
	{"UAE Dirham", "AED", "د.إ"},
	{"Pound Sterling", "GBP", "£"},
	{"US Dollar", "USD", "$"},
	{"Uruguayan Peso", "UYU", "$U"},
	{"Uzbekistani Som", "UZS", "so'm"},
	{"Venezuelan Bolívar", "VES", "Bs.S"},
	{"Vietnamese Đồng", "VND", "₫"},
	{"Yemeni Rial", "YER", "﷼"},
	{"Zambian Kwacha", "ZMW", "ZK"},
	{"Zimbabwean Dollar", "ZWL", "Z$"},
}

type fileData struct {
	prefix string
	ext    string
	mime   string
}

var (
	file = []fileData{
		{"document", "pdf", "application/pdf"},
		{"document", "doc", "application/msword"},
		{"document", "docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
		{"document", "txt", "text/plain"},
		{"document", "odt", "application/vnd.oasis.opendocument.text"},
		{"document", "rtf", "application/rtf"},
		{"spreadsheet", "xls", "application/vnd.ms-excel"},
		{"spreadsheet", "xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		{"spreadsheet", "csv", "text/csv"},
		{"spreadsheet", "ods", "application/vnd.oasis.opendocument.spreadsheet"},
		{"presentation", "ppt", "application/vnd.ms-powerpoint"},
		{"presentation", "pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
		{"presentation", "odp", "application/vnd.oasis.opendocument.presentation"},
		{"image", "jpg", "image/jpeg"},
		{"image", "jpeg", "image/jpeg"},
		{"image", "png", "image/png"},
		{"image", "gif", "image/gif"},
		{"image", "bmp", "image/bmp"},
		{"image", "svg", "image/svg+xml"},
		{"image", "webp", "image/webp"},
		{"image", "tiff", "image/tiff"},
		{"image", "ico", "image/x-icon"},
		{"audio", "mp3", "audio/mpeg"},
		{"audio", "wav", "audio/wav"},
		{"audio", "flac", "audio/flac"},
		{"audio", "aac", "audio/aac"},
		{"audio", "ogg", "audio/ogg"},
		{"video", "mp4", "video/mp4"},
		{"video", "avi", "video/x-msvideo"},
		{"video", "mkv", "video/x-matroska"},
		{"video", "mov", "video/quicktime"},
		{"video", "wmv", "video/x-ms-wmv"},
		{"archive", "zip", "application/zip"},
		{"archive", "tar", "application/x-tar"},
		{"archive", "gz", "application/gzip"},
		{"archive", "rar", "application/vnd.rar"},
		{"archive", "7z", "application/x-7z-compressed"},
		{"data", "json", "application/json"},
		{"data", "xml", "application/xml"},
		{"data", "yaml", "application/x-yaml"},
	}
)

var (
	webDomain = []string{
		"google", "youtube", "facebook", "twitter", "instagram",
		"linkedin", "github", "microsoft", "apple", "amazon",
		"netflix", "spotify", "reddit", "wikipedia", "dropbox",
		"slack", "zoom", "adobe", "salesforce", "oracle",
		"ibm", "intel", "nvidia", "samsung", "sony",
		"paypal", "stripe", "shopify", "wordpress", "medium",
		"notion", "figma", "canva", "trello", "asana",
		"mailchimp", "hubspot", "intercom", "zendesk", "twilio",
		"cloudflare", "digitalocean", "heroku", "vercel", "netlify",
		"mongodb", "postgresql", "redis", "elastic", "grafana",
		"gitlab", "bitbucket", "jira", "confluence", "jenkins",
		"docker", "kubernetes", "terraform", "ansible", "vault",
		"airbnb", "uber", "lyft", "doordash", "instacart",
		"ebay", "etsy", "aliexpress", "walmart", "target",
		"bloomberg", "reuters", "techcrunch", "theverge", "wired",
		"cnn", "bbc", "nytimes", "guardian", "forbes",
		"harvard", "mit", "stanford", "coursera", "udemy",
		"twitch", "discord", "telegram", "whatsapp", "snapchat",
		"tiktok", "pinterest", "tumblr", "quora", "stackoverflow",
		"booking", "tripadvisor", "expedia", "kayak", "airbnb",
	}

	webTLD = []string{
		"com", "org", "net", "edu", "gov",
		"io", "co", "app", "dev", "tech",
		"info", "biz", "online", "store", "shop",
		"us", "uk", "ca", "br", "de",
		"fr", "jp", "au", "in", "nl",
	}

	webSubdomain = []string{
		"www", "mail", "api", "app", "admin",
		"blog", "shop", "store", "portal", "dashboard",
		"dev", "staging", "sandbox", "test", "demo",
		"cdn", "static", "assets", "media", "images",
		"docs", "help", "support", "kb", "status",
		"auth", "login", "account", "profile", "secure",
		"m", "mobile", "web", "beta", "labs",
		"news", "press", "careers", "jobs", "about",
		"forum", "community", "chat", "talk", "social",
		"analytics", "data", "reports", "metrics", "insights",
	}

	webUserAgent = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.59",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Mobile Safari/537.36",
	}
)

var (
	loremWord = []string{
		"lorem", "ipsum", "dolor", "sit", "amet",
		"consectetur", "adipiscing", "elit", "sed", "do",
		"eiusmod", "tempor", "incididunt", "ut", "labore",
		"et", "dolore", "magna", "aliqua", "enim",
		"ad", "minim", "veniam", "quis", "nostrud",
		"exercitation", "ullamco", "laboris", "nisi", "aliquip",
		"ex", "ea", "commodo", "consequat", "duis",
		"aute", "irure", "in", "reprehenderit", "voluptate",
		"velit", "esse", "cillum", "eu", "fugiat",
		"nulla", "pariatur", "excepteur", "sint", "occaecat",
		"cupidatat", "non", "proident", "sunt", "culpa",
		"qui", "officia", "deserunt", "mollit", "anim",
		"id", "est", "laborum", "curabitur", "pretium",
		"tincidunt", "lacus", "nec", "porta", "ante",
		"sapien", "vel", "aliquet", "augue", "arcu",
		"dignissim", "risus", "diam", "viverra", "nisl",
		"accumsan", "sodales", "ornare", "cras", "nunc",
		"ligula", "metus", "dictum", "quam", "proin",
		"egestas", "urna", "volutpat", "faucibus", "purus",
		"felis", "gravida", "neque", "convallis", "massa",
	}
)

type creditCardVendorData struct {
	name   string
	bin    []string
	length []int
}

var (
	creditCardVendor = []creditCardVendorData{
		{"Visa", []string{"4"}, []int{13, 16, 19}},
		{"Mastercard", []string{"51", "52", "53", "54", "55", "2221", "2720"}, []int{16}},
		{"American Express", []string{"34", "37"}, []int{15}},
		{"Discover", []string{"6011", "622", "644", "645", "646", "647", "648", "649", "65"}, []int{16, 19}},
		{"JCB", []string{"3528", "3535", "3538", "3589"}, []int{16, 17, 18, 19}},
		{"Diners Club", []string{"300", "301", "302", "303", "304", "305", "36", "38"}, []int{14}},
		{"UnionPay", []string{"62"}, []int{16, 17, 18, 19}},
		{"Maestro", []string{"6304", "6759", "6761", "6762", "6763"}, []int{12, 13, 14, 15, 16, 17, 18, 19}},
		{"Mir", []string{"2200", "2201", "2202", "2203", "2204"}, []int{16, 17, 18, 19}},
		{"Elo", []string{"636368", "438935", "504175", "451416", "509048", "509067", "509049", "509069", "509050", "509074"}, []int{16}},
		{"Hipercard", []string{"606282", "3841"}, []int{13, 16}},
		{"RuPay", []string{"60", "6521", "6522"}, []int{16}},
		{"Troy", []string{"9792"}, []int{16}},
		{"Verve", []string{"5061", "6500", "6501", "6502"}, []int{16, 19}},
		{"UATP", []string{"1"}, []int{15}},
		{"Dankort", []string{"5019", "4571"}, []int{16}},
		{"Bancontact", []string{"6703"}, []int{16}},
		{"Aura", []string{"50"}, []int{16}},
		{"Laser", []string{"6304", "6706", "6771", "6709"}, []int{16, 17, 18, 19}},
		{"Cartes Bancaires", []string{"4", "5"}, []int{16}},
	}
)

type personPhoneData struct {
	countryCode string
	format      []string
}

var (
	personFirstName = map[string][]string{
		"en-US": {
			"James", "John", "Robert", "Michael", "William",
			"David", "Richard", "Joseph", "Thomas", "Charles",
			"Christopher", "Daniel", "Matthew", "Anthony", "Donald",
			"Mary", "Patricia", "Jennifer", "Linda", "Barbara",
			"Elizabeth", "Susan", "Jessica", "Sarah", "Karen",
			"Lisa", "Nancy", "Betty", "Margaret", "Sandra",
			"Ashley", "Emily", "Amanda", "Melissa", "Stephanie",
		},
		"pt-BR": {
			"Miguel", "Arthur", "Heitor", "Davi", "Gabriel",
			"Pedro", "Matheus", "Lucas", "Benjamim", "Nicolas",
			"Guilherme", "Rafael", "Felipe", "João", "Leonardo",
			"Sofia", "Alice", "Valentina", "Laura", "Isabella",
			"Manuela", "Júlia", "Heloísa", "Luísa", "Maria",
			"Beatriz", "Lara", "Ana", "Clara", "Lívia",
			"Fernanda", "Carla", "Bruna", "Camila", "Mariana",
		},
	}

	personMiddleName = map[string][]string{
		"en-US": {
			"Ray", "Jay", "Mae", "Ann", "Lynn",
			"Dean", "Dale", "Wayne", "Gene", "Earl",
			"Keith", "Scott", "Alan", "Todd", "Blair",
			"Grace", "Marie", "Rose", "Jean", "Claire",
			"Faith", "Hope", "Joy", "Dawn", "Faye",
			"Nicole", "Renee", "Beth", "Gail", "Lane",
			"Brooke", "Chase", "Drew", "Reid", "Paige",
		},
		"pt-BR": {
			"Maria", "José", "João", "Ana", "Luís",
			"Antônio", "Francisco", "Paulo", "Pedro", "Carlos",
			"Luiza", "Rita", "Helena", "Graça", "Fátima",
			"Aparecida", "Conceição", "Benedita", "Sebastião", "Augusto",
			"Henrique", "Eduardo", "Renato", "Roberto", "Ricardo",
			"Marcelo", "Maurício", "Cláudia", "Patrícia", "Letícia",
			"Adriana", "Denise", "Vanessa", "Cristiane", "Viviane",
		},
	}

	personLastName = map[string][]string{
		"en-US": {
			"Smith", "Johnson", "Williams", "Brown", "Jones",
			"Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
			"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
			"Thomas", "Taylor", "Moore", "Jackson", "Martin",
			"Lee", "Perez", "Thompson", "White", "Harris",
			"Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
			"Walker", "Young", "Allen", "King", "Wright",
		},
		"pt-BR": {
			"Silva", "Santos", "Oliveira", "Souza", "Rodrigues",
			"Ferreira", "Alves", "Pereira", "Lima", "Gomes",
			"Costa", "Ribeiro", "Martins", "Carvalho", "Almeida",
			"Lopes", "Sousa", "Fernandes", "Vieira", "Barbosa",
			"Rocha", "Dias", "Nascimento", "Andrade", "Moreira",
			"Nunes", "Marques", "Machado", "Mendes", "Freitas",
			"Cardoso", "Ramos", "Teixeira", "Araújo", "Campos",
		},
	}

	personPhone = []personPhoneData{
		{"1", []string{"201 ###-####", "212 ###-####", "310 ###-####", "415 ###-####", "312 ###-####"}},
		{"55", []string{"11 #####-####", "21 #####-####", "31 #####-####", "11 ####-####", "21 ####-####"}},
		{"44", []string{"20 #### ####", "113 ### ####", "131 ### ####", "7### ######"}},
		{"49", []string{"30 ########", "89 ########", "211 ########", "160 ########"}},
		{"33", []string{"1 ## ## ## ##", "4 ## ## ## ##", "6 ## ## ## ##", "7 ## ## ## ##"}},
		{"39", []string{"02 ########", "06 ########", "3## #######"}},
		{"34", []string{"91 ### ## ##", "93 ### ## ##", "6## ### ###"}},
		{"7", []string{"495 ###-##-##", "812 ###-##-##", "9## ###-##-##"}},
		{"86", []string{"10 #### ####", "21 #### ####", "139 #### ####", "186 #### ####"}},
		{"81", []string{"3 #### ####", "6 #### ####", "90 #### ####", "80 #### ####"}},
		{"82", []string{"2 #### ####", "10 #### ####", "10 #### ####"}},
		{"91", []string{"11 #### ####", "22 #### ####", "98## ######"}},
		{"52", []string{"55 #### ####", "33 #### ####", "1## ### ####"}},
		{"54", []string{"11 ####-####", "351 ####-####", "9 #### ####"}},
		{"56", []string{"2 #### ####", "9 #### ####"}},
		{"57", []string{"1 ### ####", "310 ### ####", "300 ### ####"}},
		{"61", []string{"2 #### ####", "4## ### ###", "4## ### ###"}},
		{"64", []string{"9 ### ####", "3 ### ####", "4 ### ####"}},
		{"27", []string{"11 ### ####", "21 ### ####", "8# ### ####"}},
		{"234", []string{"803 ### ####", "806 ### ####", "813 ### ####"}},
		{"20", []string{"2 #### ####", "10 #### ####", "100 ### ####"}},
		{"92", []string{"21 #### ####", "300 ### ####", "301 ### ####"}},
		{"62", []string{"21 #### ####", "22 #### ####", "81## ######"}},
		{"63", []string{"2 ### ####", "917 ### ####", "918 ### ####"}},
		{"66", []string{"2 ### ####", "8# ### ####", "9# ### ####"}},
		{"84", []string{"24 #### ####", "28 #### ####", "9# #### ####"}},
		{"90", []string{"212 ### ####", "312 ### ####", "5## ### ####"}},
		{"380", []string{"44 ### ####", "63 ### ####", "67 ### ####"}},
		{"48", []string{"22 ### ## ##", "12 ### ## ##", "5## ### ###"}},
		{"351", []string{"21 ### ####", "22 ### ####", "9# ### ####"}},
	}

	personEmailDomain = []string{
		"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com",
		"protonmail.com", "live.com", "msn.com", "aol.com", "zoho.com",
		"mail.com", "yandex.com", "gmx.com", "tutanota.com", "fastmail.com",
		"inbox.com", "hushmail.com", "runbox.com", "mailfence.com", "posteo.net",
		"uol.com.br", "bol.com.br", "ig.com.br", "terra.com.br", "globo.com",
		"zipmail.com.br", "r7.com", "oi.com.br", "yahoo.com.br", "hotmail.com.br",
		"web.de", "orange.fr", "free.fr", "laposte.net", "sfr.fr",
		"libero.it", "virgilio.it", "tin.it", "seznam.cz", "wp.pl",
	}

	personJobTitle = map[string][]string{
		"en-US": {
			"Software Engineer", "Product Manager", "Data Scientist", "UX Designer", "DevOps Engineer",
			"Marketing Manager", "Sales Representative", "Financial Analyst", "HR Manager", "Operations Manager",
			"Chief Executive Officer", "Chief Technology Officer", "Chief Financial Officer", "Vice President", "Director",
			"Project Manager", "Business Analyst", "Systems Administrator", "Network Engineer", "Security Analyst",
			"Accountant", "Legal Counsel", "Customer Success Manager", "Content Strategist", "Brand Manager",
			"Recruiter", "Data Engineer", "Machine Learning Engineer", "QA Engineer", "Technical Writer",
			"Supply Chain Manager", "Procurement Specialist", "Research Scientist", "Graphic Designer", "Full Stack Developer",
		},
		"pt-BR": {
			"Engenheiro de Software", "Gerente de Produto", "Cientista de Dados", "Designer UX", "Engenheiro DevOps",
			"Gerente de Marketing", "Representante de Vendas", "Analista Financeiro", "Gerente de RH", "Gerente de Operações",
			"Diretor Executivo", "Diretor de Tecnologia", "Diretor Financeiro", "Vice-Presidente", "Diretor Geral",
			"Gerente de Projetos", "Analista de Negócios", "Administrador de Sistemas", "Engenheiro de Redes", "Analista de Segurança",
			"Contador", "Assessor Jurídico", "Gerente de Sucesso do Cliente", "Estrategista de Conteúdo", "Gerente de Marca",
			"Recrutador", "Engenheiro de Dados", "Engenheiro de Machine Learning", "Analista de QA", "Redator Técnico",
			"Gerente da Cadeia de Suprimentos", "Especialista em Compras", "Pesquisador Científico", "Designer Gráfico", "Desenvolvedor Full Stack",
		},
	}
)
