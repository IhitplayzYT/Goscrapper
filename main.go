package main

import (
	"os"
)

var data Data

// TODO:

func search(str string, flag i8) {
	for _, v := range website_list {
		parse_html(v)
	}

}

func main() {
	args := os.Args
	keyword, flag := parse_args(args)
	if keyword == "" || flag == 0 {
		return
	}
	search(keyword, flag)
	return
}
