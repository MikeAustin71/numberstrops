package numberstr

type IBasicErrorPrefix interface {
	GetEPrefStrings() [][2]string
}

type IErrorPrefix interface {
	GetIsLastLineTerminatedWithNewLine() bool

	GetEPrefCollectionLen() int

	SetCtx(newErrContext string)

	SetCtxEmpty()

	SetEPref(newErrPrefix string)

	SetEPrefCtx(newErrPrefix string, newErrContext string)

	SetEPrefOld(oldErrPrefix string)

	SetMaxTextLineLen(maxErrPrefixTextLineLength int)

	SetIsLastLineTermWithNewLine(isLastLineTerminatedWithNewLine bool)

	String() string
}
