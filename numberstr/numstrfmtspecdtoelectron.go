package numberstr

import (
	"sync"
)

type numStrFmtSpecDtoElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numStrFmtSpecDtoElectron.
//
func (nStrFmtSpecElectron numStrFmtSpecDtoElectron) ptr() *numStrFmtSpecDtoElectron {

	if nStrFmtSpecElectron.lock == nil {
		nStrFmtSpecElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecElectron.lock.Lock()

	defer nStrFmtSpecElectron.lock.Unlock()

	newNumStrFmtSpecElectron := new(numStrFmtSpecDtoElectron)

	newNumStrFmtSpecElectron.lock = new(sync.Mutex)

	return newNumStrFmtSpecElectron
}
