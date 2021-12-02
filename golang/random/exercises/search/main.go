package main

import (
	"fmt"
	"sort"
)

/*
[
"2 BR | 2 BA | 777 Brockton Avenue, Abington MA 23501",
"3 BR | 1.5 BA | 30 Memorial Drive, Avon MA 23922",
"1 BR | 1 BA | 250 Hartford Avenue, Bellingham MA 20019",
"4 BR | 1 BA | 700 Oak Street, Brockton MA 23601",
"2 BR | 1 BA | 66-4 Parkhurst Rd, Chelmsford MA 18124",
"4 BR | 2 BA | 591 Memorial Dr, Chicopee MA 10240",
"1 BR | 1 BA | 121 Worcester Rd, Framingham MA 17001",
"2 BR | 2 BA | 301 Massachusetts Ave, Lunenburg MA 19462",
"3 BR | 2 BA | 501 Cambridge Ave, Lunenburg MA 19462"
]
*/

type Property struct {
	Address   string
	Bedrooms  float32
	Bathrooms float32
}

func printProperties(properties []Property) {
	for _, property := range properties {
		fmt.Printf("%#v\n", property)
	}
}

func makeProperty(address string, bedrooms, bathrooms float32) Property {
	return Property{
		Address:   address,
		Bedrooms:  bedrooms,
		Bathrooms: bathrooms,
	}
}

func main() {
	properties := []Property{
		makeProperty("777 Brockton Avenue, Abington MA 23501", 2, 2),
		makeProperty("30 Memorial Drive, Avon MA 23922", 3, 1.5),
		makeProperty("250 Hartford Avenue, Bellingham MA 20019", 1, 1),
		makeProperty("700 Oak Street, Brockton MA 23601", 4, 1),
		makeProperty("66-4 Parkhurst Rd, Chelmsford MA 18124", 2, 1),
		makeProperty("591 Memorial Dr, Chicopee MA 10240", 4, 2),
		makeProperty("121 Worcester Rd, Framingham MA 17001", 1, 1),
		makeProperty("301 Massachusetts Ave, Lunenburg MA 19462", 2, 2),
		makeProperty("501 Cambridge Ave, Lunenburg MA 19462", 3, 2),
	}

	// Sort by Bedrooms
	sort.SliceStable(properties, func(i, j int) bool {
		return properties[i].Bedrooms < properties[j].Bedrooms
	})

	printProperties(properties)

}
