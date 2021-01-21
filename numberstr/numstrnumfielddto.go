package numberstr

import (
	"fmt"
	"sync"
)

type NumberFieldDto struct {
	requestedNumFieldLength int
	actualNumFieldLength    int
	minimumNumFieldLength   int
	textJustifyFormat       StrOpsTextJustify
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

	nFieldDto.textJustifyFormat =
		numFieldDtoIn.textJustifyFormat

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

	newNumFieldDto.textJustifyFormat =
		nFieldDto.textJustifyFormat

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto
}

// GetActualNumFieldLength - Returns the internal member variable
// Actual Number Field Length (NumberFieldDto.actualNumFieldLength).
//
func (nFieldDto *NumberFieldDto) GetActualNumFieldLength() int {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	return nFieldDto.actualNumFieldLength
}

// GetMinimumNumFieldLength - Returns the internal member variable
// Minimum Number Field Length:
//    NumberFieldDto.minimumNumFieldLength
//
func (nFieldDto *NumberFieldDto) GetMinimumNumFieldLength() int {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	return nFieldDto.minimumNumFieldLength
}

// GetMinimumNumFieldLength - Returns the internal member variable
// Requested Number Field Length:
//    NumberFieldDto.requestedNumFieldLength
//
func (nFieldDto *NumberFieldDto) GetRequestedNumFieldLength() int {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	return nFieldDto.requestedNumFieldLength
}

// GetTextJustification - Returns the internal member variable
// Text Justification Format:
//    NumberFieldDto.textJustifyFormat
//
func (nFieldDto *NumberFieldDto) GetTextJustification() StrOpsTextJustify {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	return nFieldDto.textJustifyFormat
}

// New - Creates and returns a new instance of 'NumberFieldDto'.
//
// Be advised that this version of method 'New()' will default
// the text justification to right-justify the number string in
// the number field.
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

	// Default to Right-Justified text format
	newNumFieldDto.textJustifyFormat =
		StrOpsTextJustify(0).Right()

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto
}

// NewFromComponents - This version of the 'New' method also
// creates and returns a new instance of NumberFieldDto. However,
// this version provides the caller with the flexibility to
// configure the text justification format used when inserting a
// number string in a text number field.
//
//
//
// ------------------------------------------------------------------------
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
//
//  textJustify         StrOpsTextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to 'strToJustify' when
//       it is positioned inside of the returned output string.
//       This enumeration value must be one of the three following
//       format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumberFieldDto
//     - If the method completes successfully, this parameter will
//       represent a new instance of NumberFieldDto populated in
//       accordance with the input parameters.
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nFieldDto NumberFieldDto) NewFromComponents(
	requestedNumberFieldLen int,
	textJustify StrOpsTextJustify,
	ePrefix string) (
	NumberFieldDto,
	error) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	ePrefix += "NumberFieldDto.NewFromComponents()\n"

	if !textJustify.XIsValid() {

		return NumberFieldDto{},
		fmt.Errorf("%v\n" +
			"Error: Input parameter 'textJustify' is invalid!\n" +
			"'textJustify' integer value = '%v'\n",
			ePrefix,
			textJustify.XValueInt())

	}

	newNumFieldDto := NumberFieldDto{}

	newNumFieldDto.requestedNumFieldLength = requestedNumberFieldLen

	newNumFieldDto.minimumNumFieldLength = -1

	newNumFieldDto.actualNumFieldLength = -1

	// Default to Right-Justified text format
	newNumFieldDto.textJustifyFormat =
		textJustify

	newNumFieldDto.lock = new(sync.Mutex)

	return newNumFieldDto, nil
}

// SetActualNumFieldLength - Sets the actual number field
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
	requestedNumberFieldLen int) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.requestedNumFieldLength =
		requestedNumberFieldLen

}

// SetTextJustification - Sets the text justification format
// for this number field. Input parameter 'textJustify' provides
// the enumeration value which is used to set internal member
// variable 'NumberFieldDto.textJustifyFormat'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textJustify         StrOpsTextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to 'strToJustify' when
//       it is positioned inside of the returned output string.
//       This enumeration value must be one of the three following
//       format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nFieldDto *NumberFieldDto) SetTextJustification(
	textJustify StrOpsTextJustify) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.textJustifyFormat = textJustify
}
