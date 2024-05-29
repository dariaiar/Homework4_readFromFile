package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myFile, err := os.Open("Shoppinglist.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myFile.Close()

	var shoppingList []string

	scanner := bufio.NewScanner(myFile)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		shoppingList = append(shoppingList, line)
	}
	//shoppingList[i] = line

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("Your Shopping list slice is:  ", shoppingList)

	for b, rangedShoppingList := range shoppingList {
		fmt.Println(b, rangedShoppingList)
	}

	var findWord string
	fmt.Printf("To search, type name of the product:   ")
	fmt.Scan(&findWord)

	for n := 0; n < 10; n++ {
		if findWord == shoppingList[n] {
			fmt.Printf("FOUND MATCH -- product #%v,  %v   \n", n, shoppingList[n])
		} else {
			fmt.Println("MATCH was not found")
		}
	}
}
