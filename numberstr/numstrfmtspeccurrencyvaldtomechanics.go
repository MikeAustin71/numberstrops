package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoMechanics struct {
	lock *sync.Mutex
}

// setCurrencyValDtoFromComponents - Transfers new data to an instance of
// NumStrFmtSpecCurrencyValueDto. After completion, all the data
// fields within input parameter 'nStrFmtSpecCurrencyValDto' will be
// overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecCurrencyValDto     *NumStrFmtSpecCurrencyValueDto,
//
//     - A pointer to an instance of NumStrFmtSpecCurrencyValueDto.
//       All of the data values in this object will be overwritten
//       set to new values based on the following input parameters.
//
//
//  positiveValueFmt              string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
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
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  negativeValueFmt              string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
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
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols                []rune
//     - The authorized character symbol associated with this
//       currency specification. In the United States, the currency
//       symbol is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The minor currency name. In the United States, the minor
//       currency name is 'Cent'.
//
//
//  minorCurrencySymbols            []rune
//     - These are the unicode characters for minor currency
//       symbols. In the United States, the minor currency symbol
//       is the cent sign (¢), represented by a single unicode
//       character ('\U000000a2'). Some countries and cultures have
//       currency symbols consisting of two or more characters.
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
//                 intSeparatorRepetitions uint   // Number of times this character/group is repeated
//                                                // A zero value signals unlimited repetitions.
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
//  err                           error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValMech *numStrFmtSpecCurrencyValueDtoMechanics) setCurrencyValDtoFromComponents(
	nStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValMech.lock == nil {
		nStrFmtSpecCurrValMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValMech.lock.Lock()

	defer nStrFmtSpecCurrValMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoMechanics.setCurrencyValDtoFromComponents()")

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is invalid!\n"+
			"'nStrFmtSpecCurrencyValDto' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if currencySymbols == nil {
		currencySymbols = make([]rune, 0, 5)
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"currencySymbols is a zero length rune array!\n",
			ePrefix.String())

		return err
	}

	if minorCurrencySymbols == nil {
		minorCurrencySymbols = make([]rune, 0, 5)
	}

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			fmt.Sprintf("Input parameter"+
				" 'positiveValueFmt' = '%v'",
				positiveValueFmt)))

	if err != nil {
		return err
	}

	_,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			fmt.Sprintf("Input parameter"+
				" 'negativeValueFmt' = '%v'",
				positiveValueFmt)))

	if err != nil {
		return err
	}

	lenMinorCurrSymbols := len(minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto := NumStrFmtSpecCurrencyValueDto{}

	newNStrFmtSpecCurrencyValDto.lock =
		new(sync.Mutex)

	newNStrFmtSpecCurrencyValDto.positiveValueFmt =
		positiveValueFmt

	newNStrFmtSpecCurrencyValDto.negativeValueFmt =
		negativeValueFmt

	newNStrFmtSpecCurrencyValDto.decimalDigits =
		decimalDigits

	newNStrFmtSpecCurrencyValDto.currencyCode =
		currencyCode

	newNStrFmtSpecCurrencyValDto.currencyCodeNo =
		currencyCodeNo

	newNStrFmtSpecCurrencyValDto.currencyName =
		currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	runesCopied := copy(newNStrFmtSpecCurrencyValDto.currencySymbols,
		currencySymbols)

	if runesCopied != lenCurrencySymbols {
		err = fmt.Errorf("%v\n"+
			"Error copying currency symbols!\n"+
			"Expected to copy %v currency symbol runes.\n"+
			"However, only %v currency symbol runes were copied.\n"+
			"Source currencySymbols = '%v'\n"+
			"Destination newNStrFmtSpecCurrencyValDto.currencySymbols = '%v'\n",
			ePrefix.String(),
			lenCurrencySymbols,
			runesCopied,
			string(currencySymbols),
			string(newNStrFmtSpecCurrencyValDto.currencySymbols))

		return err
	}

	newNStrFmtSpecCurrencyValDto.minorCurrencyName =
		minorCurrencyName

	newNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenMinorCurrSymbols, 10)

	runesCopied = copy(newNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		minorCurrencySymbols)

	if runesCopied != lenMinorCurrSymbols {
		err = fmt.Errorf("%v\n"+
			"Error copying minor currency symbols!\n"+
			"Expected to copy %v minor currency symbol runes.\n"+
			"However, only %v minor currency symbol runes were copied.\n"+
			"Source minorCurrencySymbols = '%v'\n"+
			"Destination newNStrFmtSpecCurrencyValDto.minorCurrencySymbols = '%v'\n",
			ePrefix.String(),
			lenMinorCurrSymbols,
			runesCopied,
			string(minorCurrencySymbols),
			string(newNStrFmtSpecCurrencyValDto.minorCurrencySymbols))

		return err
	}

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newNStrFmtSpecCurrencyValDto.numericSeparators.CopyIn(
			&numericSeparators,
			ePrefix.XCtx("numericSeparators->newNStrFmtSpecCurrencyValDto"))

	if err != nil {
		return err
	}

	err =
		newNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
			&numFieldLenDto,
			ePrefix.XCtx(" numFieldLenDto->newNStrFmtSpecCurrencyValDto"))

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	err =
		nStrFmtSpecCurrValNanobot.copyIn(
			nStrFmtSpecCurrencyValDto,
			&newNStrFmtSpecCurrencyValDto,
			ePrefix.XCtx("newNStrFmtSpecCurrencyValDto-> "+
				"nStrFmtSpecCurrencyValDto"))

	return err
}
