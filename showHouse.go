package main

import "fmt"

func showHouse() {
	furniture := []Furniture{
		{Name: "Диван", Length: 200, Width: 100, Height: 80},
		{Name: "Кровать", Length: 180, Width: 90, Height: 60},
		{Name: "Стол", Length: 120, Width: 80, Height: 75},
		{Name: "Стул", Length: 40, Width: 40, Height: 90},
		{Name: "Полка", Length: 100, Width: 30, Height: 150},
	}

	appliances := []Appliance{
		{Name: "Микроволновая печь", Length: 50, Width: 40, Height: 30},
		{Name: "Плита", Length: 60, Width: 60, Height: 85},
		{Name: "Стиральная машина", Length: 60, Width: 60, Height: 85},
		{Name: "Телевизор", Length: 100, Width: 10, Height: 60},
		{Name: "Фен", Length: 20, Width: 10, Height: 15},
	}

	familyMembers := []FamilyMember{
		{Name: "Мама", Age: 40, Height: 160, Weight: 60},
		{Name: "Папа", Age: 45, Height: 180, Weight: 80},
		{Name: "Сестра", Age: 20, Height: 165, Weight: 55},
		{Name: "Бабушка", Age: 70, Height: 155, Weight: 70},
		{Name: "Брат", Age: 25, Height: 175, Weight: 75},
	}

	fmt.Println("Мебель:")
	for _, item := range furniture {
		fmt.Printf("%s: Длина - %.2f, Ширина - %.2f, Высота - %.2f\n", item.Name, item.Length, item.Width, item.Height)
	}

	fmt.Println("\nБытовая техника:")
	for _, item := range appliances {
		fmt.Printf("%s: Длина - %.2f, Ширина - %.2f, Высота - %.2f\n", item.Name, item.Length, item.Width, item.Height)
	}

	fmt.Println("\nЧлены семьи:")
	for _, member := range familyMembers {
		fmt.Printf("%s: Возраст - %d, Рост - %.2f, Вес - %.2f\n", member.Name, member.Age, member.Height, member.Weight)
	}
}
