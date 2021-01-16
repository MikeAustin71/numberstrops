package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoMolecule struct {
	lock *sync.Mutex
}

func (nStrFmtSpecCurrDtoMolecule *numStrFmtSpecCurrencyValueDtoMolecule) testValidityOfCurrencyValDto(
	nStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecCurrDtoMolecule.lock == nil {
		nStrFmtSpecCurrDtoMolecule.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrDtoMolecule.lock.Lock()

	defer nStrFmtSpecCurrDtoMolecule.lock.Unlock()

	ePrefix += "numStrFmtSpecCurrencyValueDtoMolecule.testValidityOfCurrencyValDto() "

	isValid = false

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	if nStrFmtSpecCurrencyValDto.currencySymbol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Symbol is missing!\n"+
			"The currency symbol is set to zero and is therefore invalid.\n"+
			"NumStrFmtSpecCurrencyValueDto.currencySymbol == 0\n",
			ePrefix)
		return isValid, err
	}

	err =
		nStrFmtSpecCurrencyValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix +
				"\nValidating 'nStrFmtSpecCurrencyValDto' Number Separators\n ")

	if err != nil {
		return isValid, err
	}

	nStrCurrencyAtom :=
		numStrFmtSpecCurrencyValueDtoAtom{}

	_,
		err = nStrCurrencyAtom.testCurrencyPositiveValueFormat(
		nStrFmtSpecCurrencyValDto,
		ePrefix+
			"\nTesting nStrFmtSpecCurrencyValDto.positiveValueFmt\n")

	if err != nil {
		return isValid, err
	}

	_,
		err = nStrCurrencyAtom.testCurrencyNegativeValueFormat(
		nStrFmtSpecCurrencyValDto,
		ePrefix+
			"\nTesting nStrFmtSpecCurrencyValDto.positiveValueFmt\n")

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
