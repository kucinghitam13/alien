package web

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var (
	tokenRegex *regexp.Regexp
)

func init() {
	tokenRegex = regexp.MustCompile(TOKEN_REGEX)
}

func (r *Repo) refreshCredentials() (err error) {
	for i := 1; i <= r.Config.RetryMaxAttempt; i++ {
		if err = r.doRefreshCredentials(); err == nil {
			break
		}
		log.Printf("[ERR] Error refreshing credentials, attempt : %d err: %s\n", i, err.Error())

		if i != r.Config.RetryMaxAttempt {
			time.Sleep(time.Duration(r.Config.RetryIntervalSeconds) * time.Second)
		}
	}
	return
}

func (r *Repo) doRefreshCredentials() (err error) {
	request := getDefaultRequest(http.MethodGet, r.Config.Host, nil)

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

	csrfToken, err := parseCSRFToken(string(body))
	if err != nil {
		log.Printf("[ERR] error getting csrf token: %s\n", err.Error())
		return
	}

	if len(response.Cookies()) == 0 {
		err = errors.New("no cookies were present")
		log.Printf("[ERR] error getting cookies: %s\n", err.Error())
		return
	}

	r.Credentials.CSRFToken = csrfToken
	r.Credentials.Cookies = response.Cookies()

	return
}

func parseCSRFToken(htmlBody string) (csrfToken string, err error) {
	csrfToken = tokenRegex.FindString(htmlBody)
	if csrfToken == "" {
		err = errors.New("error parsing csrf token from html body")
		return
	}
	return
}
