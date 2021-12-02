package main

import "fmt"

type Node struct {
	id       string
	tag      string
	text     string
	class    string
	src      string
	alt      string
	children []*Node
}

func main() {
	html := buildHTML()

	// res := searchByIDBFS(&html, "imageID")
	// fmt.Println("result:", res)
	fmt.Println("findByIdDFS: ", findByIdDFS(&html, "imageID"))
}

// start at the root
// create queue, add root to queue
// take first item in queue, set the queue to the rest
// check for ID equality, if match, return
// if not, check if current Node has children, if so, add them to the queue
func searchByIDBFS(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // take all items from the list starting at index 1, i.e next item on queue
		fmt.Printf("Currently looking at nodeID: %s\n", current.id)
		if current.id == id {
			fmt.Printf("result found, asked for %s, current is %s\n", id, current.id)
			return current
		}
		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return nil
}

func findByIdDFS(currentNode *Node, id string) *Node {
	if currentNode.id == id {
		fmt.Println("match found!", currentNode)
		return currentNode
	}

	if len(currentNode.children) > 0 {
		for _, child := range currentNode.children {
			fmt.Println("adding childID:", child.id)
			return findByIdDFS(child, id)
		}
	}
	return nil
}

func buildHTML() Node {
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
	image := Node{
		id:  "imageID",
		tag: "img",
		src: "http://example.com/logo.svg",
		alt: "Example's Logo",
	}

	p := Node{
		id:       "pID",
		tag:      "p",
		text:     "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{&image},
	}

	span := Node{
		id:   "spanID",
		tag:  "span",
		text: "2019 &copy; Ilija Eftimov",
	}

	div := Node{
		id:       "divID",
		tag:      "div",
		class:    "footer",
		text:     "This is the footer of the page.",
		children: []*Node{&span},
	}

	h1 := Node{
		id:   "h1ID",
		tag:  "h1",
		text: "This is a H1",
	}

	body := Node{
		id:       "bodyID",
		tag:      "body",
		children: []*Node{&h1, &p, &div},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}

	return html

}
