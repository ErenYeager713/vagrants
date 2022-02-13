package main

import(
		"fmt"
    "os"
    "path/filepath"
		"strings"
  	"io/ioutil"
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
				var res = strings.Contains(file, "kubearmor_policies_default_explorer_knoxautopolicy")
				if res == true {
					fmt.Println("Found knoxautopolicy")
					data, err1 := ioutil.ReadFile("file.txt")
				  if err1 != nil {
				    fmt.Println("File reading error", err1)
				    return
				  }
				  fmt.Println("Contents of file:")
				  fmt.Println(string(data))
				}
    }
}
