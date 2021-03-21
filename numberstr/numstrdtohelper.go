package numberstr

import (
	"sync"
)

type numStrDtoHelper struct {
	lock *sync.Mutex
}
