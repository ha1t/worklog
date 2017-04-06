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

	finished := make(chan bool)

	chan_count := 0
	for n, bf := range target_list {
		num := n
		before_filename := bf
		after_filename := fmt.Sprintf("%s__%05d.png", target_date, num+1)

		// 並列化のジレンマについて
		// 並列実行すると、処理順序の関係でファイルを上書きしてしまう可能性があるため
		// 書き換え先のファイル名はかぶらないようにする必要がある。
		// その場合、すべてのファイルをrenameするため、ずれてなければpassできた直列実行よりも遅くなる
		//if before_filename != after_filename {
		go func() {
			fmt.Println(before_filename + " -> " + after_filename)
			os.Rename(target_dir+before_filename, target_dir+after_filename)
			finished <- true
		}()
		chan_count += 1
		//}
	}

	// 終わるまで待つ
	for i := 1; i <= chan_count; i++ {
		<-finished
	}
}

func main() {
	renumber(os.Args[1], os.Args[2])
}
