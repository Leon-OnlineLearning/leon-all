//inspired by https://gist.github.com/toravir/aeb9c2923cd180a71bd15203d1aea4f2
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const recordings_folder = "/www/recording"

func statusCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func deleteChecker (dirprefix string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "DELETE" {
            file2Del := dirprefix+r.URL.String()
            err := os.Remove(file2Del)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusNotFound)
				return
			}
			fmt.Printf("file (%s) deleted:", file2Del)
			fmt.Fprintf(w, "sucess")
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", statusCheck)
	
	// staticly serving file with support to delete
	fs := http.FileServer(http.Dir(recordings_folder))
	mux.Handle("/",deleteChecker(recordings_folder, fs))
	
	http.ListenAndServe(":6111", mux)
}

