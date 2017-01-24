package main

import (
    "github.com/google/go-github/github"
    // "golang.org/x/oauth2"
    "golang.org/x/crypto/ssh/terminal"

    "fmt"
    "bufio"
    "strings"
    "os"
    "syscall"
)


func main() {
    //go-github example

    //Get username & password
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Username: ")
    username, _ := reader.ReadString('\n')

    fmt.Print("Password: ")
    bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
    password := string(bytePassword)

    //Authenticate the user
    tp := github.BasicAuthTransport{
        Username: strings.TrimSpace(username),
        Password: strings.TrimSpace(password),
    }

    client := github.NewClient(tp.Client())

    username = strings.TrimSpace(username)

    //messages hold one long concatenated string of all the user's commit messages
    messages := getCommitMessages(username, client)


    fmt.Print("\n")
    //Iterate through the map and print each key/value pair
    wordCountMap := wordCount(messages)
    for key, value := range wordCountMap {
        fmt.Printf("%v:\t%v\n", key, value)
    }
}

//Gets word count and return a map with Key - string Value - int
func wordCount(words string) map[string]int {
    //Split each word in words by " " and return a list of words
    wordList := strings.Fields(words)
    counts := make(map[string]int)

    for i := range wordList {
        counts[wordList[i]]++
    }

    return counts
}

func getCommitMessages(username string, client *github.Client) string {
    commitString := ""
    reposList, _, _ := client.Repositories.List("", nil)

    //Iterate through the list of repositories
    for _, repo := range reposList {
        //Get the name of each repository from the list (actually a list of structs)
        repoName := *repo.Name

        //List the commits on each repository 
        commits, _, _ := client.Repositories.ListCommits(username, repoName, nil)

        //Iterate through the array of commits
        for _, singleCommit := range commits {

            //Get the Commit field of the RepositoryCommit struct
            //commitData is a Commmit struct
            commitData := singleCommit.Commit

            commitString = commitString + " " + *commitData.Message
        }
    }
    return commitString
}