// Package sha384 accepts one file and returns a sha-384 hash

/*

The MIT License (MIT)

Copyright (c) 2016 aerth

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

*/

package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

// Read file bytes
func readFile(filename string) []byte {
	var rawBytes []byte
	var err error
	if filename != "" {
		rawBytes, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return rawBytes
}

// Return the hash
func shasum(filename string) string {
	var hashString string
	rawBytes := readFile(filename)
	sha384sum := sha512.Sum384(rawBytes)
	hashString = fmt.Sprintf("%x", sha384sum)
	return hashString
}

func main() {
	if len(os.Args) == 2 { // Single file
		filename := os.Args[1]
		sha384sum := shasum(filename)
		fmt.Printf("%s %s\n", sha384sum, filename)
		os.Exit(1)
	} else if len(os.Args) > 2 { // Multiple files

		for file := range os.Args {
			if file == 0 {
				continue
			}
			filename := os.Args[file]
			sha384sum := shasum(filename)
			fmt.Printf("%s %s\n", sha384sum, filename)

		}

	} else {
		fmt.Printf("Usage: " + os.Args[0] + " [filename] [otherfilename] \n")
		os.Exit(1)

	}
}
