/*
### 2. Калькулятор расходов
Описание: Создайте программу для учета личных расходов.
Функции:
- Добавить расход (название, сумма, категория).
- Вывести общую сумму расходов.
- Вывести расходы по категориям.
- Найти самый большой расход.
*/

package main

import "fmt"

var expenses = map[string]float64{}
var categories = map[string]float64{}

func main() {
	expenses = addConsumption("Transport", 2.50)
	expenses = addConsumption("Food", 30)
	expenses = addConsumption("Internet", 55)
	expenses = addConsumption("Transport", 2.50)
	expenses = addConsumption("Food", 20)
	expenses = addConsumption("Book", 42)
	expenses = addConsumption("Medicine", 150)

	fmt.Println(expenses)
	fmt.Println()

	total_value := totalValue(expenses)
	fmt.Println("Total value:", total_value)

	biggest_expense := biggestExpense(expenses)
	fmt.Println("The biggest expense:", biggest_expense["biggest_expense"])

	all_categories := allCategories(expenses)
	fmt.Println(all_categories)
	fmt.Println()

}

func addConsumption(category string, amount float64) map[string]float64 {
	if _, exists := expenses[category]; exists {
		expenses[category] += amount
	} else {
		expenses[category] = amount
	}
	return expenses
}

func totalValue(expenses map[string]float64) float64 {
	var sum float64
	for _, value := range expenses {
		sum += value
	}
	return sum
}

func allCategories(expenses map[string]float64) map[string]int {
	all_categories := make(map[string]int)
	counter := 0

	for key := range expenses {
		if expenses[key] != 0 {
			counter++
			all_categories[key] = counter
		}

		counter = 0
	}
	return all_categories
}

func biggestExpense(expenses map[string]float64) map[string]float64 {
	biggest_expense := map[string]float64{
		"biggest_expense": 0.0,
	}

	for key, value := range expenses {
		if value > biggest_expense["biggest_expense"] {
			biggest_expense["biggest_expense"] = value
		} else if value == biggest_expense["biggest_expense"] {
			biggest_expense[key] = value
		}

	}
	return biggest_expense

}
