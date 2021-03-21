package numberstr

import (
	"fmt"
	"sync"
)

// NumStrFmtSpecDto - Encapsulates all of the format specifications
// required to format the numeric values contained in type NumStrDto.
//
type NumStrFmtSpecDto struct {
	idNo           uint64
	idString       string
	description    string
	tag            string
	countryCulture FormatterCountry
	absoluteValue  FormatterAbsoluteValue
	currencyValue  FormatterCurrency
	signedNumValue FormatterSignedNumber
	sciNotation    NumStrFmtSpecSciNotationDto
	lock           *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecDto instance  to the data fields of the current
// instance of NumStrFmtSpecDto instance.
//
// If input parameter 'incomingFmtSpecDto' is judged to be invalid,
// this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrFmtSpecDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFmtSpecDto         *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto.
//       The data values in this object will be copied to the
//       current NumStrFmtSpecDto instance.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) CopyIn(
	incomingFmtSpecDto *NumStrFmtSpecDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.CopyIn()")

	nStrFmtSpecDtoElectron :=
		numStrFmtSpecDtoElectron{}

	return nStrFmtSpecDtoElectron.copyIn(
		fmtSpecDto,
		incomingFmtSpecDto,
		ePrefix.XCtx("incomingFmtSpecDto->fmtSpecDto"))
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrFmtSpecDto instance.
//
// If the current NumStrFmtSpecDto instance is judged to be
// invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecDto will be created and returned containing
//       all of the data values copied from the current instance of
//       NumStrFmtSpecDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.CopyOut()")

	nStrFmtSpecDtoElectron :=
		numStrFmtSpecDtoElectron{}

	return nStrFmtSpecDtoElectron.copyOut(
		fmtSpecDto,
		ePrefix.XCtx("Copy Out from 'fmtSpecDto'"))
}

// GetAbsoluteValueSpec - Returns a deep copy of the member variable
// 'absoluteValue', of type FormatterAbsoluteValue.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the FormatterAbsoluteValue
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetAbsoluteValueSpec() FormatterAbsoluteValue {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	absVal,
		_ := fmtSpecDto.absoluteValue.CopyOut(
		ErrPrefixDto{}.Ptr())

	return absVal
}

// GetCountryCulture - Returns a deep copy of the member variable
// 'countryCulture', of type FormatterCountry.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the FormatterCountry object. To
// validate the current NumStrFmtSpecDto instance reference methods
// NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCountryCulture() FormatterCountry {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	countryCulture,
		_ := fmtSpecDto.countryCulture.CopyOut(ErrPrefixDto{}.Ptr())

	return countryCulture
}

// GetCurrencyIntDigitSeparators - Returns the integer digit or
// thousands separators from the 'currency' format specification.
//
// Integer digit separators are most commonly known as the
// 'thousands' separator. In the United States, the 'thousands'
// separator character is the comma (',') and integers are shown
// in groups of three ('3').
//
//    United States Example:  1,000,000,000
//
// The actual value returned is an array of type NumStrIntSeparator
// which contains integer separator characters and grouping
// sequences. This complexity is required in order to support
// countries and cultures with integer groupings other than
// thousands.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//
//  []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation.
//
//         type NumStrIntSeparator struct {
//           intSeparatorChar     rune // Integer separator character
//           intSeparatorGrouping uint // Number of integers in a group
//           intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                        // A zero value signals unlimited repetitions.
//         }
//
//         intSeparatorChar     rune
//         - This separator is commonly known as the 'thousands'
//           separator. It is used to separate groups of integer
//           digits to the left of the decimal separator (a.k.a.
//           decimal point). In the United States, the standard
//           integer digits separator is the comma (','). Other
//           countries use periods, spaces or apostrophes to
//           separate integers.
//             United States Example:  1,000,000,000
//              numSeps.intSeparators =
//                []NumStrIntSeparator{
//                     {
//                     intSeparatorChar:   ',',
//                     intSeparatorGrouping: 3,
//                     intSeparatorRepetitions: 0,
//                     },
//                  }
//
//         intSeparatorGrouping []uint
//         - In most western countries integer digits to the left
//           of the decimal separator (a.k.a. decimal point) are
//           separated into groups of three digits representing
//           a grouping of 'thousands' like this: '1,000,000,000'.
//           In this case the intSeparatorGrouping value would be
//           set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerSeparatorsDto' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              intSeparatorRepetitions: 1,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              intSeparatorRepetitions: 0,
//              },
//           }
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCurrencyIntDigitSeparators(
	ePrefix *ErrPrefixDto) (
	[]NumStrIntSeparator,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.GetCurrencyIntDigitSeparators() ")

	return fmtSpecDto.
		currencyValue.
		numericSeparators.
		GetIntegerDigitSeparators(
			ePrefix.XCtx(
				"currencyValue.numericSeparators"))
}

// GetCurrencyNumericSeparators - Returns a deep copy of the
// NumericSeparators instance currently configured for the
// Currency Format Specification.
//
// The Numeric Separators object contains the decimal separator
// and the integer digit separators.
//
// The integer digit separators is also known as the 'thousands'
// separator. In the United States the standard integer digit
// separator character is the comma (',') and integers are shown
// in groups of three ('3').
//
//    United States Example: 1,000,000,000,000
//
// The returned NumericSeparators object represents the Numeric
// Separator values used to configure the current instance of
// FormatterCurrency.
//
// If the FormatterCurrency or Numeric Separator
// objects are judged to be invalid, this method will return an
// error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be returned through this
//       parameter. This object is a deep copy of the Numeric
//       Separator information used to configure the current
//       instance of FormatterCurrency.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
//       Be advised that if the 'NumericSeparators' is judged
//       invalid, this method will return an error.
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCurrencyNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.GetCurrencyNumericSeparators()")

	return fmtSpecDto.
		currencyValue.GetNumericSeparators(
		ePrefix.XCtx(
			"fmtSpecDto.currencyValue"))
}

// GetCurrencySpec - Returns a deep copy of the member variable
// 'currencyValue', of type FormatterCurrency. This
// is the format specification used in formatting currency values
// within number strings.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the FormatterCountry object. To
// validate the current NumStrFmtSpecDto instance reference methods
// NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCurrencySpec() FormatterCurrency {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	currencySpec,
		_ := fmtSpecDto.currencyValue.CopyOut(nil)

	return currencySpec
}

// GetScientificNotationSpec - Returns a deep copy of the member
// variable 'sciNotation', of type NumStrFmtSpecSciNotationDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecSciNotationDto
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetScientificNotationSpec() NumStrFmtSpecSciNotationDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	scientificNotationSpec,
		_ := fmtSpecDto.sciNotation.CopyOut(ErrPrefixDto{}.Ptr())

	return scientificNotationSpec
}

// GetSignedNumSpec - Returns a deep copy of the member variable
// 'signedNumValue', of type FormatterSignedNumber.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the FormatterSignedNumber
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetSignedNumSpec() FormatterSignedNumber {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	signedNumSpec,
		_ := fmtSpecDto.signedNumValue.CopyOut(ErrPrefixDto{}.Ptr())

	return signedNumSpec
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecDto instance to determine whether the current
// instance is valid in all respects.
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
//       NumStrFmtSpecDto is valid, or not. If the current
//       NumStrFmtSpecDto contains valid data, this method returns
//       'true'. If the data is invalid, this method will return
//       'false'.
//
func (fmtSpecDto *NumStrFmtSpecDto) IsValidInstance() (
	isValid bool) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	isValid,
		_ = nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
		fmtSpecDto,
		ErrPrefixDto{}.Ptr())

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrFmtSpecDto instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx(
		"NumStrFmtSpecDto.IsValidInstanceError()",
		"Testing Validity of 'fmtSpecDto'")

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	_,
		err :=
		nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
			fmtSpecDto,
			ePrefix)

	return err
}

// NewCustomFmtSpec - Creates and returns a new instance of
// NumStrFmtSpecDto. This instance based on the minimum number of
// required input parameters and uses default values to complete
// the custom format specification.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChar          rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        rune
//     - The character which will be used to delimit 'thousands' in
//       integer number strings. In the United States, the Thousands
//       separator is the comma character (',').
//           Example: '1,000,000,000'
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
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
//  currencySymbols               []rune
//     - The authorized unicode character symbol associated with
//       this currency specification. The currency symbol for the
//       United States is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The minor currency name. In the United States, the minor
//       currency name is 'Cent'.
//
//
//  minorCurrencySymbols          []rune
//     - The unicode character for minor currency symbol. In the
//       United States, the minor currency symbol is the cent sign
//       (¢). Some countries and cultures have currency symbols
//       consisting of two or more characters.
//
//
//  currencyPositiveValueFmt      string
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
//  currencyNegativeValueFmt      string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the negative value format string in order
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
//  signedNumPositiveValueFmt     string
//     - A string specifying the number string format to be used in
//       formatting positive numeric values for signed numbers in
//       text strings.
//
//       Positive Value Formatting Terminology and Placeholders:
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
//                        the default for positive number formatting.
//
//       Valid format strings for positive numeric values
//       (NOT Currency) are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Positive Value Format
//
//
//  signedNumNegativeValueFmt     string
//     - A string specifying the number string format to be used in
//       formatting negative numeric values in text strings.
//
//       Negative Value Formatting Terminology and Placeholders:
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
//
//       Valid format strings for negative numeric values
//       (NOT Currency) are listed as follows:
//
//               -127.54   The Default Negative Value Format String
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -127.54
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -NUMFIELD
//               - NUMFIELD
//               NUMFIELD-
//               NUMFIELD -
//               (-) NUMFIELD
//               (-)NUMFIELD
//               NUMFIELD(-)
//               NUMFIELD (-)
//               (NUMFIELD)
//               ( NUMFIELD )
//               (NUMFIELD)
//               ( NUMFIELD )
//
//
//  requestedNumberFieldLen       int
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
//  numberFieldTextJustify        TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to a number string when
//       it is positioned inside of a number field. This
//       enumeration value must be one of the three following
//       format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecDto will be returned containing a fully
//       configured Number String Format Specification.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto NumStrFmtSpecDto) NewCustomFmtSpec(
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	integerDigitsGroupingSequence []uint,
	turnOnIntegerDigitsSeparation bool,
	currencyCode string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	currencyPositiveValueFmt string,
	currencyNegativeValueFmt string,
	signedNumPositiveValueFmt string,
	signedNumNegativeValueFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.NewCustomFmtSpec()")

	newNumStrFmtSpecDto := NumStrFmtSpecDto{}

	newNumStrFmtSpecDto.lock = new(sync.Mutex)

	nStrFmtSpecDtoUtil := numStrFmtSpecDtoUtility{}

	err := nStrFmtSpecDtoUtil.setCustomFmtSpecDto(
		&newNumStrFmtSpecDto,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		integerDigitsGroupingSequence,
		turnOnIntegerDigitsSeparation,
		currencyCode,
		currencyName,
		currencySymbols,
		minorCurrencyName,
		minorCurrencySymbols,
		currencyPositiveValueFmt,
		currencyNegativeValueFmt,
		signedNumPositiveValueFmt,
		signedNumNegativeValueFmt,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	return newNumStrFmtSpecDto, err
}

// NewFromComponents - Creates and returns a new instance of
// NumStrFmtSpecDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  idNo                          uint64
//     - A user created identification number associated with the
//       'nStrFmtSpecDto' object.
//
//
//  idString                      string
//     - User defined text identification associated with the
//       'nStrFmtSpecDto' object.
//
//  description                   string
//     - User defined text description associated with the
//       'nStrFmtSpecDto' object.
//
//
//  tag                           string
//     - User defined tag associated with the 'nStrFmtSpecDto'
//       object.
//
//
//  countryCulture                FormatterCountry
//     - A valid and fully populated FormatterCountry
//       object. This object contains information on the country
//       or culture related to this number string format.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  absoluteValue                 FormatterAbsoluteValue
//     - A valid and fully populated FormatterAbsoluteValue
//       object. This object contains formatting specifications
//       controlling the text display of absolute numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  currencyValue                 FormatterCurrency
//     - A valid and fully populated FormatterCurrency
//       object. This object contains formatting specifications
//       controlling the text display of currency number strings.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  signedNumValue                FormatterSignedNumber
//     - A valid and fully populated FormatterSignedNumber
//       object. This object contains formatting specifications
//       controlling the text display of signed numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  sciNotation                   NumStrFmtSpecSciNotationDto
//     - A valid and fully populated FormatterCountry
//       object. This object contains formatting specifications
//       controlling the text display of scientific notation.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto NumStrFmtSpecDto) NewFromComponents(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture FormatterCountry,
	absoluteValue FormatterAbsoluteValue,
	currencyValue FormatterCurrency,
	signedNumValue FormatterSignedNumber,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.NewWithComponents()")

	newFmtSpecDto := NumStrFmtSpecDto{}

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	newFmtSpecDto.lock = new(sync.Mutex)

	err := nStrFmtSpecDtoMech.setNumStrFmtSpecDto(
		&newFmtSpecDto,
		idNo,
		idString,
		description,
		tag,
		countryCulture,
		absoluteValue,
		currencyValue,
		signedNumValue,
		sciNotation,
		ePrefix)

	return newFmtSpecDto, err
}

// NewFromFmtSpecSetupDto - Creates and returns a new
// NumStrFmtSpecDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     *NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecDto object.
//
//       type NumStrFmtSpecSetupDto struct {
//         IdNo                                      uint64
//         IdString                                  string
//         Description                               string
//         Tag                                       string
//         CountryIdNo                               uint64
//         CountryIdString                           string
//         CountryDescription                        string
//         CountryTag                                string
//         CountryCultureName                        string
//         CountryAbbreviatedName                    string
//         CountryAlternateNames                     []string
//         CountryCodeTwoChar                        string
//         CountryCodeThreeChar                      string
//         CountryCodeNumber                         string
//         AbsoluteValFmt                            string
//         AbsoluteValTurnOnIntegerDigitsSeparation  bool
//         AbsoluteValNumFieldLen                    int
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     int
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumFieldLen                       int
//         DecimalSeparator                          rune
//         IntegerDigitsSeparator                    rune
//         IntegerDigitsGroupingSequence             []uint
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumFieldLen                   int
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         Lock                                      *sync.Mutex
//       }
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecDto will be returned to the caller.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto NumStrFmtSpecDto) NewFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.NewWithFmtSpecSetupDto()")

	newNumStrFmtSpecDto :=
		NumStrFmtSpecDto{}

	newNumStrFmtSpecDto.lock = new(sync.Mutex)

	if fmtSpecSetupDto == nil {
		return newNumStrFmtSpecDto,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		&newNumStrFmtSpecDto,
		fmtSpecSetupDto,
		ePrefix.XCtx("newNumStrFmtSpecDto"))

	return newNumStrFmtSpecDto, err
}

// SetFromFmtSpecSetupDto - Overwrites and sets the data values
// for the current NumStrFmtSpecDto instance based on the input
// data values passed in a NumStrFmtSpecSetupDto structure.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup and configuration
//       information for a NumStrFmtSpecDto object.
//
//       type NumStrFmtSpecSetupDto struct {
//         IdNo                                      uint64
//         IdString                                  string
//         Description                               string
//         Tag                                       string
//         CountryIdNo                               uint64
//         CountryIdString                           string
//         CountryDescription                        string
//         CountryTag                                string
//         CountryCultureName                        string
//         CountryAbbreviatedName                    string
//         CountryAlternateNames                     []string
//         CountryCodeTwoChar                        string
//         CountryCodeThreeChar                      string
//         CountryCodeNumber                         string
//         AbsoluteValFmt                            string
//         AbsoluteValTurnOnIntegerDigitsSeparation  bool
//         AbsoluteValNumFieldLen                    int
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     int
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumFieldLen                       int
//         DecimalSeparator                          rune
//         IntegerDigitsSeparator                    rune
//         IntegerDigitsGroupingSequence             []uint
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumFieldLen                   int
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//       }
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) SetFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.NewWithFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	ePrefix.SetCtx("fmtSpecDto")

	return nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		fmtSpecDto,
		fmtSpecSetupDto,
		ePrefix)
}

// SetNumStrFmtSpecDto - Overwrites and sets the data values
// for the current NumStrFmtSpecDto instance based on component
// values passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  idNo                          uint64
//     - A user created identification number associated with the
//       'nStrFmtSpecDto' object.
//
//
//  idString                      string
//     - User defined text identification associated with the
//       'nStrFmtSpecDto' object.
//
//  description                   string
//     - User defined text description associated with the
//       'nStrFmtSpecDto' object.
//
//
//  tag                           string
//     - User defined tag associated with the 'nStrFmtSpecDto'
//       object.
//
//
//  countryCulture                FormatterCountry
//     - A valid and fully populated FormatterCountry
//       object. This object contains information on the country
//       or culture related to this number string format.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  absoluteValue                 FormatterAbsoluteValue
//     - A valid and fully populated FormatterAbsoluteValue
//       object. This object contains formatting specifications
//       controlling the text display of absolute numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  currencyValue                 FormatterCurrency
//     - A valid and fully populated FormatterCurrency
//       object. This object contains formatting specifications
//       controlling the text display of currency number strings.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  signedNumValue                FormatterSignedNumber
//     - A valid and fully populated FormatterSignedNumber
//       object. This object contains formatting specifications
//       controlling the text display of signed numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  sciNotation                   NumStrFmtSpecSciNotationDto
//     - A valid and fully populated FormatterCountry
//       object. This object contains formatting specifications
//       controlling the text display of scientific notation.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) SetNumStrFmtSpecDto(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture FormatterCountry,
	absoluteValue FormatterAbsoluteValue,
	currencyValue FormatterCurrency,
	signedNumValue FormatterSignedNumber,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.SetNumStrFmtSpecDto()")

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setNumStrFmtSpecDto(
		fmtSpecDto,
		idNo,
		idString,
		description,
		tag,
		countryCulture,
		absoluteValue,
		currencyValue,
		signedNumValue,
		sciNotation,
		ePrefix)

	return err
}

// SetToDefault - Sets the current instance of Number String Format
// Specification Data Transfer Object (NumStrFmtSpecDto) to its
// default value.
//
// The default Number String Format Specification is the United
// States decimal separator ('.'), thousands separator (',') and
// currency symbol ('$').
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) SetToDefault(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.SetToDefault() ")

	err = numStrFmtSpecDtoUtility{}.ptr().
		setDefaultFormatSpec(
			fmtSpecDto,
			ePrefix)

	return err
}

// SetToDefaultIfEmpty - If the current instance of Number String
// Format Specification Data Transfer Object (NumStrFmtSpecDto) is
// empty or invalid, this method will reset the Format
// Specification to its default value.
//
// The default Number String Format Specification is the United
// States decimal separator ('.'), thousands separator (',') and
// currency symbol ('$').
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) SetToDefaultIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.SetToDefaultIfEmpty() ")

	var isValid bool

	isValid,
		_ = numStrFmtSpecDtoQuark{}.ptr().testValidityOfNumStrFmtSpecDto(
		fmtSpecDto,
		ePrefix.XCtx(
			"Testing validity of fmtSpecDto"))

	if !isValid {
		err = numStrFmtSpecDtoUtility{}.ptr().
			setDefaultFormatSpec(
				fmtSpecDto,
				ePrefix.XCtx("fmtSpecDto"))
	}

	return err
}
