package numberstr

type INumStrFormatter interface {
	SetNumStrFormatTypeCode(fmtTypeCode NumStrFormatTypeCode)

	GetNumStrFormatTypeCode() NumStrFormatTypeCode
}
