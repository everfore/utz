package cmprs

import (
	"testing"
)

func TestUntar(t *testing.T) {
	Untar("./tar/test.tar", "./untar")
}

func TestUnzip(t *testing.T) {
	Unzip("./zip/test.zip", "./uzip")
}
