package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Repo struct {
	Id     string
	Repository string
    Source string
    Owner string
    DefaultBranch string
    IsPublic bool
    CreationDate string
}

func main() {
    url := "https://www.bridgecrew.cloud/api/v1/repositories"
     
    api := os.Getenv("BRIDGECREW_API")
    
    if api == "" {
        log.Fatal("BRIDGECREW_API is missing")
    }

    // Create a Bearer string by appending string access token
    var bearer = "Bearer " + api

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

    var repos []Repo
    json.Unmarshal([]byte(body), &repos)
	//fmt.Printf("Repository : %+v", repos)
    for _, repository := range repos {
        fmt.Println(repository.Repository, repository.IsPublic, repository.Owner)
    }
}
