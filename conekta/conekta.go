package conekta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gustavotero7/go-conekta/models"
)

var (
	ApiKey     = ""
	ApiVersion = "2.0.0"
)

const conektaUrl = "https://api.conekta.io"

func request(method, path string, v interface{}) ([]byte, error) {
	jsonPayload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, conektaUrl+path, bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/vnd.conekta-v"+ApiVersion+"+json")
	req.SetBasicAuth(ApiKey, "")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// If status isn't OK (200) then try to decode it into a conekta error
	// If decode fails then return decode (Unmarshal) error and nil response
	if res.StatusCode != 200 {
		conektaError := &models.ConektaError{}
		if err := json.Unmarshal(body, conektaError); err != nil {
			return nil, err
		}
		return nil, conektaError
	}

	return body, nil
}

func ConektaFormatAmount(value float64) (formatted int64, err error) {
	strnum := fmt.Sprintf("%.2f", value)
	strnum = strings.Replace(strnum, ".", "", -1)
	formatted, err = strconv.ParseInt(strnum, 10, 64)
	return
}

func ConektaFormatToFloat64(conektaFormatted int64) float64 {
	return float64(conektaFormatted) / 100
}
