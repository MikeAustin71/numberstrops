package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFmtSpecSignedNumValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumericSeparatorDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecSignedNumValueDto instance  to the data fields
// of the current instance of NumStrFmtSpecSignedNumValueDto
// instance.
//
// If input parameter 'incomingSignedNumValDto' is judged to be
// invalid, this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrFmtSpecSignedNumValueDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingSignedNumValDto    *NumStrFmtSpecSignedNumValueDto
//     - A pointer to an instance of NumStrFmtSpecSignedNumValueDto.
//       The data values in this object will be copied to the
//       current NumStrFmtSpecSignedNumValueDto instance.
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) CopyIn(
	incomingSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.CopyIn()")

	nStrFmtSpecSignedNumValNanobot :=
		numStrFmtSpecSignedNumValNanobot{}

	return nStrFmtSpecSignedNumValNanobot.copyIn(
		nStrFmtSpecSignedNumValueDto,
		incomingSignedNumValDto,
		ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrFmtSpecSignedNumValueDto instance.
//
// If the current NumStrFmtSpecSignedNumValueDto instance is judged
// to be invalid, this method will return an error.
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecSignedNumValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecSignedNumValueDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrFmtSpecSignedNumValueDto.
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecSignedNumValueDto,
	error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.CopyOut()")

	nStrFmtSpecSignedNumValNanobot :=
		numStrFmtSpecSignedNumValNanobot{}

	return nStrFmtSpecSignedNumValNanobot.copyOut(
		nStrFmtSpecSignedNumValueDto,
		ePrefix.XCtx(
			"Copy Out from 'nStrFmtSpecSignedNumValueDto'"))
}

// GetDecimalSeparator - Returns the decimal separator character.
// This is the character (rune) used to separate integer and
// fractional digits in a floating point number.
//
// In the United States, the Decimal Separator character is the
// decimal point or period ('.').
//
//  Example:   123.45
//
// Decimal Separator is extracted from the underlying member
// variable, 'nStrFmtSpecSignedNumValueDto.numberSeparatorsDto'.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetDecimalSeparator() rune {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.
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
// variable, 'nStrFmtSpecSignedNumValueDto.numberSeparatorsDto'.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetIntegerDigitsSeparator() rune {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.
		numberSeparatorsDto.
		GetDecimalSeparator()
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetIntegerDigitsGroupingSequence() []uint {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.
		numberSeparatorsDto.
		GetIntegerDigitsGroupingSequence()
}

// GetNegativeValueFormat - Returns the formatting string used to
// format negative signed number values in text number strings.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetNegativeValueFormat() string {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.negativeValueFmt
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Signed Number Value Dto.
//
// The NumberFieldDto details the length of the number field in which
// the signed numeric value will be displayed and right justified.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetNumberFieldLengthDto() NumberFieldDto {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.numFieldLenDto.CopyOut()
}

// GetNumberSeparatorsDto - Returns the NumericSeparatorDto
// instance currently configured for this Number String Format Specification
// Signed Number Value Dto.
//
// The Digits Separator Dto contains the decimal separator as well as the
// 'thousands' separators and the grouping sequence for separating thousands
// digits in the integer component of a number string.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetNumberSeparatorsDto() NumericSeparatorDto {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.numberSeparatorsDto.CopyOut()
}

// GetPositiveValueFormat - Returns the formatting string used to
// format positive signed number values in text number strings.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetPositiveValueFormat() string {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.positiveValueFmt
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecSignedNumValueDto instance to determine whether
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
//       NumStrFmtSpecSignedNumValueDto is valid, or not. If the
//       current NumStrFmtSpecSignedNumValueDto contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValMolecule :=
		numStrFmtSpecSignedNumValMolecule{}

	isValid,
		_ = nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
		nStrFmtSpecSignedNumValueDto,
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrFmtSpecSignedNumValueDto instance to determine
// whether the current instance is valid in all respects.
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
//     - If the current instance of NumStrFmtSpecSignedNumValueDto
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("NumStrFmtSpecSignedNumValueDto.IsValidInstanceError()",
		"Testing Validity of 'nStrFmtSpecSignedNumValueDto'")

	nStrFmtSpecSignedNumValMolecule :=
		numStrFmtSpecSignedNumValMolecule{}

	_,
		err :=
		nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
			nStrFmtSpecSignedNumValueDto,
			ePrefix)

	return err
}

// NewWithDefaults - Creates and returns a new instance of
// NumStrFmtSpecSignedNumValueDto.
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'NumStrFmtSpecSignedNumValueDto.NewFromComponents()'.
//
// The NumStrFmtSpecSignedNumValueDto type encapsulates the
// formatting parameters necessary to format signed numeric values
// for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
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
//  negativeValueFmt              string
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
//  NumStrFmtSpecSignedNumValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecSignedNumValueDto.
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
func (nStrFmtSpecSignedNumValueDto NumStrFmtSpecSignedNumValueDto) NewWithDefaults(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecSignedNumValueDto,
	error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.NewWithDefaults()")

	newNStrFmtSpecSignedNumValueDto :=
		NumStrFmtSpecSignedNumValueDto{}

	nStrFmtSpecSignedNumValDtoUtil :=
		nStrFmtSpecSignedNumValUtility{}

	err :=
		nStrFmtSpecSignedNumValDtoUtil.setSignedNumValDtoWithDefaults(
			&newNStrFmtSpecSignedNumValueDto,
			positiveValueFmt,
			negativeValueFmt,
			turnOnIntegerDigitsSeparation,
			decimalSeparatorChar,
			thousandsSeparatorChar,
			[]uint{3},
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
}

// NewFromComponents - Creates and returns a new instance of
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
//  negativeValueFmt              string
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
//  numberSeparatorsDto        NumericSeparatorDto
//     - This instance of 'NumericSeparatorDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorDto struct {
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
//  NumStrFmtSpecSignedNumValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecSignedNumValueDto.
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
func (nStrFmtSpecSignedNumValueDto NumStrFmtSpecSignedNumValueDto) NewFromComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparatorDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecSignedNumValueDto,
	error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.NewFromComponents()")

	newNStrFmtSpecSignedNumValueDto := NumStrFmtSpecSignedNumValueDto{}

	nStrFmtSpecSignedNumValMech :=
		nStrFmtSpecSignedNumValMechanics{}

	err := nStrFmtSpecSignedNumValMech.setSignedNumValDto(
		&newNStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
}

// NewFromFmtSpecSetupDto - Creates and returns a new
// NumStrFmtSpecSignedNumValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     *NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecSignedNumValueDto object. Only the following
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
//         CurrencyName                              string
//         CurrencySymbol                            rune
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
//  NumStrFmtSpecSignedNumValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecSignedNumValueDto will be returned to the
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
func (nStrFmtSpecSignedNumValueDto NumStrFmtSpecSignedNumValueDto) NewFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecSignedNumValueDto,
	error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecSignedNumValueDto.NewFromFmtSpecSetupDto()")

	newNStrFmtSpecSignedNumValueDto :=
		NumStrFmtSpecSignedNumValueDto{}

	if fmtSpecSetupDto == nil {
		return newNStrFmtSpecSignedNumValueDto,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValDtoUtil :=
		nStrFmtSpecSignedNumValUtility{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	err :=
		nStrFmtSpecSignedNumValDtoUtil.setSignedNumValDtoWithDefaults(
			&newNStrFmtSpecSignedNumValueDto,
			fmtSpecSetupDto.SignedNumValPositiveValueFmt,
			fmtSpecSetupDto.SignedNumValNegativeValueFmt,
			fmtSpecSetupDto.SignedNumValTurnOnIntegerDigitsSeparation,
			fmtSpecSetupDto.DecimalSeparator,
			fmtSpecSetupDto.IntegerDigitsSeparator,
			fmtSpecSetupDto.IntegerDigitsGroupingSequence,
			fmtSpecSetupDto.SignedNumValNumFieldLen,
			fmtSpecSetupDto.SignedNumValNumFieldTextJustify,
			ePrefix.XCtx(
				"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
}

// SetFromFmtSpecSetupDto - Sets the data values for current
// NumStrFmtSpecSignedNumValueDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup and configuration
//       information for a NumStrFmtSpecSignedNumValueDto
//       object. Only the following data fields with a prefix of
//       "SignedNumVal" are used.
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
//         CurrencyName                              string
//         CurrencySymbol                            rune
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.SetFromFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValDtoUtil :=
		nStrFmtSpecSignedNumValUtility{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	return nStrFmtSpecSignedNumValDtoUtil.setSignedNumValDtoWithDefaults(
		nStrFmtSpecSignedNumValueDto,
		fmtSpecSetupDto.SignedNumValPositiveValueFmt,
		fmtSpecSetupDto.SignedNumValNegativeValueFmt,
		fmtSpecSetupDto.SignedNumValTurnOnIntegerDigitsSeparation,
		fmtSpecSetupDto.DecimalSeparator,
		fmtSpecSetupDto.IntegerDigitsSeparator,
		fmtSpecSetupDto.IntegerDigitsGroupingSequence,
		fmtSpecSetupDto.SignedNumValNumFieldLen,
		fmtSpecSetupDto.SignedNumValNumFieldTextJustify,
		ePrefix.XCtx("Setting 'nStrFmtSpecSignedNumValueDto'"))

}

// SetNegativeValueFormat - Sets the negative value format string
// used to configure signed numeric values in text number strings.
//
// If input parameter 'negativeValueFmt' is invalid, this method
// will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  negativeValueFmt              string
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNegativeValueFormat(
	negativeValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.SetNegativeValueFormat()")

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	_,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormatStr(
		negativeValueFmt,
		ePrefix.XCtx(
			"Testing validity of 'negativeValueFmt'"))

	if err != nil {
		return err
	}

	nStrFmtSpecSignedNumValueDto.negativeValueFmt = negativeValueFmt

	return err
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current NumStrFmtSpecSignedNumValueDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLenDto   NumberFieldDto
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecSignedNumValueDto.SetNumberFieldLengthDto()")

	return nStrFmtSpecSignedNumValueDto.numFieldLenDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumberSeparatorsDto - Sets the Number Separators Dto object
// for the current NumStrFmtSpecSignedNumValueDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberSeparatorsDto        NumericSeparatorDto
//     - This instance of 'NumericSeparatorDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorDto struct {
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNumberSeparatorsDto(
	numberSeparatorsDto NumericSeparatorDto) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	_ = nStrFmtSpecSignedNumValueDto.numberSeparatorsDto.CopyIn(
		&numberSeparatorsDto,
		new(ErrPrefixDto))
}

// SetPositiveValueFormat - Sets the positive value format string
// used to configure signed numeric values in text number strings.
//
// If input parameter 'positiveValueFmt' is invalid, this method
// will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetPositiveValueFormat(
	positiveValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecSignedNumValueDto.SetPositiveValueFormat()")

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	_,
		err = nStrSignedNumElectron.testSignedNumValPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			"Testing validity of 'positiveValueFmt'"))

	if err != nil {
		return err
	}

	nStrFmtSpecSignedNumValueDto.positiveValueFmt = positiveValueFmt

	return err
}

// SetSignedNumValDto - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecSignedNumValueDto.
//
// The NumStrFmtSpecSignedNumValueDto type encapsulates the
// formatting parameters necessary for formatting signed number
// values in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt              string
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
//  negativeValueFmt              string
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
//  numberSeparatorsDto        NumericSeparatorDto
//     - This instance of 'NumericSeparatorDto' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparatorDto struct {
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetSignedNumValDto(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparatorDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecSignedNumValueDto.SetSignedNumValDto()")

	nStrFmtSpecSignedNumValMech :=
		nStrFmtSpecSignedNumValMechanics{}

	return nStrFmtSpecSignedNumValMech.setSignedNumValDto(
		nStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix.XCtx("Setting 'nStrFmtSpecSignedNumValueDto'"))
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// NumStrFmtSpecSignedNumValueDto.
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValueDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithDefaults - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecSignedNumValueDto. The input parameters
// represent the minimum information required to configure
// the data values for a NumStrFmtSpecSignedNumValueDto object.
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
//  positiveValueFmt              string
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
//  negativeValueFmt              string
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetWithDefaults(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecSignedNumValueDto.SetWithDefaults()")

	nStrFmtSpecSignedNumValDtoUtil :=
		nStrFmtSpecSignedNumValUtility{}

	return nStrFmtSpecSignedNumValDtoUtil.setSignedNumValDtoWithDefaults(
		nStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		[]uint{3},
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix.XCtx("Setting 'nStrFmtSpecSignedNumValueDto'"))

}
