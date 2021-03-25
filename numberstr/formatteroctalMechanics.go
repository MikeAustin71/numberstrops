package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterOctalMechanics.
//
func (fmtOctalMech formatterOctalMechanics) ptr() *formatterOctalMechanics {

	if fmtOctalMech.lock == nil {
		fmtOctalMech.lock = new(sync.Mutex)
	}

	fmtOctalMech.lock.Lock()

	defer fmtOctalMech.lock.Unlock()

	newFmtOctalMech :=
		new(formatterOctalMechanics)

	newFmtOctalMech.lock = new(sync.Mutex)

	return newFmtOctalMech
}

// setFmtOctalWithComponents - Overwrites and resets all the data
// values within input parameter 'formatterOctal'.
//
// Data values used to replace those within input parameter
// 'formatterOctal' will be generated from components passed as
// input parameters and detailed below.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterOctal                *FormatterOctal
//     - A pointer to an instance of FormatterOctal. Member
//       variable data values within this instance will be
//       overwritten and reset to new values based on the input
//       parameters listed below.
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
// Reference:
//    https://en.wikipedia.org/wiki/Octal
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
func (fmtOctalMech *formatterOctalMechanics) setFmtOctalWithComponents(
	formatterOctal *FormatterOctal,
	leftPrefix string,
	rightSuffix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalMech.lock == nil {
		fmtOctalMech.lock = new(sync.Mutex)
	}

	fmtOctalMech.lock.Lock()

	defer fmtOctalMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterOctalMechanics." +
			"setFmtOctalWithComponents()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

	err = integerSeparators.IsValidInstanceError(
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}

	err = numFieldDto.IsValidInstanceError(
		ePrefix.XCtx(
			"numFieldDto"))

	if err != nil {
		return err
	}

	newFormatterOctal := FormatterOctal{}

	newFormatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Octal()

	newFormatterOctal.leftPrefix =
		leftPrefix

	newFormatterOctal.rightSuffix =
		rightSuffix

	newFormatterOctal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newFormatterOctal.integerSeparators.CopyIn(
			&integerSeparators,
			ePrefix.XCtx(
				"integerSeparators->"+
					"newFormatterOctal"))

	if err != nil {
		return err
	}

	err =
		newFormatterOctal.numFieldDto.CopyIn(
			&numFieldDto,
			ePrefix.XCtx(
				"numFieldDto->"+
					"newFormatterOctal"))

	if err != nil {
		return err
	}

	err = formatterOctalAtom{}.ptr().
		copyIn(
			formatterOctal,
			&newFormatterOctal,
			ePrefix.XCtx(
				"newFormatterOctal->"+
					"formatterOctal"))

	return err
}
