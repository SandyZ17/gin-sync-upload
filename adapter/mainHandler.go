package adapter

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/SandyZ17/gin-sync-upload/common"
)

func UploadSliceFile(fileName string, fileType string, fileContext io.ReadCloser) error {
	filePath := "upload/" + fileName + "/"
	// h := md5.New()
	h := md5.Sum([]byte(fileName))
	fileHash := fmt.Sprintf("%x", h)
	// check folder exist or not
	err := common.CheckFolderExist(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	// slice request body and write them into different part
	count := 0
	for {
		buf := make([]byte, 4194304)
		n, err := fileContext.Read(buf[:])
		if n < 0 {
			log.Println("Read Slice Err")
			return err
		} else if n == 0 {
			break
		} else if n > 0 {
			sliceFileName := filePath + fileHash + "-" + strconv.Itoa(count)
			file, err := os.OpenFile(sliceFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
			if err != nil {
				log.Default().Printf("Open file %v failed", sliceFileName)
				return err
			}
			_, err = file.Write(buf[0:n])
			if err != nil {
				log.Default().Printf("Write file %v failed", file.Name())
			}
			file.Close()
			count+=1
		}
	}
	return nil
}

func DownloadFile() {

}
