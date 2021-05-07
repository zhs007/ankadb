package ankadb

// AnkaDBCoreVersion - ankadb core version
const AnkaDBCoreVersion = "v0.3.35"

type requestID string

const RequestIDKey = requestID("ankadb")

// func loadFile(filename string) ([]byte, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	fileinfo, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fileSize := fileinfo.Size()
// 	buffer := make([]byte, fileSize)

// 	bytesread, err := file.Read(buffer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if int64(bytesread) != fileSize {
// 		return nil, ErrLoadFileReadSize
// 	}

// 	return buffer, nil
// }
