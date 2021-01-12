package numberstr

import "sync"

type NumStrFormatter struct {
	valueDisplaySpec              NumStrValSpec
	idNo                          uint64
	idString                      string
	description                   string
	tag                           string
	countryName                   string // Full Name of country associated with this currency
	abbreviatedCountryName        string // Abbreviated Country Name
	alternateCountryName          string // Alternate Country Name Example: United States of America
	countryCodeTwoChar            string // ISO-3166 Two Character Country Code https://en.wikipedia.org/wiki/ISO_3166-2
	countryCodeThreeChar          string // ISO-3166 Three Character Country Code https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	countryCodeNumber             string // ISO-3166 Country Code Number: 3-Numeric Digits https://en.wikipedia.org/wiki/ISO_3166-1_numeric
	positiveValueFmt              string
	negativeValueFmt              string
	decimalSeparator              rune
	currencySymbol                rune
	currencyDecimalDigits         int    // The number of digits after the decimal separator for currency presentations
	currencyCode                  string // Wold Currency Code (3-Characters). Reference: http://www.xe.com/symbols.php
	currencyName                  string // The common name of this currency i.e 'Dollar', 'Yuan' etc.
	integerDigitsGroupingSequence []uint // Integer Digit Separation. Usually = 3 for thousands separation
	integerDigitsSeparator        rune   // a.k.a. Thousands Separator
	turnOnIntegerDigitsSeparation bool
	sciNotMantissaLength          uint // The length of the fractional digits in the significand which will be displayed
	sciNotExponentChar            rune // Usually 'e' or 'E'
	sciNotExponentUsesLeadingPlus bool // If true, positive exponent values are prefixed with a leading plus (+) sign. '2.652e+8'
	numFieldDto                   numberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies all data elements from 'nStrFormtrIncoming'
// into the current NumStrFormatter instance.
//
func (nStrFormatter *NumStrFormatter) CopyIn(
	nStrFormtrIncoming *NumStrFormatter,
	ePrefix string) error {

	if nStrFormatter.lock == nil {
		nStrFormatter.lock = new(sync.Mutex)
	}

	nStrFormatter.lock.Lock()

	defer nStrFormatter.lock.Unlock()

	ePrefix += "NumStrFormatter.CopyIn() "

	nStrFormtrQuark := numStrFormatterQuark{}

	return nStrFormtrQuark.copyIn(
		nStrFormatter,
		nStrFormtrIncoming,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current NumStrFormatter
// instance.
//
func (nStrFormatter *NumStrFormatter) CopyOut() NumStrFormatter {

	if nStrFormatter.lock == nil {
		nStrFormatter.lock = new(sync.Mutex)
	}

	nStrFormatter.lock.Lock()

	defer nStrFormatter.lock.Unlock()

	nStrFormtrQuark := numStrFormatterQuark{}

	newNStrFormtr,
		_ := nStrFormtrQuark.copyOut(
		nStrFormatter,
		"")

	return newNStrFormtr
}

// IsValidInstanceError - Analyzes the current NumStrFormatter
// instance and returns a boolean value of 'false' if any of the
// member data elements contain invalid data.
//
func (nStrFormatter *NumStrFormatter) IsValidInstance() bool {

	if nStrFormatter.lock == nil {
		nStrFormatter.lock = new(sync.Mutex)
	}

	nStrFormatter.lock.Lock()

	defer nStrFormatter.lock.Unlock()

	nStrFormatterMech := numStrFormatterMechanics{}

	isValid,
		_ := nStrFormatterMech.testValidityOfNumStrFormatter(
		nStrFormatter,
		"")

	return isValid
}

// IsValidInstanceError - Analyzes the current NumStrFormatter
// instance and returns an error if any of the member data elements
// contain invalid data.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the current NumStrFormatter instance contains invalid
//       data values, this method will return a detailed error
//       message describing the nature of the data error.
//
func (nStrFormatter *NumStrFormatter) IsValidInstanceError(
	ePrefix string) (
	err error) {

	if nStrFormatter.lock == nil {
		nStrFormatter.lock = new(sync.Mutex)
	}

	nStrFormatter.lock.Lock()

	defer nStrFormatter.lock.Unlock()

	ePrefix += "NumStrFormatter.IsValidInstanceError() "

	nStrFormatterMech := numStrFormatterMechanics{}

	_,
		err = nStrFormatterMech.testValidityOfNumStrFormatter(
		nStrFormatter,
		ePrefix)

	return err
}
