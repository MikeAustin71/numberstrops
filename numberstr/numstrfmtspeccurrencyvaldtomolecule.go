package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoMolecule struct {
	lock *sync.Mutex
}

// testValidityOfCurrencyValDto - Tests the validity of
// NumStrFmtSpecCurrencyValueDto objects.
//
func (nStrFmtSpecCurrDtoMolecule *numStrFmtSpecCurrencyValueDtoMolecule) testValidityOfCurrencyValDto(
	nStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecCurrDtoMolecule.lock == nil {
		nStrFmtSpecCurrDtoMolecule.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrDtoMolecule.lock.Lock()

	defer nStrFmtSpecCurrDtoMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoMolecule.testValidityOfCurrencyValDto()")

	isValid = false

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecCurrencyValDto.currencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Symbol is missing!\n"+
			"The currency symbol is set to zero and is therefore invalid.\n"+
			"NumStrFmtSpecCurrencyValueDto.currencySymbols == 0\n",
			ePrefix.String())

		return isValid, err
	}

	err =
		nStrFmtSpecCurrencyValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'nStrFmtSpecCurrencyValDto' Number Separators"))

	if err != nil {
		return isValid, err
	}

	nStrCurrencyAtom :=
		numStrFmtSpecCurrencyValueDtoAtom{}

	_,
		err = nStrCurrencyAtom.testCurrencyPositiveValueFormat(
		nStrFmtSpecCurrencyValDto,
		ePrefix.XCtx("Testing nStrFmtSpecCurrencyValDto.positiveValueFmt"))

	if err != nil {
		return isValid, err
	}

	_,
		err = nStrCurrencyAtom.testCurrencyNegativeValueFormat(
		nStrFmtSpecCurrencyValDto,
		ePrefix.XCtx("\nTesting nStrFmtSpecCurrencyValDto.positiveValueFmt"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
