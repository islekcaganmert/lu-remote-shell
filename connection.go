package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type shell struct {
	Input      bool              `json:"input"`
	Resolution []int             `json:"resolution"`
	Screen     []string          `json:"screen"`
	Variables  map[string]string `json:"variables"`
	Active     bool              `json:"active"`
}

func get_aas(domain string) (string, error) {
	resp, err := http.PostForm("https://aas.hereus.net/protocols/aas/get", url.Values{
		"domain": {domain},
	})
	if err != nil {
		return domain, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain, nil
	} else {
		return string(body), nil
	}

}

func get_shell(email string, password string, shellname string) (shell, error) {
	var s shell
	domain, err := get_aas(strings.Split(email, "@")[1])
	resp, err := http.PostForm("https://"+domain+"/protocols/shell", url.Values{
		"current_user_username": {strings.Split(email, "@")[0]},
		"current_user_password": {password},
		"method":                {"GET"},
		"shell":                 {shellname},
	})
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal([]byte(string(body)), &s)
		if err != nil {
			return s, err
		}
		return s, nil
	}
	return s, nil
}

func send_command(email string, password string, shellname string, command string) error {
	domain, err := get_aas(strings.Split(email, "@")[1])
	resp, err := http.PostForm("https://"+domain+"/protocols/shell", url.Values{
		"current_user_username": {strings.Split(email, "@")[0]},
		"current_user_password": {password},
		"method":                {"RUN"},
		"shell":                 {shellname},
		"command":               {command},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
