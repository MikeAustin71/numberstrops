package numberstr

import (
	"fmt"
	"sync"
)

type formatterBinaryMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryMechanics.
//
func (fmtBinaryMech formatterBinaryMechanics) ptr() *formatterBinaryMechanics {

	if fmtBinaryMech.lock == nil {
		fmtBinaryMech.lock = new(sync.Mutex)
	}

	fmtBinaryMech.lock.Lock()

	defer fmtBinaryMech.lock.Unlock()

	newBinaryMechanics :=
		new(formatterBinaryMechanics)

	newBinaryMechanics.lock = new(sync.Mutex)

	return newBinaryMechanics
}

// setWithComponents - Transfers new data to an instance of
// FormatterBinary. After completion, all the data fields within
// input parameter 'formatterBinary' will be overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterBinary                *FormatterBinary
//     - A pointer to an instance of FormatterBinary. All of the
//       data values in this object will be overwritten and set to
//       new values based on the following input parameters.
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
//  numericSeparators        NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                     In some countries and cultures other integer groupings
//                     are used. In India, for example, a number might be
//                     formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//                     Numerals have an integer grouping value of four ('4').
//                         Chinese Numerals Example: '12,3456,7890,2345'
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
//
//
//  numFieldDto                NumberFieldDto
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
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtBinaryMech formatterBinaryMechanics) setWithComponents(
	formatterBinary *FormatterBinary,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryMech.lock == nil {
		fmtBinaryMech.lock = new(sync.Mutex)
	}

	fmtBinaryMech.lock.Lock()

	defer fmtBinaryMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryMechanics." +
			"setWithComponents()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if formatterBinary.lock == nil {
		formatterBinary.lock = new(sync.Mutex)
	}

	formatterBinary.numStrFmtType =
		NumStrFormatTypeCode(0).Binary()

	formatterBinary.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err = formatterBinary.numericSeparators.CopyIn(
		&numericSeparators,
		ePrefix.XCtx(
			"numericSeparators->"+
				"formatterBinary.numericSeparators"))

	if err != nil {
		return err
	}

	err = formatterBinary.numFieldDto.CopyIn(
		&numFieldLenDto,
		ePrefix.XCtx(
			"numFieldLenDto->"+
				"formatterBinary.numFieldDto"))

	return err
}
