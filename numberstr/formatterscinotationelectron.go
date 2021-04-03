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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

	return targetFmtSciNotation.numFieldDto.CopyIn(
		&incomingFmtSciNotation.numFieldDto,
		ePrefix.XCtx(
			"incomingFmtSciNotation.numFieldDto->"+
				"targetFmtSciNotation.numFieldDto"))
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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
		newFormatterSciNotation.numFieldDto.CopyIn(
			&formatterSciNotation.numFieldDto,
			ePrefix.XCtx(
				"formatterSciNotation.numFieldDto->"+
					"newFormatterSciNotation.numFieldDto"))

	return newFormatterSciNotation, err
}

// equal - Receives two FormatterSciNotation objects and proceeds
// to determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSciNotationOne   *FormatterSciNotation
//     - A pointer to the first FormatterSciNotation object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterSciNotation object to determine if they are
//       equivalent.
//
//
//  fmtSciNotationTwo   *FormatterSciNotation
//     - A pointer to the second FormatterSciNotation object. This
//       method will compare all data elements in the first
//       FormatterSciNotation object to all corresponding data
//       elements in this second FormatterSciNotation object to
//       determine if they are equivalent.
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
//  isEqual             bool
//     - If all the data elements in 'fmtSciNotationOne' are equal to
//       all the corresponding data elements in 'fmtSciNotationTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtSciNotationOne' are equal to
//       all the corresponding data elements in 'fmtSciNotationTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtSciNotElectron *formatterSciNotationElectron) equal(
	fmtSciNotationOne *FormatterSciNotation,
	fmtSciNotationTwo *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtSciNotElectron.lock == nil {
		fmtSciNotElectron.lock = new(sync.Mutex)
	}

	fmtSciNotElectron.lock.Lock()

	defer fmtSciNotElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	isEqual = false

	ePrefix.SetEPref(
		"formatterSciNotationElectron.copyIn()")

	if fmtSciNotationOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSciNotationOne' is invalid!\n"+
			"'fmtSciNotationOne' is a 'nil' pointer.\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtSciNotationTwo == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSciNotationTwo' is invalid!\n"+
			"'fmtSciNotationTwo' is a 'nil' pointer.\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtSciNotationOne.significandUsesLeadingPlus !=
		fmtSciNotationTwo.significandUsesLeadingPlus {

		err = fmt.Errorf("%v\n"+
			"fmtSciNotationOne.significandUsesLeadingPlus!=\n"+
			"  fmtSciNotationTwo.significandUsesLeadingPlus\n"+
			"fmtSciNotationOne.significandUsesLeadingPlus='%v'\n"+
			"fmtSciNotationTwo.significandUsesLeadingPlus='%v'\n",
			ePrefix.String(),
			fmtSciNotationOne.significandUsesLeadingPlus,
			fmtSciNotationTwo.significandUsesLeadingPlus)

		return isEqual, err
	}

	if fmtSciNotationOne.mantissaLength !=
		fmtSciNotationTwo.mantissaLength {

		err = fmt.Errorf("%v\n"+
			"fmtSciNotationOne.mantissaLength!=\n"+
			"  fmtSciNotationTwo.mantissaLength\n"+
			"fmtSciNotationOne.mantissaLength='%v'\n"+
			"fmtSciNotationTwo.mantissaLength='%v'\n",
			ePrefix.String(),
			fmtSciNotationOne.mantissaLength,
			fmtSciNotationTwo.mantissaLength)

		return isEqual, err
	}

	if fmtSciNotationOne.exponentChar !=
		fmtSciNotationTwo.exponentChar {

		err = fmt.Errorf("%v\n"+
			"fmtSciNotationOne.exponentChar!=\n"+
			"  fmtSciNotationTwo.exponentChar\n"+
			"fmtSciNotationOne.exponentChar='%v'\n"+
			"fmtSciNotationTwo.exponentChar='%v'\n",
			ePrefix.String(),
			string(fmtSciNotationOne.exponentChar),
			string(fmtSciNotationTwo.exponentChar))

		return isEqual, err
	}

	if fmtSciNotationOne.exponentUsesLeadingPlus !=
		fmtSciNotationTwo.exponentUsesLeadingPlus {

		err = fmt.Errorf("%v\n"+
			"fmtSciNotationOne.exponentUsesLeadingPlus!=\n"+
			"  fmtSciNotationTwo.exponentUsesLeadingPlus\n"+
			"fmtSciNotationOne.exponentUsesLeadingPlus='%v'\n"+
			"fmtSciNotationTwo.exponentUsesLeadingPlus='%v'\n",
			ePrefix.String(),
			fmtSciNotationOne.exponentUsesLeadingPlus,
			fmtSciNotationTwo.exponentUsesLeadingPlus)

		return isEqual, err
	}

	_,
		err = fmtSciNotationOne.numFieldDto.Equal(
		fmtSciNotationTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtSciNotationOne vs fmtSciNotationTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
