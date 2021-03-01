package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFmtSpecAbsoluteValueDto struct {
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumericSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrFmtSpecAbsoluteValueDto ('inComingNStrFmtAbsValDto') to
// the data fields of the current NumStrFmtSpecAbsoluteValueDto
// instance.
//
// If 'inComingNStrFmtAbsValDto' is judged to be invalid, this
// method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrFmtSpecAbsoluteValueDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  inComingNStrFmtAbsValDto   *NumStrFmtSpecAbsoluteValueDto
//     - A pointer to an instance of NumStrFmtSpecAbsoluteValueDto.
//       The data values in this object will be copied to the
//       current NumStrFmtSpecAbsoluteValueDto instance.
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyIn(
	inComingNStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.CopyIn()")

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyIn(
		nStrFmtAbsValDto,
		inComingNStrFmtAbsValDto,
		ePrefix.XCtx("inComingNStrFmtAbsValDto -> nStrFmtAbsValDto"))
}

// CopyOut - Returns a deep copy of the current
// NumStrFmtSpecAbsoluteValueDto instance.
//
// If the current NumStrFmtSpecAbsoluteValueDto instance is judged
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
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecAbsoluteValueDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyOut(
	ePrefix *ErrPrefixDto) (NumStrFmtSpecAbsoluteValueDto, error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.CopyOut()")

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyOut(
		nStrFmtAbsValDto,
		ePrefix.XCtx(
			"Copy -> 'nStrFmtAbsValDto'"))
}

// GetAbsoluteValueFormat - Returns the formatting string used to
// format absolute numeric values in text number strings.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetAbsoluteValueFormat() string {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.absoluteValFmt
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
// variable, 'nStrFmtAbsValDto.numberSeparatorsDto'.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetDecimalSeparator() rune {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.
		numberSeparatorsDto.
		GetDecimalSeparator()
}

// GetIntegerDigitsSeparator - Returns the unicode character (rune)
// used to separate integer digits. This is typically known as the
// 'thousands' separator which is used to separate thousands in
// three digit groups. In the United States, the inter digits
// separator is the comma (',').
//
//  Example 1,000,000,000
//
// Integer Digits Separator is extracted from the underlying member
// variable, 'nStrFmtSpecCurrValDto.numberSeparatorsDto'.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetIntegerDigitsSeparator() rune {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.
		numberSeparatorsDto.
		GetIntegerDigitsSeparators()
}

// GetIntegerDigitsGroupingSequence - Returns the value of the
// integer digits grouping sequence. This refers to grouping of
// integer digits within a string of numeric digits.
//
// In most western countries integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000,000'. In this case, the integer digits
// grouping sequence would be configured as:
//        integerDigitsGroupingSequence = []uint{3}
//
// In some countries and cultures, other integer groupings are
// used. In India, for example, a number might be formatted as
// like this: '6,78,90,00,00,00,00,000'. The right most group
// has three digits and all the others are grouped by two digits.
// In this case integer digits grouping sequence would be
// configured as:
//        integerDigitsGroupingSequence = []uint{3,2}
//
// The integer digits grouping sequence is extracted from member
// variable 'nStrFmtSpecCurrValDto.numberSeparatorsDto'.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetIntegerDigitsGroupingSequence() []uint {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.
		numberSeparatorsDto.
		GetIntegerDigitsGroupingSequence()
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Absolute Value Dto.
//
// The NumberFieldDto details the length of the number field in
// which the signed numeric value will be displayed and right
// justified.
//
// If the NumberFieldDto object is judged to be invalid, this
// method will return an error.
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
//       NumberFieldDto will be created and returned containing all
//       of the data values copied from the Number Field values
//       used to configure the current instance of
//       NumStrFmtSpecAbsoluteValueDto.
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
//       Be advised - If the NumberFieldDto object is judged to be
//       invalid, this method will return an error.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecAbsoluteValueDto.GetNumberFieldLengthDto()")

	return nStrFmtAbsValDto.numFieldLenDto.CopyOut(
		ePrefix.XCtx(
			"nStrFmtAbsValDto.numFieldLenDto->"))
}

// GetNumericSeparatorsDto - Returns the NumericSeparatorsDto
// instance currently configured for this Number String Format
// Specification Signed Number Value Dto.
//
// The Numeric Separators Dto object contains the decimal
// separator, the 'thousands' separator and the integer grouping
// sequence for separating thousands digits within the integer
// component of a number string.
//
// The returned NumericSeparatorsDto object represents the Numeric
// Separator values used to configure the current instance of
// NumStrFmtSpecAbsoluteValueDto.
//
// If the Numeric Separator Dto object is judged to be invalid,
// this method will return an error.
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
//  NumericSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumericSeparatorsDto will be created and returned
//       containing all of the data values used to configure the
//       current instance of NumStrFmtSpecAbsoluteValueDto.
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
//       Be advised - If the NumericSeparatorsDto object is judged
//       to be invalid, this method will return an error.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumericSeparatorsDto(
	ePrefix *ErrPrefixDto) (
	NumericSeparatorsDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecAbsoluteValueDto.GetNumericSeparatorsDto()")

	return nStrFmtAbsValDto.numberSeparatorsDto.CopyOut(
		ePrefix.XCtx("nStrFmtAbsValDto.numberSeparatorsDto->"))
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecAbsoluteValueDto instance to determine whether
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
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumStrFmtSpecAbsoluteValueDto
//       is valid, or not. If the current NumStrFmtSpecAbsoluteValueDto
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	isValid,
		_ = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		nStrFmtAbsValDto,
		ErrPrefixDto{}.Ptr())

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecAbsoluteValueDto instance to determine whether the
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("NumStrFmtSpecAbsoluteValueDto.IsValidInstanceError()",
		"Testing Validity of 'nStrFmtAbsValDto'")

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err := nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		nStrFmtAbsValDto,
		ePrefix)

	return err
}

// NewWithDefaults - Creates and returns a new instance of
// NumStrFmtSpecAbsoluteValueDto.
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'NumStrFmtSpecAbsoluteValueDto.NewFromComponents()'.
//
// The NumStrFmtSpecAbsoluteValueDto type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
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
//               "127.54" THE DEFAULT Absolute Value Format
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
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto NumStrFmtSpecAbsoluteValueDto) NewWithDefaults(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecAbsoluteValueDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.NewWithDefaults()")

	newNumStrFmtSpecAbsValueDto :=
		NumStrFmtSpecAbsoluteValueDto{}

	nStrFmtSpecAbsValDtoUtil :=
		numStrFmtSpecAbsoluteValueDtoUtility{}

	err := nStrFmtSpecAbsValDtoUtil.setAbsValDtoWithDefaults(
		&newNumStrFmtSpecAbsValueDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		[]uint{3},
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewFromComponents - Creates and returns a new instance of
// NumStrFmtSpecAbsoluteValueDto.
//
// This type encapsulates the formatting parameters necessary to
// format absolute values for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
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
//               "127.54" THE DEFAULT Absolute Value Format
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
//  numberSeparatorsDto        NumericSeparatorsDto
//     - This instance of 'NumericSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorsDto struct {
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of
//       NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto NumStrFmtSpecAbsoluteValueDto) NewFromComponents(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecAbsoluteValueDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecAbsoluteValueDto.NewFromComponents()")

	newNStrFmtSpecAbsoluteValueDto := NumStrFmtSpecAbsoluteValueDto{}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err := nStrFmtSpecAbsValDtoMech.setAbsValDto(
		&newNStrFmtSpecAbsoluteValueDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'newNStrFmtSpecAbsoluteValueDto'"))

	return newNStrFmtSpecAbsoluteValueDto, err
}

// NewFromFmtSpecSetupDto - Creates and returns a new
// NumStrFmtSpecAbsoluteValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecAbsoluteValueDto object. Only the following
//       data fields with a prefix of "AbsoluteVal" are used.
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecAbsoluteValueDto will be returned to the
//       caller.
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
func (nStrFmtAbsValDto NumStrFmtSpecAbsoluteValueDto) NewFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecAbsoluteValueDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"nNumStrFmtSpecAbsoluteValueDto.NewFromFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return NumStrFmtSpecAbsoluteValueDto{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	newNumStrFmtSpecAbsValueDto :=
		NumStrFmtSpecAbsoluteValueDto{}

	nStrFmtSpecAbsValDtoUtil :=
		numStrFmtSpecAbsoluteValueDtoUtility{}

	err := nStrFmtSpecAbsValDtoUtil.setAbsValDtoWithDefaults(
		&newNumStrFmtSpecAbsValueDto,
		fmtSpecSetupDto.AbsoluteValFmt,
		fmtSpecSetupDto.AbsoluteValTurnOnIntegerDigitsSeparation,
		fmtSpecSetupDto.DecimalSeparator,
		fmtSpecSetupDto.IntegerDigitsSeparator,
		fmtSpecSetupDto.IntegerDigitsGroupingSequence,
		fmtSpecSetupDto.AbsoluteValNumFieldLen,
		fmtSpecSetupDto.AbsoluteValNumFieldTextJustify,
		ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// SetAbsoluteValueFormat - Sets the formatting string used to
// format absolute numeric values in text number strings.
//
// If the absolute value format string input parameter,
// 'absoluteValueFormatStr', is invalid, this method will return
// an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValueFormatStr     string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetAbsoluteValueFormat(
	absoluteValueFormatStr string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.SetAbsoluteValueFormat()")

	nStrAbsValDtoElectron :=
		numStrFmtSpecAbsoluteValueDtoElectron{}

	_,
		err = nStrAbsValDtoElectron.testAbsoluteValueFormatStr(
		absoluteValueFormatStr,
		ePrefix.XCtx(
			"Testing validity of 'absoluteValueFormatStr'"))

	if err != nil {
		return err
	}

	nStrFmtAbsValDto.absoluteValFmt = absoluteValueFormatStr

	return err
}

// SetAbsValDto - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecAbsoluteValueDto.
//
// The NumStrFmtSpecAbsoluteValueDto type encapsulates the
// formatting parameters necessary for formatting signed number
// values in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
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
//               "127.54" THE DEFAULT Absolute Value Format
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
//  numberSeparatorsDto        NumericSeparatorsDto
//     - This instance of 'NumericSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorsDto struct {
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
//  numFieldDto                NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string which will be
//       displayed within the number field.
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
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetAbsValDto(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecAbsoluteValueDto.SetAbsValDto()")

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	return nStrFmtSpecAbsValDtoMech.setAbsValDto(
		nStrFmtAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'nStrFmtAbsValDto'"))
}

// SetFromFmtSpecSetupDto - Sets the data values for current
// NumStrFmtSpecAbsoluteValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecAbsoluteValueDto object. Only the following
//       data fields with a prefix of "AbsoluteVal" are used.
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
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecCountryDto.SetFromFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoUtil :=
		numStrFmtSpecAbsoluteValueDtoUtility{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	return nStrFmtSpecAbsValDtoUtil.setAbsValDtoWithDefaults(
		nStrFmtAbsValDto,
		fmtSpecSetupDto.AbsoluteValFmt,
		fmtSpecSetupDto.AbsoluteValTurnOnIntegerDigitsSeparation,
		fmtSpecSetupDto.DecimalSeparator,
		fmtSpecSetupDto.IntegerDigitsSeparator,
		fmtSpecSetupDto.IntegerDigitsGroupingSequence,
		fmtSpecSetupDto.AbsoluteValNumFieldLen,
		fmtSpecSetupDto.AbsoluteValNumFieldTextJustify,
		ePrefix)
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current NumStrFmtSpecAbsoluteValueDto instance.
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.SetNumberFieldLengthDto()")

	return nStrFmtAbsValDto.numFieldLenDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumberSeparatorsDto - Sets the Number Separators Dto object
// for the current NumStrFmtSpecAbsoluteValueDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberSeparatorsDto        NumericSeparatorsDto
//     - This instance of 'NumericSeparatorsDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorsDto struct {
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
//  --- NONE ---
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetNumberSeparatorsDto(
	numberSeparatorsDto NumericSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	return nStrFmtAbsValDto.numberSeparatorsDto.CopyIn(
		&numberSeparatorsDto,
		ePrefix)
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	nStrFmtAbsValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithDefaults - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecAbsoluteValueDto. The input parameters
// represent the minimum information required to configure
// the data values for a NumStrFmtSpecAbsoluteValueDto object.
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetWithDefaults(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecAbsoluteValueDto.SetWithDefaults()")

	nStrFmtSpecAbsValDtoUtil :=
		numStrFmtSpecAbsoluteValueDtoUtility{}

	return nStrFmtSpecAbsValDtoUtil.setAbsValDtoWithDefaults(
		nStrFmtAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		[]uint{3},
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)
}
