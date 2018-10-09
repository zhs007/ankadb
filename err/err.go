package ankadberr

import (
	pb "github.com/zhs007/ankadb/proto"
)

// AnkaError -
type AnkaError interface {
	Error() string
	ErrCode() pb.CODE
}

// ankaError -
type ankaError struct {
	errcode pb.CODE
}

// FormatCode -
func FormatCode(code pb.CODE) pb.CODE {
	if _, ok := pb.CODE_name[int32(code)]; !ok {
		return pb.CODE_INVALID_CODE
	}

	return code
}

func (err *ankaError) Error() string {
	return "ankaError - [" + pb.CODE_name[int32(FormatCode(err.errcode))] + "]"
}

func (err *ankaError) ErrCode() pb.CODE {
	return err.errcode
}

// NewError -
func NewError(errcode pb.CODE) AnkaError {
	err := &ankaError{errcode: errcode}

	return err
}

// BuildErrorString -
func BuildErrorString(errcode pb.CODE) string {
	return "ankaError - [" + pb.CODE_name[int32(FormatCode(errcode))] + "]"
}

// GetErrCode -
func GetErrCode(err interface{}) pb.CODE {
	if ce, ok := err.(ankaError); ok {
		return ce.ErrCode()
	}

	return pb.CODE_INVALID_CODE
}
