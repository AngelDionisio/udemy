package main

import (
	"fmt"
	"strings"
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
	_ = paths

	paths2 := [][]ResourceRef{
		{{resourceType: "TOURSHEET", resourceID: "tour_abc"}},
		{{resourceType: "TOURSHEET", resourceID: "tour_125"}},
		{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_125"}},
		{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_130"}},
		{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_130"}, {resourceType: "DEAL", resourceID: "deal_130"}},
	}
	_ = paths2

	/*
		{"(TOURSHEET,tour125)"}	1
		{"(TOURSHEET,tour_125)"}	1
		{"(TOURSHEET,tour_125)","(TASK,task_125)"}	2
		{"(TOURSHEET,tour_125)","(TASK,task_130)"}	2
		{"(TOURSHEET,tour_125)","(TASK,task_130)","(DEAL,deal_130)"}	3
	*/

	// for k, v := range paths {
	// 	fmt.Printf("key: %v, value: %s\n", k, v)
	// }

	// filteredPaths := selectLongestPaths(paths)

	// fmt.Println(filteredPaths)
	// fmt.Printf("Filtered paths: %+v", filteredPaths)

	result := SelectLongestPathExtended(paths2)
	fmt.Println(result)
	result2 := SelectLongestPathExtended_OLD(paths2)
	fmt.Println(result2)

	result3 := SelectLongestPathExtended(paths)
	fmt.Println(result3)
	result4 := SelectLongestPathExtended_OLD(paths)
	fmt.Println(result4)
}

func makeKey(path []ResourceRef) string {
	var sb strings.Builder
	if len(path) == 1 {
		sb.WriteString(path[0].resourceType)
		sb.WriteString((path[0].resourceID))
		return sb.String()
	}

	for i := 0; i < len(path)-1; i++ {
		sb.WriteString(path[i].resourceType)
		sb.WriteString(path[i].resourceID)
	}

	return sb.String()
}

func makeLongerPathKey(path []ResourceRef) string {
	var sb strings.Builder
	for i := 0; i < len(path); i++ {
		sb.WriteString(path[i].resourceType)
		sb.WriteString(path[i].resourceID)
	}

	return sb.String()
}

/* Iterate through list, depth 1 paths are added to map by default
when longer paths are found, if they contain the previous path, replace it with longer path
e.g.
{{resourceType: "TOURSHEET", resourceID: "tour_abc"}},
{{resourceType: "TOURSHEET", resourceID: "tour_125"}},
{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_125"}},
{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_130"}},
{{resourceType: "TOURSHEET", resourceID: "tour_125"}, {resourceType: "TASK", resourceID: "task_130"}, {resourceType: "DEAL", resourceID: "deal_130"}}

resulting in:
[[{TOURSHEET, tour_abc}], [{TOURSHEET, tour_125} {TASK task_125}], [{TOURSHEET, tour_125} {TASK, task_130} {DEAL, deal_130}]]

*/
func SelectLongestPathExtended(paths [][]ResourceRef) [][]ResourceRef {
	longestPathsMap := make(map[string][]ResourceRef)
	visitedMap := make(map[string]bool)

	for _, path := range paths {
		k := makeKey(path)
		if _, found := longestPathsMap[k]; !found {
			if _, visited := visitedMap[k]; !visited {
				longestPathsMap[k] = path
				visitedMap[k] = true
				continue
			}
			newKey := makeLongerPathKey(path)
			longestPathsMap[newKey] = path
			visitedMap[newKey] = true
			continue
		}

		newKey := makeLongerPathKey(path)
		delete(longestPathsMap, k)
		longestPathsMap[newKey] = path

	}

	var filteredPaths [][]ResourceRef
	for _, pathList := range longestPathsMap {
		filteredPaths = append(filteredPaths, pathList)
	}
	return filteredPaths
}

func SelectLongestPathExtended_OLD(paths [][]ResourceRef) [][]ResourceRef {
	longestPathsMap := make(map[string][]ResourceRef)

	for _, path := range paths {
		k := makeKey(path)
		// fmt.Printf("path: %v, key: %v, len(path): %v\n", path, k, len(path))
		_, found := longestPathsMap[k]
		if !found {
			// fmt.Println("FIRST TIME PATH ADD:", path)
			longestPathsMap[k] = path
		}

		if found {
			newKey := makeLongerPathKey(path)
			// fmt.Printf("FOUND LONGER PATH: deleting old key: %v, adding new key: %v\n", k, newKey)
			delete(longestPathsMap, k)
			longestPathsMap[newKey] = path
		}

	}

	var filteredPaths [][]ResourceRef
	for _, pathList := range longestPathsMap {
		filteredPaths = append(filteredPaths, pathList)
	}
	return filteredPaths
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
