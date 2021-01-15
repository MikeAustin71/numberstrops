package numberstr

import "sync"

type numStrFmtSpecAbsoluteValueDtoQuark struct {
	lock *sync.Mutex
}

// getValidAbsoluteValFmtChars- Returns an array of runes which
// constitute the only valid characters allowed in Absolute Value
// format strings.
//
func (nStrFmtSpecAbsValueDtoQuark *numStrFmtSpecAbsoluteValueDtoQuark) getValidAbsoluteValFmtChars() []rune {

	if nStrFmtSpecAbsValueDtoQuark.lock == nil {
		nStrFmtSpecAbsValueDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValueDtoQuark.lock.Lock()

	defer nStrFmtSpecAbsValueDtoQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")

}
