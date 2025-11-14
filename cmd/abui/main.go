package main

import (
	"fmt"

	u "github.com/alehoppai/abui-go/pkg/abui"
)

func main() {
	tree := u.Flex(u.Flex())
	html := tree.Render()
	fmt.Printf("%s\n", html)
}
