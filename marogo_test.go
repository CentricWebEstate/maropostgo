package marogo

import "testing"

func TestGetContactsByList(t *testing.T) {
	m := Maropost{"abcd1234", "000"}
	json, err := m.GetContactsByList("112233")
	t.Logf("%+v\n\n\n%+v", json, err)
}
