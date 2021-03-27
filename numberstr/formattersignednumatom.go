package numberstr

import (
	"fmt"
	"sync"
)

type formatterSignedNumberAtom struct {
	lock *sync.Mutex
}

// testSignedNumValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
func (fmtSignedNumAtom *formatterSignedNumberAtom) testSignedNumValPositiveValueFormat(
	nStrFmtSpecSignedNumberValueDto *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSignedNumAtom.lock == nil {
		fmtSignedNumAtom.lock = new(sync.Mutex)
	}

	fmtSignedNumAtom.lock.Lock()

	defer fmtSignedNumAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("formatterSignedNumberAtom.testSignedNumValPositiveValueFormat()")

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	nStrSignedNumElectron :=
		formatterSignedNumberElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValPositiveValueFormatStr(
		nStrFmtSpecSignedNumberValueDto.positiveValueFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecSignedNumberValueDto.positiveValueFmt'"))

	return isValid, err
}

// testSignedNumValNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (fmtSignedNumAtom *formatterSignedNumberAtom) testSignedNumValNegativeValueFormat(
	nStrFmtSpecSignedNumberValueDto *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSignedNumAtom.lock == nil {
		fmtSignedNumAtom.lock = new(sync.Mutex)
	}

	fmtSignedNumAtom.lock.Lock()

	defer fmtSignedNumAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("formatterSignedNumberAtom.testSignedNumValNegativeValueFormat()")

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrSignedNumElectron :=
		formatterSignedNumberElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormatStr(
		nStrFmtSpecSignedNumberValueDto.negativeValueFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecSignedNumberValueDto.negativeValueFmt'"))

	return isValid, err
}
