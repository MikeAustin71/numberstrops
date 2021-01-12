package numberstr

import "sync"

type NumStrFmtSpecScientificNotationDto struct {
	mantissaLength          uint
	exponentChar            rune
	exponentUsesLeadingPlus bool
	numFieldLen             NumberFieldDto
	lock                    *sync.Mutex
}
