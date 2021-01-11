package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatDtoUtility struct {
	lock *sync.Mutex
}

func (nStrFmtDtoUtil *numStrFormatDtoUtility) setToDefaultsIfEmpty(
	nStrFmtDto *NumStrFormatDto,
	ePrefix string) (
	err error) {

	if nStrFmtDtoUtil.lock == nil {
		nStrFmtDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtDtoUtil.lock.Lock()

	defer nStrFmtDtoUtil.lock.Unlock()

	ePrefix += "numStrFormatDtoUtility.setToDefaultsIfEmpty() "

	if nStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	_,
		err = nStrFmtDtoNanobot.testNumStrFormatDtoValidity(
		nStrFmtDto,
		ePrefix+"Testing validity of 'nStrFmtDto' ")

	if err == nil {
		return err
	}

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	err = nStrFmtDtoMech.setToDefaults(
		nStrFmtDto,
		ePrefix+"Setting 'nStrFmtDto' to defaults ")

	if err != nil {
		return err
	}

	return err
}
