package numberstr

import (
	"fmt"
	"io"
	"sync"
)

type strOpsElectron struct {
	lock *sync.Mutex
}

// readBytes - Implements io.Reader interface for type StrOps.
// 'readBytes' reads up to len(p) bytes into 'p'. This
// method supports buffered 'read' operations.
//
// The internal member string variable, 'StrOps.stringData'
// is written into 'p'. When the end of 'StrOps.stringData'
// is written to 'p', the method returns error = 'io.EOF'.
//
// 'StrOps.stringData' can be accessed through Getter an Setter methods,
// StrOps.GetStringData() and StrOps.SetStringData()
//
func (sOpsElectron *strOpsElectron) readBytes(
	strOpsInstance *StrOps,
	p []byte,
	ePrefix string) (
	n int,
	err error) {

	if sOpsElectron.lock == nil {
		sOpsElectron.lock = new(sync.Mutex)
	}

	sOpsElectron.lock.Lock()

	defer sOpsElectron.lock.Unlock()

	ePrefix += "strOpsElectron.readBytes() "

	n = 0
	err = nil

	if strOpsInstance == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'strOpsInstance' is a 'nil' pointer!\n",
			ePrefix)

		return n, err
	}

	n = len(p)
	err = io.EOF

	if n == 0 {
		strOpsInstance.cntBytesRead = 0
		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' is zero length!\n",
			ePrefix)

		return 0, err
	}

	strData := strOpsInstance.stringData

	w := []byte(strData)

	lenW := uint64(len(w))

	cntBytesRead := strOpsInstance.cntBytesRead

	if lenW == 0 ||
		cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
		n = 0
		return n, err
	}

	startReadIdx := cntBytesRead

	remainingBytesToRead := lenW - cntBytesRead

	if uint64(n) < remainingBytesToRead {
		remainingBytesToRead = startReadIdx + uint64(n)
		err = nil
	} else {
		remainingBytesToRead += startReadIdx
		err = io.EOF
	}

	n = 0

	for i := startReadIdx; i < remainingBytesToRead; i++ {
		p[n] = w[i]
		n++
	}

	cntBytesRead += uint64(n)

	if cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
	} else {
		strOpsInstance.cntBytesRead = cntBytesRead
	}

	return n, err
}
