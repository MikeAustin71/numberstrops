package numberstr

import (
	"sync"
)

type formatterCollectionElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionElectron.
//
func (fmtCollectionElectron formatterCollectionElectron) ptr() *formatterCollectionElectron {

	if fmtCollectionElectron.lock == nil {
		fmtCollectionElectron.lock = new(sync.Mutex)
	}

	fmtCollectionElectron.lock.Lock()

	defer fmtCollectionElectron.lock.Unlock()

	newFmtCollectionElectron := new(formatterCollectionElectron)

	newFmtCollectionElectron.lock = new(sync.Mutex)

	return newFmtCollectionElectron
}
