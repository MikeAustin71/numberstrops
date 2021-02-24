package numberstr

import "sync"

var mCountryCultureCodeToString = map[int]string{
	-1:   "None",
	32:   "Argentina",
	36:   "Australia",
	40:   "Austria",
	56:   "Belgium",
	76:   "Brazil",
	124:  "Canada",
	9124: "CanadaFrench",
	152:  "Chile",
	156:  "China",
	170:  "Columbia",
	203:  "Czechia",
	208:  "Denmark",
	233:  "Estonia",
	246:  "Finland",
	250:  "France",
	276:  "Germany",
	300:  "Greece",
	372:  "Ireland",
	376:  "Israel",
	380:  "Italy",
	442:  "Luxembourg",
	528:  "Netherlands",
	578:  "Norway",
	620:  "Portugal",
	724:  "Spain",
	752:  "Sweden",
	826:  "UnitedKingdom",
	840:  "UnitedStates",
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

// Belgium - Specifies the country of The Kingdom of Belgium.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Belgium() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(56)
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

// Brazil - Specifies the country of Brazil.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Brazil() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(76)
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

// Czechia - Specifies the The Czech Republic
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Czechia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(203)
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

// Finland - Specifies the The Republic of Finland.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Finland() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(246)
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

// Luxembourg - Specifies The Grand Duchy of Luxembourg.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Luxembourg() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(442)
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

// Norway - Specifies the Kingdom of Norway.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Norway() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(578)
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

// Spain - Specifies The Kingdom of Spain.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Spain() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(724)
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

// UnitedKingdom - Specifies the United Kingdom.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) UnitedKingdom() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(826)
}

// UnitedStates - Specifies the United Kingdom.
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) UnitedStates() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(840)
}
