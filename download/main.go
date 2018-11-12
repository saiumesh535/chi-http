package download

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	fpath "path/filepath"
)


// SmallFiles small files
func SmallFiles(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	// this would give you path directory/folder/filename.pdf/jpeg**
	fileLocation := fpath.Join(workDir, "files", "somefilename.extension")
	file, _ := os.Open(fileLocation)
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err);
		w.Write([]byte("something wen't wrong!!"));
        return
    }
	w.Write(bytes)
}