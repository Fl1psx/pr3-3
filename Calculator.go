package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Calculator структура для калькулятора
type Calculator struct {
	history []string
}

// NewCalculator создает новый экземпляр калькулятора
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

// Add сложение
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.history = append(c.history, fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

// Subtract вычитание
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.history = append(c.history, fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

// Multiply умножение
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.history = append(c.history, fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
	return result
}

// Divide деление
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ошибка: деление на ноль")
	}
	result := a / b
	c.history = append(c.history, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
	return result, nil
}

// Power возведение в степень
func (c *Calculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	c.history = append(c.history, fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
	return result
}

// SquareRoot квадратный корень
func (c *Calculator) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("ошибка: нельзя извлечь корень из отрицательного числа")
	}
	result := math.Sqrt(a)
	c.history = append(c.history, fmt.Sprintf("√%.2f = %.2f", a, result))
	return result, nil
}

// Percentage процент от числа
func (c *Calculator) Percentage(number, percent float64) float64 {
	result := (number * percent) / 100
	c.history = append(c.history, fmt.Sprintf("%.2f%% от %.2f = %.2f", percent, number, result))
	return result
}

// ShowHistory показывает историю операций
func (c *Calculator) ShowHistory() {
	if len(c.history) == 0 {
		fmt.Println("История операций пуста")
		return
	}

	fmt.Println("\n--- История операций ---")
	for i, operation := range c.history {
		fmt.Printf("%d. %s\n", i+1, operation)
	}
	fmt.Println("------------------------")
}

// ClearHistory очищает историю операций
func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
	fmt.Println("История операций очищена")
}

// ShowMenu показывает меню операций
func ShowMenu() {
	fmt.Println("\n=== КАЛЬКУЛЯТОР ===")
	fmt.Println("1. Сложение (+)")
	fmt.Println("2. Вычитание (-)")
	fmt.Println("3. Умножение (*)")
	fmt.Println("4. Деление (/)")
	fmt.Println("5. Возведение в степень (^)")
	fmt.Println("6. Квадратный корень (√)")
	fmt.Println("7. Процент от числа (%)")
	fmt.Println("8. Показать историю операций")
	fmt.Println("9. Очистить историю")
	fmt.Println("0. Выход")
	fmt.Println("====================")
}

// GetNumber получает число от пользователя
func GetNumber(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число")
			continue
		}

		return value
	}
}

// GetOperation получает операцию от пользователя
func GetOperation() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Выберите операцию (1-9, 0 для выхода): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 1 && strings.Contains("0123456789", input) {
			return input
		}

		fmt.Println("Ошибка: введите число от 0 до 9")
	}
}

func main() {
	calculator := NewCalculator()

	fmt.Println("Добро пожаловать в калькулятор!")

	for {
		ShowMenu()
		choice := GetOperation()

		switch choice {
		case "0":
			fmt.Println("До свидания!")
			return

		case "1": // Сложение
			a := GetNumber("Введите первое число: ")
			b := GetNumber("Введите второе число: ")
			result := calculator.Add(a, b)
			fmt.Printf("Результат: %.2f + %.2f = %.2f\n", a, b, result)

		case "2": // Вычитание
			a := GetNumber("Введите первое число: ")
			b := GetNumber("Введите второе число: ")
			result := calculator.Subtract(a, b)
			fmt.Printf("Результат: %.2f - %.2f = %.2f\n", a, b, result)

		case "3": // Умножение
			a := GetNumber("Введите первое число: ")
			b := GetNumber("Введите второе число: ")
			result := calculator.Multiply(a, b)
			fmt.Printf("Результат: %.2f * %.2f = %.2f\n", a, b, result)

		case "4": // Деление
			a := GetNumber("Введите делимое: ")
			b := GetNumber("Введите делитель: ")
			result, err := calculator.Divide(a, b)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Результат: %.2f / %.2f = %.2f\n", a, b, result)
			}

		case "5": // Возведение в степень
			base := GetNumber("Введите основание: ")
			exponent := GetNumber("Введите степень: ")
			result := calculator.Power(base, exponent)
			fmt.Printf("Результат: %.2f ^ %.2f = %.2f\n", base, exponent, result)

		case "6": // Квадратный корень
			a := GetNumber("Введите число: ")
			result, err := calculator.SquareRoot(a)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Результат: √%.2f = %.2f\n", a, result)
			}

		case "7": // Процент от числа
			number := GetNumber("Введите число: ")
			percent := GetNumber("Введите процент: ")
			result := calculator.Percentage(number, percent)
			fmt.Printf("Результат: %.2f%% от %.2f = %.2f\n", percent, number, result)

		case "8": // Показать историю
			calculator.ShowHistory()

		case "9": // Очистить историю
			calculator.ClearHistory()

		default:
			fmt.Println("Неизвестная операция")
		}

		fmt.Print("\nНажмите Enter для продолжения...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
