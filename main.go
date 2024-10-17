package main

import "fmt"

var client = make(map[string]int)

func RegUsers() {
	var (
		name    string
		balance int
	)
	fmt.Println("Введите имя")
	fmt.Scan(&name, balance)
	client[name] = 0
	fmt.Println("Вы успешно заригестрированы")
}

func FindBalace() {
	fmt.Println("Введите имя")
	var ScanName string
	fmt.Scan(&ScanName)
	for name, balance := range client {
		if name == ScanName {
			fmt.Println(name, "= ", balance)
		} else {
			fmt.Println("Клиент не найден")
		}
	}
}

func ListClient() {
	for name := range client {
		fmt.Println(name)

	}
}

func Replenishment() {
	fmt.Println("Введите имя")
	var ScanName string
	fmt.Scan(&ScanName)
	for name, balance := range client {
		if name == ScanName {
			fmt.Println("Текущий баланс:", name, "= ", balance)
			fmt.Println("Введите сумму пополнения")
			var Replen int
			fmt.Scan(&Replen)
			client[name] = balance + Replen
			fmt.Println("Вы пополнили на сумм = ", Replen)
		} else {
			fmt.Println("Клиент не найден")
		}
	}
}

func WriteOff() {
	var ScanName string
	fmt.Scan(&ScanName)
	for name, balance := range client {
		if name == ScanName {
			fmt.Println("Текущий баланс:", name, "= ", balance)
			fmt.Println("Введите сумму списания")
			var Replen int
			fmt.Scan(&Replen)
			client[name] = balance - Replen
			fmt.Println("C вашего  = ", Replen)
		} else {
			fmt.Println("Клиент не найден")
		}
	}

}

func main() {
	for {
		fmt.Println("Выберите операцию ")
		fmt.Println("Регистрация = 1")
		fmt.Println("Проверка баланса= 2")
		fmt.Println("Список клиентов = 3")
		fmt.Println("Пополнение = 4")
		fmt.Println("Списание = 5")
		fmt.Println("Выйти = 6")

		var a int
		fmt.Scan(&a)

		if a == 1 {
			RegUsers()
		} else if a == 2 {
			FindBalace()
		} else if a == 3 {
			ListClient()

		} else if a == 4 {
			Replenishment()
		} else if a == 5 {
			WriteOff()

		} else if a == 6 {
			break

		} else {
			fmt.Println("Выберите  от 1 до 5")
		}

	}

}
