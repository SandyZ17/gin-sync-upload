package adapter

import "io"

func UploadAdapter(fileName string, fileContext io.ReadCloser) error {
	buf := make([]byte, 4096)
	for {
		fileContext.Read(buf)
	}
}
