package testCardpr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const urlApi string = "https://core.codepr.ru/api/v2/crm/user_create_or_update"

func FailCreate(body map[string]interface{}, client http.Client) {
	bodyNotValidApi := body
	bodyNotValidApi["app_key"] = "42"
	bytesRepresentation, err := json.Marshal(bodyNotValidApi)

	if err != nil {
		fmt.Print(err)
	}
	start := time.Now()
	resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	finish := time.Since(start).Seconds()
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if resp.StatusCode == 400 {

		SuccessView("failCreate", finish)
		return

	}

	fmt.Print("failCreate тест провален")
	fmt.Print(resp.StatusCode)
	fmt.Print(result)
}
func SuccessCreate(body map[string]interface{}, client http.Client) {
	bytesRepresentation, err := json.Marshal(body)

	if err != nil {
		fmt.Print(err)
	}
	start := time.Now()
	resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	finish := time.Since(start).Seconds()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if resp.StatusCode == 200 {
		if val, ok := result["success"]; ok {
			if true == val {
				SuccessView("successCreate", finish)
				return
			}
		}
	}

	fmt.Print("successCreate тест провален")
	fmt.Print(resp.StatusCode)
	fmt.Print(result)
}

func SuccessUpdate(body map[string]interface{}, client http.Client) {

	bytesRepresentation, err := json.Marshal(body)

	if err != nil {
		fmt.Print(err)
	}
	resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
	defer resp.Body.Close()

	var result map[string]interface{}
	var user_hash string

	json.NewDecoder(resp.Body).Decode(&result)
	if resp.StatusCode == 200 {
		if val, ok := result["success"]; ok {
			if true == val {
				user_hash = result["user_hash"].(string)
			}
		}
	}

	bodyWithApdeitURL := body
	newLink := "https://ya.ru"
	bodyWithApdeitURL["link"] = newLink
	bytesApdeitURL, err := json.Marshal(bodyWithApdeitURL)

	if err != nil {
		fmt.Print(err)
	}
	time.Sleep(2 * time.Second)
	start := time.Now()
	respApdeitURL, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesApdeitURL))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	finish := time.Since(start).Seconds()

	var resultUpdate map[string]interface{}
	json.NewDecoder(respApdeitURL.Body).Decode(&resultUpdate)

	if respApdeitURL.StatusCode == 200 {
		if val, ok := resultUpdate["success"]; ok {
			if true == val {
				if user_hash == resultUpdate["user_hash"].(string) && strings.Contains(resultUpdate["form_url"].(string), newLink) {

					SuccessView("successUpdate", finish)
					return

				}
			}
		}
	}

	fmt.Print("successUpdate тест провален")
	fmt.Print(resp.StatusCode)
	fmt.Print(result)
}

func PhoneFail(body map[string]interface{}, client http.Client) {
	body["phone"] = "42"
	bytesRepresentation, err := json.Marshal(body)

	if err != nil {
		fmt.Print(err)
	}
	start := time.Now()
	resp, err := client.Post(urlApi, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	finish := time.Since(start).Seconds()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if resp.StatusCode == 400 {

		SuccessView("phoneFail", finish)
		return

	}

	fmt.Print("successCreate тест провален")
	fmt.Print(resp.StatusCode)
	fmt.Print(result)

}
