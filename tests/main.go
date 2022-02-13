package main

import(
		"fmt"
    "os"
    "path/filepath"
)

func main(){
	var files []string

    root := "/home/runner/work/vagrants/vagrants"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}
