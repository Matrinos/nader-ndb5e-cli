package main

import (
	"fmt"
	"testing"
)

func TestParseProductInfo(t *testing.T) {
	var p = new(ProductInfomation)
	ParseProductInfo("4235450000000000000000500006000100E60000", p)

	fmt.Println(p)
}
