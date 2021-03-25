package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterOctalUtility.
//
func (fmtOctalUtil formatterOctalUtility) ptr() *formatterOctalUtility {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	newFmtOctalUtil :=
		new(formatterOctalUtility)

	newFmtOctalUtil.lock = new(sync.Mutex)

	return newFmtOctalUtil
}

// setFmtOctalWithDefaults - Overwrites and resets the data fields
// within input parameter 'formatterOctal'. The new data values are
// created from default parameters specified below.
//
// Member variable data fields in the input parameter
// FormatterOctal instance 'formatterOctal' are configured as
// follows:
//
// Number String Format Type (numStrFmtType) =
//     NumStrFormatTypeCode(0).Octal()
//
// Octal Number String Left Prefix String (leftPrefix) = ""
//
// Octal Number String Right Suffix String (rightSuffix) = ""
//
// Turn On Integer Digits Separation
//  (turnOnIntegerDigitsSeparation) = false
//
// Integer Separators (integerSeparators) =
//   Integer Digit Separator  = ","
//   Integer Digit Grouping Sequence = 3
//   Integer Group Repetitions = 0
//
// Number Field Dto = Off
//   Requested Number Field Length = -1
//   Text Justify Format (textJustifyFormat) =
//     TextJustify(0).Right
//
func (fmtOctalUtil *formatterOctalUtility) setFmtOctalWithDefaults(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterOctalUtility." +
			"setFmtOctalWithDefaults()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

	formatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewBasic(
		",",
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

	return formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			formatterOctal,
			"",    // leftPrefix
			"",    // rightSuffix
			false, //turnOnIntegerDigitsSeparation
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterOctal"))
}
