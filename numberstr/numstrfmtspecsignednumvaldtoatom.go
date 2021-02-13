package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSignedNumValAtom struct {
	lock *sync.Mutex
}

// testSignedNumValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
func (nStrSignedNumAtom *numStrFmtSpecSignedNumValAtom) testSignedNumValPositiveValueFormat(
	nStrFmtSpecSignedNumberValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrSignedNumAtom.lock == nil {
		nStrSignedNumAtom.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom.lock.Lock()

	defer nStrSignedNumAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecSignedNumValAtom.testSignedNumValPositiveValueFormat()")

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

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
func (nStrSignedNumAtom *numStrFmtSpecSignedNumValAtom) testSignedNumValNegativeValueFormat(
	nStrFmtSpecSignedNumberValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrSignedNumAtom.lock == nil {
		nStrSignedNumAtom.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom.lock.Lock()

	defer nStrSignedNumAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecSignedNumValAtom.testSignedNumValNegativeValueFormat()")

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormatStr(
		nStrFmtSpecSignedNumberValueDto.negativeValueFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecSignedNumberValueDto.negativeValueFmt'"))

	return isValid, err
}
