package numberstr

import "sync"

type NumStrFmtSpecAbsoluteValueDto struct {
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumStrFmtSpecDigitsSeparatorsDto
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyIn(
	inComingNStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.CopyIn()\n "

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyIn(
		nStrFmtAbsValDto,
		inComingNStrFmtAbsValDto,
		ePrefix+
			"inComingNStrFmtAbsValDto -> nStrFmtAbsValDto\n ")
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
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecAbsoluteValueDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyOut(
	ePrefix string) (NumStrFmtSpecAbsoluteValueDto, error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.CopyOut()\n "

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyOut(
		nStrFmtAbsValDto,
		ePrefix+
			"Copy -> 'nStrFmtAbsValDto'\n ")
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

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Absolute Value Dto.
//
// The NumberFieldDto details the length of the number field in which
// the signed numeric value will be displayed and right justified.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumberFieldLengthDto() NumberFieldDto {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.numFieldLenDto.CopyOut()
}

// GetNumberSeparatorsDto - Returns the NumStrFmtSpecDigitsSeparatorsDto
// instance currently configured for this Number String Format Specification
// Signed Number Value Dto.
//
// The Digits Separator Dto contains the decimal separator as well as the
// 'thousands' separators and the grouping sequence for separating thousands
// digits in the integer component of a number string.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumberSeparatorsDto() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.numberSeparatorsDto.CopyOut()
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
		"")

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
//       If this instance of NumStrFmtSpecAbsoluteValueDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.IsValidInstanceError() \n" +
		"Testing Validity of 'nStrFmtAbsValDto'\n "

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err := nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		nStrFmtAbsValDto,
		ePrefix)

	return err
}

// New() - Creates and returns a new instance of
// NumStrFmtSpecAbsoluteValueDto.
//
// This variant of the 'NumStrFmtSpecAbsoluteValueDto.New()'
// method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
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
//  NumStrFmtSpecAbsoluteValueDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecAbsoluteValueDto.
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
func (nStrFmtAbsValDto NumStrFmtSpecAbsoluteValueDto) New(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecAbsoluteValueDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	newNumStrFmtSpecAbsValueDto :=
		NumStrFmtSpecAbsoluteValueDto{}

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.New() \n" +
		"Setting data values for 'newNumStrFmtSpecAbsValueDto'\n "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	numberSeparatorsDto,
		err := NumStrFmtSpecDigitsSeparatorsDto{}.NewFromComponents(
		decimalSeparatorChar,
		thousandsSeparatorChar,
		[]uint{3},
		ePrefix)

	if err != nil {
		return newNumStrFmtSpecAbsValueDto, err
	}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err = nStrFmtSpecAbsValDtoMech.setAbsValDto(
		&newNumStrFmtSpecAbsValueDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewFromComponents() - Creates and returns a new instance of
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
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecAbsoluteValueDto,
	error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.NewFromComponents()\n "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	newNStrFmtSpecAbsoluteValueDto := NumStrFmtSpecAbsoluteValueDto{}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err := nStrFmtSpecAbsValDtoMech.setAbsValDto(
		&newNStrFmtSpecAbsoluteValueDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"Setting 'newNStrFmtSpecAbsoluteValueDto'\n ")

	return newNStrFmtSpecAbsoluteValueDto, err
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
//  ePrefix                    string
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetAbsoluteValueFormat(
	absoluteValueFormatStr string,
	ePrefix string) (
	err error) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	nStrAbsValDtoElectron :=
		numStrFmtSpecAbsoluteValueDtoElectron{}

	_,
		err = nStrAbsValDtoElectron.testAbsoluteValueFormatStr(
		absoluteValueFormatStr,
		ePrefix+"\n"+
			"Testing validity of 'absoluteValueFormatStr'\n ")

	if err != nil {
		return err
	}

	nStrFmtAbsValDto.absoluteValFmt = absoluteValueFormatStr

	return err
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
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecAbsoluteValueDto.SetNumberFieldLengthDto()\n"

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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetNumberSeparatorsDto(
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	_ = nStrFmtAbsValDto.numberSeparatorsDto.CopyIn(
		&numberSeparatorsDto,
		"")
}

// SetAbsoluteValueDto() - This method will set all of the member
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
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) SetAbsoluteValueDto(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	requestedNumberFieldLen int,
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.SetAbsoluteValueDto() "

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	return nStrFmtSpecAbsValDtoMech.setAbsValDto(
		nStrFmtAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix+
			"\n Setting 'nStrFmtAbsValDto'\n ")
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
