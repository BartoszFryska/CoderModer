package handlers

import (
	"fmt"
	"net/http"
)

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Job submitted!")
}

func JobStatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Job status queried!")
}

func ListJobsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of jobs!")
}

func CancelJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Job canceled!")
}

func ExecuteJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Job executed!")
}
