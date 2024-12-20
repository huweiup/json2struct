package main

import (
	"fmt"
	"testing"
)

func TestJson2Struct(t *testing.T) {
	jsonData := `[{"a":1,"B":"string"},{"a":2,"B":"another string"}]`
	got, err := Json2Struct(jsonData, "A")
	if err != nil {
		panic(err)
	}
	fmt.Println(got)
}
