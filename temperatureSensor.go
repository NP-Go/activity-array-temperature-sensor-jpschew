package main

import "fmt"

func main() {
	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}
	//Insert Code here
	fmt.Println(len(temperature))
	var total, average float64

	for i := 0; i < len(temperature); i++ {
		total += temperature[i]
	}
	average = total / float64(len(temperature))

	fmt.Printf("%.2f %.2f\n", total, average)

}
