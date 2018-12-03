package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

type GitLog struct {
	Commit  string
	Merge   string
	Author  string
	Date    string
	Message string
}

const (
	ServicePort = ":5000"
	apiVersion  = "/api/v1"
)

var datas []GitLog
var gitdir string
var logfile string

func handleRequests() {
	r := mux.NewRouter()
	myRouter := r.PathPrefix(apiVersion).Subrouter()
	myRouter.Path("/log").HandlerFunc(jsonToweb)
	fmt.Println("Start..")
	log.Fatal(http.ListenAndServe(ServicePort, myRouter))
}

func jsonToweb(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(datas)
}

func getGitLog() {
	cmd := exec.Command("git", "--git-dir", gitdir+"/.git", "log")
	outfile, err := os.Create(logfile)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}

func readFile() {
	file, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tmpLst []string
	var commitFlag bool
	var mergeFlag bool

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		if line != "" {
			s := strings.Split(line, " ")
			// fmt.Println(s)

			var tmpCommit string
			var tmpMerge string
			var tmpAuthor string
			var tmpDate string
			var tmpMessage string

			switch {
			case s[0] == "commit":
				tmpCommit = s[1]
				tmpLst = append(tmpLst, tmpCommit)
				commitFlag = true

			case s[0] == "Merge:" && commitFlag == true:
				tmpMerge = strings.Join(s[1:], " ")
				tmpLst = append(tmpLst, tmpMerge)
				mergeFlag = true

			case s[0] == "Author:" && commitFlag == true:
				tmpAuthor = strings.Join(s[1:], " ")
				tmpLst = append(tmpLst, tmpAuthor)

			case s[0] == "Date:" && commitFlag == true:
				tmpDate = strings.Join(s[1:], " ")
				tmpLst = append(tmpLst, tmpDate)

			// example: [    Merge branch 'regexp_destroy_customer' into 'master']
			case s[0] == "" && commitFlag == true:
				tmpMessage = strings.Join(s[0:], " ")
				tmpLst = append(tmpLst, strings.TrimSpace(tmpMessage))
				commitFlag = false
			}

			// fmt.Println(tmpLst)
			// fmt.Println(len(tmpLst))

			switch {
			case mergeFlag == true && len(tmpLst) == 5:
				setDataMerge(tmpLst)
				tmpLst = nil
				mergeFlag = false
			case mergeFlag == false && len(tmpLst) == 4:
				setDataNoMerge(tmpLst)
				tmpLst = nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func setDataMerge(input []string) {
	datas = append(datas, GitLog{
		Commit:  input[0],
		Merge:   input[1],
		Author:  input[2],
		Date:    input[3],
		Message: input[4],
	})
}

func setDataNoMerge(input []string) {
	datas = append(datas, GitLog{
		Commit:  input[0],
		Merge:   "",
		Author:  input[1],
		Date:    input[2],
		Message: input[3],
	})
}

func getFinalJson() {
	// todo: anonymous struct
	// data [{},{}]
	js, err := json.Marshal(datas)
	if err != nil {
		fmt.Println("error:", err)
	}
	getjson := fmt.Sprintf("%s", js)
	fmt.Println(getjson)
}

func main() {
	if len(os.Args) > 1 {
		gitdir = os.Args[1]
		logfile = "/tmp/gitlog.log"
		getGitLog()
	} else {
		// logfile = "logSMALL.log"
		logfile = "gitlog.log"
	}
	readFile()
	getFinalJson()
	handleRequests()
}
