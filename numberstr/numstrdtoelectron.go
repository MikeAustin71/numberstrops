package numberstr

import (
	"sync"
)

type numStrDtoElectron struct {
	lock *sync.Mutex
}
