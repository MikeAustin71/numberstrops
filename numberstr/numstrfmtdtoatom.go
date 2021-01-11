package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtDtoAtom struct {
	lock *sync.Mutex
}

// copyIn - Receives pointers to two instances of NumStrFormatDto
// labeled as input parameters, 'numStrFmtDto1' and 'numStrFmtDto2'.
// This method then proceeds to copy all of the data from
// 'numStrFmtDto2' into 'numStrFmtDto1'.  When the copy process is
// completed, the data values in 'numStrFmtDto1' and 'numStrFmtDto2'
// will be identical.
//
// IMPORTANT
//
// Be advised, this method will overwrite the data values contained
// in input parameter, 'numStrFmtDto1'.
//
func (nStrFmtDtoAtom *numStrFmtDtoAtom) copyIn(
	numStrFmtDto1 *NumStrFormatDto,
	numStrFmtDto2 *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDtoAtom.lock == nil {
		nStrFmtDtoAtom.lock = new(sync.Mutex)
	}

	nStrFmtDtoAtom.lock.Lock()

	defer nStrFmtDtoAtom.lock.Unlock()

	ePrefix += "numStrFormatDtoMechanics.copyIn() "

	if numStrFmtDto1 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto1' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if numStrFmtDto2 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto2' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	numStrFmtDto1.numStrFmtConfigs = make(map[NumStrValSpec]NumStrFormatter,
		len(numStrFmtDto2.numStrFmtConfigs))

	for idx, nStrFormtrs := range numStrFmtDto2.numStrFmtConfigs {
		numStrFmtDto1.numStrFmtConfigs[idx] = nStrFormtrs.CopyOut()
	}

	//TODO Fix copyIn()
	numStrFmtDto1.numFieldDto =
		numStrFmtDto2.numFieldDto.CopyOut()

	return err
}

// copyOut - Receives a pointer to an instance of NumStrFormatDto
// and returns a deep copy of that NumStrFormatDto instance.
//
func (nStrFmtDtoAtom *numStrFmtDtoAtom) copyOut(
	numStrFmtDto *NumStrFormatDto,
	ePrefix string) (
	newNumStrFormatDto NumStrFormatDto,
	err error) {

	if nStrFmtDtoAtom.lock == nil {
		nStrFmtDtoAtom.lock = new(sync.Mutex)
	}

	nStrFmtDtoAtom.lock.Lock()

	defer nStrFmtDtoAtom.lock.Unlock()

	ePrefix += "numStrFormatDtoMechanics.copyOut() "
	newNumStrFormatDto = NumStrFormatDto{}

	if numStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)
		return newNumStrFormatDto, err
	}

	newNumStrFormatDto.numStrFmtConfigs =
		make(map[NumStrValSpec]NumStrFormatter, 0)

	newNumStrFormatDto.numStrFmtConfigs =
		make(map[NumStrValSpec]NumStrFormatter,
			len(numStrFmtDto.numStrFmtConfigs))

	for idx, nStrFormtrs := range numStrFmtDto.numStrFmtConfigs {
		newNumStrFormatDto.numStrFmtConfigs[idx] = nStrFormtrs.CopyOut()
	}

	newNumStrFormatDto.numFieldDto =
		numStrFmtDto.numFieldDto.CopyOut()

	newNumStrFormatDto.lock = new(sync.Mutex)

	return newNumStrFormatDto, err
}
