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
func (nStrNumFieldDtoElectron *numStrNumFieldDtoElectron) copyIn(
	targetNumFieldDto *NumberFieldDto,
	incomingNumFieldDto *NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	ePrefix += "\nnumStrNumFieldDtoElectron.copyIn() "

	if targetNumFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'targetNumFieldDto' is a 'nil' pointer!\n",
			ePrefix)

		return err
	}

	if incomingNumFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'incomingNumFieldDto' is a 'nil' pointer!\n",
			ePrefix)

		return err
	}

	if targetNumFieldDto.lock == nil {
		targetNumFieldDto.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err = nStrNFldDtoQuark.testValidityNumberFieldDto(
		incomingNumFieldDto,
		ePrefix+"\n"+
			"Testing Validity of 'incomingNumFieldDto' ")

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
