package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func compile(filename string) error {
	args := []string{"-std=c++17", "-O2", "-Wall", "-Wextra", "-pedantic", "-Wshadow", "-Wformat=2", "-Wfloat-equal", "-Wconversion", "-Wlogical-op", "-Wshift-overflow=2", "-Wduplicated-cond", "-Wcast-qual", "-Wcast-align", "-Wno-unused-result", "-Wno-sign-conversion", filename}

	cmd := exec.Command("g++", args...)
	return cmd.Run()
}

func runTest(test TestCase) bool {
	cmd := exec.Command("./a.out")
	cmd.Stdin = strings.NewReader(test.Input)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if(out.String() != test.Output) {
		fmt.Println("Test failed for the following input:")
		fmt.Println(test.Input)
		fmt.Println("Expected:")
		fmt.Println(test.Output)
		fmt.Println("Received:")
		fmt.Println(out.String())
	}
	return out.String() == test.Output
}

func runTests(tests []TestCase) {
	for _, test := range(tests) {
		result := runTest(test)
		if !result {
			os.Exit(1)
		} 	
	}
	fmt.Println("All tests successful")
	os.Exit(0)
}
