package numberstr

import "sync"

type NumStrFmtSpecAbsoluteValueDto struct {
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numberSeparatorsDto           NumStrFmtSpecDigitsSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrFmtSpecAbsoluteValueDto ('inComingNStrFmtAbsValDto') to
// the data fields of the current NumStrFmtSpecAbsoluteValueDto
// instance.
//
// If 'inComingNStrFmtAbsValDto' is judged to be invalid, this
// method will return an error.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyIn(
	inComingNStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecAbsoluteValueDto.CopyIn()\n "

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	return nStrFmtSpecAbsValDtoMolecule.copyIn(
		nStrFmtAbsValDto,
		inComingNStrFmtAbsValDto,
		ePrefix+
			"inComingNStrFmtAbsValDto -> nStrFmtAbsValDto\n ")
}

// CopyOut - Returns a deep copy of the current
// NumStrFmtSpecAbsoluteValueDto instance.
//
// If the current NumStrFmtSpecAbsoluteValueDto instance is judged
// to be invalid, this method will return an error.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) CopyOut() NumStrFmtSpecAbsoluteValueDto {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	newNStrFmtSpecAbsoluteValueDto,
		_ := nStrFmtSpecAbsValDtoMolecule.copyOut(
		nStrFmtAbsValDto,
		"")

	return newNStrFmtSpecAbsoluteValueDto
}

// GetAbsoluteValueFormat - Returns the formatting string used to
// format absolute numeric values in text number strings.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetAbsoluteValueFormat() string {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.absoluteValFmt
}

// GetNumberFieldLengthDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Absolute Value Dto.
//
// The NumberFieldDto details the length of the number field in which
// the signed numeric value will be displayed and right justified.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumberFieldLengthDto() NumberFieldDto {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.numFieldLenDto.CopyOut()
}

// GetNumberSeparatorsDto - Returns the NumStrFmtSpecDigitsSeparatorsDto
// instance currently configured for this Number String Format Specification
// Signed Number Value Dto.
//
// The Digits Separator Dto contains the decimal separator as well as the
// 'thousands' separators and the grouping sequence for separating thousands
// digits in the integer component of a number string.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetNumberSeparatorsDto() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.numberSeparatorsDto.CopyOut()
}

// GetTurnOnIntegerDigitsSeparationFlag - Returns the boolean flag
// which signals whether integer digits within a number string will
// be grouped by thousands and separated by an integer digits
// separator such as a comma (',').
//
// If this flag is set to 'true', integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1,000,000,000,000
//
// If this flag is set to 'false',  integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1000000000000
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	return nStrFmtAbsValDto.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecAbsoluteValueDto instance to determine whether
// the current instance is valid in all respects.
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
//       value will signal whether the current NumStrFmtSpecAbsoluteValueDto
//       is valid, or not. If the current NumStrFmtSpecAbsoluteValueDto
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	nStrFmtSpecAbsValDtoAtom :=
		numStrFmtSpecAbsoluteValueDtoAtom{}

	isValid,
		_ = nStrFmtSpecAbsValDtoAtom.testValidityOfAbsoluteValDto(
		nStrFmtAbsValDto,
		"")

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecAbsoluteValueDto instance to determine whether the
// current instance is valid in all respects.
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
//       If this instance of NumStrFmtSpecAbsoluteValueDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (nStrFmtAbsValDto *NumStrFmtSpecAbsoluteValueDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtAbsValDto.lock == nil {
		nStrFmtAbsValDto.lock = new(sync.Mutex)
	}

	nStrFmtAbsValDto.lock.Lock()

	defer nStrFmtAbsValDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecAbsoluteValueDto.IsValidInstanceError() \n" +
		"Testing Validity of 'nStrFmtAbsValDto' "

	nStrFmtSpecAbsValDtoAtom :=
		numStrFmtSpecAbsoluteValueDtoAtom{}

	_,
		err := nStrFmtSpecAbsValDtoAtom.testValidityOfAbsoluteValDto(
		nStrFmtAbsValDto,
		ePrefix)

	return err
}
