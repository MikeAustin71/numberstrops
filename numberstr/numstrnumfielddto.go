package numberstr

import (
	"fmt"
	"sync"
)

type NumberFieldDto struct {
	requestedNumFieldLength int         // User requested number field length
	actualNumFieldLength    int         // Machine generated actual number field Length
	minimumNumFieldLength   int         // Machine generated minimum number field length
	textJustifyFormat       TextJustify // User specified text justification
	lock                    *sync.Mutex
}

// Empty - Deletes and resets data values for all member variables
// in the current NumberFieldDto instance to their initial 'zero'
// values.
//
func (nFieldDto *NumberFieldDto) Empty() {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	_ = numStrNumFieldDtoQuark{}.ptr().empty(
		nFieldDto,
		nil)

	nFieldDto.lock.Unlock()

	nFieldDto.lock = nil
}

// CopyIn - Copies the data fields from an incoming NumberFieldDto
// to the data fields of the current instance of NumberFieldDto.
//
// When this method completes processing both the current
// NumberFieldDto instance and the incoming NumberFieldDto instance
// will have identical data values.
//
// This method will overwrite the member variable data values of
// the current NumberFieldDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingNumFieldDto        *NumberFieldDto
//     - A pointer to an instance of NumberFieldDto. The member
//       variable data fields from this instance will be copied to
//       those contained in the current NumberFieldDto instance.
//
//       If this instance of NumberFieldDto is judged to be invalid,
//       an error will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nFieldDto *NumberFieldDto) CopyIn(
	incomingNumFieldDto *NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumberFieldDto.CopyIn()")

	nStrNumFieldDtoElectron :=
		numStrNumFieldDtoElectron{}

	return nStrNumFieldDtoElectron.copyIn(
		nFieldDto,
		incomingNumFieldDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current NumberFieldDto instance.
//
// If the current NumberFieldDto instance is judged to be invalid, this
// method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumberFieldDto
//     - A new instance of NumberFieldDto containing data values
//       identical to those contained in the current NumberFieldDto
//       instance.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nFieldDto *NumberFieldDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumberFieldDto.CopyOut()")

	nStrNumFieldDtoElectron :=
		numStrNumFieldDtoElectron{}

	newNumFieldDto,
		err := nStrNumFieldDtoElectron.copyOut(
		nFieldDto,
		ePrefix.XCtx("nFieldDto->"))

	return newNumFieldDto, err
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

// GetNumberField - Receives a formatted number string and
// positions that number string in a number field that is returned
// to the caller. This output number field is a text string with
// string length and justification positioning determined by the
// settings of the current NumberFieldDto instance.
//
// The output field length will be determined by internal member
// variable, 'NumberFieldDto.requestedNumFieldLength'.
//
// If the current value of 'NumberFieldDto.requestedNumFieldLength'
// is less than the length of input parameter, 'formattedNumStr',
// the value of 'NumberFieldDto.requestedNumFieldLength' will be
// automatically increased to equal the length of 'formattedNumStr'.
//
// The position in the output number field will be determined by
// internal member variable 'NumberFieldDto.textJustifyFormat'. The
// text justification format specification will support one of
// three values: 'Right-Justified', 'Left-Justified' or 'Centered'.
//
// If the member variable,'NumberFieldDto.textJustifyFormat',is
// currently set to an invalid value, such as 'None', this method
// will default 'NumberFieldDto.textJustifyFormat' to a value of
// 'Right-Justified'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  formattedNumStr     string
//     - This is the formatted number string which will be inserted
//       into, and positioned within, the output number field text
//       string. This string will remain unaltered in content. It
//       will simply be positioned within a number field of equal
//       or greater length.
//
//       If 'formattedNumStr' is passed as an empty string, this
//       method, will return an error.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumberField      string
//     - This is output text number field. The length of this field
//       will be determined by the internal member variable
//       'requestedNumFieldLength' from the current NumberFieldDto
//       instance. The position of of input parameter
//       'formattedNumStr' within the output text number field will
//       be determined by internal member variable
//       'textJustifyFormat' from the current NumberFieldDto
//       instance.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nFieldDto *NumberFieldDto) GetNumberField(
	formattedNumStr string,
	ePrefix *ErrPrefixDto) (
	newNumberField string,
	err error) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumberFieldDto.GetNumberField()")

	lenFormattedNumStr := len(formattedNumStr)

	if lenFormattedNumStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formattedNumStr' is"+
			" an empty string!\n",
			ePrefix.String())

		return newNumberField, err
	}

	if nFieldDto.requestedNumFieldLength < lenFormattedNumStr {

		nFieldDto.requestedNumFieldLength = lenFormattedNumStr
	}

	if !nFieldDto.textJustifyFormat.XIsValid() {

		nFieldDto.textJustifyFormat =
			TextJustify(0).Right()
	}

	sOpsNanobot := strOpsNanobot{}

	newNumberField,
		err = sOpsNanobot.justifyTextInStrField(
		formattedNumStr,
		nFieldDto.requestedNumFieldLength,
		nFieldDto.textJustifyFormat,
		ePrefix.String())

	return newNumberField, err
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

// GetRequestedNumFieldLength - Returns the internal member variable
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
func (nFieldDto *NumberFieldDto) GetTextJustification() TextJustify {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	return nFieldDto.textJustifyFormat
}

// IsValidInstance - Performs a diagnostic review of the current
// NumberFieldDto instance to determine whether the current
// instance is valid in all respects.
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
//     - This returned boolean value will signal whether the current
//       NumberFieldDto is valid, or not. If the current
//       NumberFieldDto contains valid data, this method returns
//       'true'. If the data is invalid, this method will return
//       'false'.
//
func (nFieldDto *NumberFieldDto) IsValidInstance() (
	isValid bool) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	isValid,
		_ = nStrNFldDtoQuark.testValidityNumberFieldDto(
		nFieldDto,
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumberFieldDto instance to determine whether the current
// instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumberFieldDto contains invalid
//       data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nFieldDto *NumberFieldDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumberFieldDto.IsValidInstanceError()")

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err := nStrNFldDtoQuark.testValidityNumberFieldDto(
		nFieldDto,
		ePrefix)

	return err
}

// New - Creates and returns a new instance of NumberFieldDto.
//
// 'requestedNumFieldLength' is set to zero.
//
// 'textJustifyFormat' is set to TextJustify(0).Right().
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumberFieldDto
//     - If this method completes successfully, a new instance of
//       NumberFieldDto will be returned to the caller.
//
func (nFieldDto NumberFieldDto) New() NumberFieldDto {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	newNumFieldDto := NumberFieldDto{
		requestedNumFieldLength: 0,
		actualNumFieldLength:    0,
		minimumNumFieldLength:   0,
		textJustifyFormat:       TextJustify(0).Right(),
		lock:                    nil,
	}

	return newNumFieldDto
}

// NewBasic - Creates and returns a new instance of NumberFieldDto.
//
// The input parameters represent the minimum information required
// to configure the data values for a NumberFieldDto object.
// Default values are used to supplement these input parameters.
//
// This method will default the text justification to
// 'right-justify'. This setting will position the number string on
// the right margin of the number field.
//            Example: "      123456.54"
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  requestedNumberFieldLen    int
//     - This is the requested length of the number field in which
//       the number string will be displayed.
//
//       If this field length is greater than the actual length of
//       the number string, the number string will be justified
//       within the number field. If the actual number string
//       length is greater than the requested number field length,
//       the number field length will be automatically expanded
//       to display the entire number string. The 'requested'
//       number field length is used to create number fields
//       of standard lengths for text presentations.
//
//       If 'requestedNumberFieldLen' is set to a value of minus
//       one (-1), the final number field length will be set to
//       the length of the actual number string.
//
//       If 'requestedNumberFieldLen' is set equal to zero or to a
//       value less than minus one (-1), it will be automatically
//       converted to minus one (-1).
//
//
//  numberFieldTextJustify        TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to a number string when
//       it is positioned inside of a number field. This
//       enumeration value must be one of the three following
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  newNumFieldDto      NumberFieldDto
//     - If this method completes successfully, a new instance of
//       NumberFieldDto will be returned to the caller.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nFieldDto NumberFieldDto) NewBasic(
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	newNumFieldDto NumberFieldDto,
	err error) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumberFieldDto.NewBasic()")

	if !numberFieldTextJustify.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberFieldTextJustify' is invalid!\n"+
			"numberFieldTextJustify integer value='%v'\n",
			ePrefix.String(),
			numberFieldTextJustify.XValueInt())

		return newNumFieldDto, err
	}

	err = numStrNumFieldDtoMechanics{}.ptr().
		setNumberFieldDto(
			&newNumFieldDto,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumFieldDto, err
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

// SetBasic - Sets the internal member variable values
// for the current instance of NumberFieldDto.
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
//  numberFieldTextJustify        TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to a number string when
//       it is positioned inside of a number field. This
//       enumeration value must be one of the three following
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nFieldDto *NumberFieldDto) SetBasic(
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumberFieldDto.SetBasic()")

	nStrNumFieldDtoMech := numStrNumFieldDtoMechanics{}

	return nStrNumFieldDtoMech.setNumberFieldDto(
		nFieldDto,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)
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
//  numberFieldTextJustify        TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to a number string when
//       it is positioned inside of a number field. This
//       enumeration value must be one of the three following
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
	numberFieldTextJustify TextJustify) {

	if nFieldDto.lock == nil {
		nFieldDto.lock = new(sync.Mutex)
	}

	nFieldDto.lock.Lock()

	defer nFieldDto.lock.Unlock()

	nFieldDto.textJustifyFormat = numberFieldTextJustify
}
