package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoUtility struct {
	lock *sync.Mutex
}

// setAbsValDtoWithDefaults() - Sets the data values for an
// instance FormatterAbsoluteValue passed as an input
// parameter.
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
//  nStrFmtSpecAbsValDto          *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All of the data values in this object will be overwritten
//       set to new values based on the following input parameters.
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
//  turnOnThousandsSeparator      bool
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
//
// -----------------------------------------------------------------
//
// Return Values
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
//       error message.
//
func (nStrFmtSpecAbsValDtoUtil *numStrFmtSpecAbsoluteValueDtoUtility) setAbsValDtoWithDefaults(
	nStrFmtSpecAbsValDto *FormatterAbsoluteValue,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecAbsValDtoUtil.lock == nil {
		nStrFmtSpecAbsValDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoUtil.lock.Lock()

	defer nStrFmtSpecAbsValDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecAbsoluteValueDtoUtility." +
			"setAbsValDtoWithDefaults()")

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is invalid!\n"+
			"'nStrFmtSpecAbsValDto' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewWithDefaults(
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	if err != nil {
		return err
	}

	var numberSeparatorsDto NumericSeparators

	intSeps := make([]NumStrIntSeparator, 1, 5)

	intSeps[0].intSeparatorChar = thousandsSeparatorChar
	intSeps[0].intSeparatorGrouping = 3
	intSeps[0].intSeparatorRepetitions = 0

	numberSeparatorsDto,
		err = NumericSeparators{}.NewWithComponents(
		decimalSeparatorChar,
		intSeps,
		ePrefix.XCtx(
			fmt.Sprintf("\n"+
				"decimalSeparatorChar='%v'\n"+
				"thousandsSeparatorChar='%v'\n",
				decimalSeparatorChar,
				thousandsSeparatorChar)))

	if err != nil {
		return err
	}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err = nStrFmtSpecAbsValDtoMech.setAbsValDtoWithComponents(
		nStrFmtSpecAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix)

	return err
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterAbsoluteValue instance
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
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecAbsValDto          *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All data values in this object will be overwritten and
//       set to United States default values for absolute numeric
//       values displayed in number strings.
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
func (nStrFmtSpecAbsValDtoUtil *numStrFmtSpecAbsoluteValueDtoUtility) setToUnitedStatesDefaults(
	nStrFmtSpecAbsValDto *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecAbsValDtoUtil.lock == nil {
		nStrFmtSpecAbsValDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoUtil.lock.Lock()

	defer nStrFmtSpecAbsValDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecAbsoluteValueDtoUtility." +
			"setToUnitedStatesDefaults()")

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is invalid!\n"+
			"'nStrFmtSpecAbsValDto' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewWithDefaults(
		-1,
		TextJustify(0).Right(),
		ePrefix)

	if err != nil {
		return err
	}

	var numberSeparatorsDto NumericSeparators

	intSeps := make([]NumStrIntSeparator, 1, 5)

	intSeps[0].intSeparatorChar = ','
	intSeps[0].intSeparatorGrouping = 3
	intSeps[0].intSeparatorRepetitions = 0

	numberSeparatorsDto,
		err = NumericSeparators{}.NewWithComponents(
		'.',
		intSeps,
		ePrefix.XCtx(
			"decimalSeparatorChar='%v' "+
				"thousandsSeparatorChar='%v'"))

	if err != nil {
		return err
	}
	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err = nStrFmtSpecAbsValDtoMech.setAbsValDtoWithComponents(
		nStrFmtSpecAbsValDto,
		"127.54",
		true,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix)

	return err
}
