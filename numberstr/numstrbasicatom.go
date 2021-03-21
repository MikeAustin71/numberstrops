package numberstr

import (
	"sync"
)

type numStrBasicAtom struct {
	lock *sync.Mutex
}
