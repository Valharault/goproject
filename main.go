package main

import (
    "fmt"
    "net/http"
    "time"
    "os"
    "bufio"
    "strings"
)


func timeHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case http.MethodGet:
        currentTime := time.Now()
        fmt.Fprintf(w, fmt.Sprintf("%s",fmt.Sprintf(currentTime.Format("15h04"))))
    }
}

func addHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case http.MethodPost:
             if err := req.ParseForm(); err != nil {
                 fmt.Println("Something went bad")
                 fmt.Fprintln(w, "Something went bad")
                 return
             }
             newEntry := make(map[string]string)
             for key, value := range req.PostForm {
                if (key == "author") {
                    newEntry["author"] = value[0]
                } else if (key == "entry") {
                    newEntry["entry"] = value[0]
                }
             }

             fmt.Fprintf(w, "%v: %s\n", newEntry["author"], newEntry["entry"])
             save(newEntry)
    }
}

func getEntriesHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case http.MethodGet:
             if err := req.ParseForm(); err != nil {
                 fmt.Println("Something went bad")
                 fmt.Fprintln(w, "Something went bad")
                 return
             }
             file, err := os.Open("./texte.txt")
             if err == nil {
                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                    if len(strings.Split(scanner.Text(), ":")) > 1 {
                        fmt.Fprintf(w, "%s\n", strings.Split(scanner.Text(), ":")[1])
                   }
                }
             }
    }
}

func save(newEntry map[string]string) {
    fmt.Println(newEntry)
    saveFile, err := os.OpenFile("./texte.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
    defer saveFile.Close()

    w := bufio.NewWriter(saveFile)
    if err == nil {
        fmt.Fprintf(w, "%s:%s\n", newEntry["author"], newEntry["entry"])
    }
    w.Flush()
}

func main() {
 http.HandleFunc("/", timeHandler)
 http.HandleFunc("/add", addHandler)
 http.HandleFunc("/entries", getEntriesHandler)
 http.ListenAndServe(":4567", nil)
}
