package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{"beer": "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}
	sdrinks := make([]string, len(drinks))
	ix := 0
	for eng := range drinks {
		sdrinks[ix] = eng
		ix++
		fmt.Println(eng)
	}

	for eng, fr := range drinks {
		fmt.Printf("the french of %s is %s\n", eng, fr)
	}

	sort.Strings(sdrinks)

	for _, eng := range sdrinks {
		fmt.Println(eng)
	}

	for _, eng := range sdrinks {
		fmt.Printf("the french of %s is %s\n", eng, drinks[eng])
	}
}
