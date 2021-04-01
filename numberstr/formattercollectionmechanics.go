package numberstr

import (
	"sync"
)

type formatterCollectionMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionMechanics.
//
func (fmtCollectionMech formatterCollectionMechanics) ptr() *formatterCollectionMechanics {

	if fmtCollectionMech.lock == nil {
		fmtCollectionMech.lock = new(sync.Mutex)
	}

	fmtCollectionMech.lock.Lock()

	defer fmtCollectionMech.lock.Unlock()

	newFmtCollectionMechanics :=
		new(formatterCollectionMechanics)

	newFmtCollectionMechanics.lock = new(sync.Mutex)

	return newFmtCollectionMechanics
}
