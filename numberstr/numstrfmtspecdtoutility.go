package numberstr

import (
	"sync"
)

type numStrFmtSpecDtoUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoUtility.
//
func (nStrFmtSpecDtoUtil numStrFmtSpecDtoUtility) ptr() *numStrFmtSpecDtoUtility {

	if nStrFmtSpecDtoUtil.lock == nil {
		nStrFmtSpecDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoUtil.lock.Lock()

	defer nStrFmtSpecDtoUtil.lock.Unlock()

	newNumStrFmtSpecDtoUtility := new(numStrFmtSpecDtoUtility)

	newNumStrFmtSpecDtoUtility.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoUtility
}
