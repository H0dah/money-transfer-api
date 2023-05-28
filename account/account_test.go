package account

import (
	"testing"
)

// TestConstructAccountsMap tests that length of GetDataFromDisk return == length of ConstructAccountsMAp return
func TestConstructAccountsMap(t *testing.T) {

	data := GetDataFromDisk()
	data_mapped, err := constructAccountsMap(data)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != len(data_mapped) {
		t.Fatal("constructAccountsMap test failed")
	}

}
