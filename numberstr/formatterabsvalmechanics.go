package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueMechanics struct {
	lock *sync.Mutex
}

// setAbsValDtoWithComponents - Sets the data value for incoming parameter
// 'nStrFmtSpecAbValDto', an instance of
// FormatterAbsoluteValue.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteVal                *FormatterAbsoluteValue
//     - A pointer to a FormatterAbsoluteValue object. All
//       of the data fields in this object will overwritten and set
//       to new values based on the following input parameters.
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
func (fmtAbsValMech *formatterAbsoluteValueMechanics) setAbsValDtoWithComponents(
	fmtAbsoluteVal *FormatterAbsoluteValue,
	absoluteValueFormat string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumericSeparators,
	numFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValMech.lock == nil {
		fmtAbsValMech.lock = new(sync.Mutex)
	}

	fmtAbsValMech.lock.Lock()

	defer fmtAbsValMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueMechanics." +
			"setAbsValDtoWithComponents()\n")

	if fmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrAbsValDtoElectron :=
		formatterAbsoluteValueElectron{}

	_,
		err =
		nStrAbsValDtoElectron.testAbsoluteValueFormatStr(
			absoluteValueFormat,
			ePrefix.XCtx(
				fmt.Sprintf(
					"absoluteValueFormat='%v'",
					absoluteValueFormat)))

	if err != nil {
		return err
	}

	newNStrFmtSpecAbsValDto :=
		FormatterAbsoluteValue{}

	newNStrFmtSpecAbsValDto.lock =
		new(sync.Mutex)

	newNStrFmtSpecAbsValDto.numStrFmtType =
		NumStrFormatTypeCode(0).AbsoluteValue()

	newNStrFmtSpecAbsValDto.absoluteValFmt =
		absoluteValueFormat

	newNStrFmtSpecAbsValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newNStrFmtSpecAbsValDto.numericSeparators.CopyIn(
			&numberSeparatorsDto,
			ePrefix.XCtx("'numericSeparators'->'newNStrFmtSpecAbsValDto'\n "))

	if err != nil {
		return err
	}

	err =
		newNStrFmtSpecAbsValDto.numFieldDto.CopyIn(
			&numFieldLenDto,
			ePrefix.XCtx("numFieldDto"))

	nStrFmtSpecAbsValDtoNanobot :=
		formatterAbsoluteValueNanobot{}

	err = nStrFmtSpecAbsValDtoNanobot.copyIn(
		fmtAbsoluteVal,
		&newNStrFmtSpecAbsValDto,
		ePrefix.XCtx(
			"newNStrFmtSpecAbsValDto -> fmtAbsoluteVal\n"))

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueMechanics.
//
func (fmtAbsValMech formatterAbsoluteValueMechanics) ptr() *formatterAbsoluteValueMechanics {

	if fmtAbsValMech.lock == nil {
		fmtAbsValMech.lock = new(sync.Mutex)
	}

	fmtAbsValMech.lock.Lock()

	defer fmtAbsValMech.lock.Unlock()

	newAbsValMechanics :=
		new(formatterAbsoluteValueMechanics)

	newAbsValMechanics.lock = new(sync.Mutex)

	return newAbsValMechanics
}
