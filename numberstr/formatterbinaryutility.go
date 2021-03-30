package numberstr

import (
	"fmt"
	"sync"
)

type formatterBinaryUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryUtility.
//
func (fmtBinaryUtil formatterBinaryUtility) ptr() *formatterBinaryUtility {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	newBinaryUtility :=
		new(formatterBinaryUtility)

	newBinaryUtility.lock = new(sync.Mutex)

	return newBinaryUtility
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterBinary instance to United
// States Default values.
//
// In the United States, Binary Number default formatting
// parameters are defined as follows:
//
//    Turn On Integer Digits Separation: false
//         Decimal Separator Character: '.'
//                 Number Field Length: -1
//          Number Field Justification: Right-Justified
//
// Note: With 'Turn On Integer Digits Separation' set to false,
// integer digit separation is not applied to binary digits
// displayed in number strings.
//
func (fmtBinaryUtil *formatterBinaryUtility) setToUnitedStatesDefaults(
	formatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"setToUnitedStatesDefaults()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewBasic(
		-1,
		TextJustify(0).Right(),
		ePrefix)

	if err != nil {
		return err
	}
	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.NewIntSepsDetailRunes(
		[]rune{'.'}, // Decimal Separators
		[]rune{','}, // Integer Digits Separators
		3,           // Integer Digit Grouping
		0,           // Integer Separator Repetitions
		ePrefix.XCtx(
			"decimalSeparatorChars='.'\n"+
				"integerDigitsSeparators=','\n"))

	if err != nil {
		return err
	}

	err = formatterBinaryMechanics{}.ptr().setWithComponents(
		formatterBinary,
		false, // turnOnIntegerDigitsSeparation
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"formatterBinary"))

	return err
}
