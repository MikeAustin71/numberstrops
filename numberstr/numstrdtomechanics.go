package numberstr

import (
	"sync"
)

type numStrDtoMechanics struct {
	lock *sync.Mutex
}
