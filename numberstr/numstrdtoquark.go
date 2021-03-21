package numberstr

import (
	"sync"
)

type numStrDtoQuark struct {
	lock *sync.Mutex
}
