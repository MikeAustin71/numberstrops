package numberstr

import (
	"fmt"
	"reflect"
	"sync"
)

// FormatterOctal - The FormatterOctal type encapsulates the
// formatting parameters necessary to format octal digits for
// display in text number strings.
//
// Reference:
//   https://en.wikipedia.org/wiki/Octal
//
//
// Member data elements are listed as follows:
//
//       type FormatterOctal struct {
//        numStrFmtType                 NumStrFormatTypeCode
//        leftPrefix                    string // A prefix added to beginning of number string
//        rightSuffix                   string // A suffix or postfix added to end of number string.
//        turnOnIntegerDigitsSeparation bool
//        integerSeparators             NumStrIntSeparatorsDto
//        numFieldDto                   NumberFieldDto
//        lock                          *sync.Mutex
//       }
//
//  numStrFmtType           NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).Octal()
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//        https://en.wikipedia.org/wiki/Octal
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
type FormatterOctal struct {
	numStrFmtType                 NumStrFormatTypeCode
	leftPrefix                    string // A prefix added to beginning of number string
	rightSuffix                   string // A suffix or postfix added to end of number string.
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming FormatterOctal
// instance to the data fields of the current FormatterOctal
// instance.
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// CopyInINumStrFormatter - Receives an incoming INumStrFormatter
// object, converts it to a FormatterOctal instance and
// proceeds to copy the the data fields into those of the
// current FormatterOctal instance.
//
// If the dynamic type of INumStrFormatter is not equal to
// FormatterOctal, an error will be returned. Likewise,
// if the data fields of input parameter 'incomingIFormatter' are
// judged to be invalid, an error will be returned.
//
// Be advised, all of the data fields in the current
// FormatterOctal instance will be overwritten.
//
// This method is required by interface INumStrFormatter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingIFormatter            INumStrFormatter
//     - An instance of Interface INumStrFormatter. If this the
//       dynamic type is not equal to FormatterOctal an error
//       will be returned.
//
//       The data values in this object will be copied to the
//       current FormatterOctal instance.
//
//       If input parameter 'incomingIFormatter' is judged to
//       be invalid, this method will return an error.
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
func (fmtOctal *FormatterOctal) CopyInINumStrFormatter(
	incomingIFormatter INumStrFormatter,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"CopyInINumStrFormatter()")

	if incomingIFormatter == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIFormatter' is "+
			"'nil'\n",
			ePrefix.String())
	}

	incomingFormatterOctal,
		isValid := incomingIFormatter.(*FormatterOctal)

	if !isValid {

		actualType :=
			reflect.TypeOf(incomingIFormatter)

		typeName := "Unknown"

		if actualType != nil {
			typeName = actualType.Name()
		}

		return fmt.Errorf("%v\n"+
			"Error: 'incomingIFormatter' is NOT Type "+
			"FormatterOctal.\n"+
			"'incomingIFormatter' is type %v",
			ePrefix.String(),
			typeName)
	}

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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// CopyOutINumStrFormatter - Creates and returns a deep copy of the current
// FormatterOctal instance as an INumStrFormatter object.
//
// If the current FormatterOctal instance is judged to be
// invalid, this method will return an error.
//
// This method is required by interface INumStrFormatter.
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
//  INumStrFormatter
//     - If this method completes successfully, a new instance of
//       FormatterOctal will be created, converted to an
//       instance of interface INumStrFormatter and returned
//       to the calling function. The returned data represents a
//       deep copy of the current FormatterOctal instance.
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
func (fmtOctal *FormatterOctal) CopyOutINumStrFormatter(
	ePrefix *ErrPrefixDto) (
	INumStrFormatter,
	error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"CopyOutINumStrFormatter()")

	newFormatterOctal,
		err := formatterOctalAtom{}.ptr().
		copyOut(
			fmtOctal,
			ePrefix.XCtx(
				"fmtOctal->"))

	return INumStrFormatter(&newFormatterOctal), err
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

}

// GetFmtNumStr - Returns a number string formatted for octal
// digits based on the configuration encapsulated in the current
// instance of FormatterOctal.
//
// This method is required by interface INumStrFormatter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absValIntegerRunes            []rune
//     - An array of runes containing the integer component of the
//       numeric value to be formatted as octal digits. This array
//       of integer digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  absValFractionalRunes         []rune
//     - An array of runes containing the fractional component of
//       the numeric value to be formatted as octal digits. This
//       array of numeric digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  signVal                       int
//     - The parameter designates the numeric sign of the final
//       formatted octal numeric value returned by this method.
//
//       Valid values for 'signVal' are listed as follows:
//         -1 = Signals a negative numeric value
//          0 = Signals that the numeric value is zero.
//         +1 = Signals a positive numeric value
//
//
//  baseNumSys                    BaseNumberSystemType
//     - Designates the base number system associated with input
//       parameters 'absValIntegerRunes' and 'absValIntegerRunes'
//       described above.
//
//       Valid values for 'baseNumSys' are listed as follows:
//         BaseNumSys.Binary(),
//         BaseNumSys.Octal(),
//         BaseNumSys.Decimal(),
//         BaseNumSys.Hexadecimal(),
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  fmtNumStr                     string
//     - If this method completes successfully, a formatted
//       octal number string will be returned.
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
func (fmtOctal *FormatterOctal) GetFmtNumStr(
	absValIntegerRunes []rune,
	absValFractionalRunes []rune,
	signVal int,
	baseNumSys BaseNumberSystemType,
	ePrefix *ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"GetFmtNumStr()")

	if baseNumSys != BaseNumberSystemType(0).Octal() {
		err = fmt.Errorf("%v\n"+
			"Error: Base Numbering System is NOT equal to Base-8!\n"+
			"baseNumSys=='%v'\n",
			ePrefix.String(),
			baseNumSys.XValueInt())
		return fmtNumStr, err
	}

	if signVal < 0 {
		fmtNumStr = "-"
	}

	fmtNumStr += string(absValIntegerRunes) +
		string(absValFractionalRunes)

	return fmtNumStr, err
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// GetNumberFieldDto - Returns the NumberFieldDto object
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
//       Be advised that if the returned NumberFieldDto object is
//       judged invalid, this method will return an error.
//
func (fmtOctal *FormatterOctal) GetNumberFieldDto(
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
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"GetNumberFieldDto()")

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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// NewDetail - Creates and returns a new instance of
// FormatterOctal, 'newFormatterOctal'.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// Data values for 'newFormatterOctal' will be generated from
// the FormatterOctal components passed as input parameters.
//
// This method differs from FormatterOctal.NewDetailRunes()
// in that this method accepts a string for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterOctal, reference method:
//   'FormatterOctal.NewWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators       string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterOctal.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions       uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//  newFormatterOctal             FormatterOctal
//     - If this method completes successfully, new instance of
//       FormatterOctal will be returned configured according
//       to the input parameters listed above.
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
//
func (fmtOctal FormatterOctal) NewDetail(
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	newFormatterOctal FormatterOctal,
	err error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"NewWithComponents()")

	newFormatterOctal.lock = new(sync.Mutex)

	err = formatterOctalUtility{}.ptr().
		setFmtOctalDetail(
			&newFormatterOctal,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			leftPrefix,
			rightSuffix,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterOctal"))

	return newFormatterOctal, err
}

// NewDetailRunes - Creates and returns a new instance of
// FormatterOctal, 'newFormatterOctal'.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// Data values for 'newFormatterOctal' will be generated from
// the FormatterOctal components passed as input parameters.
//
// This method differs from FormatterOctal.NewDetail()
// in that this method accepts a rune array for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterOctal, reference method:
//   'FormatterOctal.NewWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators       string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterOctal.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions       uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//  newFormatterOctal             FormatterOctal
//     - If this method completes successfully, new instance of
//       FormatterOctal will be returned configured according
//       to the input parameters listed above.
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
//
func (fmtOctal FormatterOctal) NewDetailRunes(
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	newFormatterOctal FormatterOctal,
	err error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"NewDetailRunes()")

	newFormatterOctal.lock = new(sync.Mutex)

	err = formatterOctalUtility{}.ptr().
		setFmtOctalDetailRunes(
			&newFormatterOctal,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			leftPrefix,
			rightSuffix,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterOctal"))

	return newFormatterOctal, err
}

// NewWithComponents - Creates and returns a new instance of
// FormatterOctal, 'newFormatterOctal'.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// Data values for 'newFormatterOctal' will be generated from
// the FormatterOctal components passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
//  newFormatterOctal             FormatterOctal
//     - If this method completes successfully, new instance of
//       FormatterOctal will be returned configured according
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
func (fmtOctal FormatterOctal) NewWithComponents(
	leftPrefix string,
	rightSuffix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	newFormatterOctal FormatterOctal,
	err error) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"NewWithComponents()")

	newFormatterOctal.lock = new(sync.Mutex)

	err = formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			&newFormatterOctal,
			leftPrefix,
			rightSuffix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"newFormatterOctal"))

	return newFormatterOctal, err
}

// NewWithDefaults - Creates and returns a new FormatterOctal format
// specification. The returned instance is generated using defaults
// for octal number string formatting.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// Member variable data fields in the returned FormatterOctal
// instance are configured as follows:
//
// Number String Format Type (numStrFmtType) =
//     NumStrFormatTypeCode(0).Octal()
//
// Octal Number String Left Prefix String (leftPrefix) = ""
//
// Octal Number String Right Suffix String (rightSuffix) = ""
//
// Turn On Integer Digits Separation
//  (turnOnIntegerDigitsSeparation) = false
//
// Integer Separators (integerSeparators) =
//   Integer Digit Separator  = ","
//   Integer Digit Grouping Sequence = 3
//   Integer Group Repetitions = 0
//
// Number Field Dto = Off
//   Requested Number Field Length = -1
//   Text Justify Format (textJustifyFormat) =
//     TextJustify(0).Right
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
//  FormatterOctal
//     - If this method completes successfully, a new instance of
//       FormatterOctal configured with default values will be
//       returned.
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
func (fmtOctal FormatterOctal) NewWithDefaults(
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
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"NewWithComponents()")

	newFormatterOctal := FormatterOctal{}

	newFormatterOctal.lock = new(sync.Mutex)

	err := formatterOctalUtility{}.ptr().
		setFmtOctalWithDefaults(
			&newFormatterOctal,
			ePrefix.XCtx(
				"newFormatterOctal"))

	return newFormatterOctal, err
}

// SetDetail - Overwrites and resets all the data values in
// the current instance of FormatterOctal.
//
// Data values used to replace those in the current FormatterOctal
// instance will be generated from components passed as input
// parameters and detailed below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// This method differs from FormatterOctal.SetDetailRunes()
// in that this method accepts a string for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterOctal, reference method:
//   'FormatterOctal.SetWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterOctal instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators       string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterOctal.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions       uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//
func (fmtOctal *FormatterOctal) SetDetail(
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetDetail()")

	return formatterOctalUtility{}.ptr().
		setFmtOctalDetail(
			fmtOctal,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			leftPrefix,
			rightSuffix,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"fmtOctal"))
}

// SetDetailRunes - Overwrites and resets all the data values in
// the current instance of FormatterOctal.
//
// Data values used to replace those in the current FormatterOctal
// instance will be generated from components passed as input
// parameters and detailed below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// This method differs from FormatterOctal.SetDetail()
// in that this method accepts a rune array for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterOctal, reference method:
//   'FormatterOctal.SetWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterOctal instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators       []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterOctal.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions       uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//
func (fmtOctal *FormatterOctal) SetDetailRunes(
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetDetailRunes()")

	return formatterOctalUtility{}.ptr().
		setFmtOctalDetailRunes(
			fmtOctal,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			leftPrefix,
			rightSuffix,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"fmtOctal"))
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// SetNumberFieldDto - Sets the Number Field Length Dto
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
func (fmtOctal *FormatterOctal) SetNumberFieldDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetNumberFieldDto()")

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

// SetWithComponents - Overwrites and resets all the data values in
// the current instance of FormatterOctal.
//
// Data values used to replace those in the current FormatterOctal
// instance will be generated from components passed as input
// parameters and detailed below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterOctal instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
func (fmtOctal *FormatterOctal) SetWithComponents(
	leftPrefix string,
	rightSuffix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetWithComponents()")

	return formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			fmtOctal,
			leftPrefix,
			rightSuffix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"newFormatterOctal"))
}

// SetWithDefaults - Overwrites and resets all the data fields
// within the current instance of FormatterOctal. The new data
// values are generated from default parameters specified below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// Member variable data fields in the current FormatterOctal
// instance are configured as follows:
//
// Number String Format Type (numStrFmtType) =
//     NumStrFormatTypeCode(0).Octal()
//
// Octal Number String Left Prefix String (leftPrefix) = ""
//
// Octal Number String Right Suffix String (rightSuffix) = ""
//
// Turn On Integer Digits Separation
//  (turnOnIntegerDigitsSeparation) = false
//
// Integer Separators (integerSeparators) =
//   Integer Digit Separator  = ","
//   Integer Digit Grouping Sequence = 3
//   Integer Group Repetitions = 0
//
// Number Field Dto = Off
//   Requested Number Field Length = -1
//   Text Justify Format (textJustifyFormat) =
//     TextJustify(0).Right
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterOctal instance.
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
func (fmtOctal *FormatterOctal) SetWithDefaults(
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterOctal." +
			"SetWithDefaults()")

	return formatterOctalUtility{}.ptr().
		setFmtOctalWithDefaults(
			fmtOctal,
			ePrefix.XCtx(
				"fmtOctal"))
}
