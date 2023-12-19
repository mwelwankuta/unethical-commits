package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mwelwankuta/unethical-commits/github"
	"github.com/mwelwankuta/unethical-commits/models"
)

func main() {
	// load
	gerr := godotenv.Load()
	if gerr != nil {
		log.Fatalf("err loading: %v", gerr)
	}
	var GitHubOwner string = os.Getenv("GITHUB_OWNER")
	var GitHubRepo string = os.Getenv("GITHUB_REPO")
	var GitHubBaseURL string = "https://api.github.com"
	var GitHubToken string = os.Getenv("GITHUB_TOKEN")

	githubClient := github.NewGitHubClient(GitHubBaseURL, GitHubToken)

	// Randomly decide whether to create an issue or a commit
	// Create an issue
	issue := models.GitHubIssue{
		Title: "Random Issue",
		Body:  "This is a randomly generated issue.",
	}

	err := githubClient.CreateContribution(GitHubOwner, GitHubRepo, issue)
	if err != nil {
		fmt.Println("Error creating GitHub issue:", err)
		return
	}

	fmt.Println("GitHub issue created successfully.")
}
