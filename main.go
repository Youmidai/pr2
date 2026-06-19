package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}

	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scan(&empl.Name)
			fmt.Println("Возраст:")
			fmt.Scan(&empl.Age)
			fmt.Println("Позиция:")
			fmt.Scan(&empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scan(&empl.Salary)

			add := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					fmt.Println("Сотрудник успешно добавлен!")
					add = true
					break
				}
			}

			if !add {
				fmt.Println("Достигнуто ограничение по кол-ву сотрудников")
			}

		case 2:
			fmt.Println("\nВведите имя сотрудника для удаления:")
			var remove string
			fmt.Scan(&remove)

			del := false
			for i := 0; i < size; i++ {
				if empls[i] != nil && empls[i].Name == remove {
					empls[i] = nil
					fmt.Println("Сотрудник", remove, "удален.")
					del = true
					break
				}
			}

			if !del {
				fmt.Println("Сотрудник с таким именем не найден.")
			}

		case 3:
			fmt.Println("\nСписок сотрудников")
			count := 0
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf("Имя: %s | Возраст: %d | Должность: %s | Зарплата: %d\n",
						empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
					count++
				}
			}

			if count == 0 {
				fmt.Println("Штат пуст.")
			}

		case 4:
			break

		default:
			fmt.Println("Неизвестная команда. Попробуйте еще раз.")
		}
	}
}
