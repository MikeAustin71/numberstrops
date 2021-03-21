package numberstr

import (
	"sync"
)

type numStrDtoUtility struct {
	lock *sync.Mutex
}
