package main

import (
	"os"
)

var data Data

func search(flag i8) {
	for _, v := range website_list {
		parse_html(v)

	}
}

func main() {
	KEYWORD, flag := parse_args(os.Args)
	init_wlist()
	boost_wlist()
	if flag == 0 {
		return
	}

	search(flag)
}
