package models

import (
	"fmt"
	"testing"
)

func Test_timeInt2Date(t *testing.T) {
	a := int64(1668847255)
	fmt.Println(timeInt2Date(a))
}