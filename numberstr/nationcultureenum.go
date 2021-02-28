package numberstr

import "sync"

var mCountryCultureCodeToString = map[int]string{
	-1:     "None",
	8:      "Albania",
	32:     "Argentina",
	36:     "Australia",
	40:     "Austria",
	48:     "Bahrain",
	999999: "Bitcoin",
	50:     "Bangladesh",
	56:     "Belgium",
	70:     "BosniaHerzegovina",
	76:     "Brazil",
	100:    "Bulgaria",
	124:    "Canada",
	9124:   "CanadaFrench",
	152:    "Chile",
	156:    "China",
	170:    "Columbia",
	188:    "CostaRica",
	191:    "Croatia",
	192:    "Cuba",
	196:    "Cyprus",
	203:    "CzechiaEuro",
	9203:   "CzechiaKoruna",
	208:    "Denmark",
	818:    "Egypt",
	233:    "Estonia",
	77777:  "Euro",
	246:    "Finland",
	250:    "France",
	276:    "Germany",
	300:    "Greece",
	344:    "HongKong",
	348:    "Hungary",
	356:    "India",
	360:    "Indonesia",
	364:    "Iran",
	372:    "Ireland",
	376:    "Israel",
	380:    "Italy",
	392:    "Japan",
	410:    "KoreaSouth",
	414:    "Kuwait",
	428:    "Latvia",
	438:    "Liechtenstein",
	440:    "Lithuania",
	442:    "Luxembourg",
	458:    "Malaysia",
	470:    "Malta",
	484:    "Mexico",
	504:    "Morocco",
	516:    "Namibia",
	524:    "Nepal",
	528:    "Netherlands",
	554:    "NewZealand",
	578:    "Norway",
	512:    "Oman",
	586:    "Pakistan",
	604:    "Peru",
	608:    "Philippines",
	616:    "Poland",
	620:    "Portugal",
	634:    "Qatar",
	642:    "Romania",
	643:    "Russia",
	682:    "SaudiArabia",
	688:    "Serbia",
	702:    "Singapore",
	703:    "Slovakia",
	710:    "SouthAfrica",
	724:    "Spain",
	144:    "SriLanka",
	752:    "Sweden",
	756:    "Switzerland",
	158:    "Taiwan",
	792:    "Turkey",
	804:    "Ukraine",
	784:    "UnitedArabEmirates",
	826:    "UnitedKingdom",
	840:    "UnitedStates",
	862:    "Venezuela",
	704:    "VietNam",
}

// CountryCultureId - Country Culture Identifier is a means of
// identifying an individual country or culture. This enumeration
// is primarily used to request and assign number and currency
// formatting for use in configuring number strings
//
// Since the Go Programming Language does not directly support
// enumerations, the 'CountryCultureId' type has been adapted to
// function in a manner similar to classic enumerations.
// 'CountryCultureId' is declared as a type 'int'. The method names
// for this type effectively represent an enumeration of countries
// and cultures. A sample of these methods is shown below:
//
//
// Method                              Integer
// Name                                 Value
// ------                              -------
//
// CountryCultureId(0).None()            -1
// CountryCultureId(0).France()         250
// CountryCultureId(0).Germany()        276
// CountryCultureId(0).UnitedKingdom()  826
// CountryCultureId(0).UnitedStates()   840
//
// The integer values for Country Codes are taken from the
// ISO 3166-1 specification:
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
// Cultures and other custom entities are assigned arbitrary
// 4-digit number codes.
//
//
// For easy access to these enumeration values, use the global
// variable 'CntryCulId'.
//     Example: CntryCulId.UnitedStates()
//
// Otherwise you will need to use the formal syntax.
//     Example: CountryCultureId(0).UnitedStates()
//
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the CountryCultureId methods in
// alphabetical order. Be advised that all 'CountryCultureId'
// methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration.
//
type CountryCultureId int

var lockCountryCultureId sync.Mutex

// None - Signals that the CountryCultureId Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) None() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(-1)
}

// Albania - Specifies The Republic of Albania.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Albania() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(8)
}

// Argentina - Specifies the country of Argentina
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Argentina() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(32)
}

// Australia - Specifies the country of Australia
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Australia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(36)
}

// Austria - Specifies the country of Austria
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Austria() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(40)
}

// Bahrain - Specifies The Kingdom of Bahrain.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Bahrain() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(48)
}

// Bitcoin - Specifies the country of Bitcoin
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Bitcoin() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(999999)
}

// Bangladesh - Specifies The People's Republic of Bangladesh.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Bangladesh() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(40)
}

// Belgium - Specifies the country of The Kingdom of Belgium.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Belgium() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(56)
}

// BosniaHerzegovina - Specifies the country of
// Bosnia and Herzegovina.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) BosniaHerzegovina() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(70)
}

// Brazil - Specifies the country of Brazil.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Brazil() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(76)
}

// Bulgaria - Specifies The Republic of Bulgaria.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Bulgaria() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(76)
}

// Canada - Specifies the country of Canada
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Canada() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(124)
}

// CanadaFrench - Specifies the culture of CanadaFrench
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) CanadaFrench() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(9124)
}

// Chile - Specifies the country of Chile
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Chile() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(152)
}

// China - Specifies the country of China
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) China() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(156)
}

// Columbia - Specifies the country of Columbia
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Columbia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(170)
}

// CostaRica - Specifies The Republic of Costa Rica.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) CostaRica() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(188)
}

// Croatia - Specifies The Republic of Croatia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Croatia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(191)
}

// Cuba - Specifies The Republic of Cuba.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Cuba() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(192)
}

// Cyprus - Specifies The Republic of Cyprus.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Cyprus() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(196)
}

// CzechiaEuro - Specifies the The Czech Republic
//
// Czechia, or The Czech Republic, is a member of the European
// Union. As such, it is legally bound to adopt the 'Euro' as its
// official currency. As of 2020, it has been hesitant to do so
// and the Czech Koruna still remains in wide spread use.
//
// This format configures number strings with the Euro currency
// symbol. For the Czech Koruna currency, reference:
//   CountryCultureId.CzechiaKoruna()
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) CzechiaEuro() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(203)
}

// CzechiaKoruna - Specifies the The Czech Republic
// and the Koruna Currency.
//
// Czechia, or The Czech Republic, is a member of the European
// Union. As such, it is legally bound to adopt the 'Euro' as its
// official currency. As of 2020, it has been hesitant to do so
// and the Czech Koruna still remains in wide spread use.
//
// This format configures number strings with the Koruna currency
// symbol. For the Czech Euro currency, reference:
//   CountryCultureId.CzechiaEuro()
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) CzechiaKoruna() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(9203)
}

// Denmark - Specifies the The Kingdom of Denmark
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Denmark() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(208)
}

// Egypt - Specifies the The Arab Republic of Egypt.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Egypt() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(818)
}

// France - Specifies the The French Republic.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) France() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(250)
}

// Estonia - Specifies the The Republic of Estonia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Estonia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(233)
}

// Euro - Specifies the The  European Union.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Euro() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(233)
}

// Finland - Specifies the The Republic of Finland.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Finland() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(77777)
}

// Germany - Specifies the Federal Republic of Germany
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Germany() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(276)
}

// Greece - Specifies The Hellenic Republic.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Greece() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(276)
}

// HongKong - Specifies The Hong Kong Special Administrative
// Region of China.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) HongKong() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(344)
}

// Hungary - Specifies the country of Hungary.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Hungary() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(348)
}

// India - Specifies The Republic of India.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) India() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(356)
}

// Indonesia - Specifies The Republic of Indonesia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Indonesia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(360)
}

// Iran - Specifies The Islamic Republic of Iran.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Iran() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(364)
}

// Ireland - Specifies the country of Ireland.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Ireland() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(372)
}

// Israel - Specifies the State of Israel.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Israel() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(376)
}

// Italy - Specifies the Italian Republic.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Italy() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(380)
}

// Japan - Specifies the country Japan.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Japan() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(392)
}

// KoreaSouth - Specifies The Republic of Korea.
// (South Korea)
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) KoreaSouth() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(410)
}

// Kuwait - Specifies The State of Kuwait.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Kuwait() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(414)
}

// Latvia - Specifies The Republic of Latvia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Latvia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(428)
}

// Liechtenstein - Specifies The Principality of Liechtenstein.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Liechtenstein() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(438)
}

// Lithuania - Specifies The Republic of Lithuania.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Lithuania() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(440)
}

// Luxembourg - Specifies The Grand Duchy of Luxembourg.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Luxembourg() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(442)
}

// Malaysia - Specifies the country of Malaysia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Malaysia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(458)
}

// Malta - Specifies the country of Malaysia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Malta() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(470)
}

// Mexico - Specifies The United Mexican States.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Mexico() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(484)
}

// Morocco - Specifies The Kingdom of Morocco.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Morocco() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(504)
}

// Namibia - Specifies The Republic of Namibia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Namibia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(516)
}

// Nepal - Specifies The Federal Democratic Republic of Nepal.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Nepal() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(524)
}

// Netherlands - Specifies the Kingdom of the Netherlands.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Netherlands() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(528)
}

// NewZealand - Specifies the country of New Zealand.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) NewZealand() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(554)
}

// Norway - Specifies the Kingdom of Norway.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Norway() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(578)
}

// Oman - Specifies The Sultanate of Oman.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Oman() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(512)
}

// Pakistan - Specifies The Islamic Republic of Pakistan.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Pakistan() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(586)
}

// Peru - Specifies The Republic of Perú.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Peru() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(586)
}

// Philippines - Specifies The Republic of the Philippines.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Philippines() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(608)
}

// Poland - Specifies The Republic of Poland.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Poland() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(616)
}

// Portugal - Specifies The Portuguese Republic.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Portugal() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(620)
}

// Qatar - Specifies The State of Qatar.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Qatar() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(620)
}

// Romania - Specifies the country of Romania.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Romania() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(642)
}

// Russia - Specifies The Russian Federation.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Russia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(643)
}

// SaudiArabia - Specifies The Kingdom of Saudi Arabia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) SaudiArabia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(682)
}

// Serbia - Specifies The Republic of Serbia.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Serbia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(688)
}

// Singapore - Specifies The Republic of Singapore.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Singapore() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(682)
}

// Slovakia - Specifies The Republic of Singapore.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Slovakia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(703)
}

// SouthAfrica - Specifies The Republic of South Africa.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) SouthAfrica() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(710)
}

// Spain - Specifies The Kingdom of Spain.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Spain() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(724)
}

// SriLanka - Sri Lanka. Specifies The Democratic Socialist
// Republic of Sri Lanka.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) SriLanka() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(144)
}

// Sweden - Specifies the Kingdom of Sweden.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Sweden() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(752)
}

// Switzerland - Specifies The Swiss Confederation.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Switzerland() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(756)
}

// Taiwan - Specifies The Republic of China.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Taiwan() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(158)
}

// Turkey - Specifies The Republic of Turkey.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Turkey() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(792)
}

// Ukraine - Specifies the country of Ukraine.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Ukraine() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(804)
}

// UnitedArabEmirates - Specifies the country of
// The United Arab Emirates.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) UnitedArabEmirates() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(784)
}

// UnitedKingdom - Specifies the United Kingdom.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) UnitedKingdom() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(826)
}

// UnitedStates - Specifies the United States.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) UnitedStates() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(840)
}

// Venezuela - Specifies The Bolivarian Republic of
// Venezuela.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Venezuela() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(862)
}

// VietNam - Specifies The Socialist Republic of Viet Nam
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) VietNam() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(704)
}
