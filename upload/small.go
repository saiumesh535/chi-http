package upload

import (
	"net/http"
	util "../utils"
)

// FileUpload to given destination
func FileUpload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close();
	r.ParseForm();
	fileResume, fileHeader, err := r.FormFile("resume");
	if err != nil {
		panic("reading file error")
	}
	defer fileResume.Close()
	uploadErr := util.UploadSmallFiles(fileHeader.Filename, fileResume);
	if uploadErr != nil {
		w.Write([]byte("something wen't wrong!!"))
	} else {
		w.Write([]byte("success"))
	}
}
