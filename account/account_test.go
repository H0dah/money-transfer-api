package account

import "testing"

// TestGetData tests that length != 0
func TestGetData(t *testing.T) {
	url := "https://git.io/Jm76h"

	data, err := getData(url)
	if len(data) == 0 {
		t.Fatal("URL don't have data")
	}
	if err != nil {
		t.Fatal(err)
	}

}

// TestConstructAccountsMap tests that length of GetData return == length of ConstructAccountsMAp return
func TestConstructAccountsMap(t *testing.T) {
	url := "https://git.io/Jm76h"

	data, _ := getData(url)
	data_mapped, err := constructAccountsMap(data)

	if len(data) != len(data_mapped) {
		t.Fatal("constructAccountsMap test failed")
	}
	if err != nil {
		t.Fatal(err)
	}
}
