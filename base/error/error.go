package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	_, err := os.Open("/test.txt")
	if err , ok := err.(*os.PathError); ok {
		fmt.Println("file at path", err.Path, "failed to open")
		//return
	}
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("matched files", files)
}