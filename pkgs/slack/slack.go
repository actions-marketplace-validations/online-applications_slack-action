package slack

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func GetCommit(commitSha, commitMessage string) (string, error) {
	// Get commit from env, if none was given, get from S3
	if commitMessage != "" {
		return commitMessage, nil
	} else {
		return GetCommitPr(commitSha)
	}
}

func GetCommitPr(commitSha string) (string, error){
	arg0 := "git log --format=%B -n 1"
	out, err := exec.Command("sh", "-c", arg0, commitSha).Output()
	if err != nil {
		log.Fatalln("Error was found while getting the latest commit message", err)
	}
	return string(out[:]), err
}

func GetBuildUrl(prBuildUrlRaw, pushBuildUrl, runId string) string{
	// Get build URL depending on push \ pr job
	if prBuildUrlRaw == ""{
		return pushBuildUrl + "/" + runId
	} else {
		return strings.TrimSuffix(prBuildUrlRaw, ".diff")
	}
}

func SendMessage(payload []byte, url string) error{
	client := &http.Client{}
	// Building request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Error was found while building HTTP request", err)
		return err
	}
	// Adding headers
	req.Header.Add("Content-Type", "application/json")

	// Sending request
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error was found while sending Slack's HTTP request", err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error was found while reading Slack's HTTP response", err)
		return err
	}
	if string(body) == "ok" {
		log.Println("Slack message was succesfuly sent")
	}
	return err
}
