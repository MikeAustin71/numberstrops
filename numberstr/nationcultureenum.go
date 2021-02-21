package numberstr

import "sync"

var mCountryCultureCodeToString = map[int]string{
	-1:   "None",
	32:   "Argentina",
	36:   "Australia",
	124:  "Canada",
	9124: "CanadaFrench",
	152:  "Chile",
	203:  "Czechia",
	156:  "China",
	170:  "Columbia",
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

// Czechia - Specifies the The Czech Republic
//
// This method is part of the standard enumeration.
//
func (cntryCulId CountryCultureId) Czechia() CountryCultureId {

	lockCountryCultureId.Lock()

	defer lockCountryCultureId.Unlock()

	return CountryCultureId(203)
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
