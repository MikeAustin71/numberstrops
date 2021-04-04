package numberstr

import "sync"

type NumStrBasic struct {
	lock *sync.Mutex
}

// GetCountryFormatters - Receives a pointer to a Formatter
// Collection and a Country/Culture ID. Using these parameters the
// formatter collection is populated with a series of number string
// formatter specifications associated with the country or culture
// designated by the input parameter 'countryCulture'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtCollection       *FormatterCollection
//     - A pointer to an instance of FormatterCollection. This
//       object encapsulates multiple format specifications (a.k.a
//       'formatters) required to format the numeric values
//       contained in type NumStrDto.
//
//       This method will populate this formatter collection with
//       a series of format specifications associated with the
//       country or culture identified by input parameter,
//       'countryCulture'.
//
//       Currently, supported formatters include, binary, signed
//       number, absolute value, currency, hexadecimal and octal.
//
//
//  countryCulture      CountryCultureId
func (numStrBasic *NumStrBasic) GetCountryFormatters(
	fmtCollection *FormatterCollection,
	countryCulture CountryCultureId,
	errorPrefix IBasicErrorPrefix) error {

	if numStrBasic.lock == nil {
		numStrBasic.lock = new(sync.Mutex)
	}

	numStrBasic.lock.Lock()

	defer numStrBasic.lock.Unlock()

	var ePrefix *ErrPrefixDto

	if errorPrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {

		ePrefix = ErrPrefixDto{}.NewFromIBasicErrorPrefix(errorPrefix)
	}

	ePrefix.SetEPref(
		"NumStrBasic." +
			"GetCountryFormatters()")

	return numStrBasicQuark{}.ptr().getCountryFormatters(
		fmtCollection,
		countryCulture,
		ePrefix)
}
