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
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrCurrencyAtom.lock == nil {
		nStrCurrencyAtom.lock = new(sync.Mutex)
	}

	nStrCurrencyAtom.lock.Lock()

	defer nStrCurrencyAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoElectron.testCurrencyPositiveValueFormat()")

	isValid = false

	if nStrFmtSpecCurrencyValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValueDto' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		nStrFmtSpecCurrencyValueDto.positiveValueFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecCurrencyValueDto.positiveValueFmt'"))

	isValid = true
	return isValid, err
}

// testCurrencyNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrCurrencyAtom *numStrFmtSpecCurrencyValueDtoAtom) testCurrencyNegativeValueFormat(
	nStrFmtSpecCurrencyValueDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrCurrencyAtom.lock == nil {
		nStrCurrencyAtom.lock = new(sync.Mutex)
	}

	nStrCurrencyAtom.lock.Lock()

	defer nStrCurrencyAtom.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoAtom.testCurrencyNegativeValueFormat()")

	isValid = false

	if nStrFmtSpecCurrencyValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValueDto' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		nStrFmtSpecCurrencyValueDto.negativeValueFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecCurrencyValueDto.negativeValueFmt'"))

	isValid = true
	return isValid, err
}
