package main

import (
	"testing"
)

// test length != 0
func TestGetData(t *testing.T) {
	data, err := getData(url)
	if len(data) == 0 {
		t.Fatal("URL don't have data")
	}
	if err != nil {
		t.Fatal(err)
	}

}

// test length of getData return == length of constructAccountsMAp return
func TestConstructAccountsMap(t *testing.T) {
	data, _ := getData(url)
	data_mapped, err := constructAccountsMap(data)

	if len(data) != len(data_mapped) {
		t.Fatal("constructAccountsMap test failed")
	}
	if err != nil {
		t.Fatal(err)
	}
}
