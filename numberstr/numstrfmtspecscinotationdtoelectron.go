package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSciNotationDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance
// of NumStrFmtSpecSciNotationDto ('incomingSciNotDto') to the
// target NumStrFmtSpecSciNotationDto instance ('targetSciNotDto').
//
// The NumStrFmtSpecSciNotationDto type encapsulates the Scientific
// Notation format specification used to format number strings for
// text displays.
//
// When this method completes processing 'targetSciNotDto' and
// 'incomingSciNotDto' will have identical data values.
//
// This method will overwrite the member variable data values of
// input parameter, 'targetSciNotDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetSciNotDto            *NumStrFmtSpecSciNotationDto
//     - A pointer to an instance of NumStrFmtSpecSciNotationDto.
//       The internal member variable data values for this instance
//       will be set equal to those of the NumStrFmtSpecSciNotationDto
//       instance, 'incomingSciNotDto'. Therefore be aware that the
//       existing data values of 'targetSciNotDto' will be
//       overwritten.
//
//  incomingSciNotDto          *NumStrFmtSpecSciNotationDto
//     - A pointer to a second instance of NumStrFmtSpecSciNotationDto.
//       The member variable data fields from this instance will be
//       copied to those contained in input parameter,'targetSciNotDto'.
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
func (nStrFmtSpecSciNotDtoElectron *numStrFmtSpecSciNotationDtoElectron) copyIn(
	targetSciNotDto *NumStrFmtSpecSciNotationDto,
	incomingSciNotDto *NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecSciNotDtoElectron.lock == nil {
		nStrFmtSpecSciNotDtoElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDtoElectron.lock.Lock()

	defer nStrFmtSpecSciNotDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecSciNotationDtoElectron.copyIn()")

	if targetSciNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if incomingSciNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingSciNotDto' is invalid!\n"+
			"'incomingSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if targetSciNotDto.lock == nil {
		targetSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotQuark :=
		numStrFmtSpecSciNotationDtoQuark{}

	_,
		err = nStrFmtSpecSciNotQuark.testValidityOfNumStrFmtSpecSciNotationDto(
		incomingSciNotDto,
		ePrefix.XCtx(
			"incomingSciNotDto"))

	if err != nil {
		return err
	}

	targetSciNotDto.significandUsesLeadingPlus =
		incomingSciNotDto.significandUsesLeadingPlus

	targetSciNotDto.mantissaLength =
		incomingSciNotDto.mantissaLength

	targetSciNotDto.exponentChar =
		incomingSciNotDto.exponentChar

	targetSciNotDto.exponentUsesLeadingPlus =
		incomingSciNotDto.exponentUsesLeadingPlus

	return targetSciNotDto.numFieldLenDto.CopyIn(
		&incomingSciNotDto.numFieldLenDto,
		ePrefix.XCtx(
			"incomingSciNotDto.numFieldLenDto->"+
				"targetSciNotDto.numFieldLenDto"))
}

// copyOut - Returns a deep copy of the NumStrFmtSpecSciNotationDto
// instance passed as input parameter, 'numStrFmtSpecSciNotDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFmtSpecSciNotDto     *NumStrFmtSpecSciNotationDto
//     - A pointer to an instance of NumStrFmtSpecSciNotationDto.
//       The internal member variable data values for this instance
//       will be copied to a new instance of
//       NumStrFmtSpecSciNotationDto, ('newNumStrFmtSpecSciNotDto'),
//       which is returned to the caller.
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
//  newNumStrFmtSpecSciNotDto  NumStrFmtSpecSciNotationDto
//     - A new instance of NumStrFmtSpecSciNotationDto containing
//       data values identical to those contained in input
//       parameter, 'numStrFmtSpecSciNotDto'. This returned instance
//       of NumStrFmtSpecSciNotationDto represents a deep copy of
//       input parameter, 'numStrFmtSpecSciNotDto'.
//
//       If 'numStrFmtSpecSciNotDto' is judged to be invalid, this
//       method will return an error.
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
func (nStrFmtSpecSciNotDtoElectron *numStrFmtSpecSciNotationDtoElectron) copyOut(
	numStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) (
	newNumStrFmtSpecSciNotDto NumStrFmtSpecSciNotationDto,
	err error) {

	if nStrFmtSpecSciNotDtoElectron.lock == nil {
		nStrFmtSpecSciNotDtoElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDtoElectron.lock.Lock()

	defer nStrFmtSpecSciNotDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecSciNotationDtoElectron.copyOut()")

	if numStrFmtSpecSciNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpecSciNotDto' is invalid!\n"+
			"'numStrFmtSpecSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return newNumStrFmtSpecSciNotDto, err
	}

	if numStrFmtSpecSciNotDto.lock == nil {
		numStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotQuark :=
		numStrFmtSpecSciNotationDtoQuark{}

	_,
		err = nStrFmtSpecSciNotQuark.testValidityOfNumStrFmtSpecSciNotationDto(
		numStrFmtSpecSciNotDto,
		ePrefix.XCtx("numStrFmtSpecSciNotDto"))

	if err != nil {
		return newNumStrFmtSpecSciNotDto, err
	}

	newNumStrFmtSpecSciNotDto.significandUsesLeadingPlus =
		numStrFmtSpecSciNotDto.significandUsesLeadingPlus

	newNumStrFmtSpecSciNotDto.mantissaLength =
		numStrFmtSpecSciNotDto.mantissaLength

	newNumStrFmtSpecSciNotDto.exponentChar =
		numStrFmtSpecSciNotDto.exponentChar

	newNumStrFmtSpecSciNotDto.exponentUsesLeadingPlus =
		numStrFmtSpecSciNotDto.exponentUsesLeadingPlus

	newNumStrFmtSpecSciNotDto.lock = new(sync.Mutex)

	err =
		newNumStrFmtSpecSciNotDto.numFieldLenDto.CopyIn(
			&numStrFmtSpecSciNotDto.numFieldLenDto,
			ePrefix.XCtx(
				"numStrFmtSpecSciNotDto.numFieldLenDto->"+
					"newNumStrFmtSpecSciNotDto.numFieldLenDto"))

	return newNumStrFmtSpecSciNotDto, err
}
