package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func renumber(target_dir string, target_date string) {
	files, err := filepath.Glob(target_dir + target_date + "_*.png")
	if err != nil {
		panic(err)
	}

	target_list := []string{}
	for _, file := range files {
		basename := filepath.Base(file)
		target_list = append(target_list, basename)
	}

	sort.Strings(target_list)

	for n, before_filename := range target_list {
		after_filename := fmt.Sprintf("%s_%05d.png", target_date, n+1)
		if before_filename != after_filename {
			fmt.Println(before_filename + " -> " + after_filename)
			os.Rename(target_dir+before_filename, target_dir+after_filename)
		}
	}
}

func main() {
	renumber(os.Args[1], os.Args[2])
}
