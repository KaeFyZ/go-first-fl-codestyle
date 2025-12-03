// go-first-fl-codestyle/main.go
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Attack возвращает текст о том, сколько урона персонаж нанес противнику.
func Attack(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(3, 5))
	}

	if class == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(5, 10))
	}

	if class == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// Defence возвращает текст о том, сколько урона персонаж смог заблокировать.
func Defence(char_name, char_class string) string {
	switch char_class {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

// Special возвращает строку с описанием специального умения персонажа.
func Special(name, class string) string {
	switch class {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)
	default:
		return "неизвестный класс персонажа"
	}
}

// StartTraining запускает тренировку персонажа и обрабатывает команды пользователя.
func StartTraining(name, class string) string {
	// Сообщение о классе персонажа
	switch class {
	case "warrior":
		fmt.Printf("%s, ты Воитель — отличный боец ближнего боя.\n", name)
	case "mage":
		fmt.Printf("%s, ты Маг — превосходный укротитель стихий.\n", name)
	case "healer":
		fmt.Printf("%s, ты Лекарь — чародей, способный исцелять раны.\n", name)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		if _, err := fmt.Scanf("%s\n", &cmd); err != nil {
			fmt.Println("Ошибка ввода команды:", err)
			continue
		}

		switch cmd {
		case "attack":
			fmt.Println(Attack(name, class))
		case "defence":
			fmt.Println(Defence(name, class))
		case "special":
			fmt.Println(Special(name, class))
		}
	}

	return "тренировка окончена"
}

// ChoiseClass позволяет игроку выбрать класс персонажа и возвращает выбранный класс.
func ChoiseClass() string {
	var approval string
	var class string

	for approval != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")

		if _, err := fmt.Scanf("%s\n", &class); err != nil {
			fmt.Println("Ошибка ввода класса:", err)
			continue
		}

		switch class {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")

		if _, err := fmt.Scanf("%s\n", &approval); err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		approval = strings.ToLower(approval)
	}

	return class
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	if _, err := fmt.Scanf("%s\n", &name); err != nil {
		fmt.Println("Ошибка ввода класса:", err)
		return
	}

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := ChoiseClass()

	fmt.Println(StartTraining(name, class))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
