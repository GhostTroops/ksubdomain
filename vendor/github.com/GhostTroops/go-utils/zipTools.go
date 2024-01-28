package go_utils

import (
	"bytes"
	"compress/gzip"
)

func GzipBytes(data []byte) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil
	}
	if err := gz.Close(); err != nil {
		return nil
	}
	return buf.Bytes()
}

//func RmZipFile(zipFile string, rmFs ...string) {
//	jarWriter, err := zip33.OpenReader(zipFile)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer jarWriter.Close()
//	for _, file := range jarWriter.File {
//		for _, x := range rmFs {
//			if file.Name == x {
//				file.Remove()
//
//			}
//		}
//	}
//}
