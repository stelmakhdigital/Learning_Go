// Теперь попробуйте сами. Сделайте map для хранения прейскуранта в рублях:

// Выведите перечень деликатесов — продуктов дороже 500 рублей.
// Заказ выдан слайсом []string{"хлеб", "буженина", "сыр", "огурцы"}. Посчитайте стоимость заказа.

package main

import "fmt"

func main() {
	foodPrice := make(map[string]int)
	foodPrice["хлеб"] = 50
	foodPrice["молоко"] = 100
	foodPrice["масло"] = 200
	foodPrice["колбаса"] = 500
	foodPrice["соль"] = 20
	foodPrice["огурцы"] = 200
	foodPrice["сыр"] = 600
	foodPrice["ветчина"] = 700
	foodPrice["буженина"] = 900
	foodPrice["помидоры"] = 250
	foodPrice["рыба"] = 300
	foodPrice["хамон"] = 1500

	// fmt.Println(foodPrice)

	for key, value := range foodPrice {
		if value > 500 {
			fmt.Println(key, value)
		}
	}

	sliceOrder := []string{"хлеб", "буженина", "сыр", "огурцы"}
	totalPrice := 0

	for _, productName := range sliceOrder {
		if foodPrice != nil {
			totalPrice += foodPrice[productName]
		}
	}

	fmt.Printf("Total price: %v", totalPrice)

}
