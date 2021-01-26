package numberstr

import (
	"fmt"
	"strings"
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

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrNumFieldDtoElectron.copyIn() "

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

// copyOut - Returns a deep copy of the NumberFieldDto instance
// passed as input parameter, 'numFieldDto'.
//
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
//
//  ePrefix                    string
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
//  newNumFieldDto             NumberFieldDto
//     - A new instance of NumberFieldDto containing data values
//       identical to those contained in input parameter,
//       'numFieldDto'. This returned instance of NumberFieldDto
//       represents a deep copy of input parameter, 'numFieldDto'.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrNumFieldDtoElectron *numStrNumFieldDtoElectron) copyOut(
	numFieldDto *NumberFieldDto,
	ePrefix string) (
	newNumFieldDto NumberFieldDto,
	err error) {

	if nStrNumFieldDtoElectron.lock == nil {
		nStrNumFieldDtoElectron.lock = new(sync.Mutex)
	}

	nStrNumFieldDtoElectron.lock.Lock()

	defer nStrNumFieldDtoElectron.lock.Unlock()

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix)

		return newNumFieldDto, err
	}

	if numFieldDto.lock == nil {
		numFieldDto.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark := numStrNumFieldDtoQuark{}

	_,
		err = nStrNFldDtoQuark.testValidityNumberFieldDto(
		numFieldDto,
		ePrefix+"\n"+
			"Testing Validity of 'numFieldDto' ")

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
