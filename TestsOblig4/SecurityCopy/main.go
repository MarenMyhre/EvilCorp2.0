package main


import (
	"fmt"
	"../Map"

)


func main() {
	hapi := Map.NewV1("ee17c954-a67a-43bf-afb4-9c54ac05ab2b") // API key fra Holidays

	holidays, err := hapi.Holidays(map[string]interface{}{
		// Required
		"country": "NO",
		"year":    "2017",
		// Optional
		// "month":    "7",
		// "day":      "4",
		// "previous": "true",
		// "upcoming": "true",
		// "public":   "true",
		// "pretty":   "true",
	})

	if err != nil {
		// Error handling...
	}

	fmt.Println("%#v\n", holidays)
}