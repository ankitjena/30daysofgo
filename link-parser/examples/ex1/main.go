package main

import (
	"strings"
	"github.com/ankitjena/30daysofgo/link-parser"
	"fmt"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
	<a href="/other-page">A link to
	 another page</a>
	<a href="/new-page">New link</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}