package main

import (
	"io/ioutil"
	"flag"
	"fmt"
)

func main() {
	inFile := flag.String("in","","Input file")
	outFile := flag.String("out","","Output file")
	pkg := flag.String("package","main","Package name")
	name := flag.String("name","File","Identifier to use for the embedded data")

	flag.Parse()

	inData, err := ioutil.ReadFile(*inFile)

	if err != nil {
		panic(err.Error())
	}

	t := "package %s\n\nvar %s []byte = %#v\n"
	s := fmt.Sprintf(t, *pkg, *name, inData)

	err = ioutil.WriteFile(*outFile, []byte(s), 0644)

	if err != nil {
		panic(err.Error())
	}
}
