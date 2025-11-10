package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Chris-Greaves/pic-portal/photo-processor-service/types"
)

func main() {
	dir := os.Args[1]
	files, err := filepath.Glob(dir + "/*.info")
	if err != nil {
		fmt.Println(err)
		return
	}

	var contents []types.FileInfo

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var info types.FileInfo
		json.Unmarshal(data, &info)
		contents = append(contents, info)
	}

	if len(contents) > 0 {
		fmt.Println("Files with .info extension found and their contents:")
		for _, content := range contents {
			var jsonString, _ = json.Marshal(content)
			fmt.Println(string(jsonString))
		}
	} else {
		fmt.Println("No files with .info extension found")
	}
}
