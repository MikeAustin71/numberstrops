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
	currencyCodeNo                string
	currencyName                  string
	currencySymbols               []rune
	minorCurrencyName             string
	minorCurrencySymbols          []rune
	turnOnIntegerDigitsSeparation bool
	numericSeparators             NumericSeparators
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
//       If input parameter 'incomingCurrencyValDto' is judged to
//       be invalid, this method will return an error.
//
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) CopyIn(
	incomingCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.CopyIn()")

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	return nStrFmtSpecCurrValNanobot.copyIn(
		nStrFmtSpecCurrValDto,
		incomingCurrencyValDto,
		ePrefix.XCtx("incomingCurrencyValDto->nStrFmtSpecCurrValDto"))
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
//  NumStrFmtSpecCurrencyValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecCurrencyValueDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrFmtSpecCurrencyValueDto.
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.CopyOut()")

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	return nStrFmtSpecCurrValNanobot.copyOut(
		nStrFmtSpecCurrValDto,
		ePrefix.XCtx("Copy Out from 'nStrFmtSpecCurrValDto'"))
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

// GetCurrencyCodeNo - Returns the ISO 4217 Currency Code Number
// associated with this currency specification. The Currency Code
// Number is stored as a string per ISO 4217.
//
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetCurrencyCodeNo() string {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.currencyCodeNo
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

// GetCurrencySymbols - Returns the currency symbols associated
// with this currency. The currency symbol for the United States is
// the dollar sign ('$'). Some countries and cultures have currency
// symbols consisting of two or more characters.
//
// Reference:
//        https://www.xe.com/symbols.php
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetCurrencySymbols() []rune {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.currencySymbols
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

// GetDecimalSeparator - Returns the unicode character (rune) used
// to separate integer and fractional digits in a floating point
// number.
//
// In the United States, the Decimal Separator character is the
// decimal point or period ('.').
//
//  Example:   123.45
//
// Decimal Separator is extracted from the underlying member
// variable, 'nStrFmtSpecCurrValDto.numericSeparators'.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetDecimalSeparator() rune {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.
		numericSeparators.
		GetDecimalSeparator()
}

// GetIntegerDigitSeparators - Returns an array of type
// NumStrIntSeparator. The data contained in type
// NumStrIntSeparator is used to separate integer digits.
//
// The returned integer digit separators are those configured
// for the current instance of NumStrFmtSpecCurrencyValueDto.
//
// The integer digit separators is also known as the 'thousands'
// separator. In the United States the standard integer digit
// separator character is the comma (',') and integers are shown
// in groups of three ('3').
//
//    United States Example: 1,000,000,000,000
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
//  []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation.
//
//        type NumStrIntSeparator struct {
//          intSeparatorChar     rune
//          intSeparatorGrouping uint
//        }
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
//       this case 'integerSeparators' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              },
//           }
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetIntegerDigitSeparators(
	ePrefix *ErrPrefixDto) (
	[]NumStrIntSeparator,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"GetIntegerDigitSeparators()")

	return nStrFmtSpecCurrValDto.
		numericSeparators.
		GetIntegerDigitSeparators(
			ePrefix.XCtx(
				"nStrFmtSpecCurrValDto.numericSeparators"))
}

// GetMinorCurrencySymbols - Returns the minor currency symbols.
//
// The authorized unicode character symbols associated with the
// minor currency specification. The minor currency symbol for
// the United States is the cent sign ('¢'). Some countries and
// cultures have minor currency symbols consisting of two or more
// characters.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetMinorCurrencySymbols() []rune {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	return nStrFmtSpecCurrValDto.minorCurrencySymbols
}

// GetNumberFieldLengthDto - Returns a deep copy of the
// NumberFieldDto object/ currently configured for this Number
// String Format Specification Currency Value Dto.
//
// The NumberFieldDto details the length of the number field in
// which the signed numeric value will be displayed and right
// justified.
//
// If the NumberFieldDto object is judged to be invalid, this
// method will return an error.
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
//  NumberFieldDto
//     - If this method completes successfully, a new instance of
//       NumberFieldDto will be returned through this parameter.
//       This object is deep copy of the Number Field information
//       used to configure the current instance of
//       NumStrFmtSpecCurrencyValueDto.
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
//       Be advised that if the returned 'NumberFieldDto' object is
//       judged invalid, this method will return an error.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.GetNumberFieldLengthDto()")

	return nStrFmtSpecCurrValDto.numFieldLenDto.CopyOut(
		ePrefix.XCtx("nStrFmtSpecCurrValDto.numFieldLenDto->"))
}

// GetNumericSeparators - Returns a deep copy of the
// NumericSeparators instance currently configured for this
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
// NumStrFmtSpecCurrencyValueDto.
//
// If the NumStrFmtSpecCurrencyValueDto or NumericSeparators object
// is judged to be invalid, this method will return an error.
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
//       instance of NumStrFmtSpecCurrencyValueDto.
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
//       If the NumStrFmtSpecCurrencyValueDto or NumericSeparators
//       object is judged to be invalid, this method will return
//       an error.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) GetNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"GetNumericSeparators()")

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err := nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
		nStrFmtSpecCurrValDto,
		ePrefix)

	if err != nil {
		return NumericSeparators{}, err
	}

	return nStrFmtSpecCurrValDto.numericSeparators.CopyOut(
		ePrefix.XCtx("nStrFmtSpecCurrValDto.numericSeparators"))
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
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecCurrencyValueDto instance to determine whether the
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
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
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("NumStrFmtSpecCurrencyValueDto.IsValidInstanceError()",
		"Testing Validity of 'nStrFmtSpecCurrValDto'")

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err := nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
		nStrFmtSpecCurrValDto,
		ePrefix)

	return err
}

// NewWithComponents - Creates and returns a new instance of
// NumStrFmtSpecCurrencyValueDto.
//
// The NumStrFmtSpecCurrencyValueDto type encapsulates the
// configuration parameters necessary to format numeric currency
// values for display in text number strings.
//
// This method requires detailed input parameters which provide
// granular control over all data fields contained in the returned
// new instance of NumStrFmtSpecCurrencyValueDto.
//
// For a 'New' method using minimum input parameters coupled
// with default values, see:
//      NumStrFmtSpecCurrencyValueDto.NewWithDefaults()
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
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols               []rune
//     - The authorized unicode character symbols associated with
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
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       included in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparator              rune
//         integerSeparators []NumStrIntSeparator
//        }
//
//        decimalSeparator              rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerSeparators             []NumStrIntSeparator
//           - An array of NumStrIntSeparator elements used to specify
//             the integer separation operation.
//
//              type NumStrIntSeparator struct {
//                intSeparatorChar     rune
//                intSeparatorGrouping uint
//              }
//
//               intSeparatorChar     rune
//               - This separator is commonly known as the 'thousands'
//                 separator. It is used to separate groups of integer
//                 digits to the left of the decimal separator (a.k.a.
//                 decimal point). In the United States, the standard
//                 integer digits separator is the comma (','). Other
//                 countries use periods, spaces or apostrophes to
//                 separate integers.
//                   United States Example:  1,000,000,000
//                    numSeps.intSeparators =
//                      []NumStrIntSeparator{
//                           {
//                           intSeparatorChar:   ',',
//                           intSeparatorGrouping: 3,
//                           },
//                        }
//
//               intSeparatorGrouping []uint
//               - In most western countries integer digits to the left
//                 of the decimal separator (a.k.a. decimal point) are
//                 separated into groups of three digits representing
//                 a grouping of 'thousands' like this: '1,000,000,000'.
//                 In this case the intSeparatorGrouping value would be
//                 set to three ('3').
//
//             In some countries and cultures other integer groupings are
//             used. In India, for example, a number might be formatted
//             like this: '6,78,90,00,00,00,00,000'. The right most group
//             has three digits and all the others are grouped by two. In
//             this case 'integerSeparators' would be configured as
//             follows:
//             as:
//
//             numSeps.intSeparators =
//               []NumStrIntSeparator{
//                    {
//                    intSeparatorChar:   ',',
//                    intSeparatorGrouping: 3,
//                    },
//                    {
//                    intSeparatorChar:     ',',
//                    intSeparatorGrouping: 2,
//                    },
//                 }
//
//
//  numFieldDto                   NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string which will be
//       inserted and displayed within the number field.
//
//       The NumberFieldDto object also contains specifications
//       for positioning or alignment of the number string within
//       the number field. This alignment dynamic is described as
//       text justification. The member variable '
//       NumberFieldDto.textJustifyFormat' is used to specify one
//       of three possible alignment formats. One of these formats
//       will be selected to control the alignment of the number
//       string within the number field. These optional alignment
//       formats are shown below with examples:
//
//       (1) 'Right-Justification' - "       NumberString"
//       (2) 'Left-Justification' - "NumberString        "
//       (3) 'Centered'           - "    NumberString    "
//
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//
//  ePrefix                       *ErrPrefixDto
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
//  NumStrFmtSpecCurrencyValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of
//       NumStrFmtSpecCurrencyValueDto.
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
func (nStrFmtSpecCurrValDto NumStrFmtSpecCurrencyValueDto) NewWithComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"NewWithComponents()")

	newNStrFmtSpecCurrencyValueDto := NumStrFmtSpecCurrencyValueDto{}

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	err := nStrFmtSpecCurrValMech.setCurrencyValDtoFromComponents(
		&newNStrFmtSpecCurrencyValueDto,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyCodeNo,
		currencyName,
		currencySymbols,
		minorCurrencyName,
		minorCurrencySymbols,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx("Setting 'newNStrFmtSpecCurrencyValueDto'"))

	return newNStrFmtSpecCurrencyValueDto, err
}

// NewWithDefaults - Creates and returns a new instance of
// NumStrFmtSpecCurrencyValueDto. This method specifies the minimum
// number of input parameters required to construct a new instance
// of NumStrFmtSpecCurrencyValueDto. Default values are used to
// supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of NumStrFmtSpecCurrencyValueDto,
// reference method:
//   'NumStrFmtSpecCurrencyValueDto.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'NumStrFmtSpecCurrencyValueDto.NewWithComponents()'.
//
// The NumStrFmtSpecCurrencyValueDto type encapsulates the
// formatting parameters necessary to format numeric currency
// values for display in text number strings.
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
//       integer number strings. In the United States, the
//       Thousands separator is the comma character (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       NumStrFmtSpecCurrencyValueDto.NewWithComponents().
//
//
//  turnOnThousandsSeparator      bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnThousandsSeparator' is a boolean flag
//       used to control the 'Thousands Separator'. When set to
//       'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
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
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto NumStrFmtSpecCurrencyValueDto) NewWithDefaults(
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	turnOnThousandsSeparator bool,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto.NewWithDefaults()")

	newNStrFmtSpecCurrencyValDto :=
		NumStrFmtSpecCurrencyValueDto{}

	nStrFmtSpecCurrValDtoUtil :=
		numStrFmtSpecCurrencyValueDtoUtility{}

	err :=
		nStrFmtSpecCurrValDtoUtil.setCurrValDtoWithDefaults(
			&newNStrFmtSpecCurrencyValDto,
			decimalSeparatorChar,
			thousandsSeparatorChar,
			turnOnThousandsSeparator,
			positiveValueFmt,
			negativeValueFmt,
			decimalDigits,
			currencyCode,
			currencyCodeNo,
			currencyName,
			currencySymbols,
			minorCurrencyName,
			minorCurrencySymbols,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx("Setting 'newNStrFmtSpecCurrencyValDto'"))

	return newNStrFmtSpecCurrencyValDto, err
}

// NewWithFmtSpecSetupDto - Creates and returns a new
// NumStrFmtSpecCurrencyValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     *NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecCurrencyValueDto object. Only the following
//       data fields with a prefix of "Currency" are used.
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
//         AbsoluteValNumSeps                        NumericSeparators
//         AbsoluteValNumField                       NumberFieldDto
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     uint
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumSeps                           NumericSeparators
//         CurrencyNumField                          NumberFieldDto
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumSeps                       NumericSeparators
//         SignedNumValNumField                      NumberFieldDto
//         SciNotSignificandUsesLeadingPlus          bool
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         SciNotNumFieldTextJustify                 TextJustify
//         Lock                                      *sync.Mutex
//       }
//
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
//  NumStrFmtSpecCurrencyValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecCurrencyValueDto will be returned to the
//       caller.
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
func (nStrFmtSpecCurrValDto NumStrFmtSpecCurrencyValueDto) NewWithFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecCurrencyValueDto,
	error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"NewWithFmtSpecSetupDto()")

	newNStrFmtSpecCurrencyValDto :=
		NumStrFmtSpecCurrencyValueDto{}

	if fmtSpecSetupDto == nil {
		return newNStrFmtSpecCurrencyValDto,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	err := nStrFmtSpecCurrValMech.setCurrencyValDtoFromComponents(
		&newNStrFmtSpecCurrencyValDto,
		fmtSpecSetupDto.CurrencyPositiveValueFmt,
		fmtSpecSetupDto.CurrencyNegativeValueFmt,
		fmtSpecSetupDto.CurrencyDecimalDigits,
		fmtSpecSetupDto.CurrencyCode,
		fmtSpecSetupDto.CurrencyCodeNo,
		fmtSpecSetupDto.CurrencyName,
		fmtSpecSetupDto.CurrencySymbols,
		fmtSpecSetupDto.MinorCurrencyName,
		fmtSpecSetupDto.MinorCurrencySymbols,
		fmtSpecSetupDto.CurrencyTurnOnIntegerDigitsSeparation,
		fmtSpecSetupDto.CurrencyNumSeps,
		fmtSpecSetupDto.CurrencyNumField,
		ePrefix)

	return newNStrFmtSpecCurrencyValDto, err
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

// SetCurrencyCodeNo - Sets the Currency Code Number associated
// with this currency. Currency Code Numbers are designated by the
// ISO 4217 standard.
//
// Currency Code Numbers are stored as strings per the ISO 4217
// standard.
//
// Reference:
//       https://en.wikipedia.org/wiki/ISO_4217
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencyCodeNo(
	currencyCodeNo string) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	nStrFmtSpecCurrValDto.currencyCodeNo = currencyCodeNo
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

// SetCurrencySymbols - Sets the currency name and currency symbols
// associated with this currency.
//
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  currencyName        string
//     - The official name for this currency.
//
//
//  currencySymbols     []rune
//     - The authorized unicode character symbols associated with
//       this currency specification. The currency symbol for the
//       United States is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetCurrencySymbols(
	currencyName string,
	currencySymbols []rune,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.SetCurrencySymbols()")

	err = nil

	if currencySymbols == nil {
		currencySymbols = make([]rune, 0, 5)
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"The currency symbols rune array has a zero length.\n"+
			"len(currencySymbols)==0\n",
			ePrefix.String())
		return err
	}

	nStrFmtSpecCurrValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		nStrFmtSpecCurrValDto.currencySymbols,
		currencySymbols)

	nStrFmtSpecCurrValDto.currencyName =
		currencyName

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

// SetMinorCurrencySymbol - Sets the minor currency name and
// minor currency symbols associated with this currency.
//
// Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  ePrefix                       *ErrPrefixDto
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetMinorCurrencySymbol(
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.SetCurrencySymbols()")

	err = nil

	if minorCurrencySymbols == nil {
		minorCurrencySymbols = make([]rune, 0, 5)
	}

	lenMinorCurrencySymbols := len(minorCurrencySymbols)

	if lenMinorCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'minorCurrencySymbols' is invalid!\n"+
			"The currency symbols rune array has a zero length.\n"+
			"len(minorCurrencySymbols)==0\n",
			ePrefix.String())
		return err
	}

	nStrFmtSpecCurrValDto.minorCurrencySymbols =
		make([]rune, lenMinorCurrencySymbols, 10)

	_ = copy(
		nStrFmtSpecCurrValDto.minorCurrencySymbols,
		minorCurrencySymbols)

	nStrFmtSpecCurrValDto.minorCurrencyName =
		minorCurrencyName

	return err
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetNegativeValueFormat(
	negativeValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.SetNegativeValueFormat()")

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		negativeValueFmt,
		ePrefix.XCtx("Testing validity of 'negativeValueFmt'"))

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
//  numberFieldLenDto   NumberFieldDto
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto.SetNumberFieldLengthDto()")

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
//  numericSeparators        NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparators struct {
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
	numberSeparatorsDto NumericSeparators) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	_ = nStrFmtSpecCurrValDto.numericSeparators.CopyIn(
		&numberSeparatorsDto,
		new(ErrPrefixDto))
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetPositiveValueFormat(
	positiveValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.SetPositiveValueFormat()")

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx("Testing validity of 'positiveValueFmt'"))

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

// SetWithComponents - This method will overwrite and set all data
// data values for the current instance of
// NumStrFmtSpecCurrencyValueDto.
//
// The NumStrFmtSpecCurrencyValueDto type encapsulates the
// formatting parameters necessary to format numeric currency
// values for display in text number strings.
//
// This method requires detailed input parameters to control
// configuration for all member variables in the current instance
// of NumStrFmtSpecCurrencyValueDto. For a similar method using
// minimum input parameters coupled with default values, see:
//      NumStrFmtSpecCurrencyValueDto.SetWithDefaults()
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current NumStrFmtSpecCurrencyValueDto instance.
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
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols               []rune
//     - The authorized unicode character symbols associated with
//       this currency specification. The currency symbol for the
//       United States is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The unicode character for minor currency symbol. In the
//       United States, the minor currency symbol is the cent sign
//       (¢). Some countries and cultures have currency symbols
//       consisting of two or more characters.
//
//
//  minorCurrencySymbols          []rune
//     - The unicode character for minor currency symbol. In the
//       United States, the minor currency symbol is the cent sign
//       (¢).
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
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       included in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparator              rune
//         integerSeparators []NumStrIntSeparator
//        }
//
//        decimalSeparator              rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerSeparators             []NumStrIntSeparator
//           - An array of NumStrIntSeparator elements used to specify
//             the integer separation operation.
//
//              type NumStrIntSeparator struct {
//                intSeparatorChar     rune
//                intSeparatorGrouping uint
//              }
//
//               intSeparatorChar     rune
//               - This separator is commonly known as the 'thousands'
//                 separator. It is used to separate groups of integer
//                 digits to the left of the decimal separator (a.k.a.
//                 decimal point). In the United States, the standard
//                 integer digits separator is the comma (','). Other
//                 countries use periods, spaces or apostrophes to
//                 separate integers.
//                   United States Example:  1,000,000,000
//                    numSeps.intSeparators =
//                      []NumStrIntSeparator{
//                           {
//                           intSeparatorChar:   ',',
//                           intSeparatorGrouping: 3,
//                           },
//                        }
//
//               intSeparatorGrouping []uint
//               - In most western countries integer digits to the left
//                 of the decimal separator (a.k.a. decimal point) are
//                 separated into groups of three digits representing
//                 a grouping of 'thousands' like this: '1,000,000,000'.
//                 In this case the intSeparatorGrouping value would be
//                 set to three ('3').
//
//             In some countries and cultures other integer groupings are
//             used. In India, for example, a number might be formatted
//             like this: '6,78,90,00,00,00,00,000'. The right most group
//             has three digits and all the others are grouped by two. In
//             this case 'integerSeparators' would be configured as
//             follows:
//             as:
//
//             numSeps.intSeparators =
//               []NumStrIntSeparator{
//                    {
//                    intSeparatorChar:   ',',
//                    intSeparatorGrouping: 3,
//                    },
//                    {
//                    intSeparatorChar:     ',',
//                    intSeparatorGrouping: 2,
//                    },
//                 }
//
//
//  numFieldDto                NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string which will be
//       inserted and displayed within the number field.
//
//       The NumberFieldDto object also contains specifications
//       for positioning or alignment of the number string within
//       the number field. This alignment dynamic is described as
//       text justification. The member variable '
//       NumberFieldDto.textJustifyFormat' is used to specify one
//       of three possible alignment formats. One of these formats
//       will be selected to control the alignment of the number
//       string within the number field. These optional alignment
//       formats are shown below with examples:
//
//       (1) 'Right-Justification' - "       NumberString"
//       (2) 'Left-Justification' - "NumberString        "
//       (3) 'Centered'           - "    NumberString    "
//
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetWithComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"SetWithComponents()")

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	return nStrFmtSpecCurrValMech.setCurrencyValDtoFromComponents(
		nStrFmtSpecCurrValDto,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyCodeNo,
		currencyName,
		currencySymbols,
		minorCurrencyName,
		minorCurrencySymbols,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'nStrFmtSpecCurrValDto'"))
}

// SetWithDefaults - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecCurrencyValueDto. The input parameters represent
// the minimum information required to configure a
// NumStrFmtSpecCurrencyValueDto object.
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current NumStrFmtSpecCurrencyValueDto instance.
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
//       integer number strings. In the United States, the
//       Thousands separator is the comma character (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       NumStrFmtSpecCurrencyValueDto.SetWithComponents().
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
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetWithDefaults(
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	turnOnIntegerDigitsSeparation bool,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCurrencyValueDto.SetWithDefaults()")

	nStrFmtSpecCurrValDtoUtil :=
		numStrFmtSpecCurrencyValueDtoUtility{}

	return nStrFmtSpecCurrValDtoUtil.setCurrValDtoWithDefaults(
		nStrFmtSpecCurrValDto,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		turnOnIntegerDigitsSeparation,
		positiveValueFmt,
		negativeValueFmt,
		decimalDigits,
		currencyCode,
		currencyCodeNo,
		currencyName,
		currencySymbols,
		minorCurrencyName,
		minorCurrencySymbols,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix.XCtx("Setting 'nStrFmtSpecCurrValDto'"))
}

// SetWithFmtSpecSetupDto - Sets the data values for current
// NumStrFmtSpecCurrencyValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current NumStrFmtSpecCurrencyValueDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecCurrencyValueDto object. Only the following
//       data fields with a prefix of "Currency" are used.
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
//         AbsoluteValNumSeps                        NumericSeparators
//         AbsoluteValNumField                       NumberFieldDto
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     uint
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumSeps                           NumericSeparators
//         CurrencyNumField                          NumberFieldDto
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumSeps                       NumericSeparators
//         SignedNumValNumField                      NumberFieldDto
//         SciNotSignificandUsesLeadingPlus          bool
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         SciNotNumFieldTextJustify                 TextJustify
//         Lock                                      *sync.Mutex
//       }
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetWithFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecCurrencyValueDto." +
			"SetWithFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValMech :=
		numStrFmtSpecCurrencyValueDtoMechanics{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	return nStrFmtSpecCurrValMech.setCurrencyValDtoFromComponents(
		nStrFmtSpecCurrValDto,
		fmtSpecSetupDto.CurrencyPositiveValueFmt,
		fmtSpecSetupDto.CurrencyNegativeValueFmt,
		fmtSpecSetupDto.CurrencyDecimalDigits,
		fmtSpecSetupDto.CurrencyCode,
		fmtSpecSetupDto.CurrencyCodeNo,
		fmtSpecSetupDto.CurrencyName,
		fmtSpecSetupDto.CurrencySymbols,
		fmtSpecSetupDto.MinorCurrencyName,
		fmtSpecSetupDto.MinorCurrencySymbols,
		fmtSpecSetupDto.CurrencyTurnOnIntegerDigitsSeparation,
		fmtSpecSetupDto.CurrencyNumSeps,
		fmtSpecSetupDto.CurrencyNumField,
		ePrefix)
}
