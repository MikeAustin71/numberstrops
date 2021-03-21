package numberstr

import (
	"sync"
)

type numStrBasicMechanics struct {
	lock *sync.Mutex
}
