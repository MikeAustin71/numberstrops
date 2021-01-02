package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFormatDto struct {
	valueDisplaySpec NumStrValSpec
	positiveValueFmt NumStrPosValFmtMode
	negativeValueFmt NumStrNegValFmtMode

	lock *sync.Mutex
}

// New - Creates and returns a new instance of
// NumStrFormatDto.
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  valueDisplaySpec    NumStrValSpec
//     - An enumeration value for Number String Value Specification.
//       Used to specify display of positive and and negative values,
//       or absolute values.
//
//
//  positiveValueFmt    NumStrPosValFmtMode
//     - An enumeration value for Number String Positive Value Format
//       Mode. Used to specify formatting for positive numeric values.
//
//
//  negativeValueFmt    NumStrNegValFmtMode
//     - An enumeration value for Number String Negative Value Format
//       Mode. Used to specify formatting for negative numeric values.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newFmtDto           NumStrFormatDto
//     - If this method completes successfully, a new instance of NumStrFormatDto
//       be returned to calling function.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto NumStrFormatDto) New(
	valueDisplaySpec NumStrValSpec,
	positiveValueFmt NumStrPosValFmtMode,
	negativeValueFmt NumStrNegValFmtMode,
	ePrefix string) (
	newFmtDto NumStrFormatDto,
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.New() "

	if !valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			valueDisplaySpec.XValueInt())
		return newFmtDto, err
	}

	if !positiveValueFmt.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'positiveValueFmt' is invalid!\n"+
			"positiveValueFmt='%v'\n",
			positiveValueFmt.XValueInt())
		return newFmtDto, err
	}

	if !negativeValueFmt.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'negativeValueFmt' is invalid!\n"+
			"negativeValueFmt='%v'\n",
			negativeValueFmt.XValueInt())
		return newFmtDto, err
	}

	newFmtDto.valueDisplaySpec = valueDisplaySpec
	newFmtDto.negativeValueFmt = negativeValueFmt
	newFmtDto.positiveValueFmt = positiveValueFmt

	return newFmtDto, err
}

// GetValueDisplaySpec - Returns the NumStrValSpec
// value.
//
func (nStrFmtDto *NumStrFormatDto) GetValueDisplaySpec() (
	valueDisplaySpec NumStrValSpec) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	valueDisplaySpec = nStrFmtDto.valueDisplaySpec

	return valueDisplaySpec
}

// GetPositiveValueFormatMode - Returns the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) GetPositiveValueFormatMode() (
	positiveValueFmtMode NumStrPosValFmtMode) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	positiveValueFmtMode = nStrFmtDto.positiveValueFmt

	return positiveValueFmtMode
}

// GetNegativeValueFormatMode - Returns the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) GetNegativeValueFormatMode() (
	negativeValueFmtMode NumStrNegValFmtMode) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	negativeValueFmtMode = nStrFmtDto.negativeValueFmt

	return negativeValueFmtMode
}

// SetValueDisplaySpec - Sets the NumStrValSpec
// value.
//
func (nStrFmtDto *NumStrFormatDto) SetValueDisplaySpec(
	valueDisplaySpec NumStrValSpec,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetValueDisplaySpec() "

	if !valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			valueDisplaySpec.XValueInt())
		return err
	}

	nStrFmtDto.valueDisplaySpec = valueDisplaySpec

	return err
}

// SetPositiveValueFormatMode - Sets the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) SetPositiveValueFormatMode(
	positiveValueFmtMode NumStrPosValFmtMode,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetPositiveValueFormatMode() "

	if !positiveValueFmtMode.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'positiveValueFmtMode' is invalid!\n"+
			"positiveValueFmtMode='%v'\n",
			positiveValueFmtMode.XValueInt())
		return err
	}

	nStrFmtDto.positiveValueFmt = positiveValueFmtMode

	return err
}

// SetNegativeValueFormatMode - Sets the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) SetNegativeValueFormatMode(
	negativeValueFmtMode NumStrNegValFmtMode,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNegativeValueFormatMode() "

	if !negativeValueFmtMode.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'negativeValueFmtMode' is invalid!\n"+
			"negativeValueFmtMode='%v'\n",
			negativeValueFmtMode.XValueInt())
		return err
	}

	negativeValueFmtMode = nStrFmtDto.negativeValueFmt

	return err
}
