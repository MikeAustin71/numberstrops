package numberstr

import (
	"sync"
)

type numStrBasicNanobot struct {
	lock *sync.Mutex
}
