package numberstr

import "sync"

type formatterCurrencyQuark struct {
	lock *sync.Mutex
}

// getValidCurrencyPositiveValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Positive Value format
// strings for currency displays.
//
func (fmtCurrQuark *formatterCurrencyQuark) getValidCurrencyPositiveValFmtChars() []rune {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")
}

// getValidCurrencyNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for currency displays.
//
func (fmtCurrQuark *formatterCurrencyQuark) getValidCurrencyNegativeValFmtChars() []rune {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	return []rune("(-) $127.54NUMFIELD")
}
