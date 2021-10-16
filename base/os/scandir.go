package main

import (
	"fmt"
	"io/ioutil"
)

func scanDir(dirname string, level int)  {
	s := "|--"
	for i := 0; i < level; i++{
		s = "|  " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, fi.Name())
		if fi.IsDir() {
			scanDir(filename, level + 1)
		}
	}
}


func main(){
	scanDir("./", 2)
}