package main

import "fmt"

// Задача 1: Создание и использование структуры
// Создайте структуру Student, которая будет содержать следующие поля:

var first = Student{Name: "Abraham Lincoln", Age: 56, Grade: 4}

// Создайте экземпляр структуры Rectangle и вызовите оба метода, чтобы проверить их работу

var rect = Rectangle{Height: 3, Width: 5}

// Создайте экземпляр структуры Person и выведите полную информацию о человеке, включая адрес

var user1 = Person{Name: "John Watson", Age: 28, HomeAddress: Address{City: "London", Street: "Baker Street", HouseNumber: "221B "}}

var point_one = Point{X: 4.2, Y: 6.9}
var point_two = Point{X: 4.2, Y: 6.9}

// Создайте срез из нескольких экземпляров структуры Book
var books = []Book{
	{Title: "Alchemist", Author: "Paulo Coelho", Year: 1988},
	{Title: "Maktub", Author: "Paulo Coelho", Year: 1994},
	{Title: "The Silmarillion", Author: "J. R. R. Tolkien", Year: 1977},
	{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien", Year: 1954},
	{Title: "Fahrenheit 451", Author: "Ray Bradbury", Year: 1953},
	{Title: "The Martian Chronicles", Author: "Ray Bradbury", Year: 1950},
	{Title: "The Hitchhiker’s Guide to the Galaxy", Author: "Douglas Adams", Year: 1979},
	{Title: "The Restaurant at the End of the Universe", Author: "Douglas Adams", Year: 1980},
	{Title: "Марги Судхур", Author: "Sadriddin Ayni", Year: 1956},
	{Title: "Мактаби Кӯҳна", Author: "Sadriddin Ayni", Year: 1935},
}

// Напишите функцию FindBooksByAuthor, которая принимает срез книг и имя автора, а затем возвращает срез книг, написанных этим автором

func FindBooksByAuthor(author string) []Book {
	books_slice := []Book{}
	for _, book := range books {
		if author == book.Author {
			books_slice = append(books_slice, book)
		}
	}
	return books_slice
}

// Напишите функцию PrintStudentInfo, которая принимает экземпляр структуры Student и выводит информацию о студенте в формате:
// "Студент: [Name], Возраст: [Age], Курс: [Grade]"

func PrintStudentInfo(s Student) {
	fmt.Println("Student:", s.Name, "Age:", s.Age, "Grade:", s.Grade)
}

// Напишите функцию ArePointsEqual, которая принимает два экземпляра структуры Point и возвращает true, если их координаты совпадают, и false в противном случае
func ArePointsEqual(p1, p2 Point) bool {
	if p1 == p2 {
		return true
	} else {
		return false
	}
}

func main() {
	//fmt.Println("Student Info:")
	// PrintStudentInfo(first)
	fmt.Println()
	// fmt.Println("Area:", rect.area(), "Perimeter:", rect.perimeter())
	// fmt.Println(user1)
	// points := ArePointsEqual(point_one, point_two)
	// fmt.Println(points)
	fmt.Println()
	books = FindBooksByAuthor("Paulo Coelho")
	fmt.Println(books)

}
