package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {

	//move everything to a subdir to skip
	err := filepath.WalkDir("/Users/erdal.cinar/personal/bootcamp/week-4-homework-1-nowo/ListGoSourceCode", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("walk failed: %v\n", err)
	}
}
