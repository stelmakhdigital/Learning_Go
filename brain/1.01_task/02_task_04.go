//- найдите радиус окружности, если её длина равна 35.
//- радиус окружности radiusCycle объявите как указатель на `float64`.
//- вычислите площадь круга, используя при расчёте разыменовывание указателя radiusCycle.
//- При необходимости дробные значения округлите до двух знаков после запятой.

package main

import (
	"fmt"
	"math"
)

func main() {
	length := 35
	var radius float64 = float64(length) / (2 * math.Pi)

	var radiusCycle *float64

	radiusCycle = &radius

	areaCycle := math.Pi * (*radiusCycle * *radiusCycle)

	fmt.Printf("Радиус окружности: %.2f\n", *radiusCycle)
	fmt.Printf("Площадь круга: %.2f\n", areaCycle)
}