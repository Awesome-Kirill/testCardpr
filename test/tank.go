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

func Shooting(count int, treads int) {
	jobsChanel := make(chan Signal, count)
	for i := 0; i < count; i++ {
		jobsChanel <- Signal{
			id:       i,
			httpCode: 0,
		}

	}
	signalChanel := make(chan Signal)
	for i := 0; i < treads; i++ {

		go shoot(signalChanel, jobsChanel)
	}

	var awt int
	var sukkses_shots int
	stataHttpCode := make(map[int]int)
	for {
		if len(jobsChanel) < 1 {
			fmt.Print("Finish shoot \n")
			fmt.Print("Success shoots ", sukkses_shots)
			fmt.Print(stataHttpCode)
			return
		}
		select {
		case msg1 := <-signalChanel:

			fmt.Print("Read frome chanel")
			if _, ok := stataHttpCode[msg1.httpCode]; !ok {
				stataHttpCode[msg1.httpCode] = 0
			}
			stataHttpCode[msg1.httpCode] = stataHttpCode[msg1.httpCode] + 1
			sukkses_shots++

		default:
			time.Sleep(1 * time.Second)
			awt++
			if awt > 2000 {
				fmt.Print("Finish shoot \n")
				fmt.Print("Success shoots ", sukkses_shots)
				fmt.Print(stataHttpCode)
				return
			}
			fmt.Print("Await \n", awt)

		}

	}

}

func shoot(signal chan Signal, jobsChan chan Signal) error {

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

	if err != nil {
		fmt.Print("Marshal failed \n")
	}
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for job := range jobsChan {

		fmt.Println("thread", job.id)
		resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
		//defer resp.Body.Close()
		if err != nil {
			fmt.Println("thread error", job.id)
			continue

		}

		fmt.Println("thread", job.id, "code=>", resp.StatusCode)
		signal <- Signal{
			id:       job.id,
			httpCode: resp.StatusCode,
		}

	}

	return nil

}
