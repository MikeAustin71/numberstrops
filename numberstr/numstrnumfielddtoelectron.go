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
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
