package numberstr

import "sync"

// NumStrIntSeparatorsDto is used to transport and manage an array
// of NumStrIntSeparator. These integer separators are in turn
// used to control the grouping and separation characters
// employed in separating digits within an integer number.
//
//  United States Example:        1,000,000,000
//  European Example:             1 000 000 000
//  Indian Number System Example: 6,78,90,00,00,00,00,000
//
type NumStrIntSeparatorsDto struct {
	intSeparators []NumStrIntSeparator
	lock          *sync.Mutex
}
