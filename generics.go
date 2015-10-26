package marogo

import "sync"
import "encoding/json"
import "net/http"
import "bytes"
import "fmt"

func MakeRequest(address string, method string, data interface{}) (*http.Response, error) {
	address = API_URL + address
	jsob, err := json.Marshal(data)
	fmt.Print("%v\n", string(jsob))
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(jsob)
	request, err := http.NewRequest(method, address, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func MakeAsyncRequest(address string, method string, data interface{}, wg *sync.WaitGroup) (bool, error) {
	MakeRequest(address, method, data)
	wg.Done()

	return true, nil
}
