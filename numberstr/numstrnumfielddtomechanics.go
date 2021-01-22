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
	textJustify TextJustify,
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

	if !textJustify.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textJustify' is invalid!\n"+
			"'textJustify' integer value = '%v'\n",
			ePrefix,
			textJustify.XValueInt())
		return err
	}

	numFieldDto.requestedNumFieldLength =
		requestedNumberFieldLen

	numFieldDto.minimumNumFieldLength = -1

	numFieldDto.actualNumFieldLength = -1

	numFieldDto.textJustifyFormat =
		textJustify

	return err
}
