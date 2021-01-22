package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDtoMechanics struct {
	lock *sync.Mutex
}

func (nStrFmtSpecDtoMech *numStrFmtSpecDtoMechanics) setNumStrFmtSpecDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture NumStrFmtSpecCountryDto,
	absoluteValue NumStrFmtSpecAbsoluteValueDto,
	currencyValue NumStrFmtSpecCurrencyValueDto,
	signedNumValue NumStrFmtSpecSignedNumValueDto,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix string) (err error) {

	if nStrFmtSpecDtoMech.lock == nil {
		nStrFmtSpecDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech.lock.Lock()

	defer nStrFmtSpecDtoMech.lock.Unlock()

	ePrefix += "numStrFmtSpecDtoMechanics.setNumStrFmtSpecDto() "

	if nStrFmtSpecDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDto' is nil pointer!\n",
			ePrefix)
		return err
	}

	if nStrFmtSpecDto.lock == nil {
		nStrFmtSpecDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDto.idNo = idNo
	nStrFmtSpecDto.idString = idString
	nStrFmtSpecDto.description = description
	nStrFmtSpecDto.tag = tag

	err = nStrFmtSpecDto.countryCulture.CopyIn(
		&countryCulture,
		ePrefix)

	if err != nil {
		return err
	}

	err = absoluteValue.IsValidInstanceError(
		ePrefix + "\nTesting 'absoluteValue' validity\n")

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.absoluteValue.CopyIn(
		&absoluteValue,
		ePrefix+
			"\nabsoluteValue->nStrFmtSpecDto.absoluteValue\n")

	if err != nil {
		return err
	}

	err = currencyValue.IsValidInstanceError(
		ePrefix + "\nTesting 'currencyValue' validity\n")

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.currencyValue.CopyIn(
		&currencyValue,
		ePrefix+
			"\ncurrencyValue->nStrFmtSpecDto.currencyValue\n")

	if err != nil {
		return err
	}

	err = signedNumValue.IsValidInstanceError(
		ePrefix + "\nTesting 'signedNumValue' validity\n")

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.signedNumValue.CopyIn(
		&signedNumValue,
		ePrefix+
			"\nsignedNumValue->nStrFmtSpecDto.signedNumValue\n")

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.sciNotation.CopyIn(
		&sciNotation,
		ePrefix+
			"\nsciNotation->nStrFmtSpecDto.sciNotation\n")

	return err
}
