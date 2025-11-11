package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Chris-Greaves/pic-portal/photo-processor-service/types"
)

var (
	uploadsDir string
	outputDir  string
)

func main() {
	uploadsDir = os.Args[1]
	outputDir = os.Args[2]
	files, err := filepath.Glob(uploadsDir + "/*.info")
	if err != nil {
		fmt.Println(err)
		return
	}

	var uploads []types.FileInfo

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var info types.FileInfo
		json.Unmarshal(data, &info)
		uploads = append(uploads, info)
	}

	if len(uploads) > 0 {
		fmt.Println("Files with .info extension found and their contents:")
		for _, upload := range uploads {
			uploadStat, err := os.Lstat(filepath.Join(uploadsDir, upload.ID))
			if err != nil {
				fmt.Println(err)
				continue
			}

			if upload.Size != uploadStat.Size() || upload.MetaData["processed"] == "true" {
				fmt.Println("Skipping upload " + upload.ID)
				continue
			}
			err = processUpload(upload)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	} else {
		fmt.Println("No files with .info extension found")
	}
}

func processUpload(upload types.FileInfo) error {
	uploadBytes, err := os.ReadFile(filepath.Join(uploadsDir, upload.ID))
	if err != nil {
		return err
	}

	outputDir2 := filepath.Join(outputDir, upload.MetaData["uploaded_by"])
	exists, err := dirExists(outputDir2)
	if err != nil {
		return err
	}

	if !exists {
		err = os.MkdirAll(outputDir2, 0666)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filepath.Join(outputDir2, upload.MetaData["name"]), uploadBytes, 0666)
	if err != nil {
		return err
	}
	upload.MetaData["processed"] = "true"
	updatedInfo, err := json.Marshal(upload)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(uploadsDir, upload.ID)+".info", updatedInfo, 0666)
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
