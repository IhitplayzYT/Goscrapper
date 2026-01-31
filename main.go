package main

import (
	"os"
	"sync"
)

var data Data
var wg sync.WaitGroup

// TODO:

func search(str string, flag i8) {
	for _, v := range website_list {
		parse_html(v, str)
	}
	wg.Wait()

}

func main() {
	args := os.Args
	keyword, flag := parse_args(args)
	if flag == 0 {
		return
	}
	search(keyword, flag)
}
