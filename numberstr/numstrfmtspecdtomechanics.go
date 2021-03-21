package numberstr

import (
	"sync"
)

type numStrFmtSpecDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoUtility.
//
func (nStrFmtSpecDtoMech numStrFmtSpecDtoMechanics) ptr() *numStrFmtSpecDtoMechanics {

	if nStrFmtSpecDtoMech.lock == nil {
		nStrFmtSpecDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech.lock.Lock()

	defer nStrFmtSpecDtoMech.lock.Unlock()

	newNumStrFmtSpecDtoMechanics := new(numStrFmtSpecDtoMechanics)

	newNumStrFmtSpecDtoMechanics.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoMechanics
}
