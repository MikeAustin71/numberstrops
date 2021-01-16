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
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumAtom.lock == nil {
		nStrSignedNumAtom.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom.lock.Lock()

	defer nStrSignedNumAtom.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecSignedNumValAtom.testSignedNumValPositiveValueFormat() "

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValPositiveValueFormatStr(
		nStrFmtSpecSignedNumberValueDto.positiveValueFmt,
		ePrefix+
			"\nTesting validity of 'nStrFmtSpecSignedNumberValueDto.positiveValueFmt'\n ")

	return isValid, err
}

// testSignedNumValNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrSignedNumAtom *numStrFmtSpecSignedNumValAtom) testSignedNumValNegativeValueFormat(
	nStrFmtSpecSignedNumberValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumAtom.lock == nil {
		nStrSignedNumAtom.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom.lock.Lock()

	defer nStrSignedNumAtom.lock.Unlock()

	ePrefix += "numStrFmtSpecSignedNumValAtom.testSignedNumValNegativeValueFormat()\n "

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormatStr(
		nStrFmtSpecSignedNumberValueDto.negativeValueFmt,
		ePrefix+
			"Testing validity of 'nStrFmtSpecSignedNumberValueDto.negativeValueFmt'\n ")

	return isValid, err
}
