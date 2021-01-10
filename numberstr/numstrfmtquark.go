package numberstr

import "sync"

type numStrFormatQuark struct {
	lock *sync.Mutex
}

// getValidFormatChars - Returns an array of runes which designate
// the valid formatting characters which may used in number string
// formats.
func (nStrFmtQuark *numStrFormatQuark) getValidFormatChars() []rune {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	validFmtChars := []rune("(-) +$127.54NUMFIELD")

	return validFmtChars
}

// getDefaultPositiveNumStrFormat - Returns the default format for
// number strings displaying positive values.

func (nStrFmtQuark *numStrFormatQuark) getDefaultPositiveNumStrFormat() string {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	return "127.54"
}

// getDefaultPositiveNumStrFormat - Returns the default format for
// number strings displaying negative values.

func (nStrFmtQuark *numStrFormatQuark) getDefaultNegativeNumStrFormat() string {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	return "-127.54"
}

// getDefaultPositiveNumStrFormat - Returns the default decimal
// separator character for number string formatting.
//
// The default character is the decimal point ('.').
//   Example: 123.456
//
func (nStrFmtQuark *numStrFormatQuark) getDefaultDecimalSeparator() rune {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	return '.'
}

// getDefaultThousandsSeparator - Returns the default thousands
// separator character for number string formatting.
//
// The default character is the comma (',').
//   Example: 1,000,000,000,000
//
func (nStrFmtQuark *numStrFormatQuark) getDefaultThousandsSeparator() rune {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	return '.'
}

// getDefaultCurrencySymbol - Returns the default Currency information
// encapsulated in a new instance of CurrencySymbolDto. The default
// for currency is the United States Dollar.
//
func (nStrFmtQuark *numStrFormatQuark) getDefaultCurrencySymbol() CurrencySymbolDto {

	if nStrFmtQuark.lock == nil {
		nStrFmtQuark.lock = new(sync.Mutex)
	}

	nStrFmtQuark.lock.Lock()

	defer nStrFmtQuark.lock.Unlock()

	defaultCurrSym := CurrencySymbolDto{}

	defaultCurrSym.idNo = 1
	defaultCurrSym.currencySymbol = '$'
	defaultCurrSym.currencyCode = "USD"
	defaultCurrSym.currencyName = "Dollar"
	defaultCurrSym.countryName = "United States"
	defaultCurrSym.abbreviatedCountryName = "USA"
	defaultCurrSym.alternateCountryName = "United States of America"
	defaultCurrSym.countryCodeTwoChar = "US"
	defaultCurrSym.countryCodeThreeChar = "USA"
	defaultCurrSym.countryCodeNumber = "840"
	defaultCurrSym.positiveValCurrFormat = "$127.54"
	defaultCurrSym.negativeValCurrFormat = "$-127.54"

	defaultCurrSym.lock = new(sync.Mutex)

	return defaultCurrSym
}
