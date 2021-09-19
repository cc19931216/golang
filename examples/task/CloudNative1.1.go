package main

import "fmt"

func main() {
	strA := [5]string{"i", "am", "stupid", "and", "weak"}
	fmt.Println(strA)

	for k, _ := range strA {
		switch k {
		case 2:
			strA[k] = "smart"
		case 4:
			strA[k] = "strong"
		default:
		}
	}
	fmt.Println(strA)
}
