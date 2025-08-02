package main

import (
	"bytes"
	"fmt"
	"io"
	"maps"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const DateFormat string = "02/01/2006"

func main() {
	rawUrl := "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req="
	today := time.Now().Format(DateFormat)
	rawUrl += today

	regexpCharCode, err := regexp.Compile("<CharCode>([A-Z])+<\\/CharCode>")
	if err != nil {
		panic(err)
	}
	regexpVunitRate, err := regexp.Compile("<VunitRate>([0-9]+,[0-9]+)<\\/VunitRate>")
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	var buf []byte
	req, _ := http.NewRequest(http.MethodGet, rawUrl, bytes.NewBuffer(buf))
	req.Header.Set(
		"User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	readBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	CharCodes := regexpCharCode.FindAll(readBody, -1)
	VunitRates := regexpVunitRate.FindAll(readBody, -1)
	valutes := make(map[string]float64)
	for i := range CharCodes {
		CharCode := strings.ReplaceAll(strings.ReplaceAll(string(CharCodes[i]), "<CharCode>", ""), "</CharCode>", "")

		VunitRates[i] = []byte(strings.ReplaceAll(strings.ReplaceAll(string(VunitRates[i]), "<VunitRate>", ""), "</VunitRate>", ""))
		VunitRate, err := strconv.ParseFloat(strings.Join(strings.Split(string(VunitRates[i]), ","), "."), 64)
		if err != nil {
			fmt.Println(err)
		}

		valutes[CharCode] = VunitRate
	}
	maps.Values(valutes)
	maximum := 0.0
	maxChar := ""
	for i, v := range valutes {
		if v > maximum {
			maximum = v
			maxChar = i
		}
	}
	minimum := 999999999999.9
	minChar := ""
	for i, v := range valutes {
		if v < minimum {
			minimum = v
			minChar = i
		}
	}
	for _, v := range valutes {
		fmt.Println(v)
	}

	fmt.Println("max: ", valutes[maxChar])
	fmt.Println("min: ", valutes[minChar])
	return
}
