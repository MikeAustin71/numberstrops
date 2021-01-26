package numberstr

import (
	"fmt"
	"sync"
)

type numStrNumFieldDtoMechanics struct {
	lock *sync.Mutex
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
func (nStrNumFieldDtoMech *numStrNumFieldDtoMechanics) setNumberFieldDto(
	numFieldDto *NumberFieldDto,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix string) (err error) {

	if nStrNumFieldDtoMech.lock == nil {
		nStrNumFieldDtoMech.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoMech.lock.Lock()

	defer nStrNumFieldDtoMech.lock.Unlock()

	ePrefix += "numStrNumFieldDtoMechanics.setNumberFieldDto()\n"
	err = nil

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix)

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
			ePrefix,
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
			ePrefix+
				"Testing validity of final 'numFieldDto'\n")

	return err
}
