package numberstr

import (
	"sync"
)

type numStrDtoAtom struct {
	lock *sync.Mutex
}
