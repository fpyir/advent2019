package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func fuelFor(moduleMass float64) float64 {
	return math.Floor(moduleMass/3) - 2
}

func main() {
	modules, err := os.Open("01_1__input.txt")
	if err != nil {
		fmt.Println("Error reading file 01_1__input.txt")
		fmt.Println(err.Error())
	}
	scanner := bufio.NewScanner(modules)
	var total float64 = 0
	for scanner.Scan() {
		module, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fuel := fuelFor(module)
		for fuel > 0 {
			total += fuel
			fuel = fuelFor(fuel)
		}
	}
	formattedTotal := fmt.Sprintf("%.2f", total)
	fmt.Println("total fuel required: ", formattedTotal)
}
