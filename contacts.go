package marogo

import "sync"
import "encoding/json"

func (m Maropost) NewContact(first_name string, last_name string, email string) *Contact {
	contact := Contact{m, first_name, last_name, email, "", "", make(map[string]interface{}), false}
	return &contact
}

func (c *Contact) SubscribeToLists(lists []string) (bool, error) {
	wg := &sync.WaitGroup{}
	for _, v := range lists {
		wg.Add(1)
		go MakeAsyncRequest(c.Account+"/lists/"+v+"/contacts.json?auth_token="+c.AuthToken, "POST", c, wg)
	}
	wg.Wait()
	return true, nil
}

func (m *Maropost) GetContactsByList(list string) (*[]map[string]interface{}, error) {
	response, err := MakeRequest(m.Account+"/lists/"+list+"/contacts.json?auth_token="+m.AuthToken, "GET", nil)
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	json_encoder := json.NewDecoder(response.Body)
	err = json_encoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
