package numberstr

import (
	"sync"
)

// FormatterCollection - Encapsulates all of the format specifications
// required to format the numeric values contained in type NumStrDto.
//
type FormatterCollection struct {
	idNo          uint64
	idString      string
	description   string
	tag           string
	fmtCollection []INumStrFormatter
	lock          *sync.Mutex
}
