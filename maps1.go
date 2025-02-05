package main

import "fmt"

func main() {

	//1 как добавить элемент в map
	money := map[string]int{
		"dollars": 1000,
		"euros": 1200,
		"som": 100,
	}
	money["yen"] = 500
	fmt.Println(money)

	//2 как удалить элемент из map не зная его ключ
	var kay string
	fmt.Print("Enter key:")
	fmt.Scan(&kay)

	delete(money, kay)
	fmt.Println(money)

	//3 как перебрать все элементы map
	for key, value := range money{
		fmt.Println("Key:", key,"| Value:", value)
	}
	//4 Как проверить, существует ли ключ в map, не изменяя значение переменной?
	var key1 string
	fmt.Print("Enter key:")
	fmt.Scan(&key1)

	el, status := money[key1] 
	fmt.Println(el, status)
	if status == false{
		fmt.Println("ключ не найден")
	}else{
		fmt.Println("ключ найден")
	}
	
	// 5 Как создать пустую map и добавить в нее несколько значений?
	ingresients := make(map[string]int)
	var ingresient1 string
	var ingresient2 string
	var ingresient3 string

	fmt.Print("Enter ingresient1:")
	fmt.Scan(&ingresient1)
	fmt.Print("Enter ingresient2:")
	fmt.Scan(&ingresient2)
	fmt.Print("Enter ingresient3:")
	fmt.Scan(&ingresient3)

	var count1 int
	var count2 int
	var count3 int

	fmt.Print("Enter number:")
	fmt.Scan(&count1)
	fmt.Print("Enter number:")
	fmt.Scan(&count2)
	fmt.Print("Enter number:")
	fmt.Scan(&count3)

	ingresients[ingresient1] = count1
	ingresients[ingresient2] = count2
	ingresients[ingresient3] = count3
	fmt.Println(ingresients)	
}