package numberstr

import "sync"

type numStrFmtSpecSignedNumValQuark struct {
	lock *sync.Mutex
}

// getValidSignedNumNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for signed number values.
//
func (nStrFmtSpecSignedNumValQuark *numStrFmtSpecSignedNumValQuark) getValidSignedNumPositiveValFmtChars() []rune {

	if nStrFmtSpecSignedNumValQuark.lock == nil {
		nStrFmtSpecSignedNumValQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValQuark.lock.Lock()

	defer nStrFmtSpecSignedNumValQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")
}

// getValidSignedNumNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for signed number values.
//
func (nStrFmtSpecSignedNumValQuark *numStrFmtSpecSignedNumValQuark) getValidSignedNumNegativeValFmtChars() []rune {

	if nStrFmtSpecSignedNumValQuark.lock == nil {
		nStrFmtSpecSignedNumValQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValQuark.lock.Lock()

	defer nStrFmtSpecSignedNumValQuark.lock.Unlock()

	return []rune("(-) $127.54NUMFIELD")
}
