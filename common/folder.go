package common

import (
	"log"
	"os"
)

func CheckFolderExist(filePath string) error {
	// 如不存在则创建文件夹
	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Printf("create file %v failed", filePath)
			return err
		} else {
			log.Printf("create file %v successed", filePath)
			return nil
		}
	}
	return err
}
