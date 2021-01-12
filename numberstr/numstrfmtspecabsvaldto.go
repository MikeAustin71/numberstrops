package numberstr

import "sync"

type NumStrFmtSpecAbsoluteValueDto struct {
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numberSeparators              NumStrFmtSpecDigitsSeparatorsDto
	numFieldLen                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming NumStrFmtSpecAbsoluteValueDto
// to the data fields of the current instance of NumStrFmtSpecAbsoluteValueDto.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyIn(
	inComingNStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return

}
