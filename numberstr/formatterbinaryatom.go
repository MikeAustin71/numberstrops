package numberstr

import "sync"

type formatterBinaryAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryAtom.
//
func (fmtBinaryAtom formatterBinaryAtom) ptr() *formatterBinaryAtom {

	if fmtBinaryAtom.lock == nil {
		fmtBinaryAtom.lock = new(sync.Mutex)
	}

	fmtBinaryAtom.lock.Lock()

	defer fmtBinaryAtom.lock.Unlock()

	newBinaryAtom :=
		new(formatterBinaryAtom)

	newBinaryAtom.lock = new(sync.Mutex)

	return newBinaryAtom
}
