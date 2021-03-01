package numberstr

import (
	"sync"
)

// NumStrIntSeparator is designed to be used as an array
// element in an array which will be used to insert integer
// separators, primarily thousands separators, into number
// strings. Some countries/cultures do not use thousands
// separation and instead rely on multiple integer separation
// characters and grouping sequences for a single number
// string. One notable example of this found in the 'Indian
// Number System'.
//
// An array of NumStrIntSeparator elements provides the flexibility
// necessary to process these complex number formats.
//
type NumStrIntSeparator struct {
	intSeparatorChar     rune
	intSeparatorGrouping uint
	lock                 *sync.Mutex
}

// IsValidInstanceError
func (nStrIntSep *NumStrIntSeparator) IsValidInstanceError(
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	ePrefix.SetEPrefCtx(
		"NumStrIntSeparator.IsValidInstanceError()",
		"Testing Validity of 'nStrIntSep'")

	_,
		err =
		numStrIntSeparatorQuark{}.Ptr().
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				ePrefix)

	return err
}
