package numberstr

import "sync"

type NumberFieldDto struct {
	requestedNumFieldLength int
	actualNumFieldLength    int
	minimumNumFieldLength   int
	lock                    *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming NumberFieldDto
// to the data fields of the current instance of NumberFieldDto.
//
func (nFieldDto *NumberFieldDto) CopyIn(
	numFieldDtoIn *NumberFieldDto) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.requestedNumFieldLength =
		numFieldDtoIn.requestedNumFieldLength

	nFieldDto.actualNumFieldLength =
		numFieldDtoIn.actualNumFieldLength

	nFieldDto.minimumNumFieldLength =
		numFieldDtoIn.minimumNumFieldLength

}

// CopyOut - Returns a deep copy of the current NumberFieldDto instance.
//
func (nFieldDto *NumberFieldDto) CopyOut() NumberFieldDto {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	newNumFieldDto := NumberFieldDto{}

	newNumFieldDto.requestedNumFieldLength =
		nFieldDto.requestedNumFieldLength
	newNumFieldDto.actualNumFieldLength =
		nFieldDto.actualNumFieldLength
	newNumFieldDto.minimumNumFieldLength =
		nFieldDto.minimumNumFieldLength

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto
}

// New - Creates and returns a new instance of 'NumberFieldDto'.
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  requestedNumberFieldLen    int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the the
//       number field. If the actual number string length is greater
//       than the requested number field length, the number field
//       length will be automatically expanded to display the entire
//       number string. The 'requested' number field length is used
//       to create number fields of standard lengths for text
//       presentations.
//
func (nFieldDto NumberFieldDto) New(
	requestedNumberFieldLen int) NumberFieldDto {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	newNumFieldDto := NumberFieldDto{}

	newNumFieldDto.requestedNumFieldLength = requestedNumberFieldLen

	newNumFieldDto.minimumNumFieldLength = -1

	newNumFieldDto.actualNumFieldLength = -1

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto
}

// SetRequestedNumFieldLength - Sets the 'Requested' Number
// Field length. This is the requested length of the number field
// in which the number string will be displayed. If this field
// length is greater than the actual length of the number string,
// the number string will be right justified within the the number
// field. If the actual number string length is greater than the
// number field length, the number field field length will be
// automatically expanded to display the entire number string. The
// 'requested' number field length is used to create number fields
// of standard length.
//
func (nFieldDto *NumberFieldDto) SetRequestedNumFieldLength(
	numFieldLen int) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.requestedNumFieldLength =
		numFieldLen

}

// SetMinimumNumFieldLength - Sets the minimum number field
// length.
//
func (nFieldDto *NumberFieldDto) SetMinimumNumFieldLength(
	minimumNumFieldLen int) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.minimumNumFieldLength =
		minimumNumFieldLen

}

// SetMinimumNumFieldLength - Sets the actual number field
// length.
//
func (nFieldDto *NumberFieldDto) SetActualNumFieldLength(
	actualNumFieldLen int) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.actualNumFieldLength =
		actualNumFieldLen

}
