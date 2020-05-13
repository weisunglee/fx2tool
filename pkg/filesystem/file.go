package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
)

// OpenFile : Read file into byte array
func OpenFile(path string) ([]byte, error) {

	fileHandle, err := os.Open(path)

	if err != nil {
		fmt.Println(path, "open failed")
		return nil, err
	}

	defer fileHandle.Close()
	data, err := ioutil.ReadAll(fileHandle)

	if err != nil {
		fmt.Println(path + "read failed")
		return nil, err
	}

	return data, nil
}
