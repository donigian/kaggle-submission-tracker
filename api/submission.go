package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Submission struct {
	// specify how JSON marshalling should be done
	SubmissionNumber string `json:SubmissionNumber`
	DateSubmitted int64 `json:DateSubmitted`
	ModelDescription string `json:ModelDescription`
	LocalCrossValidationScore float32 `json:LocalCrossValidationScore`
	PublicLeaderBoardScore float32 `json:PublicLeaderBoardScore`
	PrivateLeaderBoardScore float32 `json:PrivateLeaderBoardScore`
}

func (s Submission) ToJSON()[]byte {
	ToJSON, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Submission {
	submission := Submission{}
	err := json.Unmarshal(data, &submission)
	if err != nil {
		panic(err)
	}
	return submission
}

var submissions = map[string]Submission{
	"1": Submission{
		SubmissionNumber: "1",
		DateSubmitted: 651387237,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	},
	"2": Submission{
		SubmissionNumber: "2",
		DateSubmitted: 651387250,
		ModelDescription: "Model trained with XGBoost, no feature engineering, no tuning",
		LocalCrossValidationScore: 0.6611,
		PublicLeaderBoardScore: 0.06623,
		PrivateLeaderBoardScore: 0.0647,
	},
}

func AllSubmissions() []Submission {
	values := make([]Submission, len(submissions))
	idx := 0
	for _, submission := range submissions {
		values[idx] = submission
		idx++
	}
	return values
}

func SubmissionsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		submissions := AllSubmissions()
		writeJSON(w, submissions)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		submission := FromJSON(body)
		submissionNumber, created := CreateSubmission(submission)
		if created {
			w.Header().Add("Location", "/api/submissions/"+submissionNumber)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func SubmissionHandleFunc(w http.ResponseWriter, r *http.Request){
	submissionNumber := r.URL.Path[len("/api/submissions/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		submission, found := GetSubmission(submissionNumber)
		if found {
			writeJSON(w, submission)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		submission := FromJSON(body)
		exists := UpdateSubmission(submissionNumber, submission)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteSubmission(submissionNumber)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// GetSubmission returns the submission for a given submission number
func GetSubmission(submissionNumber string) (Submission, bool) {
	submission, found := submissions[submissionNumber]
	return submission, found
}

// CreateSubmission creates a new Submission if it does not exist
func CreateSubmission(submission Submission) (string, bool) {
	_, exists := submissions[submission.SubmissionNumber]
	if exists {
		return "", false
	}
	submissions[submission.SubmissionNumber] = submission
	return submission.SubmissionNumber, true
}

// UpdateSubmission updates an existing submission
func UpdateSubmission(submissionNumber string, submission Submission) bool {
	_, exists := submissions[submissionNumber]
	if exists {
		submissions[submissionNumber] = submission
	}
	return exists
}

// DeleteSubmission removes a submission from the map by submission id
func DeleteSubmission(submissionNumber string) {
	delete(submissions, submissionNumber)
}