package marogo

import "sync"
import "errors"
import "github.com/jeffail/gabs"
import "encoding/json"
import "io/ioutil"
import "fmt"

var ErrNotParsed = errors.New("Could not get gabs to parse json buffer")

func (m Maropost) NewContact(first_name string, last_name string, email string) *Contact {
	contact := Contact{m, first_name, last_name, email, "", "", make(map[string]interface{}), false}
	return &contact
}

func (c *Contact) SubscribeToLists(lists []string) (bool, error) {
	wg := &sync.WaitGroup{}
	for _, v := range lists {
		wg.Add(1)
		go MakeAsyncRequest(c.Account+"/lists/"+v+"/contacts.json?auth_token="+c.AuthToken, "POST", c, wg, false)
	}
	wg.Wait()
	return true, nil
}

func (m *Maropost) GetContactsByList(list string, page string) (*gabs.Container, error) {
	// Make our request
	response, err := MakeRequest(m.Account+"/lists/"+list+"/contacts.json?page="+page+"&auth_token="+m.AuthToken, "GET", nil, false)
	if err != nil {
		return nil, err
	}

	var object interface{}
	jsonDecoder := json.NewDecoder(response.Body)
	if err = jsonDecoder.Decode(&object); err != nil {
		return nil, err
	}

	jsonObject := gabs.New()
	jsonObject.SetP(object, "array")

	return jsonObject.S("array"), nil
}

func (m *Maropost) UpdateContact(id string, listId string, data interface{}) (*gabs.Container, error) {
	object := make(map[string]interface{})
	object["contact"] = data
	url := m.Account + "/lists/" + listId + "/contacts/" + id + ".json?auth_token=" + m.AuthToken
	response, err := MakeRequest(url, "PUT", object, true)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonParsed, err := gabs.ParseJSON(jsonBytes)
	if err != nil {
		fmt.Printf("%s\n%s\n\n", url, string(jsonBytes))
	}
	return jsonParsed, err
}
