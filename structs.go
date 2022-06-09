package main

type Data struct {
	Name        string
	Group       string
	Url         string
	Interactive bool
	MemoryLimit uint32
	TimeLimit   uint32
	Tests       []TestCase
	TestType    string
}

type TestCase struct {
	Input  string
	Output string
}

