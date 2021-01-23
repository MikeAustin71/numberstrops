package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFmtSpecCurrencyValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	decimalDigits                 uint
	currencyCode                  string
	currencyName                  string
	currencySymbol                rune
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumStrFmtSpecDigitsSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecCurrencyValueDto instance to the data fields
// of the current NumStrFmtSpecCurrencyValueDto instance.
//
// If input parameter 'incomingCurrencyValDto' is judged to be
// invalid, this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrFmtSpecCurrencyValueDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingCurrencyValDto     *NumStrFmtSpecCurrencyValueDto
//     - A pointer to an instance of NumStrFmtSpecCurrencyValueDto.
//       The data values in this object will be copied to the
//       current NumStrFmtSpecCurrencyValueDto instance.
//
//
//  ePrefix                    string
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will encapsulate
//       an error message. The input parameter, 'ePrefix', will be
//       prefixed and inserted at the beginning of the returned
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) CopyIn(
	incomingCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix string) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecCurrencyValueDto.CopyIn() "

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	return nStrFmtSpecCurrValNanobot.copyIn(
		nStrFmtSpecCurrValDto,
		incomingCurrencyValDto,
		ePrefix+
			"\nincomingCurrencyValDto->nStrFmtSpecCurrValDto ")
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrFmtSpecCurrencyValueDto instance.
//
// If the current NumStrFmtSpecCurrencyValueDto instance is judged
// to be invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
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
//  NumStrFmtSpecCurrencyValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecCurrencyValueDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrFmtSpecCurrencyValueDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will encapsulate
//       an error message. The input parameter, 'ePrefix', will be
//       prefixed and inserted at the beginning of the returned
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) CopyOut(
	ePrefix string) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecCurrencyValueDto.CopyOut()\n "

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	return nStrFmtSpecCurrValNanobot.copyOut(
		nStrFmtSpecCurrValDto,
		ePrefix+
			"Copy Out from 'nStrFmtSpecCurrValDto'\n ")
}

// GetNegativeValueFormat - Returns the formatting string used to
// format negative currency values in text number strings.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetNegativeValueFormat() string {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.negativeValueFmt
}

// GetCurrencyCode - Returns the ISO 4217 Currency Code associated
// with this currency specification.
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetCurrencyCode() string {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.currencyCode
}

// GetCurrencyName - Returns the official name of this
// currency. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetCurrencyName() string {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.currencyName
}

// GetCurrencySymbol - Returns the currency symbol associated
// with this currency. Reference:
//        https://www.xe.com/symbols.php
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetCurrencySymbol() rune {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.currencySymbol
}

// GetDecimalDigits - Returns the standard number of
// digits to the right of the decimal point formatted
// in currency number strings.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetDecimalDigits() uint {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.decimalDigits
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Currency Value Dto.
//
// The NumberFieldDto details the length of the number field in which
// the signed numeric value will be displayed and right justified.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetNumberFieldLengthDto() NumberFieldDto {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.numFieldLenDto.CopyOut()
}

// GetNumberSeparatorsDto - Returns the NumStrFmtSpecDigitsSeparatorsDto
// instance currently configured for this Number String Format Specification
// Currency Value Dto.
//
// The Digits Separator Dto contains the decimal separator as well as the
// 'thousands' separators and the grouping sequence for separating thousands
// digits in the integer component of a number string.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetNumberSeparatorsDto() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.numberSeparatorsDto.CopyOut()
}

// GetPositiveValueFormat - Returns the formatting string used to
// format positive currency values in text number strings.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetPositiveValueFormat() string {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.positiveValueFmt
}

// GetTurnOnIntegerDigitsSeparationFlag - Returns the boolean flag
// which signals whether integer digits within a number string will
// be grouped by thousands and separated by an integer digits
// separator such as a comma (',').
//
// If this flag is set to 'true', integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1,000,000,000,000
//
// If this flag is set to 'false',  integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1000000000000
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecCurrencyValueDto instance to determine whether
// the current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the current
//       NumStrFmtSpecCurrencyValueDto is valid, or not. If the
//       current NumStrFmtSpecCurrencyValueDto contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	isValid,
		_ = nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
		nStrFmtSpecCurrValDto,
		"")

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecSignedNumValueDto instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumStrFmtSpecCurrencyValueDto
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecCurrencyValueDto.IsValidInstanceError()\n" +
		"Testing Validity of 'nStrFmtSpecCurrValDto'\n"

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err := nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
		nStrFmtSpecCurrValDto,
		ePrefix)

	return err
}

// New() - Creates and returns a new instance of
// NumStrFmtSpecCurrencyValueDto.
//
// This variant of the 'NumStrFmtSpecCurrencyValueDto.New()'
// method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'NumStrFmtSpecCurrencyValueDto.NewFromComponents()'.
//
// The NumStrFmtSpecCurrencyValueDto type encapsulates the
// formatting parameters necessary to format currency values for
// display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  negativeValueFmt              string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbol                rune
//     - The authorized character symbol associated with this
//       currency specification.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  decimalSeparatorChar       rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar     rune
//     - The character which will be used to delimit 'thousands' in
//       integer number strings. In the United States, the Thousands
//       separator is the comma character (',').
//           Example: '1,000,000,000'
//
//
//  requestedNumberFieldLen    int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the the
//       number field. If the actual number string length is greater
//       than the requested number field length, the number field
//       length will be automatically expanded to display the entire
//       number string. The 'requested' number field length is used
//       to create number fields of standard lengths for text
//       presentations.
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecCurrencyValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecCurrencyValueDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto NumStrFmtSpecCurrencyValueDto) New(
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyName string,
	currencySymbol rune,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.New() "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	newNStrFmtSpecCurrencyValDto := NumStrFmtSpecCurrencyValueDto{}

	numberSeparatorsDto,
		err := NumStrFmtSpecDigitsSeparatorsDto{}.New(
		decimalSeparatorChar,
		thousandsSeparatorChar,
		[]uint{3},
		ePrefix)

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	err = nStrFmtSpecCurrValMech.setCurrencyValDto(
		&newNStrFmtSpecCurrencyValDto,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyName,
		currencySymbol,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'newNStrFmtSpecCurrencyValDto'\n ")

	return newNStrFmtSpecCurrencyValDto, err
}

// NewFromComponents() - Creates and returns a new instance of
// NumStrFmtSpecSignedNumValueDto.
//
// This type encapsulates the formatting parameters necessary to
// format signed numeric values for display in text number
// strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  negativeValueFmt              string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbol                rune
//     - The authorized character symbol associated with this
//       currency specification.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
//  numberSeparatorsDto        NumStrFmtSpecDigitsSeparatorsDto
//     - This instance of 'NumStrFmtSpecDigitsSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumStrFmtSpecDigitsSeparatorsDto struct {
//         decimalSeparator              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparator rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerDigitsSeparator rune
//
//        This type also encapsulates the integer digits separator, often
//        referred to as the 'Thousands Separator'. This is used to
//        separate thousands digits within the integer component of a
//        number string.
//
//        integerDigitsGroupingSequence []uint
//
//        Related to the integer digits separator, the integer digits
//        grouping sequence is also encapsulated in this type. The integer
//        digits grouping sequence is used to identify the digits which
//        will be grouped and separated by the integer digits separator.
//
//        In most western countries integer digits to the left of the
//        decimal separator (a.k.a. decimal point) are separated into
//        groups of three digits representing a grouping of 'thousands'
//        like this: '1,000,000,000,000'. In this case the parameter
//        integer digits grouping sequence would be configured as:
//                     integerDigitsGroupingSequence = []uint{3}
//
//        In some countries and cultures other integer groupings are used.
//        In India, for example, a number might be formatted as like this:
//                      '6,78,90,00,00,00,00,000'
//        The right most group has three digits and all the others are
//        grouped by two. In this case integer digits grouping sequence
//        would be configured as:
//                     integerDigitsGroupingSequence = []uint{3,2}
//
//
//  requestedNumberFieldLen    int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the the
//       number field. If the actual number string length is greater
//       than the requested number field length, the number field
//       length will be automatically expanded to display the entire
//       number string. The 'requested' number field length is used
//       to create number fields of standard lengths for text
//       presentations.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecSignedNumValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecSignedNumValueDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto NumStrFmtSpecCurrencyValueDto) NewFromComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyName string,
	currencySymbol rune,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.NewFromComponents() "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	newNStrFmtSpecCurrencyValueDto := NumStrFmtSpecCurrencyValueDto{}

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	err := nStrFmtSpecCurrValMech.setCurrencyValDto(
		&newNStrFmtSpecCurrencyValueDto,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyName,
		currencySymbol,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'newNStrFmtSpecCurrencyValueDto'\n ")

	return newNStrFmtSpecCurrencyValueDto, err
}

// SetCurrencyValDto() - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecCurrencyValueDto.
//
// The NumStrFmtSpecCurrencyValueDto type encapsulates the
// formatting parameters necessary for formatting currency values
// in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  negativeValueFmt              string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbol                rune
//     - The authorized character symbol associated with this
//       currency specification.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
//  numberSeparatorsDto        NumStrFmtSpecDigitsSeparatorsDto
//     - This instance of 'NumStrFmtSpecDigitsSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumStrFmtSpecDigitsSeparatorsDto struct {
//         decimalSeparator              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparator rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerDigitsSeparator rune
//
//        This type also encapsulates the integer digits separator, often
//        referred to as the 'Thousands Separator'. This is used to
//        separate thousands digits within the integer component of a
//        number string.
//
//        integerDigitsGroupingSequence []uint
//
//        Related to the integer digits separator, the integer digits
//        grouping sequence is also encapsulated in this type. The integer
//        digits grouping sequence is used to identify the digits which
//        will be grouped and separated by the integer digits separator.
//
//        In most western countries integer digits to the left of the
//        decimal separator (a.k.a. decimal point) are separated into
//        groups of three digits representing a grouping of 'thousands'
//        like this: '1,000,000,000,000'. In this case the parameter
//        integer digits grouping sequence would be configured as:
//                     integerDigitsGroupingSequence = []uint{3}
//
//        In some countries and cultures other integer groupings are used.
//        In India, for example, a number might be formatted as like this:
//                      '6,78,90,00,00,00,00,000'
//        The right most group has three digits and all the others are
//        grouped by two. In this case integer digits grouping sequence
//        would be configured as:
//                     integerDigitsGroupingSequence = []uint{3,2}
//
//
//  requestedNumberFieldLen    int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the the
//       number field. If the actual number string length is greater
//       than the requested number field length, the number field
//       length will be automatically expanded to display the entire
//       number string. The 'requested' number field length is used
//       to create number fields of standard lengths for text
//       presentations.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencyValDto(
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyName string,
	currencySymbol rune,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.SetCurrencyValDto() "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	return nStrFmtSpecCurrValMech.setCurrencyValDto(
		nStrFmtSpecCurrValDto,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyName,
		currencySymbol,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'nStrFmtSpecCurrValDto'\n ")
}

// SetCurrencyCode - Sets the currency code associated with
// this currency. Currency codes are designated by the ISO
// 4217 standard. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencyCode(
	currencyCode string) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrValDto.currencyCode = currencyCode
}

// SetCurrencyName - Sets the currency name associated with
// this currency.
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencyName(
	currencyName string) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrValDto.currencyName = currencyName
}

// SetCurrencySymbol - Sets the currency name associated with
// this currency.
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencySymbol(
	currencySymbol rune,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecCurrencyValueDto.SetCurrencySymbol()"

	err = nil

	if currencySymbol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbol' is invalid!\n"+
			"The currency symbol is empty or missing.\n"+
			"currencySymbol==0\n",
			ePrefix)
		return err
	}

	nStrFmtSpecCurrValDto.currencySymbol = currencySymbol

	return err
}

// SetDecimalDigits - Sets the number of digits to the right of the
// decimal point which are typically included in currency text
// number string displays. For example, in the United States,
// currency is typically displayed with two digits to the right
// of the decimal point.
//   Example: $27.94
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetDecimalDigits(
	decimalDigits uint) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrValDto.decimalDigits = decimalDigits

}

// SetNegativeValueFormat - Sets the currency negative value format
// string. This string contains the formatting symbols used to
// format negative currency values in number strings.
//
// If the negative value format string input parameter,
// 'negativeValueFmt', is invalid, this method will return an
// error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  negativeValueFmt    string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetNegativeValueFormat(
	negativeValueFmt string,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.SetNegativeValueFormat()\n "

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		negativeValueFmt,
		ePrefix+
			"\nTesting validity of 'negativeValueFmt'\n ")

	if err != nil {
		return err
	}

	nStrFmtSpecCurrValDto.negativeValueFmt =
		negativeValueFmt

	return err
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current NumStrFmtSpecCurrencyValueDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLenDto  NumberFieldDto
//     - The NumberFieldDto details the length of the number field
//       in which the signed numeric value will be displayed and
//       right justified.
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int
//         actualNumFieldLength    int
//         minimumNumFieldLength   int
//       }
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.SetNumberFieldLengthDto()\n"

	err =
		nStrFmtSpecCurrValDto.numFieldLenDto.CopyIn(
			&numberFieldLenDto,
			ePrefix)

	return err
}

// SetNumberSeparatorsDto - Sets the Number Separators Dto object
// for the current NumStrFmtSpecCurrencyValueDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberSeparatorsDto        NumStrFmtSpecDigitsSeparatorsDto
//     - This instance of 'NumStrFmtSpecDigitsSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumStrFmtSpecDigitsSeparatorsDto struct {
//         decimalSeparator              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparator rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerDigitsSeparator rune
//
//        This type also encapsulates the integer digits separator, often
//        referred to as the 'Thousands Separator'. This is used to
//        separate thousands digits within the integer component of a
//        number string.
//
//        integerDigitsGroupingSequence []uint
//
//        Related to the integer digits separator, the integer digits
//        grouping sequence is also encapsulated in this type. The integer
//        digits grouping sequence is used to identify the digits which
//        will be grouped and separated by the integer digits separator.
//
//        In most western countries integer digits to the left of the
//        decimal separator (a.k.a. decimal point) are separated into
//        groups of three digits representing a grouping of 'thousands'
//        like this: '1,000,000,000,000'. In this case the parameter
//        integer digits grouping sequence would be configured as:
//                     integerDigitsGroupingSequence = []uint{3}
//
//        In some countries and cultures other integer groupings are used.
//        In India, for example, a number might be formatted as like this:
//                      '6,78,90,00,00,00,00,000'
//        The right most group has three digits and all the others are
//        grouped by two. In this case integer digits grouping sequence
//        would be configured as:
//                     integerDigitsGroupingSequence = []uint{3,2}
//
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetNumberSeparatorsDto(
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	_ = nStrFmtSpecCurrValDto.numberSeparatorsDto.CopyIn(
		&numberSeparatorsDto,
		"")
}

// SetPositiveValueFormat - Sets the currency positive value format
// string. This string contains the formatting symbols used to
// format positive currency values in number strings.
//
// If the positive value string input parameter, 'positiveValueFmt',
// is invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt    string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetPositiveValueFormat(
	positiveValueFmt string,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.SetPositiveValueFormat()\n "

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix+
			"\nTesting validity of 'positiveValueFmt'\n ")

	if err != nil {
		return err
	}

	nStrFmtSpecCurrValDto.positiveValueFmt =
		positiveValueFmt

	return err
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// NumStrFmtSpecCurrencyValueDto.
//
// This boolean flag signals whether integer digits within a number
// string will be grouped by thousands and separated by an integer
// digits separator such as a comma (',').
//
// If this flag is set to 'true', integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1,000,000,000,000
//
// If this flag is set to 'false',  integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1000000000000
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}
