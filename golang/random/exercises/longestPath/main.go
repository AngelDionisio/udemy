package main

import (
	"fmt"
)

type ResourceRef struct {
	resourceType string
	resourceID   string
}

func main() {
	paths := [][]ResourceRef{
		{{resourceType: "COLLECTION", resourceID: "654"}},
		{{resourceType: "DEAL", resourceID: "777"}},
		{{resourceType: "CONTACT", resourceID: "6785"}},
		{{resourceType: "COLLECTION", resourceID: "654"}, {resourceType: "DEAL", resourceID: "777"}},
		{{resourceType: "CONTACT", resourceID: "6785"}, {resourceType: "COLLECTION", resourceID: "8888"}},
		{{resourceType: "CONTACT", resourceID: "6785"}, {resourceType: "COLLECTION", resourceID: "8888"}, {resourceType: "DEAL", resourceID: "777"}},
	}

	for k, v := range paths {
		fmt.Printf("key: %v, value: %s\n", k, v)
	}

	filteredPaths := selectLongestPaths(paths)

	fmt.Println(filteredPaths)
	fmt.Printf("Filtered paths: %+v", filteredPaths)
}

func selectLongestPaths(paths [][]ResourceRef) [][]ResourceRef {
	type key struct{ resourceType, resourceID string }
	longestPathsMap := make(map[key][]ResourceRef)

	for _, resourceList := range paths {
		k := key{
			resourceType: resourceList[0].resourceType,
			resourceID:   resourceList[0].resourceID,
		}
		// key seen for the first time, add it to map
		if _, found := longestPathsMap[k]; !found {
			longestPathsMap[k] = resourceList
			continue
		}
		// item seen again, replace it with longest path
		longestPathsMap[k] = resourceList
	}

	var filteredPaths [][]ResourceRef
	for _, pathList := range longestPathsMap {
		filteredPaths = append(filteredPaths, pathList)
	}
	return filteredPaths
}
