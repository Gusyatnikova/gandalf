package main

import (
	"flag"
	"fmt"
)

type Request struct {
	urlFirst string
	urlSecond string
	query string
}

var req Request

func init() {
	flag.StringVar(&req.urlFirst,"f", "", "Path to search page, for example: https://leroymerlin.ru/search/")
	flag.StringVar(&req.urlSecond, "s", "", "Path to another search page, for example: https://www.ikea.com/ru/ru/search/products/")
	flag.StringVar(&req.query, "q", "", "String to find, for example: Зеркало")
}

func main() {
	flag.Parse()
	fmt.Println("Hi from gendalf!\n", req)
}