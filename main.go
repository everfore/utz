package main

import (
	"flag"
	"fmt"
	"github.com/everfore/utz/cmprs"
	"os"
	"path/filepath"
	"strings"
)

var (
	target   = ""
	filename = ""
	pwd, _   = os.Getwd()
)

func init() {
	flag.StringVar(&target, "o", "/cmprs/untar", "-o ../untar")
	flag.StringVar(&filename, "i", "./cmprs/tar/test.tar", "-i test.tar")
	flag.Parse()
	target = filepath.Join(pwd, target)
}

func main() {
	fmt.Println(target, filename)
	if strings.HasSuffix(filename, ".tar") {
		cmprs.Untar(filename, target)
	} else if strings.HasSuffix(filename, ".zip") {
		cmprs.Unzip(filename, target)
	}
}
