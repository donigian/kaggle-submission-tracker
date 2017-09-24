package api

import (
	"encoding/json"
	"net/http"
)

type Submission struct {
	// specify how JSON marshalling should be done
	SubmissionNumber int `json:SubmissionNumber`
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

var Submissions = []Submission{
	Submission{
		SubmissionNumber: 1,
		DateSubmitted: 651387237,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	},
	Submission{
		SubmissionNumber: 2,
		DateSubmitted: 651387250,
		ModelDescription: "Model trained with XGBoost, no feature engineering, no tuning",
		LocalCrossValidationScore: 0.6611,
		PublicLeaderBoardScore: 0.06623,
		PrivateLeaderBoardScore: 0.0647,
	},
}

func SubmissionHandleFunc(w http.ResponseWriter, r *http.Request){
	s, err := json.Marshal(Submissions)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(s)
}