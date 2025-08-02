package htmlparser

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Header 1</h1>
    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur aliquid hic nesciunt quam et quibusdam facilis deserunt? Inventore, optio voluptate?</p>
    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi odio repellat accusamus similique nihil voluptatibus?</p>
    <img src="xxx.jpg" width="104" height="142"/>
</body>
</html>
`



func visitNode(node *html.Node, wordsp, imagesp *int){

	if node.Type == html.TextNode {
		*wordsp+=len(strings.Fields(node.Data))

	} else if node.Type == html.ElementNode && node.Data == "img"{
		*imagesp++
	}

	for c:= node.FirstChild; c != nil; c=c.NextSibling{
		visitNode(c, wordsp, imagesp)

	}
}

func countWordsNImages(doc *html.Node)(int, int){
	var words, images int 

	visitNode(doc, &words, &images)
	return words, images

}
func HTMLParser() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured: %s", err)
		os.Exit(-1)
	}

	words, images := countWordsNImages(doc)

	fmt.Printf("%d words and %d images were found in the raw HTML", words, images)

	

}