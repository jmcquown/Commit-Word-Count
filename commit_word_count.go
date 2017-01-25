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
    "sort"
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
    wordCountMap := wordCount(messages)
    sortedMap := sortWordCount(wordCountMap)
    //Iterate through the PairList
    for _, value := range sortedMap {
        fmt.Printf("%v:\t%v\n", value.Key, value.Value)
    }
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


//http://stackoverflow.com/a/18695740
//Returns a map that is sorted by value
func sortWordCount(counts map[string]int) PairList {
    p := make(PairList, len(counts))
    index := 0
    for key, value := range counts {
        p[index] = Pair{key, value}
        index++
    }

    sort.Sort(sort.Reverse(p))
    return p
}

//All of the below is necessary in order to sort the wordcount by value
type Pair struct {
    Key string
    Value int
}

type PairList []Pair

func (p PairList) Len() int {return len(p)}
func (p PairList) Less(i, j int) bool {return p[i].Value < p[j].Value}
func (p PairList) Swap(i, j int) {p[i], p[j] = p[j], p[i]}



