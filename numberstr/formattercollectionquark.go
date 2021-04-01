package numberstr

import (
	"sync"
)

type formatterCollectionQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionQuark.
//
func (fmtCollectionQuark formatterCollectionQuark) ptr() *formatterCollectionQuark {

	if fmtCollectionQuark.lock == nil {
		fmtCollectionQuark.lock = new(sync.Mutex)
	}

	fmtCollectionQuark.lock.Lock()

	defer fmtCollectionQuark.lock.Unlock()

	newFmtCollectionQuark := new(formatterCollectionQuark)

	newFmtCollectionQuark.lock = new(sync.Mutex)

	return newFmtCollectionQuark
}
