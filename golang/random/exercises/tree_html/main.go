package main

import (
	"fmt"
	"strings"
)

// Node models an HTML element
type Node struct {
	tag      string
	id       string
	class    string
	text     string
	src      string
	alt      string
	children []*Node
}

const strTemplate = `
type: %v,
tag: %v,
ID: %v,
hasChildren: %v
`

// PrintNode logs representation of tree
func (n *Node) PrintNode() {
	fmt.Printf(strTemplate, "node", n.tag, n.id, n.HasChildren())
}

// HasChildren returns bool if node has children
func (n *Node) HasChildren() bool {
	if len(n.children) > 0 {
		return true
	}
	return false
}

// HasClass returns bool if node has a class
func (n *Node) HasClass(className string) bool {
	classes := strings.Fields(n.class)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

// printChildren logs all children for a given node
func (n *Node) printChildren() {
	if n.HasChildren() == false {
		fmt.Printf("%v has no children\n", n)
	}

	for _, v := range n.children {
		fmt.Printf(strTemplate, "child", v.tag, v.id, v.HasChildren())

	}
}

/*
<html>
  <h1>Hello, World!</h1>
  <p>
	  This is a simple HTML document.
	  <img src="http://example.com/logo.svg" alt="Example's Logo"/>
  </p>
</html>
*/

/*
<html>
  <body>
    <h1>This is a H1</h1>
    <p>
      And this is some text in a paragraph. And next to it there's an image.
      <img src="http://example.com/logo.svg" alt="Example's Logo"/>
    </p>
    <div class='footer'>
      This is the footer of the page.
      <span id='copyright'>2019 &copy; Ilija Eftimov</span>
    </div>
  </body>
</html>
*/
func main() {
	// image := Node{
	// 	tag: "img",
	// 	src: "http://example.com/logo.svg",
	// 	alt: "Example's Logo",
	// 	id:  "data-qaid=img-bx",
	// }

	// p := Node{
	// 	tag:      "p",
	// 	text:     "This is a simple HTML document.",
	// 	id:       "123",
	// 	children: []*Node{&image},
	// }

	// h1 := Node{
	// 	tag:  "h1",
	// 	text: "Hello, World!",
	// 	id:   "456",
	// }

	// html := Node{
	// 	tag:      "html",
	// 	children: []*Node{&p, &h1},
	// }

	image := Node{
		id:  "img_id_1",
		tag: "img",
		src: "http://example.com/logo.svg",
		alt: "Example's Logo",
	}

	p := Node{
		id:       "p_id_1",
		tag:      "p",
		text:     "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{&image},
	}

	span := Node{
		id:   "copyright",
		tag:  "span",
		text: "2019 &copy; Ilija Eftimov",
	}

	div := Node{
		id:       "div_id_1",
		tag:      "div",
		class:    "footer",
		text:     "This is the footer of the page.",
		children: []*Node{&span},
	}

	h1 := Node{
		id:   "h1_id_1",
		tag:  "h1",
		text: "This is a H1",
	}

	body := Node{
		id:       "body_id_1",
		tag:      "body",
		children: []*Node{&h1, &p, &div},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}

	// html.PrintNode()
	// body.printChildren()

	dfsResult := findByIDBFS(&html, "img_id_1")
	fmt.Println("BDF search result:", dfsResult)

	bfsResult := findByIDBFS(&html, "img_id_1")
	fmt.Println("BFS search result:", bfsResult)
}

// findByIDBFS BFS (breath first search) on type *Node
func findByIDBFS(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		// get zero idx item in queue list.
		current := queue[0]
		// remove idx = 0 from list by setting list to contain
		// all items in the list starting (inclusive) index 1.
		queue = queue[1:]
		if current.id == id {
			return current
		}
		idx := 0
		if len(current.children) > 0 {
			for _, child := range current.children {
				idx++
				queue = append(queue, child)
			}

		}
	}
	return nil
}

// findByIDDFS depth first tree lookup implementation
func findByIDDFS(node *Node, id string) *Node {
	if node.id == id {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIDDFS(child, id)
		}
	}
	return nil
}

func findAllByClassName(root *Node, className string) []*Node {
	result := make([]*Node, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.HasClass(className) {
			result = append(result, current)
		}
		if len(current.children) > 0 {
			for _, child := range current.children {
				queue = append(queue, child)
			}
		}
	}
	return result
}
