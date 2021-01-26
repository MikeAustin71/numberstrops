package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFmtSpecDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecDto' to input parameter
// 'targetNStrFmtSpecDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecDto'
// will be overwritten.
//
// If input parameter 'inComingNStrFmtSpecDto' is judged
// to be invalid, this method will return an error.
//
func (nStrFmtSpecDtoElectron *numStrFmtSpecDtoElectron) copyIn(
	targetNStrFmtSpecDto *NumStrFmtSpecDto,
	inComingNStrFmtSpecDto *NumStrFmtSpecDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecDtoElectron.lock == nil {
		nStrFmtSpecDtoElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoElectron.lock.Lock()

	defer nStrFmtSpecDtoElectron.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecDtoElectron.copyIn()\n "

	if targetNStrFmtSpecDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if targetNStrFmtSpecDto.lock == nil {
		targetNStrFmtSpecDto.lock = new(sync.Mutex)
	}

	if inComingNStrFmtSpecDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	_,
		err =
		nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
			inComingNStrFmtSpecDto,
			ePrefix+
				"Testing validity of 'inComingNStrFmtSpecDto'\n ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecDto.idNo =
		inComingNStrFmtSpecDto.idNo

	targetNStrFmtSpecDto.idString =
		inComingNStrFmtSpecDto.idString

	targetNStrFmtSpecDto.description =
		inComingNStrFmtSpecDto.description

	targetNStrFmtSpecDto.tag =
		inComingNStrFmtSpecDto.tag

	err = targetNStrFmtSpecDto.countryCulture.CopyIn(
		&inComingNStrFmtSpecDto.countryCulture,
		ePrefix+
			"inComingNStrFmtSpecDto->targetNStrFmtSpecDto\n")

	if err != nil {
		return err
	}

	err = targetNStrFmtSpecDto.absoluteValue.CopyIn(
		&inComingNStrFmtSpecDto.absoluteValue,
		ePrefix+
			"inComingNStrFmtSpecDto->targetNStrFmtSpecDto\n")

	if err != nil {
		return err
	}

	err = targetNStrFmtSpecDto.currencyValue.CopyIn(
		&inComingNStrFmtSpecDto.currencyValue,
		ePrefix+
			"inComingNStrFmtSpecDto->targetNStrFmtSpecDto\n")

	if err != nil {
		return err
	}

	err = targetNStrFmtSpecDto.signedNumValue.CopyIn(
		&inComingNStrFmtSpecDto.signedNumValue,
		ePrefix+
			"inComingNStrFmtSpecDto->targetNStrFmtSpecDto\n")

	if err != nil {
		return err
	}

	err = targetNStrFmtSpecDto.sciNotation.CopyIn(
		&inComingNStrFmtSpecDto.sciNotation,
		ePrefix+
			"inComingNStrFmtSpecDto->targetNStrFmtSpecDto\n")

	return err
}

// copyOut - Returns a deep copy of input parameter 'nStrFmtSpecDto'
// styled as a new instance of NumStrFmtSpecDto.
//
// If input parameter 'nStrFmtSpecDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecDtoElectron *numStrFmtSpecDtoElectron) copyOut(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	ePrefix string) (
	newNStrFmtSpecDto NumStrFmtSpecDto,
	err error) {

	if nStrFmtSpecDtoElectron.lock == nil {
		nStrFmtSpecDtoElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoElectron.lock.Lock()

	defer nStrFmtSpecDtoElectron.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecDtoElectron.copyOut()\n "

	if nStrFmtSpecDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newNStrFmtSpecDto, err
	}

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	_,
		err =
		nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
			nStrFmtSpecDto,
			ePrefix+
				"Testing validity of 'nStrFmtSpecDto'\n ")

	if err != nil {
		return newNStrFmtSpecDto, err
	}

	newNStrFmtSpecDto.idNo =
		nStrFmtSpecDto.idNo

	newNStrFmtSpecDto.idString =
		nStrFmtSpecDto.idString

	newNStrFmtSpecDto.description =
		nStrFmtSpecDto.description

	newNStrFmtSpecDto.tag =
		nStrFmtSpecDto.tag

	err = newNStrFmtSpecDto.countryCulture.CopyIn(
		&nStrFmtSpecDto.countryCulture,
		ePrefix+
			"nStrFmtSpecDto->newNStrFmtSpecDto\n")

	if err != nil {
		return newNStrFmtSpecDto, err
	}

	err = newNStrFmtSpecDto.absoluteValue.CopyIn(
		&nStrFmtSpecDto.absoluteValue,
		ePrefix+
			"nStrFmtSpecDto->newNStrFmtSpecDto\n")

	if err != nil {
		return newNStrFmtSpecDto, err
	}

	err = newNStrFmtSpecDto.currencyValue.CopyIn(
		&nStrFmtSpecDto.currencyValue,
		ePrefix+
			"nStrFmtSpecDto->newNStrFmtSpecDto\n")

	if err != nil {
		return newNStrFmtSpecDto, err
	}

	err = newNStrFmtSpecDto.signedNumValue.CopyIn(
		&nStrFmtSpecDto.signedNumValue,
		ePrefix+
			"nStrFmtSpecDto->newNStrFmtSpecDto\n")

	if err != nil {
		return newNStrFmtSpecDto, err
	}

	err = newNStrFmtSpecDto.sciNotation.CopyIn(
		&nStrFmtSpecDto.sciNotation,
		ePrefix+
			"nStrFmtSpecDto->newNStrFmtSpecDto\n")

	return newNStrFmtSpecDto, err
}
