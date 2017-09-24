package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubmission_ToJSON(t *testing.T) {
	submission := Submission{
		SubmissionNumber: 1,
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}
	json := submission.ToJSON()

	assert.Equal(t, `{"SubmissionNumber":1,"DateSubmitted":651387237,"ModelDescription":"Baseline model","LocalCrossValidationScore":0.6513,"PublicLeaderBoardScore":0.06423,"PrivateLeaderBoardScore":0.0622}`, string(json),
								"Submission JSON marshalling incorrectly.")
}

func TestSubmissionFromJSON(t*testing.T){
	json := []byte(`{"SubmissionNumber":1,"DateSubmitted":651387237,"ModelDescription":"Baseline model","LocalCrossValidationScore":0.6513,"PublicLeaderBoardScore":0.06423,"PrivateLeaderBoardScore":0.0622}`)
	submission := FromJSON(json)
	assert.Equal(t, Submission{
		SubmissionNumber: 1,
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}, submission, "Submission JSON un-marshalling incorrectly.")
}