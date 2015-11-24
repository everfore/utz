package cmprs

import (
	"testing"
)

// func TestUntar(t *testing.T) {
// 	Untar("./tar/all.tar")
// }

func TestUnzip(t *testing.T) {
	Unzip("./zip/all.zip", "/uzip")
}
