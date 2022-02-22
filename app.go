package main

import (
	"fmt"

	"github.com/indranandjha1993/nummanip/calc"
	calcNew "github.com/indranandjha1993/nummanip/v2/calc"
)

func main() {
	result := calc.Add(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum = : ", result)

	err, result1 := calcNew.Add(4, 5)

	if err != nil {
		fmt.Println("Error ::", err)
	}
	fmt.Println("New Sum = : ", result1)

}
