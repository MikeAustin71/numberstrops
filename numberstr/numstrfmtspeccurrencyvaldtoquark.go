package numberstr

import "sync"

type numStrFmtSpecCurrencyValueDtoQuark struct {
	lock *sync.Mutex
}

// getValidCurrencyPositiveValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Positive Value format
// strings for currency displays.
//
func (nStrFmtSpecCurrValDtoQuark *numStrFmtSpecCurrencyValueDtoQuark) getValidCurrencyPositiveValFmtChars() []rune {

	if nStrFmtSpecCurrValDtoQuark.lock == nil {
		nStrFmtSpecCurrValDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDtoQuark.lock.Lock()

	defer nStrFmtSpecCurrValDtoQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")
}

// getValidCurrencyNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for currency displays.
//
func (nStrFmtSpecCurrValDtoQuark *numStrFmtSpecCurrencyValueDtoQuark) getValidCurrencyNegativeValFmtChars() []rune {

	if nStrFmtSpecCurrValDtoQuark.lock == nil {
		nStrFmtSpecCurrValDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDtoQuark.lock.Lock()

	defer nStrFmtSpecCurrValDtoQuark.lock.Unlock()

	return []rune("(-) $127.54NUMFIELD")
}
