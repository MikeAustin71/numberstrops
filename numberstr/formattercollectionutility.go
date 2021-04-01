package numberstr

import (
	"sync"
)

type formatterCollectionUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionUtility.
//
func (fmtCollectionUtil formatterCollectionUtility) ptr() *formatterCollectionUtility {

	if fmtCollectionUtil.lock == nil {
		fmtCollectionUtil.lock = new(sync.Mutex)
	}

	fmtCollectionUtil.lock.Lock()

	defer fmtCollectionUtil.lock.Unlock()

	newFmtCollectionUtility := new(formatterCollectionUtility)

	newFmtCollectionUtility.lock = new(sync.Mutex)

	return newFmtCollectionUtility
}
