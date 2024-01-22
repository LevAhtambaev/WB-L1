package main

import "fmt"

// Menu целевой интерфейс
type Menu interface {
	GetItems() []string
}

// CafeMenu представляет адаптер, который преобразует интерфейс FoodMenu в интерфейс Menu.
type CafeMenu struct {
	FoodMenu *FoodMenu
}

// GetItems вызывает адаптированный метод GetFoodItems на объекте FoodMenu.
func (c *CafeMenu) GetItems() []string {
	return c.FoodMenu.GetFoodItems()
}

// FoodMenu представляет собой существующий класс, который несовместим с интерфейсом Menu.
type FoodMenu struct {
	Items []string
}

// GetFoodItems представляет собой специфический метод, который нужно адаптировать.
func (f *FoodMenu) GetFoodItems() []string {
	return f.Items
}

func main() {
	foodMenu := &FoodMenu{Items: []string{"Burger", "Pizza", "Salad", "Fries"}}
	adapter := &CafeMenu{FoodMenu: foodMenu}

	result := adapter.GetItems()
	fmt.Println(result)
}
