package slack

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"fmt"
)

func GetCommit(commitSha, commitMessage, jobStatus string) (string, error) {
	// Get commit from env, if none was given, get from S3
	if commitMessage != "" || strings.HasPrefix(jobStatus, "cron") {
		return commitMessage, nil
	} else {
		return GetCommitPr(commitSha)
	}
}

func GetCommitPr(commitSha string) (string, error){
	log.Println("Fetching pr commit...")
	arg0 := "git log --format=%B -n 1"

	cmd := exec.Command("sh", "-c", arg0, commitSha)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String() + "In GetCommitPr")
	}
	log.Println("Fetched commit:", string(out.String()))
	return out.String(), err
}

func AddSafeDirectory() (string, error){
	log.Println("Exporting git ceiling...")

	cmd := exec.Command("sh", "-c", "git config --global --add safe.directory /github/workspace")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String() + "In AddSafeDirectory")
	}
	log.Println("AddSafeDirectory:", string(out.String()))
	return out.String(), err
}

func GetBuildUrl(prBuildUrlRaw, pushBuildUrl, runId string) string{
	// Get build URL depending on push \ pr job
	if prBuildUrlRaw == ""{
		return pushBuildUrl + "/actions/runs/" + runId
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
