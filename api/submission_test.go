package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubmission_ToJSON(t *testing.T) {
	submission := Submission{
		SubmissionNumber: "1",
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}
	json := submission.ToJSON()

	assert.Equal(t, `{"SubmissionNumber":"1","DateSubmitted":1506271081,"ModelDescription":"Baseline model","LocalCrossValidationScore":0.6513,"PublicLeaderBoardScore":0.06423,"PrivateLeaderBoardScore":0.0622}`, string(json),
								"Submission JSON marshalling incorrectly.")
}

func TestSubmissionFromJSON(t*testing.T){
	json := []byte(`{"SubmissionNumber":"1","DateSubmitted":1506271081,"ModelDescription":"Baseline model","LocalCrossValidationScore":0.6513,"PublicLeaderBoardScore":0.06423,"PrivateLeaderBoardScore":0.0622}`)
	submission := FromJSON(json)
	assert.Equal(t, Submission{
		SubmissionNumber: "1",
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}, submission, "Submission JSON un-marshalling incorrectly.")
}


func TestAllSubmissions(t *testing.T) {
	submissions := AllSubmissions()
	assert.Len(t, submissions, 2, "Wrong number of submissions.")
}

func TestCreateNewSubmission(t *testing.T) {
	submission := Submission{
		SubmissionNumber: "",
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}
	_, created := CreateSubmission(submission)
	assert.True(t, created, "Submission was not created.")
}


func TestDoNotCreateExistingSubmission(t *testing.T) {
	submission := Submission{SubmissionNumber: "1"}
	_, created := CreateSubmission(submission)
	assert.False(t, created, "Submission was created.")
}


func TestUpdateExistingSubmission(t *testing.T) {
	submission := Submission{
		SubmissionNumber: "",
		DateSubmitted: 1506271081,
		ModelDescription: "Baseline model",
		LocalCrossValidationScore: 0.6513,
		PublicLeaderBoardScore: 0.06423,
		PrivateLeaderBoardScore: 0.0622,
	}
	updated := UpdateSubmission("1", submission)
	assert.True(t, updated, "Submission not updated.")

	submission, _ = GetSubmission("1")
	assert.Equal(t, "", submission.SubmissionNumber, "Submission not updated.")
	assert.Equal(t, "Baseline model", submission.ModelDescription, "Author not updated.")
}

func TestDeleteSubmission(t *testing.T) {
	DeleteSubmission("1")
	assert.Len(t, AllSubmissions(), 2, "Wrong number of submissions after delete.")
}