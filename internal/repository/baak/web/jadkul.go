package web

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kucinghitam13/alien/internal/model/baak"
)

var (
	tdRegex      *regexp.Regexp
	tdSplitRegex *regexp.Regexp
)

func init() {
	tdRegex = regexp.MustCompile(`<\/*td.*<\/td>`)
	tdSplitRegex = regexp.MustCompile(`</*td.*?>`)
}

func (r *Repo) GetJadkul(query string) (jadkulList baak.JadkulList, err error) {
	for i := 1; i <= r.Config.RetryMaxAttempt; i++ {
		jadkulList, err = r.doGetJadkul(query)
		if err == nil {
			break
		}

		log.Printf("[ERR] Error getting jadwal kuliah, attempt : %d err: %s\n", i, err.Error())

		if i != r.Config.RetryMaxAttempt {
			errCred := r.refreshCredentials()
			if errCred != nil {
				log.Printf("[ERR] Error attempting to refresh credentials: %s\n", errCred.Error())
				return
			}
			time.Sleep(time.Duration(r.Config.RetryIntervalSeconds) * time.Second)
		}
	}
	return
}

func (r *Repo) doGetJadkul(query string) (jadkulList baak.JadkulList, err error) {
	urlReq := r.Config.Host + r.Config.JadkulEndpoint
	requestBody := url.Values{
		"_token": {r.Credentials.CSRFToken},
		"teks":   {url.QueryEscape(query)},
	}

	request := getDefaultRequest(http.MethodPost, urlReq, strings.NewReader(requestBody.Encode()))
	for _, c := range r.Credentials.Cookies {
		request.AddCookie(c)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(requestBody.Encode())))
	request.Header.Add("Pragma", "no-cache")
	request.Header.Add("Origin", r.Config.Host)
	request.Header.Add("Referer", urlReq)

	response, err := r.HTTPClient.Do(request)
	if err != nil {
		log.Printf("[ERR] error fetching from source: %s\n", err.Error())
		return
	}

	if response.StatusCode > 299 {
		err = errors.New("response is not 2xx")
		log.Printf("[ERR] error fetching from source: %s\n", err.Error())
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	jadkulList, _ = parseJadkul(string(body))
	return
}

func parseJadkul(htmlBody string) (jadkulList baak.JadkulList, err error) {
	tds := tdRegex.FindAllString(htmlBody, -1)

	for i := 0; i < len(tds)/2; i += 6 {
		jadkul := baak.Jadkul{
			Class:    tdSplitRegex.Split(tds[i], -1)[1],
			Day:      tdSplitRegex.Split(tds[i+1], -1)[1],
			Matkul:   tdSplitRegex.Split(tds[i+2], -1)[1],
			Period:   tdSplitRegex.Split(tds[i+3], -1)[1],
			Room:     tdSplitRegex.Split(tds[i+4], -1)[1],
			Lecturer: tdSplitRegex.Split(tds[i+5], -1)[1],
		}
		jadkulList.Jadkuls = append(jadkulList.Jadkuls, jadkul)
	}
	jadkulList.Total = len(jadkulList.Jadkuls)
	return
}
