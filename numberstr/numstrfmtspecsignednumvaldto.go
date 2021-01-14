package numberstr

import "sync"

type NumStrFmtSpecSignedNumValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumStrFmtSpecDigitsSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecSignedNumValueDto instance  to the data fields
// of the current instance of NumStrFmtSpecSignedNumValueDto
// instance.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) CopyIn(
	incomingSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecSignedNumValueDto.CopyIn() "

	nStrFmtSpecSignedNumValAtom :=
		numStrFmtSpecSignedNumValAtom{}

	return nStrFmtSpecSignedNumValAtom.copyIn(
		nStrFmtSpecSignedNumValueDto,
		incomingSignedNumValDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current
// NumStrFmtSpecSignedNumValueDto instance.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) CopyOut() NumStrFmtSpecSignedNumValueDto {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValAtom :=
		numStrFmtSpecSignedNumValAtom{}

	newNStrFmtSpecSignedNumberValueDto,
		_ := nStrFmtSpecSignedNumValAtom.copyOut(
		nStrFmtSpecSignedNumValueDto,
		"")

	return newNStrFmtSpecSignedNumberValueDto
}

// GetNegativeValueFormat - Returns the formatting string used to
// format negative numeric values in text number strings.
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

// GetNumberSeparatorsDto - Returns the NumStrFmtSpecDigitsSeparatorsDto
// instance currently configured for this Number String Format Specification
// Signed Number Value Dto.
//
// The Digits Separator Dto contains the decimal separator as well as the
// 'thousands' separators and the grouping sequence for separating thousands
// digits in the integer component of a number string.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) GetNumberSeparatorsDto() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	return nStrFmtSpecSignedNumValueDto.numberSeparatorsDto.CopyOut()
}

// GetPositiveValueFormat - Returns the formatting string used to
// format positive numeric values in text number strings.
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
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumStrFmtSpecSignedNumValueDto
//       is valid, or not. If the current NumStrFmtSpecSignedNumValueDto
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValAtom :=
		numStrFmtSpecSignedNumValAtom{}

	isValid,
		_ = nStrFmtSpecSignedNumValAtom.testValidityOfSignedNumValDto(
		nStrFmtSpecSignedNumValueDto,
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
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
//       If this instance of NumStrFmtSpecSignedNumValueDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecSignedNumValueDto.IsValidInstanceError() \n" +
		"Testing Validity of 'nStrFmtSpecSignedNumValueDto' "

	nStrFmtSpecSignedNumValAtom :=
		numStrFmtSpecSignedNumValAtom{}

	_,
		err := nStrFmtSpecSignedNumValAtom.testValidityOfSignedNumValDto(
		nStrFmtSpecSignedNumValueDto,
		ePrefix)

	return err
}

// New() - Creates and returns a new instance of
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
func (nStrFmtSpecSignedNumValueDto NumStrFmtSpecSignedNumValueDto) New(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecSignedNumValueDto,
	error) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSignedNumValueDto.SetSignedNumValDto() "

	numFieldDto := NumberFieldDto{}.New(requestedNumberFieldLen)

	newNStrFmtSpecSignedNumValueDto := NumStrFmtSpecSignedNumValueDto{}

	nStrFmtSpecSignedNumValMolecule :=
		nStrFmtSpecSignedNumValMolecule{}

	err := nStrFmtSpecSignedNumValMolecule.setSignedNumValDto(
		&newNStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'nStrFmtSpecSignedNumValueDto'\n ")

	return newNStrFmtSpecSignedNumValueDto, err
}

// SetNegativeValueFormat - Sets the negative value format string
// used to configure signed numeric values in text number strings.
//
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
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNegativeValueFormat(
	negativeValueFmt string) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValueDto.negativeValueFmt = negativeValueFmt
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
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValueDto.numFieldLenDto.CopyIn(
		&numberFieldLenDto)
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetNumberSeparatorsDto(
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	_ = nStrFmtSpecSignedNumValueDto.numberSeparatorsDto.CopyIn(
		&numberSeparatorsDto,
		"")
}

// SetPositiveValueFormat - Sets the positive value format string
// used to configure signed numeric values in text number strings.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetPositiveValueFormat(
	positiveValueFmt string) {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	nStrFmtSpecSignedNumValueDto.positiveValueFmt = positiveValueFmt
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag.
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

// SetSignedNumValDto() - This method will set all of the member
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
func (nStrFmtSpecSignedNumValueDto *NumStrFmtSpecSignedNumValueDto) SetSignedNumValDto(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) error {

	if nStrFmtSpecSignedNumValueDto.lock == nil {
		nStrFmtSpecSignedNumValueDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValueDto.lock.Lock()

	defer nStrFmtSpecSignedNumValueDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSignedNumValueDto.SetSignedNumValDto() "

	numFieldDto := NumberFieldDto{}.New(requestedNumberFieldLen)

	nStrFmtSpecSignedNumValMolecule :=
		nStrFmtSpecSignedNumValMolecule{}

	return nStrFmtSpecSignedNumValMolecule.setSignedNumValDto(
		nStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'nStrFmtSpecSignedNumValueDto'\n ")
}