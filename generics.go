package marogo

import "sync"
import "encoding/json"
import "net/http"
import "bytes"

func MakeRequest(address string, method string, data interface{}, needsHeader bool) (*http.Response, error) {
	address = API_URL + address
	jsob, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(jsob)
	request, err := http.NewRequest(method, address, body)
	if err != nil {
		return nil, err
	}

	if needsHeader {
		request.Header.Add("Content-Type", "application/json")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func MakeAsyncRequest(address string, method string, data interface{}, wg *sync.WaitGroup, needsHeader bool) (bool, error) {
	MakeRequest(address, method, data, needsHeader)
	wg.Done()

	return true, nil
}
