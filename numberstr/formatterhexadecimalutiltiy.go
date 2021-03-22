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

// setFmtHexadecimalWithDefaults - Receives a pointer to an
// instance of FormatterHexadecimal and proceeds to set the data
// values to standard defaults.
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
//  integerSeparators = None
//
//  numFieldLenDto.requestedNumFieldLength = -1
//                   Number Field = Number String Length
//
//  numFieldLenDto.textJustifyFormat = TextJustify(0).Right()

func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexadecimalWithDefaults(
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
			"setFmtHexadecimalWithDefaults()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	formatterHex.useUpperCaseLetters = false

	formatterHex.leftPrefix = "0x"

	formatterHex.turnOnIntegerDigitsSeparation = false

	formatterHex.integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewBasic(
		" ",
		nil)

	if err != nil {
		return err
	}

	formatterHex.numFieldLenDto,
		err =
		NumberFieldDto{}.NewBasic(
			-1,
			TextJustify(0).Right(),
			nil)

	return err
}
