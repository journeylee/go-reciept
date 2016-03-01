package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/go-init/init"
)

type Goods struct {
	Name       string `ini:"name"`
	Quantifier string `ini:"quantifier"`
	Price      string `ini:"price"`
	Category   string `ini:"category"`
	BarCode    string `ini:"-"`
}

func main() {
	var goods map[string]*Goods
	var buyList []string

	cfg, err := ini.Load(os.Args[1])
	if err != nil {
		panic(err)
	}

	ipt, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(ipt, &buyList)
	if err != nil {
		panic(err)
	}

	sections := cfg.SectionStrings()
	total := len(sections) - 1
	goods = make(map[string]*Goods, total)

	for _, barcode := range sections {
		if barcode != "DEFAULT" {
			goods[barcode] = &Goods{}
			err = cfg.Section(barcode).MapTo(goods[barcode])
			if err != nil {
				panic(err)
			}
		}
	}

	for _, buyed := range buyList {

		sku := strings.Split(buyed, "-")

		count := 1
		if len(sku) > 1 {
			count, err = strconv.Atoi(sku[1])
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(count)
		for count > 0 {
			fmt.Println(goods[sku[0]])
			count--
		}
	}
}
