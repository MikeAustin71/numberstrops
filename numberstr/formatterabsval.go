package numberstr

import (
	"sync"
)

type FormatterAbsoluteValue struct {
	numStrFmtType                 NumStrFormatTypeCode
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numericSeparators             NumericSeparators
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// FormatterAbsoluteValue ('inComingNStrFmtAbsValDto') to
// the data fields of the current FormatterAbsoluteValue
// instance.
//
// If 'inComingNStrFmtAbsValDto' is judged to be invalid, this
// method will return an error.
//
// Be advised, all of the data fields in the current
// FormatterAbsoluteValue instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterAbsVal    *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       The data values in this object will be copied to the
//       current FormatterAbsoluteValue instance.
//
//
//  ePrefix                    *ErrPrefixDto
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyIn(
	incomingFormatterAbsVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.CopyIn()")

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyIn(
		fmtAbsVal,
		incomingFormatterAbsVal,
		ePrefix.XCtx("incomingFormatterAbsVal -> fmtAbsVal"))
}

// CopyOut - Returns a deep copy of the current
// FormatterAbsoluteValue instance.
//
// If the current FormatterAbsoluteValue instance is judged
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, a new instance of
//       FormatterAbsoluteValue will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyOut(
	ePrefix *ErrPrefixDto) (FormatterAbsoluteValue, error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.CopyOut()")

	nStrFmtSpecAbsValDtoNanobot :=
		numStrFmtSpecAbsoluteValueDtoNanobot{}

	return nStrFmtSpecAbsValDtoNanobot.copyOut(
		fmtAbsVal,
		ePrefix.XCtx(
			"Copy -> 'fmtAbsVal'"))
}

// GetAbsoluteValueFormat - Returns the formatting string used to
// format absolute numeric values in text number strings.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetAbsoluteValueFormat() string {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.absoluteValFmt
}

// GetDecimalSeparators - Returns the unicode character(s) (runes)
// used to separate integer and fractional digits in a floating
// point number.
//
// In the United States, the Decimal Separator character is the
// decimal point or period ('.').
//
//  Example:   123.45
//
// Decimal Separator is extracted from the underlying member
// variable, 'nStrFmtAbsValDto.numericSeparators'.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetDecimalSeparators() []rune {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.
		numericSeparators.
		GetDecimalSeparators()
}

// GetIntegerDigitSeparators - Returns an array of type
// NumStrIntSeparator. The data contained in type
// NumStrIntSeparator is used to separate integer digits.
//
// The returned integer digit separators are those configured
// for the current instance of FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) GetIntegerDigitSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"GetIntegerDigitSeparators() ")

	return fmtAbsVal.
		numericSeparators.
		GetIntegerDigitSeparators(
			ePrefix.XCtx(
				"fmtAbsVal.numericSeparators"))
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
//       FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"GetNumberFieldLengthDto()")

	return fmtAbsVal.numFieldLenDto.CopyOut(
		ePrefix.XCtx(
			"fmtAbsVal.numFieldLenDto->"))
}

// GetNumericSeparators - Returns a deep copy of the
// NumericSeparators instance currently configured for this
// Absolute Value Format Specification.
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
// FormatterAbsoluteValue.
//
// If the FormatterAbsoluteValue or NumericSeparators object
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
//       instance of FormatterAbsoluteValue.
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
//       If the FormatterAbsoluteValue or NumericSeparators
//       object is judged to be invalid, this method will return
//       an error.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.GetNumericSeparators()")

	return fmtAbsVal.numericSeparators.CopyOut(
		ePrefix.XCtx("fmtAbsVal.numericSeparators->"))
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
func (fmtAbsVal *FormatterAbsoluteValue) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterAbsoluteValue instance to determine whether
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
//       value will signal whether the current FormatterAbsoluteValue
//       is valid, or not. If the current FormatterAbsoluteValue
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (fmtAbsVal *FormatterAbsoluteValue) IsValidInstance() (
	isValid bool) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	isValid,
		_ = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		fmtAbsVal,
		ErrPrefixDto{}.Ptr())

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// FormatterAbsoluteValue instance to determine whether the
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
func (fmtAbsVal *FormatterAbsoluteValue) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterAbsoluteValue.IsValidInstanceError()",
		"Testing Validity of 'fmtAbsVal'")

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err := nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		fmtAbsVal,
		ePrefix)

	return err
}

// NewBasic - Creates and returns a new instance of
// FormatterAbsoluteValue. This method specifies the minimum
// number of input parameters required to construct a new instance
// of FormatterAbsoluteValue. Default values are used to
// supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterAbsoluteValue.NewWithComponents()'.
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChar          string
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        string
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
//       FormatterAbsoluteValue.NewWithComponents().
//
//
//  turnOnThousandsSeparator   bool
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewBasic()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	nStrFmtSpecAbsValDtoUtil :=
		formatterAbsoluteValueUtility{}

	err := nStrFmtSpecAbsValDtoUtil.
		setBasic(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewBasicRunes - Creates and returns a new instance of
// FormatterAbsoluteValue. This method specifies the minimum
// number of input parameters required to construct a new instance
// of FormatterAbsoluteValue. Default values are used to
// supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterAbsoluteValue.NewWithComponents()'.
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChar          []rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        []rune
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
//       FormatterAbsoluteValue.NewWithComponents().
//
//
//  turnOnThousandsSeparator   bool
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewBasic()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	nStrFmtSpecAbsValDtoUtil :=
		formatterAbsoluteValueUtility{}

	err := nStrFmtSpecAbsValDtoUtil.
		setBasicRunes(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewWithComponents - Creates and returns a new instance of
// FormatterAbsoluteValue.
//
// The FormatterAbsoluteValue type encapsulates the
// configuration parameters necessary to format absolute numeric
// values for display in text number strings.
//
// This method requires detailed input parameters which provide
// granular control over all data fields contained in the returned
// new instance of FormatterAbsoluteValue.
//
// For a 'New' method using minimum input parameters coupled
// with default values, see:
//      FormatterAbsoluteValue.NewBasicRunes()
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
//                 intSeparatorRepetitions uint  // Number of times this character/group is repeated
//                                               // A zero value signals unlimited repetitions.
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of
//       FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewWithComponents(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewWithComponents()")

	newNumStrFmtSpecAbsValueDto := FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err := nStrFmtSpecAbsValDtoMech.setAbsValDtoWithComponents(
		&newNumStrFmtSpecAbsValueDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(

			"Setting 'newNStrFmtSpecAbsoluteValueDto'"))

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
func (fmtAbsVal *FormatterAbsoluteValue) SetAbsoluteValueFormat(
	absoluteValueFormatStr string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.SetAbsoluteValueFormat()")

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

	fmtAbsVal.absoluteValFmt = absoluteValueFormatStr

	return err
}

// SetBasic - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// The input parameters represent the minimum information required
// to configure the data values for a FormatterAbsoluteValue
// object. Default values are used to supplement these input
// parameters.
//
// This method differs from FormatterAbsoluteValue.SetBasicRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChars       string
//     - The character or characters which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterAbsoluteValue.SetWithComponents().
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
//       FormatterAbsoluteValue.NewWithComponents().
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
func (fmtAbsVal *FormatterAbsoluteValue) SetBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetBasicRunes()")

	return formatterAbsoluteValueUtility{}.ptr().
		setBasic(
			fmtAbsVal,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)
}

// SetBasicRunes - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// The input parameters represent the minimum information required
// to configure the data values for a FormatterAbsoluteValue
// object. Default values are used to supplement these input
// parameters.
//
// This method differs from FormatterAbsoluteValue.SetBasic() in
// that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         []rune
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChars       []rune
//     - The character or characters which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterAbsoluteValue.SetWithComponents().
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
//       FormatterAbsoluteValue.NewWithComponents().
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
func (fmtAbsVal *FormatterAbsoluteValue) SetBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetBasicRunes()")

	nStrFmtSpecAbsValDtoUtil :=
		formatterAbsoluteValueUtility{}

	return nStrFmtSpecAbsValDtoUtil.setBasicRunes(
		fmtAbsVal,
		decimalSeparatorChars,
		thousandsSeparatorChars,
		turnOnThousandsSeparator,
		absoluteValFmt,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.SetNumberFieldLengthDto()")

	return fmtAbsVal.numFieldLenDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumericSeparators - Sets the Number Separators Dto object
// for the current FormatterAbsoluteValue instance.
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
//  --- NONE ---
//
func (fmtAbsVal *FormatterAbsoluteValue) SetNumericSeparators(
	numberSeparators NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetNumericSeparators()")

	return fmtAbsVal.numericSeparators.CopyIn(
		&numberSeparators,
		ePrefix)
}

// SetToUnitedStatesDefaults - Sets the member variable data
// values for the current FormatterAbsoluteValue instance
// to United States Default values.
//
// In the United States, Absolute Value default formatting
// parameters are defined as follows:
//
//    Absolute Value Number format: "127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetToUnitedStatesDefaults()")

	absValUtil := formatterAbsoluteValueUtility{}

	err = absValUtil.setToUnitedStatesDefaults(
		fmtAbsVal,
		ePrefix.XCtx("fmtAbsVal"))

	return err
}

// SetToUnitedStatesDefaultsIfEmpty - If the current
// FormatterAbsoluteValue instance is empty or invalid,
// this method will set the member variable data values to United
// States default values.
//
// If the current FormatterAbsoluteValue instance is valid
// and NOT empty, this method will take no action and exit.
//
// In the United States, Absolute Value default formatting
// parameters are defined as follows:
//
//    Absolute Value Number format: "127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetToUnitedStatesDefaults()")

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	isValid,
		_ := nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		fmtAbsVal,
		ErrPrefixDto{}.Ptr())

	absValUtil := formatterAbsoluteValueUtility{}

	if isValid {
		return err
	}

	err = absValUtil.setToUnitedStatesDefaults(
		fmtAbsVal,
		ePrefix.XCtx("fmtAbsVal"))

	return err
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	fmtAbsVal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithComponents - This method will set all of the member
// variable data values for the current instance of
// FormatterAbsoluteValue.
//
// The FormatterAbsoluteValue type encapsulates the
// configuration parameters necessary to format absolute numeric
// values for display in text number strings.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
//                 intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                              // A zero value signals unlimited repetitions.
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
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtAbsVal *FormatterAbsoluteValue) SetWithComponents(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetWithComponents()")

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	return nStrFmtSpecAbsValDtoMech.setAbsValDtoWithComponents(
		fmtAbsVal,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"Setting 'fmtAbsVal'"))
}
