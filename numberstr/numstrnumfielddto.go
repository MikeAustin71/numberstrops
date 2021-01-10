package numberstr

import "sync"

type numberFieldDto struct {
	requestedNumFieldLength int
	actualNumFieldLength    int
	minimumNumFieldLength   int
	lock                    *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming numberFieldDto
// to the data fields of the current instance of numberFieldDto.
//
func (nFieldDto *numberFieldDto) CopyIn(
	numFieldDtoIn *numberFieldDto) {

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

// CopyOut - Returns a deep copy of the current numberFieldDto instance.
//
func (nFieldDto *numberFieldDto) CopyOut() numberFieldDto {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	newNumFieldDto := numberFieldDto{}

	newNumFieldDto.requestedNumFieldLength =
		nFieldDto.requestedNumFieldLength
	newNumFieldDto.actualNumFieldLength =
		nFieldDto.actualNumFieldLength
	newNumFieldDto.minimumNumFieldLength =
		nFieldDto.minimumNumFieldLength

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto
}

// New - Creates and returns a new instance of 'numberFieldDto'.
//
// The only input parameter is the 'requested' number field length.
// This is the requested length of the number field in which the
// number string will be displayed. If this field length is greater
// than the actual length of the number string, the number string
// will be right justified within the the number field. If the
// actual number string length is greater than the number field
// length, the number field field length will be automatically
// expanded to display the entire number string. The 'requested'
// number field length is used to create number fields of standard
// length.
//
func (nFieldDto numberFieldDto) New(
	requestedNumberFieldLen int) numberFieldDto {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	newNumFieldDto := numberFieldDto{}

	newNumFieldDto.requestedNumFieldLength = requestedNumberFieldLen

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
func (nFieldDto *numberFieldDto) SetRequestedNumFieldLength(
	numFieldLen int) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.requestedNumFieldLength =
		numFieldLen

}
