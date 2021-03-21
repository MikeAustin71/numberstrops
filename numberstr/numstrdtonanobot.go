package numberstr

import (
	"sync"
)

type numStrDtoNanobot struct {
	lock *sync.Mutex
}
