package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoAtom struct {
	lock *sync.Mutex
}

// testCurrencyPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrCurrencyAtom *numStrFmtSpecCurrencyValueDtoAtom) testCurrencyPositiveValueFormat(
	nStrFmtSpecCurrencyValueDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrCurrencyAtom.lock == nil {
		nStrCurrencyAtom.lock = new(sync.Mutex)
	}

	nStrCurrencyAtom.lock.Lock()

	defer nStrCurrencyAtom.lock.Unlock()

	ePrefix += "numStrFmtSpecCurrencyValueDtoElectron.testCurrencyPositiveValueFormat() "

	isValid = false

	if nStrFmtSpecCurrencyValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		nStrFmtSpecCurrencyValueDto.positiveValueFmt,
		ePrefix+
			"Testing validity of 'nStrFmtSpecCurrencyValueDto.positiveValueFmt'\n ")

	isValid = true
	return isValid, err
}

// testCurrencyNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrCurrencyAtom *numStrFmtSpecCurrencyValueDtoAtom) testCurrencyNegativeValueFormat(
	nStrFmtSpecCurrencyValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrCurrencyAtom.lock == nil {
		nStrCurrencyAtom.lock = new(sync.Mutex)
	}

	nStrCurrencyAtom.lock.Lock()

	defer nStrCurrencyAtom.lock.Unlock()

	ePrefix += "numStrFmtSpecCurrencyValueDtoElectron.testCurrencyNegativeValueFormat() "

	isValid = false

	if nStrFmtSpecCurrencyValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		nStrFmtSpecCurrencyValueDto.negativeValueFmt,
		ePrefix+
			"Testing validity of 'nStrFmtSpecCurrencyValueDto.negativeValueFmt'\n ")

	isValid = true
	return isValid, err
}
