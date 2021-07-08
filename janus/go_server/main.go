package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const recordings_folder = "/www/recording"
	


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func deleteRecording(w http.ResponseWriter, r *http.Request) {
	if r.Method == "delete" {
		lectureId := mux.Vars(r)["lectureId"]

		file_to_delete := fmt.Sprintf("%s.wav", lectureId)

		files, err := ioutil.ReadDir(recordings_folder)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = findFileAndDelete(files, file_to_delete)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "sucess")
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func main() {
	// endpoint to get the files
	fs := http.FileServer(http.Dir(recordings_folder))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// endpoint to delete the lecture
	r := mux.NewRouter()
	r.HandleFunc("/lecture/{lectureId}", deleteRecording)
	r.HandleFunc("/hi", helloWorld)
    http.Handle("/", r)

	http.ListenAndServe(":6111", nil)
}

func findFileAndDelete(files []os.FileInfo, target_name string) error{
	for _, f := range files {
		if f.Name() == target_name{
			fmt.Println(f.Name())
			return nil
		}	
	}
	return fmt.Errorf("file (%s) not found ",target_name)
}
