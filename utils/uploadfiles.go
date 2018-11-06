package utils

import (
	"os"
	"io"
	"mime/multipart"
	fpath "path/filepath"
)

// UploadFiles to given path
func UploadFiles(filename string, file multipart.File) (err error) {
	workDir, _ := os.Getwd()
	filesDir := fpath.Join(workDir, "files", filename)
	out, err := os.Create(filesDir)
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(out, file)
	return nil
}