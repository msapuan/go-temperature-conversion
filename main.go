package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("====================================")
	fmt.Println("**__** APLIKASI KONVERSI SUHU **__**")
	fmt.Println("====================================")

	var userContinue string = "y"
	for userContinue == "y" {
		fmt.Print("Input suhu dalam celcius: ")
		fmt.Scanln(&celcius)

		fmt.Println("1. Fahrenheit")
		fmt.Println("2. Reamur")
		fmt.Println("3. Kelvin")

		var choice int
		fmt.Print("Pilih satuan suhu yang diinginkan: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			// (c * 9/5) + 32
			result := (celcius * (9.0 / 5.0)) + 32.0
			fmt.Println("Konversi suhu dalam Fahrenheit : " + fmt.Sprint(result))
		} else if choice == 2 {
			// c * 4/5
			result := celcius * (4.0 / 5.0)
			fmt.Println("Konversi suhu dalam Reamur : " + fmt.Sprint(result))
		} else if choice == 3 {
			// c + 273.15
			result := celcius + 273.15
			fmt.Println("Konversi suhu dalam Kelvin : " + fmt.Sprint(result))
		}

		fmt.Println("")

		fmt.Print("Apakah anda ingin melanjutkan konversi suhu ? (y/n) : ")
		fmt.Scanln(&userContinue)

	}

}
