package numberstr

import (
	"sync"
)

type numStrDtoMolecule struct {
	lock *sync.Mutex
}
