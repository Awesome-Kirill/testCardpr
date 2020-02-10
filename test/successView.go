package testCardpr

import "fmt"

func SuccessView(testName string, finish float64) {

	fmt.Print("\n"+testName+" "+"тест успешно пройден"+"|время запроса: = ", finish)
}
