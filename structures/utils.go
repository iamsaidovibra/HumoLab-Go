package main

// Задача 1: Создание и использование структуры
// Создайте структуру Student, которая будет содержать следующие поля:

type Student struct {
	Name  string
	Age   uint8
	Grade uint8
}

// Задача 2: Методы структур
// Создайте структуру Rectangle, которая будет содержать поля:
// - Width (float64) — ширина прямоугольника.
// - Height (float64) — высота прямоугольника.

type Rectangle struct {
	Width  float64
	Height float64
}

// Добавьте метод Area, который вычисляет и возвращает площадь прямоугольника.
// Добавьте метод Perimeter, который вычисляет и возвращает периметр прямоугольника.

func (r Rectangle) area() float64 {
	area := (r.Width * r.Height)
	return area
}

// Добавьте метод Perimeter, который вычисляет и возвращает периметр прямоугольника

func (r Rectangle) perimeter() float64 {
	perimeter := 2 * (r.Width * r.Height)
	return perimeter
}

// ### Задача 3: Вложенные структуры
// Создайте структуру Address, которая будет содержать поля:
// - City (строка) — город.
// - Street (строка) — улица.
// - HouseNumber (строка) — номер дома

type Address struct {
	City        string
	Street      string
	HouseNumber string
}

// Затем создайте структуру Person, которая будет содержать:
// - Name (строка) — имя человека.
// - Age (целое число) — возраст человека.
// - HomeAddress (тип Address) — адрес проживания

type Person struct {
	Name        string
	Age         uint8
	HomeAddress Address
}

// ### Задача 4: Сравнение структур
// Создайте структуру Point, которая будет содержать поля:
// - X (float64) — координата X.
// - Y (float64) — координата Y.

// Напишите функцию ArePointsEqual, которая принимает два экземпляра структуры Point и возвращает true, если их координаты совпадают, и false в противном случае

type Point struct {
	X float64
	Y float64
}

// ### Задача 5: Срезы структур
// Создайте структуру Book, которая будет содержать поля:
// - Title (строка) — название книги.
// - Author (строка) — автор книги.
// - Year (целое число) — год издания.

// Создайте срез из нескольких экземпляров структуры Book. Напишите функцию FindBooksByAuthor, которая принимает срез книг и имя автора, а затем возвращает срез книг, написанных этим автором

type Book struct {
	Title  string
	Author string
	Year   int
}
