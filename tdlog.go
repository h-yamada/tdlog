package tdlog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TDLog struct {
	EndPoint string //http://in.ybi.idcfcloud.net/postback/v3/event/<db>/<table>
	ApiKey   string //Write-Only API keys
}

func NewTDLog(endpoint string, apikey string) *TDLog {
	return &TDLog{EndPoint: endpoint, ApiKey: apikey}
}

func (t *TDLog) SendLog(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", t.EndPoint, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-TD-Write-Key", t.ApiKey)

	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
