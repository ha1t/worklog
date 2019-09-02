package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func renumber(targetDir string, targetDate string) {
	files, err := filepath.Glob(targetDir + targetDate + "_*.png")
	if err != nil {
		panic(err)
	}

	targetList := []string{}
	for _, file := range files {
		basename := filepath.Base(file)
		targetList = append(targetList, basename)
	}

	sort.Strings(targetList)

	finished := make(chan bool)

	chanCount := 0
	for n, bf := range targetList {
		num := n
		beforeFilename := bf
		afterFilename := fmt.Sprintf("%s__%05d.png", targetDate, num+1)

		// 並列化のジレンマについて
		// 並列実行すると、処理順序の関係でファイルを上書きしてしまう可能性があるため
		// 書き換え先のファイル名はかぶらないようにする必要がある。
		// その場合、すべてのファイルをrenameするため、ずれてなければpassできた直列実行よりも遅くなる
		//if beforeFilename != afterFilename {
		go func() {
			fmt.Println(beforeFilename + " -> " + afterFilename)
			os.Rename(targetDir+beforeFilename, targetDir+afterFilename)
			finished <- true
		}()
		chanCount++
		//}
	}

	// 終わるまで待つ
	for i := 1; i <= chanCount; i++ {
		<-finished
	}
}

func main() {
	renumber(os.Args[1], os.Args[2])
}
