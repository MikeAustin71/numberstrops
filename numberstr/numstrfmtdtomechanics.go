package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatDtoMechanics struct {
	lock *sync.Mutex
}

// setToDefaults - Receives a pointer to an instance of
// NumStrFormatDto and proceeds to set all of the internal
// member variables to their default values.
//
// The Number String Format default values represent formatting
// parameters used in the United States.
//
func (nStrFmtDtoMech *numStrFormatDtoMechanics) setToDefaults(
	nStrFmtDto *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDtoMech.lock == nil {
		nStrFmtDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtDtoMech.lock.Lock()

	defer nStrFmtDtoMech.lock.Unlock()

	ePrefix += "numStrFormatDtoMechanics.setToDefaults() "

	if nStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter nStrFmtDto is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	numStrFormatterCollection :=
		NumStrFormatCountry{}.NewPtr().UnitedStates()

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	err =
		nStrFmtDtoNanobot.setNumberStringFormatters(
			nStrFmtDto,
			numStrFormatterCollection,
			ePrefix+
				"\nSetting 'nStrFmtDto' FormatSetters to US Defaults ")

	if err != nil {
		return err
	}

	err = nStrFmtDtoNanobot.setNumberFieldLength(
		nStrFmtDto,
		-1,
		ePrefix+"\nSetting Default 'nStrFmtDto' number string length to -1 ")

	return err
}
