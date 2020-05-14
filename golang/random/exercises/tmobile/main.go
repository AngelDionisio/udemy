package main

import (
	"fmt"
)

func addListOfFloats(n []float32) float32 {
	var sum float32
	for _, v := range n {
		sum += float32(v)
	}
	return sum
}

func computeTmobileBill(data map[string]float32, people []string) map[string]float32 {
	angelExtra := data["angelExtra"]
	julioExtra := data["julioExtra"]
	joseExtra := data["joseExtra"]
	totalExtras := addListOfFloats([]float32{
		angelExtra,
		julioExtra,
		joseExtra,
	})

	var totalPerPerson float32
	totalPerPerson = (data["total"] - totalExtras) / float32(len(people))

	var evenSplit float32
	if data["extra"] != 0 {
		evenSplit = data["extra"] / 3
	}

	angelTotal := addListOfFloats([]float32{totalPerPerson * 2, angelExtra, evenSplit})
	julioTotal := addListOfFloats([]float32{totalPerPerson * 2, julioExtra, evenSplit})
	joseTotal := addListOfFloats([]float32{totalPerPerson * 2, joseExtra, evenSplit})
	totalBillAmount := addListOfFloats([]float32{angelTotal, julioTotal, joseTotal})

	m := map[string]float32{
		"totalPerPerson":  totalPerPerson,
		"angelTotal":      angelTotal,
		"julioTotal":      julioTotal,
		"joseTotal":       joseTotal,
		"totalBillAmount": totalBillAmount,
	}

	return m
}

func main() {
	peopleInPlan := []string{"Angel", "Luz", "Julio", "Artemis", "Jose", "Mami"}

	data := map[string]float32{
		"total":      202.00,
		"extra":      0.0, // any extra charges to be evenly divided
		"angelExtra": 1.0, // netflix extra charge
		"julioExtra": 1.0, // netflix extra charge
		"joseExtra":  0.0,
	}

	results := computeTmobileBill(data, peopleInPlan)

	fmt.Println(results)
	fmt.Printf("%#v\n", results)

	for k, v := range results {
		fmt.Println(k, ":", v)
	}
}
