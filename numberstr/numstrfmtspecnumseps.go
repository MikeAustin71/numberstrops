package numberstr

import (
	"sync"
)

type NumStrFmtSpecDigitsSeparatorsDto struct {
	decimalSeparator              rune
	integerDigitsSeparator        rune
	integerDigitsGroupingSequence []uint
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming NumberFieldDto
// to the data fields of the current instance of NumberFieldDto.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) CopyIn(
	incomingSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	nStrFmtSpecDigitsSepDto.decimalSeparator =
		incomingSpecDigitsSepDto.decimalSeparator

	nStrFmtSpecDigitsSepDto.integerDigitsSeparator =
		incomingSpecDigitsSepDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(incomingSpecDigitsSepDto.integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
		return
	}

	nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence[i] =
			incomingSpecDigitsSepDto.integerDigitsGroupingSequence[i]
	}

	return
}

// CopyOut - Returns a deep copy of the current NumberFieldDto instance.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) CopyOut() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	newDigitsSepDto := NumStrFmtSpecDigitsSeparatorsDto{}

	newDigitsSepDto.decimalSeparator =
		nStrFmtSpecDigitsSepDto.decimalSeparator

	newDigitsSepDto.integerDigitsSeparator =
		nStrFmtSpecDigitsSepDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		newDigitsSepDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
		return newDigitsSepDto
	}

	newDigitsSepDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		newDigitsSepDto.integerDigitsGroupingSequence[i] =
			nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence[i]
	}

	return newDigitsSepDto
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFormatDto instance to determine whether the current instance
// is a valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumStrFormatDto is
//       valid, or not. If the current NumStrFormatDto contains
//       valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
//

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecDigitsSeparatorsDto instance to determine whether the current instance
// is a valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
//       If this instance of NumStrFmtSpecDigitsSeparatorsDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecDigitsSeparatorsDto.IsValidInstanceError() \n" +
		"Testing Validity of 'nStrFmtSpecDigitsSepDto' "

	nStrFmtSpecLibQuark := numStrFmtSpecLibQuark{}

	_,
		err := nStrFmtSpecLibQuark.testValidityOfNumSepsDto(
		nStrFmtSpecDigitsSepDto,
		ePrefix)

	return err
}
