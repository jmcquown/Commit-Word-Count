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

    tp := github.BasicAuthTransport{
        Username: strings.TrimSpace(username),
        Password: strings.TrimSpace(password),
    }

    client := github.NewClient(tp.Client())

    username = strings.TrimSpace(username)

    reposList, _, _ := client.Repositories.List("", nil)

    for _, repo := range reposList {
        // repoName := github.Stringify(repo.Name)
        repoName := *repo.Name

        //List the commits on each repository 
        commits, _, _ := client.Repositories.ListCommits(username, repoName, nil)
        // fmt.Printf("\n%v\n", github.Stringify(commits))

        //Iterate through the array of commits
        for _, singleCommit := range commits {

            //Guess this isn't needed :/
            //Get a single commit's data from each index in the array
            //singleCommit is a RepositoryCommit struct
            // singleCommit := commits[i]


            //Get the Commit field of the RepositoryCommit struct
            //commitData is a Commmit struct
            commitData := singleCommit.Commit
            //+ flag adds field names
            fmt.Printf("\n%v\n", github.Stringify(commitData.Message))
        }
    }

    
    

}