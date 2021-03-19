package numberstr

import (
	"fmt"
	"sync"
)

type numStrNumFieldDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numStrNumFieldDtoMechanics.
//
func (nStrNumFieldDtoMech numStrNumFieldDtoMechanics) ptr() *numStrNumFieldDtoMechanics {

	if nStrNumFieldDtoMech.lock == nil {
		nStrNumFieldDtoMech.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoMech.lock.Lock()

	defer nStrNumFieldDtoMech.lock.Unlock()

	newNStrNumFieldDtoMech :=
		new(numStrNumFieldDtoMechanics)

	newNStrNumFieldDtoMech.lock = new(sync.Mutex)

	return newNStrNumFieldDtoMech
}

// setNumberFieldDto - Sets the internal member variable values
// for an instance of NumberFieldDto passed as an input parameter.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numFieldDto                *NumberFieldDto
//     - A pointer to an instance of NumberFieldDto. The internal
//       member variable data values for this instance will be set
//       according to the following input parameters.
//
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                           error
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
func (nStrNumFieldDtoMech *numStrNumFieldDtoMechanics) setNumberFieldDto(
	numFieldDto *NumberFieldDto,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (err error) {

	if nStrNumFieldDtoMech.lock == nil {
		nStrNumFieldDtoMech.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoMech.lock.Lock()

	defer nStrNumFieldDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrNumFieldDtoMechanics.setNumberFieldDto()")
	err = nil

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if numFieldDto.lock == nil {
		numFieldDto.lock = new(sync.Mutex)
	}

	if requestedNumberFieldLen < 1 {
		requestedNumberFieldLen = -1
	}

	if !numberFieldTextJustify.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberFieldTextJustify' is invalid!\n"+
			"'numberFieldTextJustify' integer value = '%v'\n",
			ePrefix.String(),
			numberFieldTextJustify.XValueInt())
		return err
	}

	numFieldDto.requestedNumFieldLength =
		requestedNumberFieldLen

	numFieldDto.minimumNumFieldLength = -1

	numFieldDto.actualNumFieldLength = -1

	numFieldDto.textJustifyFormat =
		numberFieldTextJustify

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err =
		nStrNFldDtoQuark.testValidityNumberFieldDto(
			numFieldDto,
			ePrefix.XCtx("Testing validity of final 'numFieldDto'"))

	return err
}
