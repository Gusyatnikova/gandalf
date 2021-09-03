package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
	"log"
)

type Request struct {
	LM string "https://leroymerlin.ru/search/"
	IKEA string "https://www.ikea.com/ru/ru/search/products/"
	query string
}

var req Request

func init() {
	flag.StringVar(&req.query, "q", "", "String to find, for example: Зеркало")
	flag.Parse()
}

type Item struct {
	Title string `json:"title"`
	Price float32 `json:"price"`
	URL string `json:"url"`
}


func main() {
	//searchStr := "?q=" + req.query
	ferretQry := `
		LET page = DOCUMENT("https://www.ikea.com/ru/ru/search/products/?q=Зеркало")
		FOR elem IN ELEMENTS(page, ".serp-grid__item")
			RETURN {
			    title: ATTR_GET(elem, "data-product-name"),
			    price: ATTR_GET(elem, "data-price"),
			    url: ATTR_GET(ELEMENT(elem, "a"), "href"),
			}
	`
	comp := compiler.New()
	progr, err := comp.Compile(ferretQry)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx := context.Background()
	ctx = drivers.WithContext(ctx, cdp.NewDriver())
	ctx = drivers.WithContext(ctx, http.NewDriver(), drivers.AsDefault())

	out, err := progr.Run(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	res := make([]*Item, 0)
	err = json.Unmarshal(out, &res)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}