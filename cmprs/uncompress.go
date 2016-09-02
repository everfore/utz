package cmprs

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"github.com/toukii/goutils"
	"io"
	"os"
	"path/filepath"
)

var (
	Spor = string(filepath.Separator)
)

func Untar(tar_filename, target string) (err error) {
	file, err := os.OpenFile(tar_filename, os.O_RDONLY, 0444)
	defer file.Close()
	if goutils.CheckErr(err) {
		return
	}
	tr := tar.NewReader(file)
	var target_filename string
	for {
		hr, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(hr.FileInfo().Name())
		target_filename = filepath.Join(target, hr.Name)
		if hr.FileInfo().IsDir() {
			os.MkdirAll(target_filename, 0666)
			continue
		}
		of, err := os.OpenFile(target_filename, os.O_CREATE|os.O_WRONLY, 0666)
		defer of.Close()
		if goutils.CheckErr(err) {
			continue
		}
		_, err = io.Copy(of, tr)
		if goutils.CheckErr(err) {
			continue
		}
	}
	return nil
}

func Unzip(zip_filename, target string) (err error) {
	zr, err := zip.OpenReader(zip_filename)
	defer zr.Close()
	if goutils.CheckErr(err) {
		return
	}
	var target_filename string
	for _, it := range zr.File {
		fmt.Println(it.Name)
		target_filename = filepath.Join(target, it.Name)
		if it.FileInfo().IsDir() {
			os.MkdirAll(target_filename, 0666)
			continue
		}
		of, err := os.OpenFile(target_filename, os.O_CREATE|os.O_WRONLY, 0666)
		defer of.Close()
		if goutils.CheckErr(err) {
			continue
		}
		rit, err := it.Open()
		defer rit.Close()
		_, err = io.Copy(of, rit)
		if goutils.CheckErr(err) {
			continue
		}
	}
	return nil
}
