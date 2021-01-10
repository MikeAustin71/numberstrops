package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatDtoUtility struct {
	lock *sync.Mutex
}

// setToDefaults - Receives a pointer to an instance of
// NumStrFormatDto and proceeds to set all of the internal
// member variables to their default values.
//
// The Number String Format default values represent formatting
// parameters used in the United States.
//
func (nStrFmtDtoUtil *numStrFormatDtoUtility) setToDefaults(
	nStrFmtDto *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDtoUtil.lock == nil {
		nStrFmtDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtDtoUtil.lock.Lock()

	defer nStrFmtDtoUtil.lock.Unlock()

	ePrefix += "numStrFormatDtoUtility.setToDefaults() "

	if nStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter nStrFmtDto is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()

	nStrFmtQuark := numStrFormatQuark{}

	nStrFmtDto.positiveValueFmt =
		nStrFmtQuark.getDefaultPositiveNumStrFormat()

	nStrFmtDto.negativeValueFmt =
		nStrFmtQuark.getDefaultNegativeNumStrFormat()

	nStrFmtDto.decimalSeparator =
		nStrFmtQuark.getDefaultDecimalSeparator()

	nStrFmtDto.thousandsSeparator =
		nStrFmtQuark.getDefaultThousandsSeparator()

	nStrFmtDto.currencyFmt =
		nStrFmtQuark.getDefaultCurrencySymbol()

	return err
}
