package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func calcSquare(x int, y int) int {
	result := x * y
	return result
}

func calcRect(width, height int) (int, int) {
	perimeter := (width + height) * 2
	square := width * height
	return perimeter, square
}

func calcCubePerimeter(side int) int {
	return side * 12
}

func calcCubeArea(side int) int {
	oneFace := side * side
	return oneFace * 6
}

func calcCube(side int) (fullLength int, fullArea int) {
	oneCubePerimeter := calcCubePerimeter(side)
	fullLength = oneCubePerimeter * 8
	oneCubeArea := calcCubeArea(side)
	fullArea = oneCubeArea * 8
	return
}

func main() {
	fmt.Println("Я домашка")

	square := calcSquare(16, 9)
	fmt.Println(square)

	perimeter, square := calcRect(4, 9)
	fmt.Println("Периметр:", perimeter, "Площадь:", square)

	fullLength, fullArea := calcCube(3)
	fmt.Println("Для 8 кубов понадобится палок (м):", fullLength,
		"и стекла (кв. м):", fullArea)
}
