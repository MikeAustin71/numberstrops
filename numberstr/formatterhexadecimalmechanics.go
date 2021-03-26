package numberstr

import (
	"fmt"
	"sync"
)

type formatterHexadecimalMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterHexadecimalMechanics.
//
func (fmtHexadecimalMech formatterHexadecimalMechanics) ptr() *formatterHexadecimalMechanics {

	if fmtHexadecimalMech.lock == nil {
		fmtHexadecimalMech.lock = new(sync.Mutex)
	}

	fmtHexadecimalMech.lock.Lock()

	defer fmtHexadecimalMech.lock.Unlock()

	newFmtHexadecimalMechanics :=
		new(formatterHexadecimalMechanics)

	newFmtHexadecimalMechanics.lock = new(sync.Mutex)

	return newFmtHexadecimalMechanics
}

// setFmtHexadecimalWithComponents - This method will overwrite and
// reset all the member data variable data values in input parameter
// 'formatterHex'.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// The new data values configured for 'formatterHex' will be
// generated from the FormatterHexadecimal components passed as
// input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterHex                  *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal. All the
//       member data variable data values in this instance will
//       overwritten and reset according to the following input
//       parameters.
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
func (fmtHexadecimalMech formatterHexadecimalMechanics) setFmtHexadecimalWithComponents(
	formatterHex *FormatterHexadecimal,
	useUpperCaseLetters bool,
	leftPrefix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalMech.lock == nil {
		fmtHexadecimalMech.lock = new(sync.Mutex)
	}

	fmtHexadecimalMech.lock.Lock()

	defer fmtHexadecimalMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexadecimalWithComponents()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	err = integerSeparators.IsValidInstanceError(
		ePrefix.XCtx("Checking integerSeparators"))

	if err != nil {
		return err
	}

	err = numFieldDto.IsValidInstanceError(
		ePrefix.XCtx(
			"Checking numFieldDto"))

	if err != nil {
		return err
	}

	newFormatterHexadecimal := FormatterHexadecimal{}

	newFormatterHexadecimal.lock = new(sync.Mutex)

	newFormatterHexadecimal.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	newFormatterHexadecimal.useUpperCaseLetters =
		useUpperCaseLetters

	newFormatterHexadecimal.leftPrefix =
		leftPrefix

	newFormatterHexadecimal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newFormatterHexadecimal.integerSeparators.
			CopyIn(
				&integerSeparators,
				ePrefix.XCtx(
					"integerSeparators->"+
						"newFormatterHexadecimal.integerSeparators"))

	if err != nil {
		return err
	}

	err = newFormatterHexadecimal.numFieldDto.
		CopyIn(
			&numFieldDto,
			ePrefix.XCtx(
				"numFieldDto->"+
					"newFormatterHexadecimal.numFieldDto"))

	if err != nil {
		return err
	}

	err =
		formatterHex.CopyIn(
			&newFormatterHexadecimal,
			ePrefix.XCtx(
				"newFormatterHexadecimal->"+
					"formatterHex"))

	return err
}
