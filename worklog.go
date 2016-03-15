package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

var now = time.Now()
var counter = 1

func main() {

	tick := time.Tick(10 * time.Second)
	for _ = range tick {
		capture()
	}

}

func capture() {

	if now.Day() != time.Now().Day() {
		counter = 1
		now = time.Now()
	} else {
		counter = getCounter(counter)
	}

	filename := fmt.Sprintf("%d%02d%02d_%05d.png", now.Year(), now.Month(), now.Day(), counter)

	cmd := "/usr/sbin/screencapture"
	cmd_option := "/Users/tsuruoka-n/screencapture/" + filename
	// fmt.Println(cmd + cmd_option) // 表示

	_, err := exec.Command(cmd, "-x", "-C", cmd_option).Output()
	if err != nil {
		log.Fatal(err)
	}

}

func getCounter(r_counter int) int {

	if r_counter == 1 {
		// すでにファイルがあるか調べてあるなら一番大きい値を使う
		filename := fmt.Sprintf("%d%02d%02d_*", now.Year(), now.Month(), now.Day())
		items, err := filepath.Glob("/Users/tsuruoka-n/screencapture/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		// アイテムないなら最初から
		if len(items) == 0 {
			return r_counter
		}
		sort.Strings(items)
		string_suffix := strings.Split(items[len(items)-1], "_")[1]
		latest_number, _ := strconv.Atoi(strings.Split(string_suffix, ".")[0])

		r_counter = latest_number + 1
	} else {
		r_counter++
	}

	return r_counter
}
