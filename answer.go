package main

import (
	"bytes"
	"log"
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
	return out.String() == test.Output
}

func runTests(tests []TestCase) {
	for i, test := range(tests) {
		result := runTest(test)
		if !result {
			log.Printf("Test %d failed\n", i)
			log.Println(test)
			return;
		} 	
	}
	log.Println("All tests successful")
}
