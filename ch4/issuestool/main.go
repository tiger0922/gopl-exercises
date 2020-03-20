package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

const issuesURL = "https://api.github.com/repos"
const owner = "tiger0922"
const repo = "gopl-exercises"

// CreateIssue creates a github issue ...
func CreateIssue(token string) error {
	posturl := issuesURL + "/" + owner + "/" + repo + "/issues"
	jsonStr := []byte(`{"title":"Testing Issue1", "body":"This is a testing issue."}`)
	body := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", posturl, body)
	req.Header.Set("Authorization", token)
	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		resp.Body.Close()
		return fmt.Errorf("search query failed: %s", resp.Status)
	}

	return nil
}

func main() {
	file, err := os.Open("token")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	token := ""
	for scanner.Scan() {
		token = scanner.Text()
	}
	err = CreateIssue(token)
	if err != nil {
		log.Fatal(err)
	}
}
