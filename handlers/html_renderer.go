package handlers

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"golang.org/x/net/html"
)

func RenderHtml(htmlStr string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}

	replaceScriptTags(doc)

	// Render the modified HTML
	var sb strings.Builder
	if err := html.Render(&sb, doc); err != nil {
		fmt.Println("Error rendering HTML:", err)
		return ""
	}

	renderedHTML := sb.String()
	return renderedHTML
}

func replaceScriptTags(n *html.Node) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "side" && attr.Val == "server" {
				vm := goja.New()
				script := "(() => {" + n.FirstChild.Data + "})()"
				vm.Set("ssr", "hello world")
				result, err := vm.RunString(script)
				if err != nil {
					fmt.Println(err)
					break
				}
				n.FirstChild = &html.Node{
					Type: html.TextNode,
					Data: result.String(),
				}
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceScriptTags(c)
	}
}
