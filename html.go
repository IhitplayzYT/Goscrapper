package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var link_map sync.Map
var visited map[string]bool = make(map[string]bool)
var mutex sync.Mutex

type Media_Type int

const (
	Invalid_T Media_Type = iota
	Image_T
	Video_T
	Audio_T
	Files_T
	Misc_T
)

// Returns the entire html for a link
func get_html(link string) string {
	response, err := http.Get(link)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	retu, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	return string(retu)
}

func parse_html(link string) {
	mutex.Lock()
	if visited[link] {
		mutex.Unlock()
		return
	}
	visited[link] = true
	mutex.Unlock()
	html := get_html(link)
	if len(html) == 0 {
		return
	}
	go parse_html_helper(link, html)
}

func parse_html_helper(link string, html string) {
	if len(html) == 0 {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		base, _ := url.Parse(link)
		ref, err := url.Parse(href)
		if err != nil {
			return
		}
		urlpath := base.ResolveReference(ref).String()
		for k, _ := range assetExt {
			if strings.Contains(href, k) {
				append_to_lmap(link, urlpath)
			}
		}
		parse_html(urlpath)
	})

	doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		append_to_lmap(link, src)
		install(src, Image_T)
	})

	doc.Find("video source").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("src"); ok {
			append_to_lmap(link, src)
			install(src, Video_T)
		}
	})
	doc.Find("audio source").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("src"); ok {
			append_to_lmap(link, src)
			install(src, Audio_T)
		}
	})

}

func append_to_lmap(key string, val string) {

	old, _ := link_map.Load(key)
	var slice []string
	if old != nil {
		slice = old.([]string)
	}
	newslice := append(slice, val)
	if link_map.CompareAndSwap(key, old, newslice) {
		return
	}

}

func install(src string, M_type Media_Type) {
	home, _ := os.UserHomeDir()
	dwnld_dir := filepath.Join(home, "Goscrapper", "output")
	switch M_type {

	case Image_T:
		{
			dwnld_dir += "images"
			break
		}
	case Audio_T:
		{
			dwnld_dir += "Audio"
			break
		}
	case Video_T:
		{
			dwnld_dir += "videos"
			break
		}
	case Files_T:
		{
			dwnld_dir += "files"
			break
		}
	default:
		{
			dwnld_dir += "misc"
			return
		}

	}

	bin := filepath.Join(home, "Godownloader", "main")
	cmd := exec.Command(bin, src, dwnld_dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Download Failed")
	}

}
