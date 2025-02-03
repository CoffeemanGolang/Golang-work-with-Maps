package main

import "fmt"

func main() {
	money := map[string]int{
		"dollars": 1000,
		"euros": 1200,
		"apples": 2,
	}
	money["dollars"]=5000

	delete(money, "apples")
	fmt.Println(money)

	//если ключа нет, то вернется 0
	fmt.Println(money["aaa"])
	//проверка на наличие ключа
	el, status := money["dollars"]
	el, status =5, true 
	fmt.Println(el, status)
}