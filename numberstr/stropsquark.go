package numberstr

import "sync"

type strOpsQuark struct {
	lock *sync.Mutex
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (sOpsQuark *strOpsQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}
