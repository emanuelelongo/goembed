package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type options struct {
	name     string
	pkg      string
	prefix   string
	noheader bool
}

func embed(path string, opt options) {
	switch stat, err := os.Stat(path); {
	case os.IsNotExist(err):
		fmt.Println("File or directory not found", err)
	case stat.IsDir():
		embedDirectory(opt.pkg, opt.prefix, opt.noheader, path)
	default:
		embedFile(opt.name, opt.pkg, opt.prefix, opt.noheader, path)
	}
}

func embedDirectory(pkg, prefix string, noheader bool, dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		fmt.Println("Can't read directory", err)
		return
	}

	omitHeader := noheader
	for _, file := range files {
		filename := file.Name()
		embedFile("", pkg, prefix, omitHeader, path.Join(dirpath, filename))
		omitHeader = true
	}
}

func embedFile(name, pkg, prefix string, noheader bool, filepath string) {
	if !noheader {
		fmt.Printf("// auto-generated\n\npackage %s\n\n", pkg)
	}
	fmt.Println()

	if name == "" {
		name = buildVarName(path.Base(filepath), prefix)
	}
	writeResource(filepath, name)
}

func buildVarName(file, prefix string) string {
	titled := strings.Title(file)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	name := reg.ReplaceAllString(titled, "")

	return prefix + name
}

func writeResource(file, name string) {
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Can't open file"+file, err)
		return
	}
	fmt.Printf("var %s = %#v\n", name, res)
}
