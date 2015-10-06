package marogo

func (m Maropost) NewContact(first_name string, last_name string, email string) *Contact {
	contact := Contact{m, first_name, last_name, email}
	return &contact
}

func (c *Contact) SubscribeToLists(lists []string) (bool, error) {
	for i, v := range lists {
		go MakeRequest(c.Account+"/lists/"+v+"/contacts.json", "POST", c)
	}
}
