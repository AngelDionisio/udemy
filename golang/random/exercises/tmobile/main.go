package main

import (
	"fmt"
)

func main() {
	peopleInPlan := []string{"Angel", "Luz", "Julio", "Artemis", "Jose", "Mami"}

	data := map[string]float64{
		"total":      202.00,
		"extra":      0.0,
		"angelExtra": 1.0, // netflix extra charge
		"julioExtra": 1.0, // netflix extra charge
		"joseExtra":  0.0,
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
	angelExtra := data["angelExtra"]
	julioExtra := data["julioExtra"]
	joseExtra := data["joseExtra"]
	totalExtras := addFloatList([]float64{
		angelExtra,
		julioExtra,
		joseExtra,
		data["extra"],
	})

	var totalPerPerson float64
	l := float64(len(people))
	totalPerPerson = (data["total"] - totalExtras) / l

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
