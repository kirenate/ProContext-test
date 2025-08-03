package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"main.go/repositories"
	"net/http"
	"time"
)

type Service struct {
	repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{repository: repository}
}

const DateFormat = "02/01/2006"

const rawUrl = "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req="

func createURL(rawDate time.Time) string {
	date := rawDate.Format(DateFormat)
	url := rawUrl + date
	return url
}

func (r *Service) ProcessData() error {
	end := time.Now().Unix()
	start := end - 7776000
	day := start + 86400
	for day != end {
		timeDay := time.Unix(day, 0)
		url := createURL(timeDay)
		resp, err := sendReq(url)
		if err != nil {
			return errors.Wrap(err, "failed to send request")
		}

		var valcurs *repositories.ValCurs
		err = xml.Unmarshal(resp, &valcurs)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal response")
		}

		fmt.Println(valcurs)
		log.Info().Interface("v", valcurs).Msg("")

		day += 86400
	}
	return nil
}

func sendReq(rawUrl string) ([]byte, error) {

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
