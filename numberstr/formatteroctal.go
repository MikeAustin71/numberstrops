package numberstr

import "sync"

// FormatterOctal - This structure stores formatting data
// for octal numeric values displayed in text number strings.
// Member data elements are listed as follows:
//
//
// Reference:
//
//   https://en.wikipedia.org/wiki/Octal
//
type FormatterOctal struct {
	numStrFmtType                 NumStrFormatTypeCode
	leftPrefix                    string // A prefix added to beginning of number string
	rightSuffix                   string // A suffix or postfix added to end of number string.
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// FormatterOctal instance to the data fields of the current
// FormatterOctal instance.
//
// If input parameter 'incomingFormatterOctal' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current FormatterOctal
// instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterOctal     *FormatterOctal
//     - A pointer to an instance of FormatterOctal.
//       The data values in this object will be copied to those in
//       the current FormatterOctal instance.
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
func (fmtOctal *FormatterOctal) CopyIn(
	incomingFormatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"CopyIn()")

	return formatterOctalAtom{}.ptr().
		copyIn(
			fmtOctal,
			incomingFormatterOctal,
			ePrefix.XCtx(
				"incomingFormatterOctal->"+
					"fmtOctal"))
}

// CopyOut - Creates and returns a deep copy of the current
// FormatterOctal instance.
//
// If the current FormatterOctal instance is judged to be invalid,
// this method will return an error.
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
//  FormatterOctal
//     - If this method completes successfully, a new instance of
//       FormatterOctal will be created and returned containing all
//       of the data values copied from the current
//       FormatterHexadecimal instance.
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
func (fmtOctal *FormatterOctal) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterOctal,
	error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"CopyIn()")

	return formatterOctalAtom{}.ptr().
		copyOut(
			fmtOctal,
			ePrefix.XCtx(
				"fmtOctal->"))
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterOctal instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) Empty() {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	_ = formatterOctalQuark{}.ptr().
		empty(fmtOctal, nil)

	fmtOctal.lock.Unlock()

	fmtOctal.lock = nil

	return
}

// GetIntegerSeparators - Returns the NumStrIntSeparatorsDto
// object configured for the current FormatterOctal instance.
// This object constitutes the specification for integer digit
// separation applied to Octal Digits in a number string.
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
//       configured for the current FormatterOctal instance.
//
//       The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       Taken as a whole, this collection represents the
//       specification controlling the integer separation
//       operation in octal number strings.
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
func (fmtOctal *FormatterOctal) GetIntegerSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"GetIntegerSeparators()")

	return fmtOctal.integerSeparators.CopyOut(
		ePrefix.XCtx(
			"fmtOctal"))
}

// GetLeftPrefix - Returns the 'Left Prefix' string configured for
// the current FormatterOctal instance. Occasionally, octal numeric
// digits displayed in a text string are prefixed with leading
// characters such as the digit zero ('0'), the letters 'o' or 'q',
// the digit–letter combination zero-'o' ('0o'), the ampersand
// symbol ('&'), the dollar sign symbol ('$'), the At Sign
// ('@') or the back slash symbol ('\').
//
// The returned string identifies this prefix and as such, it will
// be prefixed to the beginning of octal number strings in text
// displays.
//
// Reference:
//    https://en.wikipedia.org/wiki/Octal
//
func (fmtOctal *FormatterOctal) GetLeftPrefix() string {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	return fmtOctal.leftPrefix
}

// GetRightSuffix - Returns the 'Right Suffix' string configured
// for the current FormatterOctal instance. Occasionally, octal numeric
// digits displayed in a text string are suffixed with trailing
// characters such as the the upper and lower case letters ('o','O',
// 'q' or 'Q').
//
// The returned string identifies this suffix and as such, it will
// be appended to the end of octal number strings in text displays.
//
// Reference:
//    https://en.wikipedia.org/wiki/Octal
//
func (fmtOctal *FormatterOctal) GetRightSuffix() string {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	return fmtOctal.rightSuffix
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Octal Number String Format
// Specification.
//
// The NumberFieldDto details the length of the number field in
// which the octal numeric value will be displayed and justified
// left, right or center according to the specification.
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
//       FormatterOctal.
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
func (fmtOctal *FormatterOctal) GetNumberFieldLengthDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"GetNumberFieldLengthDto()")

	return fmtOctal.numFieldDto.CopyOut(
		ePrefix.XCtx(
			"fmtOctal"))
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for FormatterOctal
// objects is NumStrFormatTypeCode(0).Octal().
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	return fmtOctal.numStrFmtType

}

// GetTurnOnInterDigitsSeparation - Returns the 'Turn On Integer
// Digits Separation" flag.
//   FormatterOctal.turnOnIntegerDigitsSeparation
//
// This boolean flag determines whether integer separation is
// applied to formatted octal digits in a number string.
//
// If set to 'true', this flag signals that octal digits will be
// grouped and separated according to the specification contained
// in the internal member variable, 'integerSeparators'.
//
// When set to 'false', this flag ensures that no integer digit
// separation will occur in formatted octal number strings.
//
func (fmtOctal *FormatterOctal) GetTurnOnInterDigitsSeparation() bool {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	return fmtOctal.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterOctal instance to determine whether that instance is
// valid in all respects.
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
//       current FormatterOctal is valid, or not. If the current
//       FormatterOctal instance contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
func (fmtOctal *FormatterOctal) IsValidInstance() (
	isValid bool) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	isValid,
		_ = formatterOctalQuark{}.ptr().
		testValidityOfFormatterOctal(
			fmtOctal,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterOctal instance to determine whether that
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
//     - If the current instance of FormatterOctal contains invalid
//       data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (fmtOctal *FormatterOctal) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterOctal."+
		"IsValidInstanceError()",
		"Testing Validity of 'formatterHex'")

	_,
		err :=
		formatterOctalQuark{}.ptr().
			testValidityOfFormatterOctal(
				fmtOctal,
				ePrefix)

	return err
}

// SetIntegerSeparators - Sets the internal member variable:
//    FormatterOctal.integerSeparators
//
// This object is of type NumStrIntSeparatorsDto. The
// NumStrIntSeparatorsDto object contains the specification for
// integer digit separation as applied to Octal Digits in a
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
//       operation in octal number strings.
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
func (fmtOctal *FormatterOctal) SetIntegerSeparators(
	integerSeparators NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetIntegerSeparators()")

	err := integerSeparators.IsValidInstanceError(
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}

	err =
		fmtOctal.integerSeparators.CopyIn(
			&integerSeparators,
			ePrefix.XCtx(
				"integerSeparators->"+
					"fmtOctal.integerSeparators"))

	return err
}

// SetLeftPrefix - Sets the 'Left Prefix' string configured for the
// current FormatterOctal instance.
//
// 'Left Prefix' string characters are appended or concatenated to
// the beginning of octal digits formatted in a text number string.
//
// Occasionally, octal numeric digits displayed in a text string
// are prefixed with leading characters such as the digit zero
// ('0'), the letters 'o' or 'q', the digit–letter combination
// zero-'o' ('0o'), the ampersand symbol ('&'), the dollar sign
// symbol ('$'), the At Sign ('@') or the back slash symbol ('\').
//
// The 'Left Prefix' parameter allows for custom configuration of
// these leading prefix characters.
//
// Input parameter 'leftPrefix', shown below, controls this prefix
// string and specifies the characters which will be prefixed to
// the beginning of octal number strings.
//
// Reference:
//    https://en.wikipedia.org/wiki/Octal
//
func (fmtOctal *FormatterOctal) SetLeftPrefix(
	leftPrefix string) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	fmtOctal.leftPrefix = leftPrefix
}

// SetRightSuffix - Sets the 'Right Suffix' string configured
// for the current FormatterOctal instance.
//
// Occasionally, octal numeric digits displayed in a text string
// are suffixed with trailing characters such as the the upper and
// lower case letters ('o','O', 'q' or 'Q').
//
// The 'Right Suffix' parameter allows for custom configuration of
// these trailing suffix characters.
//
// Input parameter 'rightSuffix', shown below, controls this suffix
// string and specifies the characters which will be appended to
// the end of octal number strings.
//
// Reference:
//    https://en.wikipedia.org/wiki/Octal
//
func (fmtOctal *FormatterOctal) SetRightSuffix(
	rightSuffix string) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	fmtOctal.rightSuffix = rightSuffix
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto
// object for the current FormatterHexadecimal instance.
//
// The Number Field Length Dto object is used to specify the length
// and string justification characteristics used to display
// octal number strings within a number field.
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
func (fmtOctal *FormatterOctal) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetNumberFieldLengthDto()")

	return fmtOctal.numFieldDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterOctal object. For Octal Number
// formatters, the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).Octal().
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) SetNumStrFormatTypeCode() {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	fmtOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Octal()

	return
}

// SetTurnOnInterDigitsSeparation - Sets the 'Turn On Integer
// Digits Separation" flag.
//   FormatterOctal.turnOnIntegerDigitsSeparation
//
// This boolean flag determines whether integer separation is
// applied to formatted octal digits in a number string.
//
// If set to 'true', this flag signals that octal digits
// will be grouped and separated according to the specification
// contained in the internal member variable, 'integerSeparators'.
//
// When set to 'false', this flag ensures that no integer digit
// separation will occur in formatted octal number strings.
//
func (fmtOctal *FormatterOctal) SetTurnOnInterDigitsSeparation(
	turnOnIntegerDigitsSeparation bool) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	fmtOctal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	return
}
