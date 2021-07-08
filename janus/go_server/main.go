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

// accept UUID and UUID.wav
func IsValidFile(uuid string) bool {
	// TODO support diffrent extentions
    r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}(\.wav)?$`)
    return r.MatchString(uuid)
}

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UUID := r.URL.String()
		
		isValidRequest := IsValidFile(UUID) || len(UUID) == 0
		if !isValidRequest {
			w.WriteHeader(http.StatusForbidden)
			
			return 
		}
        
		if r.Method == "DELETE" || r.Method == "delete" {
			if !IsValidFile(UUID){
				fmt.Fprint(w,"delete query must be UUID")
			}
            err := os.Remove(file2Del)
			if err != nil {
				log.Print(err)
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

