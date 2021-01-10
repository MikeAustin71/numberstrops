package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatDtoMechanics struct {
	lock *sync.Mutex
}

// copyIn - Receives pointers to two instances of NumStrFormatDto
// labeled as input parameters, 'numStrFmtDto1' and 'numStrFmtDto2'.
// This method then proceeds to copy all of the data from
// 'numStrFmtDto2' into 'numStrFmtDto1'.  When the copy process is
// completed, the data values in 'numStrFmtDto1' and 'numStrFmtDto2'
// will be identical.
//
// IMPORTANT
//
// Be advised, this method will overwrite the data values contained
// in input parameter, 'numStrFmtDto1'.
//
func (nStrFmtDtoMech *numStrFormatDtoMechanics) copyIn(
	numStrFmtDto1 *NumStrFormatDto,
	numStrFmtDto2 *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDtoMech.lock == nil {
		nStrFmtDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtDtoMech.lock.Lock()

	defer nStrFmtDtoMech.lock.Unlock()

	ePrefix += "numStrFormatDtoMechanics.copyIn() "

	if numStrFmtDto1 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto1' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if numStrFmtDto2 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto2' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	numStrFmtDto1.valueDisplaySpec =
		numStrFmtDto2.valueDisplaySpec

	numStrFmtDto1.positiveValueFmt =
		numStrFmtDto2.positiveValueFmt

	numStrFmtDto1.negativeValueFmt =
		numStrFmtDto2.negativeValueFmt

	numStrFmtDto1.currencyFmt =
		numStrFmtDto2.currencyFmt.CopyOut()

	numStrFmtDto1.decimalSeparator =
		numStrFmtDto2.decimalSeparator

	numStrFmtDto1.thousandsSeparator =
		numStrFmtDto2.thousandsSeparator

	numStrFmtDto1.turnOnThousandsSeparator =
		numStrFmtDto2.turnOnThousandsSeparator

	numStrFmtDto1.numberFieldDto =
		numStrFmtDto2.numberFieldDto.CopyOut()

	return err
}

// copyOut - Receives a pointer to an instance of NumStrFormatDto
// and returns a deep copy of that NumStrFormatDto instance.
//
func (nStrFmtDtoMech *numStrFormatDtoMechanics) copyOut(
	numStrFmtDto *NumStrFormatDto,
	ePrefix string) (
	newNumStrFormatDto NumStrFormatDto,
	err error) {

	if nStrFmtDtoMech.lock == nil {
		nStrFmtDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtDtoMech.lock.Lock()

	defer nStrFmtDtoMech.lock.Unlock()

	ePrefix += "numStrFormatDtoMechanics.copyOut() "
	newNumStrFormatDto = NumStrFormatDto{}

	if numStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)
		return newNumStrFormatDto, err
	}

	newNumStrFormatDto.valueDisplaySpec =
		numStrFmtDto.valueDisplaySpec

	newNumStrFormatDto.positiveValueFmt =
		numStrFmtDto.positiveValueFmt

	newNumStrFormatDto.negativeValueFmt =
		numStrFmtDto.negativeValueFmt

	newNumStrFormatDto.currencyFmt =
		numStrFmtDto.currencyFmt.CopyOut()

	newNumStrFormatDto.decimalSeparator =
		numStrFmtDto.decimalSeparator

	newNumStrFormatDto.thousandsSeparator =
		numStrFmtDto.thousandsSeparator

	newNumStrFormatDto.turnOnThousandsSeparator =
		numStrFmtDto.turnOnThousandsSeparator

	newNumStrFormatDto.numberFieldDto =
		numStrFmtDto.numberFieldDto.CopyOut()

	newNumStrFormatDto.lock = new(sync.Mutex)

	return newNumStrFormatDto, err
}

// testNumStrFormatDtoValidity - Receives an instance of
// NumStrFormatDto and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (nStrFmtDtoMech *numStrFormatDtoMechanics) testNumStrFormatDtoValidity(
	numStrFmtDto *NumStrFormatDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtDtoMech.lock == nil {
		nStrFmtDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtDtoMech.lock.Lock()

	defer nStrFmtDtoMech.lock.Unlock()

	ePrefix += "NumStrFormatDto.IsValidInstanceError() "

	isValid = false

	if numStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	if !numStrFmtDto.valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Member variable 'NumStrFormatDto.valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			numStrFmtDto.valueDisplaySpec.XValueInt())
		return isValid, err
	}

	nStrFmtNanobot := numStrFormatNanobot{}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		numStrFmtDto.positiveValueFmt,
		ePrefix+"numStrFmtDto.positiveValueFmt ")

	if err != nil {
		return isValid, err
	}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		numStrFmtDto.negativeValueFmt,
		ePrefix+"numStrFmtDto.negativeValueFmt ")

	if err != nil {
		return isValid, err
	}

	err = numStrFmtDto.currencyFmt.IsValidInstanceError(
		ePrefix + "Testing numStrFmtDto.currencyFmt validity. ")

	if err != nil {
		return isValid, err
	}

	if numStrFmtDto.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: NumStrFormatDto.decimalSeparator is a zero value and\n"+
			"invalid!\n",
			ePrefix)
		return isValid, err
	}

	if numStrFmtDto.thousandsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: NumStrFormatDto.thousandsSeparator is a zero value and\n"+
			"invalid!\n",
			ePrefix)
		return isValid, err
	}

	isValid = true
	return isValid, err
}
