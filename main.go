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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt string `json:"pushed_at"`
	Size int `json:"size"`
	Watchers int `json:"watchers"`
	DefaultBranch string `json:"default_branch"`
	HasIssues bool `json:"has_issues"`
	HasDownloads bool `json:"has_downloads"`
	WatchersCount int `json:"watchers_count"`


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
	fmt.Printf("\n\n")

	fmt.Printf("CreatedAt: %s\n", repoData.CreatedAt)
	fmt.Printf("UpdatedAt: %s\n", repoData.UpdatedAt)
	fmt.Printf("PushedAt: %s\n", repoData.PushedAt)
	fmt.Printf("Size: %d\n", repoData.Size)
	fmt.Printf("Watchers: %d\n", repoData.Watchers)
	fmt.Printf("DefaultBranch: %s\n", repoData.DefaultBranch)
	fmt.Printf("HasIssues: %t\n", repoData.HasIssues)
	fmt.Printf("HasDownloads: %t\n", repoData.HasDownloads)
	fmt.Printf("WatchersCount: %d\n", repoData.WatchersCount)


}
