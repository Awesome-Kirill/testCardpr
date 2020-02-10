package testCardpr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"net/http"
)

type Signal struct {
	id       int
	httpCode int
}

func Shooting(count int) {

	signalChanel := make(chan Signal)
	for i := 0; i < count; i++ {

		go shoot(signalChanel, i)
	}
	var awt int
	var sukkses_shots int
	stataHttpCode := make(map[int]int)
	for {
		select {
		case msg1 := <-signalChanel:
			fmt.Println("thread", msg1.id, "=> http code:", msg1.httpCode)

			if _, ok := stataHttpCode[msg1.httpCode]; !ok {
				stataHttpCode[msg1.httpCode] = 0
			}
			stataHttpCode[msg1.httpCode] = stataHttpCode[msg1.httpCode] + 1
			sukkses_shots++
		default:
			time.Sleep(1 * time.Second)
			awt++
			if awt > 10 {
				fmt.Print("Finish shoot \n")
				fmt.Print("Success shoots ", sukkses_shots)
				fmt.Print(stataHttpCode)
				return
			}
			fmt.Print("Await \n", awt)
		}

	}

}

func shoot(signal chan Signal, id int) error {

	successBody := map[string]interface{}{
		"app_key":    "5240f691-60b0-4360-ac1f-601117c5408f",
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

	bytesRepresentation, err := json.Marshal(successBody)

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {

		return err
	}
	defer resp.Body.Close()

	signal <- Signal{
		id:       id,
		httpCode: resp.StatusCode,
	}
	return nil

}
