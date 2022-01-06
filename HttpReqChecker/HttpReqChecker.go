package httpreqchecker

import (
	"io/ioutil"
	"net/http"
)

func GetReq(host string) (string, int, error) {
	resp, err := http.Get(host)
	if err != nil {
		return "", resp.StatusCode, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	return string(bodyBytes), resp.StatusCode, nil
}
