package main

/* Создать приложение, которое сначала выдает меню:
- 1. Посмотреть закладки
- 2. Добавть закладки
- 3. Удалить закладку
- 4.ВЫход
При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавлеие
ПРи 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

import "fmt"

func main() {
	storage := map[string]string{}
	fmt.Println("___ Приложение для хранения закладок ___")
Menu:
	for {
		fmt.Println("Выберите опцию: 1 - Просмотр всех 2 - Добавить закладку 3 - Удалить закладку 4 - Выход")
		var op int
		fmt.Scan(&op)
		switch op {
		case 4:
			break Menu
		case 1: 
			printAll(storage)
		case 2:
			add(storage)
		case 3:
			remove(storage)
		default:
			fmt.Println("Неправильно введен номер операции")
		}
	}	
}

func add(storage map[string]string) {
	fmt.Println("Функция добавления закладки")
	fmt.Println("Введите название адреса")
	var name string
	fmt.Scan(&name)
	fmt.Println("Введите адрес")
	var address string
	fmt.Scan(&address)
	storage[name] = address
	fmt.Println("Закладка добавлена")
}

func remove(storage map[string]string) {
	fmt.Println("Функция удаления закладки")
	fmt.Println("Введите название адреса")
	var name string
	fmt.Scan(&name)
	delete(storage, name)
	fmt.Println("Закладка удалена")
}

func printAll(storage map[string]string) {
	fmt.Println("Функция вывода всех закладок")
	fmt.Printf("Количество закладок: %v\n", len(storage))
	for key, value := range storage {
		fmt.Printf("Имя: %v Адрес: %v\n", key, value)
	}
}