package main

import "fmt"

const tshirt = 43.99
const short = 50.00
const sneaker = 130.00
const hoodie = 199.99
const jeans = 149.99
const sandal = 19.99

func main() {
	fmt.Println("----------------SHOP----------------")
	fmt.Println("T-SHIRT 		R$ 43,99")
	fmt.Println("SHORT   		R$ 50,00")
	fmt.Println("SNEAKER 		R$ 130,00")
	fmt.Println("HOODIE 			R$ 199,99")
	fmt.Println("JEANS			R$ 149,99")
	fmt.Println("SANDAL			R$ 19,99")
	fmt.Println("----------------END-----------------")

	buy := func(item string) {

		if item == "tshirt" {
			fmt.Println("You purchased the t-shirt with R$ ", tshirt)
		}

		if item == "short" {
			fmt.Println("You purchased the short with R$ ", short)
		}

		if item == "sneaker" {
			fmt.Println("You purchased the sneaker with R$ ", sneaker)
		}

		if item == "hoodie" {
			fmt.Println("You purchased the hoodie with R$ ", hoodie)
		}

		if item == "jeans" {
			fmt.Println("You purchased the jeans with R$ ", jeans)
		}

		if item == "sandal" {
			fmt.Println("You purchased the sandal with R$ ", sandal)
		}

	}

	buy("jeans")

}
