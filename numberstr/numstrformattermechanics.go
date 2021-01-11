package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatterMechanics struct {
	lock *sync.Mutex
}

func (nStrFormatterMech *numStrFormatterMechanics) testValidityOfNumStrFormatter(
	nStrFormatter *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormatterMech.lock == nil {
		nStrFormatterMech.lock = new(sync.Mutex)
	}

	nStrFormatterMech.lock.Lock()

	defer nStrFormatterMech.lock.Unlock()

	ePrefix += "numStrFormatterMechanics.testValidityOfNumStrFormatter() "

	isValid = false

	if nStrFormatter == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFormatter' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrNanobot := numStrFormatterNanobot{}

	isValid = true
	err = nil

	switch nStrFormatter.valueDisplaySpec {

	case NumStrValSpec(0).CurrencyValue():

		isValid,
			err = nStrFormtrNanobot.testValidityOfCurrencyFormatter(
			nStrFormatter,
			ePrefix)

	case NumStrValSpec(0).AbsoluteValue():

		isValid,
			err = nStrFormtrNanobot.testValidityOfAbsoluteValueFormatter(
			nStrFormatter,
			ePrefix)

	case NumStrValSpec(0).SignedNumberValue():

		isValid,
			err = nStrFormtrNanobot.testValidityOfSignedNumberValueFormatter(
			nStrFormatter,
			ePrefix)
	default:
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrFormatter.valueDisplaySpec' is invalid!\n"+
			"nStrFormatter.valueDisplaySpec== '%v'\n",
			ePrefix,
			nStrFormatter.valueDisplaySpec.XValueInt())
	}

	return isValid, err
}
