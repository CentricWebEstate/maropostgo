package marogo

func MakeRequest(address string, method string, data interface{}) (bool, error) {
	address = API_URL + address
	jsob, err := json.Marshall(data)
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest(method, address, jsob)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
}
