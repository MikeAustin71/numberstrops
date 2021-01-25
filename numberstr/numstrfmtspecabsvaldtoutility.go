package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoUtility struct {
	lock *sync.Mutex
}

// setAbsValDtoWithDefaults() - Sets the data values for an
// instance NumStrFmtSpecAbsoluteValueDto passed as an input
// parameter.
//
// The NumStrFmtSpecAbsoluteValueDto type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecAbsValDto          *NumStrFmtSpecAbsoluteValueDto
//     - A pointer to an instance of NumStrFmtSpecAbsoluteValueDto.
//       All of the data values in this object will be overwritten
//       set to new values based on the following input parameters.
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
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  decimalSeparatorChar       rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar     rune
//     - The character which will be used to delimit 'thousands' in
//       integer number strings. In the United States, the Thousands
//       separator is the comma character (',').
//           Example: '1,000,000,000'
//
//
//  integerDigitsGroupingSequence  []uint
//     - Sets the integer digit grouping sequence for this instance
//       of NumStrFmtSpecAbsoluteValueDto. This grouping is
//       referred to as 'thousands' grouping when the integer
//       grouping is set to three digits (1,000,000,000).
//
//       In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecAbsValDtoUtil *numStrFmtSpecAbsoluteValueDtoUtility) setAbsValDtoWithDefaults(
	nStrFmtSpecAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	integerDigitsGroupingSequence []uint,
	requestedNumberFieldLen int,
	ePrefix string) (
	err error) {

	if nStrFmtSpecAbsValDtoUtil.lock == nil {
		nStrFmtSpecAbsValDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoUtil.lock.Lock()

	defer nStrFmtSpecAbsValDtoUtil.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecAbsoluteValueDtoUtility.setAbsValDtoWithDefaults() "

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is invalid!\n"+
			"'nStrFmtSpecAbsValDto' is a 'nil' pointer.\n",
			ePrefix)
		return err
	}

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	var numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto

	numberSeparatorsDto,
		err = NumStrFmtSpecDigitsSeparatorsDto{}.NewFromComponents(
		decimalSeparatorChar,
		thousandsSeparatorChar,
		integerDigitsGroupingSequence,
		ePrefix+
			fmt.Sprintf("\n"+
				"decimalSeparatorChar='%v'\n"+
				"thousandsSeparatorChar='%v'\n",
				decimalSeparatorChar,
				thousandsSeparatorChar))

	if err != nil {
		return err
	}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err = nStrFmtSpecAbsValDtoMech.setAbsValDto(
		nStrFmtSpecAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix)

	return err
}