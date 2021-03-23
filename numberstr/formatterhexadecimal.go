package numberstr

import "sync"

// FormatterHexadecimal - This structure stores formatting data
// for hexadecimal numeric values displayed in text number strings.
// Member data elements are listed as follows:
//
//  numStrFmtType           NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).Hexadecimal()
//
//
//  useUpperCaseLetters           bool
//     - Sets the 'Use Upper Case Letters' flag. This boolean flag
//       determines whether Hexadecimal Digits 'A' through 'F' will
//       be expressed as Upper Case Letters or Lower Case Letters
//       ('a' through 'f').
//
//       If input parameter 'useUpperCaseLetters' is set to 'true',
//       hexadecimal alphabetic digits will be expressed as upper
//       case letters. Conversely, if the parameter is 'false',
//       hexadecimal alphabetic digits will be expressed as lower
//       case letters.
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string which is concatenated to the
//       beginning of hexadecimal digits formatted in a number
//       string.
//
//       Often, hexadecimal digits displayed in a text string are
//       prefixed with the characters, "0x". Example:  '0x14AF'
//
//       Input parameter 'leftPrefix' controls this prefix string.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Sets the 'Turn On Integer Digits Separation" flag.
//          FormatterHexadecimal.turnOnIntegerDigitsSeparation
//
//       This boolean flag determines whether integer separation is
//       applied to formatted hexadecimal digits in a number
//       string.
//
//       If set to 'true', this flag signals that hexadecimal
//       digits will be grouped and separated according to the
//       specification contained in the internal member variable,
//       'integerSeparators'. When set to 'false', this flag
//       ensures that no integer digit separation will occur in
//       formatted hexadecimal number strings.
//
//
//  integerSeparators             NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
// Reference
//
//   https://en.wikipedia.org/wiki/Hexadecimal
type FormatterHexadecimal struct {
	numStrFmtType                 NumStrFormatTypeCode
	useUpperCaseLetters           bool
	leftPrefix                    string // '0x'
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// FormatterHexadecimal instance  to the data fields
// of the current instance of FormatterHexadecimal
// instance.
//
// If input parameter 'incomingFormatterHex' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current
// FormatterHexadecimal instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterHex    *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal.
//       The data values in this object will be copied to the
//       current FormatterHexadecimal instance.
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
func (formatterHex *FormatterHexadecimal) CopyIn(
	incomingFormatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"CopyIn()")

	return formatterHexadecimalElectron{}.ptr().
		copyIn(
			formatterHex,
			incomingFormatterHex,
			ePrefix.XCtx(
				"formatterHex"))
}

// CopyOut - Creates and returns a deep copy of the current
// FormatterHexadecimal instance.
//
// If the current FormatterHexadecimal instance is judged
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
//  FormatterHexadecimal
//     - If this method completes successfully, a new instance of
//       FormatterHexadecimal will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterHexadecimal.
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
func (formatterHex *FormatterHexadecimal) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterHexadecimal,
	error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"CopyOut()")

	return formatterHexadecimalElectron{}.ptr().
		copyOut(
			formatterHex,
			ePrefix.XCtx(
				"formatterHex"))
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterHexadecimal instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) Empty() {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	_ = formatterHexadecimalQuark{}.ptr().
		empty(formatterHex, nil)

	formatterHex.lock.Unlock()

	formatterHex.lock = nil
}

// GetIntegerSeparators - Returns the NumStrIntSeparatorsDto
// object configured for the current FormatterHexadecimal instance.
// This object constitutes the specification for integer digit
// separation applied to Hexadecimal Digits in a number string.
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
//     - Returns a deep copy of the NumStrIntSeparatorsDto object
//       configured for the current FormatterHexadecimal instance.
//
//       The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       Taken as a whole, this collection represents the
//       specification controlling the integer separation
//       operation in hexadecimal number strings.
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
func (formatterHex *FormatterHexadecimal) GetIntegerSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"GetIntegerSeparators()")

	return formatterHex.integerSeparators.CopyOut(
		ePrefix.XCtx(
			"formatterHex"))
}

// GetLeftPrefix - Returns the 'Left Prefix' string configured for
// the current FormatterHexadecimal instance. Often, hexadecimal
// digits displayed in a text string are prefixed with the
// characters, "0x".
//   Example:  '0x14AF'
//
// The returned string identifies this prefix and as such, it will
// be prefixed to the beginning of hexadecimal number strings in
// text displays.
//
// Reference:
//    https://en.wikipedia.org/wiki/Hexadecimal
//
func (formatterHex *FormatterHexadecimal) GetLeftPrefix() string {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.leftPrefix
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Hexadecimal Number String Format
// Specification.
//
// The NumberFieldDto details the length of the number field in
// which the hexadecimal numeric value will be displayed and
// justified left, right or center according to the specification.
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
//       FormatterHexadecimal.
//
//       The NumberFieldDto object contains formatting instructions
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
func (formatterHex *FormatterHexadecimal) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"GetNumberFieldLengthDto()")

	return formatterHex.numFieldDto.CopyOut(
		ePrefix.XCtx(
			"formatterHex"))
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for
// FormatterHexadecimal objects is
// NumStrFormatTypeCode(0).Hexadecimal().
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.numStrFmtType
}

// GetTurnOnInterDigitsSeparation - Returns the 'Turn On Integer
// Digits Separation" flag.
//   FormatterHexadecimal.turnOnIntegerDigitsSeparation
//
// This boolean flag determines whether integer separation is
// applied to formatted hexadecimal digits in a number string.
//
// If set to 'true', this flag signals that hexadecimal digits
// will be grouped and separated according to the specification
// contained in the internal member variable, 'integerSeparators'.
//
// When set to 'false', this flag ensures that no integer digit
// separation will occur in formatted octal number strings.
//
func (formatterHex *FormatterHexadecimal) GetTurnOnInterDigitsSeparation() bool {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.turnOnIntegerDigitsSeparation
}

// GetUseUpperCaseLetters - Returns the 'Use Upper Case Letters'
// flag. This boolean flag determines whether Hexadecimal Digits
// 'A' through 'F' will be expressed as Upper Case Letters or Lower
// Case Letters ('a' through 'f').
//
func (formatterHex *FormatterHexadecimal) GetUseUpperCaseLetters() bool {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.useUpperCaseLetters
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterHexadecimal instance to determine whether that instance
// is valid in all respects.
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
//     - This returned boolean value will signal whether the
//       current FormatterHexadecimal instance is valid, or not. If
//       the current FormatterHexadecimal contains valid data, this
//       method returns 'true'. If the data is invalid, this method
//       will return 'false'.
//
func (formatterHex *FormatterHexadecimal) IsValidInstance() (
	isValid bool) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	isValid,
		_ = formatterHexadecimalQuark{}.ptr().
		testValidityOfFormatterHexadecimal(
			formatterHex,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterHexadecimal instance to determine whether that
// instance is valid in all respects.
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
//     - If the current instance of FormatterHexadecimal contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (formatterHex *FormatterHexadecimal) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterHexadecimal.IsValidInstanceError()",
		"Testing Validity of 'formatterHex'")

	_,
		err :=
		formatterHexadecimalQuark{}.ptr().
			testValidityOfFormatterHexadecimal(
				formatterHex,
				ePrefix)

	return err
}

// NewWithComponents - Creates and returns a new instance of
// FormatterHexadecimal, 'newFormatterHexadecimal' .
//
// Data values for 'newFormatterHexadecimal' will be generated from
// the FormatterHexadecimal components passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  useUpperCaseLetters           bool
//     - Sets the 'Use Upper Case Letters' flag. This boolean flag
//       determines whether Hexadecimal Digits 'A' through 'F' will
//       be expressed as Upper Case Letters or Lower Case Letters
//       ('a' through 'f').
//
//       If input parameter 'useUpperCaseLetters' is set to 'true',
//       hexadecimal alphabetic digits will be expressed as upper
//       case letters. Conversely, if the parameter is 'false',
//       hexadecimal alphabetic digits will be expressed as lower
//       case letters.
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string which is concatenated to the
//       beginning of hexadecimal digits formatted in a number
//       string.
//
//       Often, hexadecimal digits displayed in a text string are
//       prefixed with the characters, "0x". Example:  '0x14AF'
//
//       Input parameter 'leftPrefix' controls this prefix string.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Sets the 'Turn On Integer Digits Separation" flag.
//          FormatterHexadecimal.turnOnIntegerDigitsSeparation
//
//       This boolean flag determines whether integer separation is
//       applied to formatted hexadecimal digits in a number
//       string.
//
//       If set to 'true', this flag signals that hexadecimal
//       digits will be grouped and separated according to the
//       specification contained in the internal member variable,
//       'integerSeparators'. When set to 'false', this flag
//       ensures that no integer digit separation will occur in
//       formatted hexadecimal number strings.
//
//
//  integerSeparators             NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//  newFormatterHexadecimal       FormatterHexadecimal
//     - If this method completes successfully, new instance of
//       FormatterHexadecimal will be returned configured according
//       to the input parameters listed above.
//
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
func (formatterHex FormatterHexadecimal) NewWithComponents(
	useUpperCaseLetters bool,
	leftPrefix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	newFormatterHexadecimal FormatterHexadecimal,
	err error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal.NewWithComponents()")

	err =
		formatterHexadecimalMechanics{}.ptr().
			setFmtHexadecimalWithComponents(
				&newFormatterHexadecimal,
				useUpperCaseLetters,
				leftPrefix,
				turnOnIntegerDigitsSeparation,
				integerSeparators,
				numFieldDto,
				ePrefix.XCtx(
					"newFormatterHexadecimal"))

	return newFormatterHexadecimal, err
}

// NewWithDefaults - Creates and returns a new FormatterHexadecimal format
// specification. The returned instance is generated using defaults
// for hexadecimal number string formatting.
//
// Member variable data fields in the returned FormatterHexadecimal
// instance are configured as follows:
//
//  numStrFmtType = NumStrFormatTypeCode(0).Hexadecimal()
//  useUpperCaseLetters = false (Alphabetic letters will be
//                              displayed as 'a' through 'f')
//  leftPrefix = "0x" All hexadecimal number strings will be
//                    prefixed with "0x". Example: '0x2b'
//
//  turnOnIntegerDigitsSeparation = false No integer digit
//                                        separation will be
//                                        applied.
//
//  integerSeparators = None
//
//  numFieldDto.requestedNumFieldLength = -1
//                   Number Field = Number String Length
//
//  numFieldDto.textJustifyFormat = TextJustify(0).Right()
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  FormatterHexadecimal
//     - If this method completes successfully, a new instance of
//       FormatterHexadecimal configured with default values will
//       be returned.
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
func (formatterHex FormatterHexadecimal) NewWithDefaults(
	ePrefix *ErrPrefixDto) (
	FormatterHexadecimal,
	error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"NewWithDefaults()")

	newFmtHexadecimal := FormatterHexadecimal{}

	newFmtHexadecimal.lock = new(sync.Mutex)

	err := formatterHexadecimalUtility{}.ptr().
		setFmtHexadecimalWithDefaults(
			&newFmtHexadecimal,
			ePrefix.XCtx(
				"newFmtHexadecimal"))

	return newFmtHexadecimal, err
}

// SetIntegerSeparators - Sets the internal member variable:
//    FormatterHexadecimal.integerSeparators
//
// This object is of type NumStrIntSeparatorsDto. The
// NumStrIntSeparatorsDto object contains the specification for
// integer digit separation as applied to Hexadecimal Digits in a
// number string.
//
// If input parameter 'integerSeparators' is judged to be invalid,
// this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerSeparators             NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       Taken as a whole, this collection represents the
//       specification controlling the integer separation
//       operation in hexadecimal number strings.
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
func (formatterHex *FormatterHexadecimal) SetIntegerSeparators(
	integerSeparators NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"SetIntegerSeparators()")

	err := integerSeparators.IsValidInstanceError(
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}

	err =
		formatterHex.integerSeparators.CopyIn(
			&integerSeparators,
			ePrefix.XCtx(
				"integerSeparators->"+
					"formatterHex.integerSeparators"))

	return err
}

// SetLeftPrefix - Sets the 'Left Prefix' string which is
// concatenated to the beginning of the hexadecimal digits
// when they are formatted in a number string.
//
// Often, hexadecimal digits displayed in a text string are
// prefixed with the characters, "0x".
//   Example:  0x14AF
//
// Input parameter 'leftPrefix' controls this prefix string.
//
func (formatterHex *FormatterHexadecimal) SetLeftPrefix(
	leftPrefix string) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.leftPrefix = leftPrefix
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto
// object for the current FormatterHexadecimal instance.
//
// The Number Field Length Dto object is used to specify the length
// and string justification characteristics used to display
// hexadecimal number strings within a number field.
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
func (formatterHex *FormatterHexadecimal) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterCurrency." +
			"SetNumberFieldLengthDto()")

	return formatterHex.numFieldDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterHexadecimal object. For Signed Number
// formatters, the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).Hexadecimal().
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) SetNumStrFormatTypeCode() {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	return
}

// SetTurnOnInterDigitsSeparation - Sets the 'Turn On Integer
// Digits Separation" flag.
//   FormatterHexadecimal.turnOnIntegerDigitsSeparation
//
// This boolean flag determines whether integer separation is
// applied to formatted hexadecimal digits in a number string.
//
// If set to 'true', this flag signals that hexadecimal digits
// will be grouped and separated according to the specification
// contained in the internal member variable, 'integerSeparators'.
//
// When set to 'false', this flag ensures that no integer digit
// separation will occur in formatted hexadecimal number strings.
//
func (formatterHex *FormatterHexadecimal) SetTurnOnInterDigitsSeparation(
	turnOnIntegerDigitsSeparation bool) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	return
}

// SetUseUpperCaseLetters - Sets the 'Use Upper Case Letters' flag.
// This boolean flag determines whether Hexadecimal Digits 'A'
// through 'F' will be expressed as Upper Case Letters or Lower
// Case Letters ('a' through 'f').
//
// If input parameter 'useUpperCaseLetters' is set to 'true',
// hexadecimal alphabetic digits will be expressed as upper case
// letters. Conversely, if the parameter is 'false', hexadecimal
// alphabetic digits will be expressed as lower case letters.
//
func (formatterHex *FormatterHexadecimal) SetUseUpperCaseLetters(
	useUpperCaseLetters bool) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.useUpperCaseLetters =
		useUpperCaseLetters

	return
}

// SetWithComponents - This method will overwrite and reset all the
// member data variable data values in the current instance of
// FormatterHexadecimal.
//
// The new data values configured for the current
// FormatterHexadecimal instance will be generated from the basic
// components passed as input parameters and described below.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  useUpperCaseLetters           bool
//     - Sets the 'Use Upper Case Letters' flag. This boolean flag
//       determines whether Hexadecimal Digits 'A' through 'F' will
//       be expressed as Upper Case Letters or Lower Case Letters
//       ('a' through 'f').
//
//       If input parameter 'useUpperCaseLetters' is set to 'true',
//       hexadecimal alphabetic digits will be expressed as upper
//       case letters. Conversely, if the parameter is 'false',
//       hexadecimal alphabetic digits will be expressed as lower
//       case letters.
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string which is concatenated to the
//       beginning of hexadecimal digits formatted in a number
//       string.
//
//       Often, hexadecimal digits displayed in a text string are
//       prefixed with the characters, "0x". Example:  '0x14AF'
//
//       Input parameter 'leftPrefix' controls this prefix string.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Sets the 'Turn On Integer Digits Separation" flag.
//          FormatterHexadecimal.turnOnIntegerDigitsSeparation
//
//       This boolean flag determines whether integer separation is
//       applied to formatted hexadecimal digits in a number
//       string.
//
//       If set to 'true', this flag signals that hexadecimal
//       digits will be grouped and separated according to the
//       specification contained in the internal member variable,
//       'integerSeparators'. When set to 'false', this flag
//       ensures that no integer digit separation will occur in
//       formatted hexadecimal number strings.
//
//
//  integerSeparators             NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//  newFormatterHexadecimal       FormatterHexadecimal
//     - If this method completes successfully, new instance of
//       FormatterHexadecimal will be returned configured according
//       to the input parameters listed above.
//
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
func (formatterHex *FormatterHexadecimal) SetWithComponents(
	useUpperCaseLetters bool,
	leftPrefix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"SetWithComponents()")

	return formatterHexadecimalMechanics{}.ptr().
		setFmtHexadecimalWithComponents(
			formatterHex,
			useUpperCaseLetters,
			leftPrefix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterHex"))
}

// SetWithDefaults - This method will overwrite and reset all the
// member variable data values in the current FormatterHexadecimal
// instance.
//
// Default values will be applied and member variable data values
// will be configured as follows:
//
//  numStrFmtType = NumStrFormatTypeCode(0).Hexadecimal()
//
//  useUpperCaseLetters = false (Alphabetic letters will be
//                              displayed as 'a' through 'f')
//
//  leftPrefix = "0x" All hexadecimal number strings will be
//                    prefixed with "0x". Example: '0x2b'
//
//  turnOnIntegerDigitsSeparation = false No integer digit
//                                        separation will be
//                                        applied.
//
//  integerSeparators = Set to blank space (" ")
//
//  numFieldDto.requestedNumFieldLength = -1
//                   Number Field = Number String Length
//
//  numFieldDto.textJustifyFormat = TextJustify(0).Right()
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
func (formatterHex *FormatterHexadecimal) SetWithDefaults(
	ePrefix *ErrPrefixDto) error {

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"SetWithDefaults()")

	return formatterHexadecimalUtility{}.ptr().
		setFmtHexadecimalWithDefaults(
			formatterHex,
			ePrefix.XCtx(
				"formatterHex"))
}
