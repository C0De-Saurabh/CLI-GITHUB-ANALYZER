package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
)

type Repo struct {
    FullName     string `json:"full_name"`
    Description  string `json:"description"`
    Stars        int    `json:"stargazers_count"`
    Forks        int    `json:"forks_count"`
    OpenIssues   int    `json:"open_issues_count"`
    Subscribers  int    `json:"subscribers_count"`
	Language string `json:"language"`
}

func fetchRepoData(repo string) (*Repo, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s", repo)
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        return nil, err
    }
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
    }

    var repoData Repo
    if err := json.NewDecoder(resp.Body).Decode(&repoData); err != nil {
        return nil, err
    }

    return &repoData, nil
}

func main() {
    repo := flag.String("repo", "", "GitHub repository in the format owner/repo")
    flag.Parse()

    if *repo == "" {
        log.Fatal("Please provide a GitHub repository in the format owner/repo")
    }

    

    repoData, err := fetchRepoData(*repo)
    if err != nil {
        log.Fatalf("Error fetching repository data: %v", err)
    }

    fmt.Printf("Repository: %s\n", repoData.FullName)
    fmt.Printf("Description: %s\n", repoData.Description)
    fmt.Printf("Stars: %d\n", repoData.Stars)
    fmt.Printf("Forks: %d\n", repoData.Forks)
    fmt.Printf("Open Issues: %d\n", repoData.OpenIssues)
    fmt.Printf("Subscribers: %d\n", repoData.Subscribers)
	fmt.Printf("Language: %s\n", repoData.Language)
}
