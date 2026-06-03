package main

import(
	"bufio"
	"os"
	"fmt"
	"strings"
)

func manageSpots(){
	count := 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Панель администратора запущена. Доступные команды:\nadd [название]\nstatus\nexit")

	for{
		fmt.Print(">")
		if !scanner.Scan(){
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(input, "add "){
			if count == 5{
				fmt.Println("База спотов переполнена")
				continue
			}
			spotName := strings.TrimSpace(strings.TrimPrefix(input, "add "))

			if spotName == ""{
				fmt.Println("Ошибка: вы не ввели название спота")
				continue
			}
			count++
			fmt.Printf("Спот %s добавлен. Места осталось: %d\n", spotName, (5 - count))
			continue
		}
		switch input{
		case "status":
			switch count{
				case 0:
					fmt.Println("База пуста")
				case 1, 2, 3, 4:
					fmt.Println("База заполняется")
				case 5:
					fmt.Println("База заполнена")
			}
		case "exit":
			fmt.Println("Завершение работы...")
			return
		case "":
			continue
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}