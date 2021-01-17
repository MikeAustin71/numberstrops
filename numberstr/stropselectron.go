package numberstr

import (
	"sync"
)

type strOpsElectron struct {
	lock *sync.Mutex
}
