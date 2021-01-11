package numberstr

import (
	"sync"
)

type numStrFormatterElectron struct {
	lock *sync.Mutex
}
