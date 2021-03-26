package numberstr

import (
	"fmt"
	"sync"
)

type formatterHexadecimalUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterHexadecimalUtility.
//
func (fmtHexadecimalUtil formatterHexadecimalUtility) ptr() *formatterHexadecimalUtility {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	newFmtHexadecimalUtility :=
		new(formatterHexadecimalUtility)

	newFmtHexadecimalUtility.lock = new(sync.Mutex)

	return newFmtHexadecimalUtility
}

// setFmtHexDetail - This method will set all of the member
// variable data values for the FormatterHexadecimal input
// parameter, 'formatterHex'. New data values will be generated
// from the input parameters described below.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// This method differs from formatterHexadecimalUtility.setFmtHexDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterCurrency, reference method:
//   'FormatterHexadecimal.SetWithComponents()'
//
// The member variable 'FormatterHexadecimal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Hexadecimal()
//
func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexDetail(
	formatterHex *FormatterHexadecimal,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	useUpperCaseLetters bool,
	leftPrefix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexWithDefaults()")

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

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewDetail(
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}
	var numFieldDto NumberFieldDto

	numFieldDto,
		err =
		NumberFieldDto{}.NewBasic(
			requestedNumberFieldLen,
			numberFieldTextJustify,
			nil)

	if err != nil {
		return err
	}

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

// setFmtHexWithDefaults - Receives a pointer to an
// instance of FormatterHexadecimal and proceeds to set the data
// values to standard defaults.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// Member variable data fields within input parameter
// 'formatterHex' are configured as follows:
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
//  formatterHex                  *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal. All the
//       member data variable data values in this instance will
//       overwritten and reset according to the following input
//       parameters.
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
func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexWithDefaults(
	formatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexWithDefaults()")

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

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewBasic(
		" ",
		nil)

	if err != nil {
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err =
		NumberFieldDto{}.NewBasic(
			-1,
			TextJustify(0).Right(),
			nil)

	if err != nil {
		return err
	}

	return formatterHexadecimalMechanics{}.ptr().
		setFmtHexadecimalWithComponents(
			formatterHex,
			false, // useUpperCaseLetters
			"0x",  // leftPrefix
			false, // turnOnIntegerDigitsSeparation
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterHex"))
}
