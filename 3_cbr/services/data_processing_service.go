package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
	"main.go/repositories"
	"net/http"
	"strconv"
	"strings"
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

		var valcurs repositories.ValCurs

		decoder := xml.NewDecoder(resp.Body)
		decoder.CharsetReader = charset.NewReaderLabel
		err = decoder.Decode(&valcurs)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal response")
		}

		for _, v := range valcurs.Valutes {
			v.Date = valcurs.Date
			err = r.repository.SaveValute(&v)
			if err != nil {
				return errors.Wrap(err, "failed to save valute")
			}
		}

		day += 86400
	}

	return nil
}

func sendReq(rawUrl string) (*http.Response, error) {

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

	return resp, nil
}

func (r *Service) GetNeededInfo() error {

	maxim, err := r.repository.GetMaxValute()
	if err != nil {
		return errors.Wrap(err, "failed to get max valute")
	}

	minim, err := r.repository.GetMinValute()
	if err != nil {
		return errors.Wrap(err, "failed to get min valute")
	}
	all, err := r.repository.GetAllRecords()
	if err != nil {
		return errors.Wrap(err, "failed to get all records")
	}
	avg := 0.0
	for _, v := range *all {

		tmp := strings.Join(strings.Split(v.Value, ","), ".")

		value, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return errors.Wrap(err, "failed to parse float")
		}

		avg += value

	}
	avg = avg / float64(len(*all))
	fmt.Println(maxim.Value, maxim.Name, maxim.Date)
	fmt.Println(minim.Value, minim.Name, minim.Date)
	fmt.Println(avg)
	return nil
}
