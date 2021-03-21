package numberstr

import (
	"fmt"
	"sync"
)

type formatterSciNotationElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance
// of FormatterSciNotation ('incomingFmtSciNotation') to the
// target FormatterSciNotation instance ('targetFmtSciNotation').
//
// The FormatterSciNotation type encapsulates the Scientific
// Notation format specification used to format number strings for
// text displays.
//
// When this method completes processing 'targetFmtSciNotation' and
// 'incomingFmtSciNotation' will have identical data values.
//
// This method will overwrite the member variable data values of
// input parameter, 'targetFmtSciNotation'.
//
// If input parameter 'incomingFmtSciNotation' is judged to be
// invalid, this method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetFmtSciNotation          *FormatterSciNotation
//     - A pointer to an instance of FormatterSciNotation.
//       The internal member variable data values for this instance
//       will be set equal to those of the FormatterSciNotation
//       instance, 'incomingSciNotDto'. Therefore be aware that the
//       existing data values of 'targetSciNotDto' will be
//       overwritten.
//
//  incomingFmtSciNotation        *FormatterSciNotation
//     - A pointer to a second instance of FormatterSciNotation.
//       The member variable data fields from this instance will be
//       copied to those contained in input parameter,'targetSciNotDto'.
//
//       If 'incomingFmtSciNotation' is judged to be invalid, this
//       method will return an error.
//
//
//  ePrefix                       *ErrPrefixDto
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
func (fmtSciNotElectron *formatterSciNotationElectron) copyIn(
	targetFmtSciNotation *FormatterSciNotation,
	incomingFmtSciNotation *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotElectron.lock == nil {
		fmtSciNotElectron.lock = new(sync.Mutex)
	}

	fmtSciNotElectron.lock.Lock()

	defer fmtSciNotElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationElectron.copyIn()")

	if targetFmtSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFmtSciNotation' is invalid!\n"+
			"'targetFmtSciNotation' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if incomingFmtSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFmtSciNotation' is invalid!\n"+
			"'incomingFmtSciNotation' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if targetFmtSciNotation.lock == nil {
		targetFmtSciNotation.lock = new(sync.Mutex)
	}

	_,
		err = formatterSciNotationQuark{}.ptr().
		testValidityOfFormatterSciNotation(
			incomingFmtSciNotation,
			ePrefix.XCtx(
				"incomingFmtSciNotation"))

	if err != nil {
		return err
	}

	targetFmtSciNotation.numStrFmtType =
		incomingFmtSciNotation.numStrFmtType

	targetFmtSciNotation.significandUsesLeadingPlus =
		incomingFmtSciNotation.significandUsesLeadingPlus

	targetFmtSciNotation.mantissaLength =
		incomingFmtSciNotation.mantissaLength

	targetFmtSciNotation.exponentChar =
		incomingFmtSciNotation.exponentChar

	targetFmtSciNotation.exponentUsesLeadingPlus =
		incomingFmtSciNotation.exponentUsesLeadingPlus

	return targetFmtSciNotation.numFieldLenDto.CopyIn(
		&incomingFmtSciNotation.numFieldLenDto,
		ePrefix.XCtx(
			"incomingFmtSciNotation.numFieldLenDto->"+
				"targetFmtSciNotation.numFieldLenDto"))
}

// copyOut - Returns a deep copy of the FormatterSciNotation
// instance passed as input parameter, 'formatterSciNotation'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  formatterSciNotation          *FormatterSciNotation
//     - A pointer to an instance of FormatterSciNotation.
//       The internal member variable data values for this instance
//       will be copied to a new instance of
//       FormatterSciNotation, ('newFormatterSciNotation'),
//       which is returned to the caller.
//
//       If 'formatterSciNotation' is judged to be invalid, this
//       method will return an error.
//
//
//  ePrefix                       *ErrPrefixDto
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
//  newFormatterSciNotation       FormatterSciNotation
//     - A new instance of FormatterSciNotation containing
//       data values identical to those contained in input
//       parameter, 'formatterSciNotation'. This returned instance
//       of FormatterSciNotation represents a deep copy of
//       input parameter, 'formatterSciNotation'.
//
//       If 'formatterSciNotation' is judged to be invalid, this
//       method will return an error.
//
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
func (fmtSciNotElectron *formatterSciNotationElectron) copyOut(
	formatterSciNotation *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	newFormatterSciNotation FormatterSciNotation,
	err error) {

	if fmtSciNotElectron.lock == nil {
		fmtSciNotElectron.lock = new(sync.Mutex)
	}

	fmtSciNotElectron.lock.Lock()

	defer fmtSciNotElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationElectron.copyOut()")

	if formatterSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterSciNotation' is invalid!\n"+
			"'formatterSciNotation' is a 'nil' pointer.\n",
			ePrefix.String())

		return newFormatterSciNotation, err
	}

	if formatterSciNotation.lock == nil {
		formatterSciNotation.lock = new(sync.Mutex)
	}

	_,
		err = formatterSciNotationQuark{}.ptr().
		testValidityOfFormatterSciNotation(
			formatterSciNotation,
			ePrefix.XCtx("formatterSciNotation"))

	if err != nil {
		return newFormatterSciNotation, err
	}

	newFormatterSciNotation.numStrFmtType =
		formatterSciNotation.numStrFmtType

	newFormatterSciNotation.significandUsesLeadingPlus =
		formatterSciNotation.significandUsesLeadingPlus

	newFormatterSciNotation.mantissaLength =
		formatterSciNotation.mantissaLength

	newFormatterSciNotation.exponentChar =
		formatterSciNotation.exponentChar

	newFormatterSciNotation.exponentUsesLeadingPlus =
		formatterSciNotation.exponentUsesLeadingPlus

	newFormatterSciNotation.lock = new(sync.Mutex)

	err =
		newFormatterSciNotation.numFieldLenDto.CopyIn(
			&formatterSciNotation.numFieldLenDto,
			ePrefix.XCtx(
				"formatterSciNotation.numFieldLenDto->"+
					"newFormatterSciNotation.numFieldLenDto"))

	return newFormatterSciNotation, err
}

// ptr - Returns a pointer to a new instance of
// formatterSciNotationElectron.
//
func (fmtSciNotElectron formatterSciNotationElectron) ptr() *formatterSciNotationElectron {

	if fmtSciNotElectron.lock == nil {
		fmtSciNotElectron.lock = new(sync.Mutex)
	}

	fmtSciNotElectron.lock.Lock()

	defer fmtSciNotElectron.lock.Unlock()

	newSciNotElectron :=
		new(formatterSciNotationElectron)

	newSciNotElectron.lock = new(sync.Mutex)

	return newSciNotElectron
}
