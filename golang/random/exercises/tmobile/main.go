package main

import (
	"fmt"
)

func main() {
	peopleInPlan := []string{"Angel", "Luz", "Julio", "Artemis", "Jose", "Mami"}

	data := map[string]float64{
		"total":             208.00,
		"extra":             0.0,
		"angelNetflixExtra": 1.0,
		"julioNetflixExtra": 1.0,
		"angelExtra":        0.0,
		"julioExtra":        0.0,
		"joseExtra":         0.0,
	}

	results := computeTmobileBill(data, peopleInPlan)

	fmt.Println(results)

	for k, v := range results {
		fmt.Println(k, ":", v)
	}
}

func addFloatList(n []float64) float64 {
	var sum float64
	for _, v := range n {
		sum += float64(v)
	}
	return sum
}

func computeTmobileBill(data map[string]float64, people []string) map[string]float64 {
	angelExtraList := []float64{data["angelExtra"], data["angelNetflixExtra"]}
	julioExtraList := []float64{data["julioExtra"], data["julioNetflixExtra"]}
	joseExtraList := []float64{data["joseExtra"]}

	angelExtra := addFloatList(angelExtraList)
	julioExtra := addFloatList(julioExtraList)
	joseExtra := addFloatList(joseExtraList)

	var totalPerPerson float64
	l := float64(len(people))
	totalPerPerson = (data["total"] - addFloatList([]float64{angelExtra, julioExtra, joseExtra, data["extra"]})) / l

	angelTotal := addFloatList([]float64{totalPerPerson * 2, angelExtra})
	julioTotal := addFloatList([]float64{totalPerPerson * 2, julioExtra})
	joseTotal := addFloatList([]float64{totalPerPerson * 2, joseExtra})
	totalBillAmount := addFloatList([]float64{angelTotal, julioTotal, joseTotal})

	m := map[string]float64{
		"totalPerPerson":  totalPerPerson,
		"angelTotal":      angelTotal,
		"julioTotal":      julioTotal,
		"joseTotal":       joseTotal,
		"totalBillAmount": totalBillAmount,
	}

	return m
}
