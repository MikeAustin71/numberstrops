package numberstr

import (
	"fmt"
	"sync"
)

type numStrNumFieldDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance
// of NumberFieldDto to the target NumberFieldDto instance.
//
// When this method completes processing 'targetNumFieldDto'
// and 'incomingNumFieldDto' will have identical data values.
//
// This method will overwrite the member variable data values of
// input parameter, 'targetNumFieldDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetNumFieldDto          *NumberFieldDto
//     - A pointer to an instance of NumberFieldDto. The internal
//       member variable data values for this instance will be set
//       equal to those of NumberFieldDto instance,
//       'incomingNumFieldDto'.
//
//  incomingNumFieldDto        *NumberFieldDto
//     - A pointer to a second instance of NumberFieldDto. The
//       member variable data fields from this instance will be
//       copied to those contained in input parameter,
//       'targetNumFieldDto'.
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
//  err                        error
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
func (nStrNumFieldDtoElectron *numStrNumFieldDtoElectron) copyIn(
	targetNumFieldDto *NumberFieldDto,
	incomingNumFieldDto *NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("numStrNumFieldDtoElectron.copyIn()")

	if targetNumFieldDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNumFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNumFieldDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingNumFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if targetNumFieldDto.lock == nil {
		targetNumFieldDto.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err = nStrNFldDtoQuark.testValidityNumberFieldDto(
		incomingNumFieldDto,
		ePrefix.XCtx("Testing Validity of 'incomingNumFieldDto'"))

	if err != nil {
		return err
	}

	targetNumFieldDto.requestedNumFieldLength =
		incomingNumFieldDto.requestedNumFieldLength

	targetNumFieldDto.actualNumFieldLength =
		incomingNumFieldDto.actualNumFieldLength

	targetNumFieldDto.minimumNumFieldLength =
		incomingNumFieldDto.minimumNumFieldLength

	targetNumFieldDto.textJustifyFormat =
		incomingNumFieldDto.textJustifyFormat

	return err
}

// copyOut - Returns a deep copy of the NumberFieldDto instance
// passed as input parameter, 'numFieldDto'.
//
// If the NumberFieldDto instance is judged to be invalid, this
// method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numFieldDto                *NumberFieldDto
//     - A pointer to an instance of NumberFieldDto. The internal
//       member variable data values for this instance will be
//       copied to a new instance of NumberFieldDto,
//       ('newNumFieldDto'), which is returned to the caller.
//
//       The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
//
//
//  ePrefix                    *ErrPrefixDto
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
//  newNumFieldDto             NumberFieldDto
//     - A new instance of NumberFieldDto containing data values
//       identical to those contained in input parameter,
//       'numFieldDto'. This returned instance of NumberFieldDto
//       represents a deep copy of input parameter, 'numFieldDto'.
//
//
//  err                        error
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
func (nStrNumFieldDtoElectron *numStrNumFieldDtoElectron) copyOut(
	numFieldDto *NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	newNumFieldDto NumberFieldDto,
	err error) {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("numStrNumFieldDtoElectron.copyOut()")

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return newNumFieldDto, err
	}

	if numFieldDto.lock == nil {
		numFieldDto.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err = nStrNFldDtoQuark.testValidityNumberFieldDto(
		numFieldDto,
		ePrefix.XCtx("Testing Validity of 'numFieldDto'"))

	if err != nil {
		return newNumFieldDto, err
	}

	newNumFieldDto.requestedNumFieldLength =
		numFieldDto.requestedNumFieldLength

	newNumFieldDto.actualNumFieldLength =
		numFieldDto.actualNumFieldLength

	newNumFieldDto.minimumNumFieldLength =
		numFieldDto.minimumNumFieldLength

	newNumFieldDto.textJustifyFormat =
		numFieldDto.textJustifyFormat

	return newNumFieldDto, err
}

// equal - Receives two NumberFieldDto objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numFieldDtoOne      *NumberFieldDto
//     - A pointer to the first NumberFieldDto object. This method
//       will compare all data elements in this object to
//       corresponding data elements in the second NumberFieldDto
//       object in order determine equivalency.
//
//
//  numFieldDtoTwo      *NumberFieldDto
//     - A pointer to the second NumberFieldDto object. This method
//       will compare all data elements in the first NumberFieldDto
//       object to corresponding data elements in this second
//       NumberFieldDto object in order determine equivalency.
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
//  equal             bool
//     - If all the data elements in 'numFieldDtoOne' are equal to
//       all the corresponding data elements in 'numFieldDtoTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'numFieldDtoOne' are equal to
//       all the corresponding data elements in 'numFieldDtoTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (nStrNumFieldDtoElectron *numStrNumFieldDtoElectron) equal(
	numFieldDtoOne *NumberFieldDto,
	numFieldDtoTwo *NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numStrNumFieldDtoElectron.equal()")

	isEqual = false

	if numFieldDtoOne == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDtoOne' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if numFieldDtoTwo == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDtoTwo' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if numFieldDtoOne.requestedNumFieldLength !=
		numFieldDtoTwo.requestedNumFieldLength {
		err = fmt.Errorf("%v\n"+
			"numFieldDtoOne.requestedNumFieldLength !=\n"+
			"numFieldDtoTwo.requestedNumFieldLength\n"+
			"numFieldDtoOne.requestedNumFieldLength='%v'\n"+
			"numFieldDtoTwo.requestedNumFieldLength='%v'\n",
			ePrefix.String(),
			numFieldDtoOne.requestedNumFieldLength,
			numFieldDtoTwo.requestedNumFieldLength)

		return isEqual, err
	}

	if numFieldDtoOne.actualNumFieldLength !=
		numFieldDtoTwo.actualNumFieldLength {
		err = fmt.Errorf("%v\n"+
			"numFieldDtoOne.actualNumFieldLength !=\n"+
			"numFieldDtoTwo.actualNumFieldLength\n"+
			"numFieldDtoOne.actualNumFieldLength='%v'\n"+
			"numFieldDtoTwo.actualNumFieldLength='%v'\n",
			ePrefix.String(),
			numFieldDtoOne.actualNumFieldLength,
			numFieldDtoTwo.actualNumFieldLength)

		return isEqual, err
	}

	if numFieldDtoOne.minimumNumFieldLength !=
		numFieldDtoTwo.minimumNumFieldLength {
		err = fmt.Errorf("%v\n"+
			"numFieldDtoOne.minimumNumFieldLength !=\n"+
			"numFieldDtoTwo.minimumNumFieldLength\n"+
			"numFieldDtoOne.minimumNumFieldLength='%v'\n"+
			"numFieldDtoTwo.minimumNumFieldLength='%v'\n",
			ePrefix.String(),
			numFieldDtoOne.minimumNumFieldLength,
			numFieldDtoTwo.minimumNumFieldLength)

		return isEqual, err
	}

	if numFieldDtoOne.textJustifyFormat !=
		numFieldDtoTwo.textJustifyFormat {
		err = fmt.Errorf("%v\n"+
			"numFieldDtoOne.textJustifyFormat !=\n"+
			"numFieldDtoTwo.textJustifyFormat\n"+
			"numFieldDtoOne.textJustifyFormat='%v'\n"+
			"numFieldDtoTwo.textJustifyFormat='%v'\n",
			ePrefix.String(),
			numFieldDtoOne.textJustifyFormat.String(),
			numFieldDtoTwo.textJustifyFormat.String())

		return isEqual, err
	}

	isEqual = true

	return isEqual, err
}

// ptr - Returns a pointer to a new instance of
// numStrNumFieldDtoElectron.
//
func (nStrNumFieldDtoElectron numStrNumFieldDtoElectron) ptr() *numStrNumFieldDtoElectron {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	newNumFieldDtoElectron :=
		new(numStrNumFieldDtoElectron)

	newNumFieldDtoElectron.lock = new(sync.Mutex)

	return newNumFieldDtoElectron
}
