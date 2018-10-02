package ankadb

import (
	"os"

	"github.com/zhs007/jarviscore/err"
	pb "github.com/zhs007/jarviscore/proto"
)

func loadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileinfo.Size()
	buffer := make([]byte, fileSize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	if int64(bytesread) != fileSize {
		return nil, jarviserr.NewError(pb.CODE_INVALID_FILEREADSIZE)
	}

	return buffer, nil
}
