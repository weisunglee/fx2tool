package filesystem

import (
	"io/ioutil"
	"log"
)

// GetSubFolderList : get subfolder list, depth = 1
func GetSubFolderList(root string) (folderList []string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			folderList = append(folderList, file.Name())
		}
	}
	return
}
