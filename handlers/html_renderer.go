package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func RenderHtml(html string) string {
	result, last := RenderHtml_frag(html)
	if last {
		return result
	}
	return RenderHtml(result)
}

func RenderHtml_frag(html string) (string, bool) {
	s := strings.Index(html, "<{")
	if s == -1 {
		fmt.Println("string start error")
		return html, true
	}
	s += 2
	e := strings.Index(html, "}>")
	if e == -1 {
		fmt.Println("string end error")
		return html, true
	}
	url := html[s:e]
	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		fmt.Println("request error")
		return html[:s-2] + "request error" + html[e+2:], false
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("body parse error")
		return html[:s-2] + html[e+2:], false
	}
	return html[:s-2] + string(result) + html[e+2:], false
}
