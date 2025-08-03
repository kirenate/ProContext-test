package main

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

const DateFormat = "02/01/2006"

func main() {
	rawUrl := "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req="
	today := time.Now().Format(DateFormat)
	rawUrl += today

	resp, err := SendReq(rawUrl)
	if err != nil {
		log.Error().Stack().Err(err)
	}
	return
}

func SendReq(rawUrl string) ([]byte, error) {

	client := http.DefaultClient

	var buf []byte
	req, err := http.NewRequest(http.MethodGet, rawUrl, bytes.NewBuffer(buf))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make new request")
	}

	req.Header.Set(
		"User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send GET request")
	}

	readBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close request body")
	}

	return readBody, nil
}
