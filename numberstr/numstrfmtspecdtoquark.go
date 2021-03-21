package numberstr

import (
	"sync"
)

type numStrFmtSpecDtoQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoQuark.
//
func (nStrFmtSpecDtoQuark numStrFmtSpecDtoQuark) ptr() *numStrFmtSpecDtoQuark {

	if nStrFmtSpecDtoQuark.lock == nil {
		nStrFmtSpecDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoQuark.lock.Lock()

	defer nStrFmtSpecDtoQuark.lock.Unlock()

	newNumStrFmtSpecDtoQuark := new(numStrFmtSpecDtoQuark)

	newNumStrFmtSpecDtoQuark.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoQuark
}
