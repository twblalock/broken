package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
)

func main() {
    start := time.Now()

    // TODO there must be a decent flag parser or something

    if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
        fmt.Println("Usage: broken port_number")
	fmt.Println("Valid params:")
	fmt.Println("\tstatus=status_code: responds with status_code")
	fmt.Println("\tsleep=sleep_time: responds after sleep_time milliseconds")
        fmt.Println("\tbody=body: response includes body")
	fmt.Println("\t/up: responds with uptime in milliseconds")
	os.Exit(1)
    }

    if len(os.Args) > 2 {
        logfile, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
        if err != nil {
            log.Fatalf("failed to open log file at %s: error: %s\n", os.Args[2], err)
        } else {
            log.SetOutput(logfile)
        }
    }

    log.Printf("START on port: %s\n", os.Args[1])

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(r)
	if err != nil {
	    log.Println(err)
        } else {
	    log.Println(string(j))
        }

	query := r.URL.Query()
	status := 200
	sleep := 0
	body := query.Get("body")

	if len(query["status"]) != 0 {
	    parsedStatus, err := strconv.Atoi(query.Get("status"))
	    if nil != err {
	        fmt.Println(err)
	        http.Error(w, fmt.Sprint(err), 500)
		return
	    } else {
	        status = parsedStatus
	    }
	}

	if len(query["sleep"]) != 0 {
            parsedSleep, err := strconv.Atoi(query.Get("sleep"))
	    if nil != err {
	        fmt.Println(err)
	        http.Error(w, fmt.Sprint(err), 500)
		return
	    } else {
		sleep = parsedSleep
            }
        }

	time.Sleep(time.Duration(sleep) * time.Millisecond)
	http.Error(w, body, status)
	return
    })

    http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL)
	fmt.Fprintf(w, "%d\n", time.Since(start) / time.Millisecond)
    })

    http.ListenAndServe(":" + os.Args[1], nil)
}
