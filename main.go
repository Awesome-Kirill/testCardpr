package main

import (
	"fmt"
	testCardpr "github.com/Awesome-Kirill/testCardpr/test"
	"net/http"
	"time"
)

func main() {

	testCardpr.Shooting(100000, 110)

	apiKey := "5240f691-60b0-4360-ac1f-601117c5408f"
	var timeOutSec time.Duration = 2
	successBody := map[string]interface{}{
		"app_key":    apiKey,
		"phone":      "+79111111112",
		"email":      "asd1d@ivan.ru",
		"name":       "Кирилл",
		"surname":    "Петров",
		"middlename": "Иванович",
		"birthday":   "11.12.1990",
		"discount":   "5",
		"bonus":      "0",
		"balance":    "0",
		"link":       "https://testphp.codepr.ru",
		"sms":        "Предлагаем установить карту: %link%",
	}

	client := http.Client{
		Timeout: timeOutSec * time.Second,
	}

	testCardpr.SuccessCreate(successBody, client)
	time.Sleep(2 * time.Second)
	// Create the target map

	filedApi := testCardpr.CopyMap(successBody)
	testCardpr.FailCreate(filedApi, client)
	time.Sleep(2 * time.Second)
	updateBody := testCardpr.CopyMap(successBody)
	testCardpr.SuccessUpdate(updateBody, client)
	time.Sleep(2 * time.Second)
	filedPhone := testCardpr.CopyMap(successBody)
	testCardpr.PhoneFail(filedPhone, client)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {

		return
	}
}
