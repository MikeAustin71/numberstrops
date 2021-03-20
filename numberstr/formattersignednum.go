package numberstr

import (
	"fmt"
	"sync"
)

type FormatterSignedNumber struct {
	numStrFormatterType           NumStrFormatTypeCode
	positiveValueFmt              string
	negativeValueFmt              string
	turnOnIntegerDigitsSeparation bool
	numericSeparators             NumericSeparators
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// FormatterSignedNumber instance  to the data fields
// of the current instance of FormatterSignedNumber
// instance.
//
// If input parameter 'incomingSignedNumValDto' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current
// FormatterSignedNumber instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingSignedNumValDto    *FormatterSignedNumber
//     - A pointer to an instance of FormatterSignedNumber.
//       The data values in this object will be copied to the
//       current FormatterSignedNumber instance.
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
func (fmtSignedNum *FormatterSignedNumber) CopyIn(
	incomingSignedNumValDto *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"CopyIn()")

	nStrFmtSpecSignedNumValNanobot :=
		formatterSignedNumberNanobot{}

	return nStrFmtSpecSignedNumValNanobot.copyIn(
		fmtSignedNum,
		incomingSignedNumValDto,
		ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// FormatterSignedNumber instance.
//
// If the current FormatterSignedNumber instance is judged
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
//  FormatterSignedNumber
//     - If this method completes successfully, a new instance of
//       FormatterSignedNumber will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterSignedNumber.
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
func (fmtSignedNum *FormatterSignedNumber) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterSignedNumber,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber.CopyOut()")

	nStrFmtSpecSignedNumValNanobot :=
		formatterSignedNumberNanobot{}

	return nStrFmtSpecSignedNumValNanobot.copyOut(
		fmtSignedNum,
		ePrefix.XCtx(
			"Copy Out from 'fmtSignedNum'"))
}

// Empty - Deletes and resets the data values of all member
// variables within a FormatterSignedNumber instance to their
// initial 'zero' values.
//
func (fmtSignedNum *FormatterSignedNumber) Empty() {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	_ = formatterSignedNumberQuark{}.ptr().
		empty(fmtSignedNum, nil)

	fmtSignedNum.lock.Unlock()

	fmtSignedNum.lock = nil

	return
}

// GetDecimalSeparator - Returns the decimal separator character(s).
// This is the character (runes) used to separate integer and
// fractional digits in a floating point number.
//
// In the United States, the Decimal Separator character is the
// decimal point or period ('.').
//
//  Example:   123.45
//
// Decimal Separator is extracted from the underlying member
// variable, 'nStrFmtSpecSignedNumValueDto.numericSeparators'.
//
func (fmtSignedNum *FormatterSignedNumber) GetDecimalSeparator() []rune {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	return fmtSignedNum.
		numericSeparators.
		GetDecimalSeparators()
}

// GetIntegerDigitSeparators - Returns an array of type
// NumStrIntSeparator. The data contained in type
// NumStrIntSeparator is used to separate integer digits.
//
// The returned integer digit separators are those configured
// for the current instance of FormatterSignedNumber.
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
//  NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       the integer separation operation in number strings.
//
//        type NumStrIntSeparator struct {
//            intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//            intSeparatorGrouping    uint    // Number of integer digits in a group
//            intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                            // A zero value signals unlimited repetitions.
//            restartIntGroupingSequence bool // If true, the entire grouping sequence is repeated
//                                            //  beginning at array index zero.
//        }
//
//        intSeparatorChars          []rune
//           - A series of runes or characters used to separate integer
//             digits in a number string. These characters are commonly
//             known as the 'thousands separator'. A 'thousands
//             separator' is used to separate groups of integer digits to
//             the left of the decimal separator (a.k.a. decimal point).
//             In the United States, the standard integer digits
//             separator is the single comma character (','). Other
//             countries and cultures use periods, spaces, apostrophes or
//             multiple characters to separate integers.
//                   United States Example:  1,000,000,000
//
//        intSeparatorGrouping       uint
//           - This unsigned integer values specifies the number of
//             integer digits within a group. This value is used to group
//             integers within a number string.
//
//             In most western countries integer digits to the left of
//             the decimal separator (a.k.a. decimal point) are separated
//             into groups of three digits representing a grouping of
//             'thousands' like this: '1,000,000,000'. In this case the
//             intSeparatorGrouping value would be set to three ('3').
//
//             In some countries and cultures other integer groupings are
//             used. In India, for example, a number might be formatted like
//             this: '6,78,90,00,00,00,00,000'. Chinese Numerals have an
//             integer grouping value of four ('4').
//                Chinese Numerals Example: '12,3456,7890,2345'
//
//        intSeparatorRepetitions    uint
//           - This unsigned integer value specifies the number of times
//             this integer grouping is repeated. A value of zero signals
//             that this integer grouping will be repeated indefinitely.
//
//        restartIntGroupingSequence bool
//           - If the NumStrIntSeparator is the last element in an array
//             of NumStrIntSeparator objects, this boolean flag signals
//             whether the entire integer grouping sequence will be
//             restarted from array element zero.
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
func (fmtSignedNum *FormatterSignedNumber) GetIntegerDigitSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"GetIntegerDigitSeparators()")

	return fmtSignedNum.
		numericSeparators.
		GetIntegerDigitSeparators(
			ePrefix.XCtx(
				"nStrFmtSpecCurrValDto.numericSeparators"))
}

// GetNegativeValueFormat - Returns the formatting string used to
// format negative signed number values in text number strings.
//
func (fmtSignedNum *FormatterSignedNumber) GetNegativeValueFormat() string {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	return fmtSignedNum.negativeValueFmt
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Signed Number Value Dto.
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
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
//       FormatterSignedNumber.
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
func (fmtSignedNum *FormatterSignedNumber) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("FormatterSignedNumber.GetNumberFieldLengthDto()")

	return fmtSignedNum.numFieldLenDto.CopyOut(
		ePrefix.XCtx(
			"fmtSignedNum.numFieldLenDto=>"))
}

// GetNumericSeparators - Returns a deep copy of the
// NumericSeparators instance currently configured for this
// Signed Number Format Specification.
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
// FormatterSignedNumber.
//
// If the FormatterSignedNumber or NumericSeparators object
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
//       instance of FormatterSignedNumber.
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
//       If the FormatterSignedNumber or NumericSeparators
//       object is judged to be invalid, this method will return
//       an error.
//
func (fmtSignedNum *FormatterSignedNumber) GetNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"GetNumericSeparators()")

	return fmtSignedNum.numericSeparators.CopyOut(
		ePrefix.XCtx(
			"fmtSignedNum.numericSeparators ->"))
}

// GetPositiveValueFormat - Returns the formatting string used to
// format positive signed number values in text number strings.
//
func (fmtSignedNum *FormatterSignedNumber) GetPositiveValueFormat() string {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	return fmtSignedNum.positiveValueFmt
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
func (fmtSignedNum *FormatterSignedNumber) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	return fmtSignedNum.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterSignedNumber instance to determine whether
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
//       FormatterSignedNumber is valid, or not. If the
//       current FormatterSignedNumber contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (fmtSignedNum *FormatterSignedNumber) IsValidInstance() (
	isValid bool) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	nStrFmtSpecSignedNumValMolecule :=
		formatterSignedNumberMolecule{}

	isValid,
		_ = nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
		fmtSignedNum,
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterSignedNumber instance to determine
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
//     - If the current instance of FormatterSignedNumber
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (fmtSignedNum *FormatterSignedNumber) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterSignedNumber.IsValidInstanceError()",
		"Testing Validity of 'fmtSignedNum'")

	nStrFmtSpecSignedNumValMolecule :=
		formatterSignedNumberMolecule{}

	_,
		err :=
		nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
			fmtSignedNum,
			ePrefix)

	return err
}

// NewWithComponents - Creates and returns a new instance of
// FormatterSignedNumber.
//
// The FormatterSignedNumber type encapsulates the
// configuration parameters necessary to format signed numeric
// values for display in text number strings.
//
// This method requires detailed input parameters to control
// all configuration parameters for the returned new instance
// of FormatterSignedNumber.
//
// For a 'New' method using minimum input parameters coupled
// with default values, see:
//      FormatterSignedNumber.NewWithDefaults()
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
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       included in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparators              rune
//         integerSeparatorsDto []NumStrIntSeparator
//        }
//
//        decimalSeparators              rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerSeparatorsDto             []NumStrIntSeparator
//           - An array of NumStrIntSeparator elements used to specify
//             the integer separation operation.
//
//               type NumStrIntSeparator struct {
//                 intSeparatorChar     rune // Integer separator character
//                 intSeparatorGrouping uint // Number of integers in a group
//                 intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                             // A zero value signals unlimited repetitions.
//               }
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
//                           intSeparatorRepetitions: 0,
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
//             this case 'integerSeparatorsDto' would be configured as
//             follows:
//             as:
//
//             numSeps.intSeparators =
//               []NumStrIntSeparator{
//                    {
//                    intSeparatorChar:   ',',
//                    intSeparatorGrouping: 3,
//                    intSeparatorRepetitions: 1,
//                    },
//                    {
//                    intSeparatorChar:     ',',
//                    intSeparatorGrouping: 2,
//                    intSeparatorRepetitions: 0,
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
//  FormatterSignedNumber
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterSignedNumber.
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
func (fmtSignedNum FormatterSignedNumber) NewWithComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	FormatterSignedNumber,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"NewWithComponents()")

	newNStrFmtSpecSignedNumValueDto := FormatterSignedNumber{}

	nStrFmtSpecSignedNumValMech :=
		formatterSignedNumberMechanics{}

	err := nStrFmtSpecSignedNumValMech.setFmtSignedNumWithComponents(
		&newNStrFmtSpecSignedNumValueDto,
		positiveValueFmt,
		negativeValueFmt,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
}

// NewWithDefaults - Creates and returns a new instance of
// FormatterSignedNumber. This method specifies the
// minimum number of input parameters required to construct a new
// instance of FormatterSignedNumber. Default values are
// used to supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterSignedNumber,
// reference method:
//   'FormatterSignedNumber.NewWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterSignedNumber.NewWithComponents()'.
//
// The FormatterSignedNumber type encapsulates the
// formatting parameters necessary to format signed numeric values
// for display in text number strings.
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
//       FormatterSignedNumber.NewWithComponents().
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
//  FormatterSignedNumber
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterSignedNumber.
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
func (fmtSignedNum FormatterSignedNumber) NewWithDefaults(
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	turnOnThousandsSeparator bool,
	positiveValueFmt string,
	negativeValueFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterSignedNumber,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"NewBasic()")

	newNStrFmtSpecSignedNumValueDto :=
		FormatterSignedNumber{}

	nStrFmtSpecSignedNumValDtoUtil :=
		formatterSignedNumberUtility{}

	err :=
		nStrFmtSpecSignedNumValDtoUtil.setBasicRunesSignedNumber(
			&newNStrFmtSpecSignedNumValueDto,
			decimalSeparatorChar,
			thousandsSeparatorChar,
			turnOnThousandsSeparator,
			positiveValueFmt,
			negativeValueFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
}

// NewWithFmtSpecSetupDto - Creates and returns a new
// FormatterSignedNumber instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     *NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       FormatterSignedNumber object. Only the following
//       data fields with a prefix of "SignedNumVal" are used.
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
//  FormatterSignedNumber
//     - If this method completes successfully, a new instance of
//       FormatterSignedNumber will be returned to the
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
func (fmtSignedNum FormatterSignedNumber) NewWithFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	FormatterSignedNumber,
	error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"NewWithFmtSpecSetupDto()")

	newNStrFmtSpecSignedNumValueDto :=
		FormatterSignedNumber{}

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

	nStrFmtSpecSignedNumValMech :=
		formatterSignedNumberMechanics{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	err :=
		nStrFmtSpecSignedNumValMech.
			setFmtSignedNumWithComponents(
				&newNStrFmtSpecSignedNumValueDto,
				fmtSpecSetupDto.SignedNumValPositiveValueFmt,
				fmtSpecSetupDto.SignedNumValNegativeValueFmt,
				fmtSpecSetupDto.SignedNumValTurnOnIntegerDigitsSeparation,
				fmtSpecSetupDto.SignedNumValNumSeps,
				fmtSpecSetupDto.SignedNumValNumField,
				ePrefix.XCtx(
					"Setting 'newNStrFmtSpecSignedNumValueDto'"))

	return newNStrFmtSpecSignedNumValueDto, err
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
func (fmtSignedNum *FormatterSignedNumber) SetNegativeValueFormat(
	negativeValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("FormatterSignedNumber.SetNegativeValueFormat()")

	nStrSignedNumElectron :=
		formatterSignedNumberElectron{}

	_,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormatStr(
		negativeValueFmt,
		ePrefix.XCtx(
			"Testing validity of 'negativeValueFmt'"))

	if err != nil {
		return err
	}

	fmtSignedNum.negativeValueFmt = negativeValueFmt

	return err
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current FormatterSignedNumber instance.
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
func (fmtSignedNum *FormatterSignedNumber) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber.SetNumberFieldLengthDto()")

	return fmtSignedNum.numFieldLenDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumberSeparatorsDto - Sets the Number Separators Dto object
// for the current FormatterSignedNumber instance.
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
//         decimalSeparators              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparators rune
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
func (fmtSignedNum *FormatterSignedNumber) SetNumberSeparatorsDto(
	numberSeparatorsDto NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("FormatterSignedNumber.SetNumericSeparators()")

	return fmtSignedNum.numericSeparators.CopyIn(
		&numberSeparatorsDto,
		ePrefix.XCtx("numericSeparators->fmtSignedNum.numericSeparators"))
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
func (fmtSignedNum *FormatterSignedNumber) SetPositiveValueFormat(
	positiveValueFmt string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("FormatterSignedNumber.SetPositiveValueFormat()")

	nStrSignedNumElectron :=
		formatterSignedNumberElectron{}

	_,
		err = nStrSignedNumElectron.testSignedNumValPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			"Testing validity of 'positiveValueFmt'"))

	if err != nil {
		return err
	}

	fmtSignedNum.positiveValueFmt = positiveValueFmt

	return err
}

// SetToUnitedStatesDefaults - Sets the member variable data values
// for this FormatterSignedNumber instance to United
// States default values.
//
// In the United States, Signed Number default formatting
// parameters are defined as follows:
//
//   Positive Signed Number format: "127.54"
//   Negative Signed Number format: "-127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterSignedNumber instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  err                           error
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
func (fmtSignedNum *FormatterSignedNumber) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"SetToUnitedStatesDefaults()")

	signedNumUtil := formatterSignedNumberUtility{}

	err = signedNumUtil.setToUnitedStatesDefaults(
		fmtSignedNum,
		ePrefix.XCtx(
			"fmtSignedNum"))

	return err
}

// SetToUnitedStatesDefaultsIfEmpty - If the current
// FormatterSignedNumber instance is empty or invalid,
// this method will set the member variable data values to United
// States default values.
//
// If the current FormatterSignedNumber instance is valid
// and NOT empty, this method will take no action and exit.
//
// In the United States, Signed Number default formatting values
// are defined as follows:
//
//   Positive Signed Number format: "127.54"
//   Negative Signed Number format: "-127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// IMPORTANT
//
// If the current FormatterSignedNumber instance is empty
// or invalid, this method will overwrite all pre-existing data
// values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  err                           error
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
func (fmtSignedNum *FormatterSignedNumber) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"SetToUnitedStatesDefaults()")

	nStrFmtSpecSignedNumValMolecule :=
		formatterSignedNumberMolecule{}

	isValid,
		_ := nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
		fmtSignedNum,
		new(ErrPrefixDto))

	if isValid {
		return err
	}

	signedNumUtil := formatterSignedNumberUtility{}

	err = signedNumUtil.setToUnitedStatesDefaults(
		fmtSignedNum,
		ePrefix)

	return err
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// FormatterSignedNumber.
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
func (fmtSignedNum *FormatterSignedNumber) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	fmtSignedNum.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithComponents - This method will set all of the member
// variable data values for the current instance of
// FormatterSignedNumber.
//
// The FormatterSignedNumber type encapsulates the
// formatting parameters necessary for formatting signed number
// values in text number strings.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterSignedNumber instance.
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
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       included in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparators              rune
//         integerSeparatorsDto []NumStrIntSeparator
//        }
//
//        decimalSeparators              rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerSeparatorsDto             []NumStrIntSeparator
//           - An array of NumStrIntSeparator elements used to specify
//             the integer separation operation.
//
//               type NumStrIntSeparator struct {
//                 intSeparatorChar     rune   // Integer separator character
//                 intSeparatorGrouping uint   // Number of integers in a group
//                 intSeparatorRepetitions uint   // Number of times this character/group is repeated
//                                                // A zero value signals unlimited repetitions.
//               }
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
//                           intSeparatorRepetitions: 0,
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
//             this case 'integerSeparatorsDto' would be configured as
//             follows:
//             as:
//
//             numSeps.intSeparators =
//               []NumStrIntSeparator{
//                    {
//                    intSeparatorChar:   ',',
//                    intSeparatorGrouping: 3,
//                    intSeparatorRepetitions: 1,
//                    },
//                    {
//                    intSeparatorChar:     ',',
//                    intSeparatorGrouping: 2,
//                    intSeparatorRepetitions: 0,
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
func (fmtSignedNum *FormatterSignedNumber) SetWithComponents(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"SetWithComponents()")

	nStrFmtSpecSignedNumValMech :=
		formatterSignedNumberMechanics{}

	return nStrFmtSpecSignedNumValMech.
		setFmtSignedNumWithComponents(
			fmtSignedNum,
			positiveValueFmt,
			negativeValueFmt,
			turnOnIntegerDigitsSeparation,
			numberSeparatorsDto,
			numFieldDto,
			ePrefix.XCtx("Setting 'fmtSignedNum'"))
}

// SetWithDefaults - This method will set all of the member
// variable data values for the current instance of
// FormatterSignedNumber. The input parameters
// represent the minimum information required to configure
// the data values for a FormatterSignedNumber object.
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterSignedNumber instance.
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
func (fmtSignedNum *FormatterSignedNumber) SetWithDefaults(
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
			"SetBasicRunes()")

	nStrFmtSpecSignedNumValDtoUtil :=
		formatterSignedNumberUtility{}

	return nStrFmtSpecSignedNumValDtoUtil.setBasicRunesSignedNumber(
		fmtSignedNum,
		decimalSeparatorChar,
		thousandsSeparatorChar,
		turnOnIntegerDigitsSeparation,
		positiveValueFmt,
		negativeValueFmt,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix.XCtx("Setting 'fmtSignedNum'"))
}

// SetWithFmtSpecSetupDto - Sets the data values for current
// FormatterSignedNumber instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterSignedNumber instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup and configuration
//       information for a FormatterSignedNumber
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
func (fmtSignedNum *FormatterSignedNumber) SetWithFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	fmtSignedNum.lock.Lock()

	defer fmtSignedNum.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterSignedNumber." +
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

	nStrFmtSpecSignedNumValMech :=
		formatterSignedNumberMechanics{}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	return nStrFmtSpecSignedNumValMech.
		setFmtSignedNumWithComponents(
			fmtSignedNum,
			fmtSpecSetupDto.SignedNumValPositiveValueFmt,
			fmtSpecSetupDto.SignedNumValNegativeValueFmt,
			fmtSpecSetupDto.SignedNumValTurnOnIntegerDigitsSeparation,
			fmtSpecSetupDto.SignedNumValNumSeps,
			fmtSpecSetupDto.SignedNumValNumField,
			ePrefix.XCtx("Setting 'fmtSignedNum'"))
}