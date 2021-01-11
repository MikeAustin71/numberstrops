package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtDtoMolecule struct {
	lock *sync.Mutex
}

// testValidityOfNumStrFmtCollection - Inspects the collection of Number
// String Formatters and returns an error if any data elements are found
// to be invalid.
//
// If this collection is valid, this method will return a boolean value of
// 'true' and set the returned error object to 'nil'.
//
func (nStrFmtDtoMolecule *numStrFmtDtoMolecule) testValidityOfNumStrFmtCollection(
	numStrFormatterCollection map[NumStrValSpec]NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtDtoMolecule.lock == nil {
		nStrFmtDtoMolecule.lock = new(sync.Mutex)
	}

	nStrFmtDtoMolecule.lock.Lock()

	defer nStrFmtDtoMolecule.lock.Unlock()

	ePrefix += "numStrFmtDtoNanobot.testValidityOfNumStrFmtCollection() "

	isValid = false

	if len(numStrFormatterCollection) != 3 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter numStrFormatterCollection should\n"+
			"consist of 3-Configurations. Instead, this collection\n"+
			"has %v-Configurations.\n",
			ePrefix,
			len(numStrFormatterCollection))
	}

	currencyCnt := 0
	absValCnt := 0
	signedNumCnt := 0

	for idx, nStrFormtr := range numStrFormatterCollection {

		err =
			nStrFormtr.IsValidInstanceError(ePrefix +
				fmt.Sprintf("Collection Index='%v' ", idx))

		if err != nil {
			isValid = false
			return isValid, err
		}

		switch nStrFormtr.valueDisplaySpec {
		case NumStrValSpec(0).CurrencyValue():
			currencyCnt++
		case NumStrValSpec(0).AbsoluteValue():
			absValCnt++
		case NumStrValSpec(0).SignedNumberValue():
			signedNumCnt++
		default:
			isValid = false
			err = fmt.Errorf("%v\n"+
				"Error: Invalid 'NumStrFormatter.valueDisplaySpec'\n"+
				"Unknown NumStrValSpec Value='%v' \n"+
				"NumStrFormatter Collection Index='%v'\n",
				ePrefix,
				nStrFormtr.valueDisplaySpec.XValueInt(),
				idx)
			return isValid, err
		}

	}

	if currencyCnt != 1 {

		isValid = false

		err = fmt.Errorf("%v\n"+
			"Error: The Number String Formatter Collection is invalid!\n"+
			"Expected the number of Currency Formatters to equal '1'.\n"+
			"Instead, the number of Currency Formatters='%v' \n"+
			ePrefix,
			currencyCnt)

		return isValid, err
	}

	if absValCnt != 1 {

		isValid = false

		err = fmt.Errorf("%v\n"+
			"Error: The Number String Formatter Collection is invalid!\n"+
			"Expected the number of Absolute Value Formatters to equal '1'.\n"+
			"Instead, the number of Absolute Value Formatters='%v' \n"+
			ePrefix,
			absValCnt)

		return isValid, err
	}

	if signedNumCnt != 1 {

		isValid = false

		err = fmt.Errorf("%v\n"+
			"Error: The Number String Formatter Collection is invalid!\n"+
			"Expected the number of Signed Number Value Formatters to equal '1'.\n"+
			"Instead, the number of Signed Number Value Formatters='%v' \n"+
			ePrefix,
			absValCnt)

		return isValid, err
	}

	isValid = true

	return isValid, err
}
