# Kaggle Submission Tracker
Competitive machine learning competition sites (like [Kaggle](kaggle.com)) require that you perform several iterations of model train, hyper-parameter tuning during the model build process. From my experience, it's profoundly important to be disciplined as you keep track of each experiment you perform and the results you observe. 

Here's an example of how I keep track of my experiments:


Submission Number | Date Submitted (Unix epoch) | Model Description       | Local CV (Std Dev)        |  Public LeaderBoard   |   Private LeaderBoard   |  
|---|---| ------------- |:-------------:| -----:| -----:|
|1|1506271081 | Baseline: GBM w/ 1st order feature engineering, no tuning  | 0.705219 (0.064116)| - | - |
|2|1506272081 | XGBoost w/ 1st order feature engineering, no tuning    | 0.754843 (0.060699) | - | - |
|3|1506273081 | Ridge w/ 1st order feature engineering, no tuning  | 0.729557 (0.052752) | - | - |
|4|1506281081 | LDA w/ 1st order feature engineering   | 0.724856 (0.056668)  | - | - |
|5|1506291081 | XGBoost w/ 1st order feature engineering & tuning   |0.759926 | - | - |
|6|1506272081 | GBM w/ 1st order feature engineering & tuning |  0.724743 | - | - |
|7|1506273081 | Ridge w/ 1st order feature engineering & tuning |  0.730535  | - | - |
|8|1506274081 | **Ensemble XGBoost, LDA, GBM**   | 		**.7900** | - | - |
|9|1506275081 | Ridge w/ 2nd order feature engineering, no tuning  | 0.728539 (0.052660) | - | - |
|10|1506277081 | XGBoost w/ 2nd order feature engineering, no tuning    | 0.746104 (0.057687) | - | - |
|11|1506278081 | GBM w/ 2nd order feature engineering, no tuning   | 0.702764 (0.046115) | - | - |
|12|1506279081 | XGBoost w/ 2nd order feature engineering & tuning   |0.752732 | - | - |
|13|1506281081 | GBM w/ 2nd order feature engineering & tuning |  0.715590 | - | - |
|14|1506291081 | Ridge w/ 2nd order feature engineering & tuning |  0.730804  | - | - |
|15|1506371081 | Ensemble XGBoost, LDA, GBM w/ 2nd order feature engineering  | 		.7725 | - | - |


This project is a Kaggle Submission tracker written in Go. 

To build the project:
> `go build`

To run tests:
> `go test ./api -v`

Here's a list of endpoint available:

### Health endpoint for microservice
```
GET http://localhost:8080/api/status
```

### Get list of all submissions
```
GET http://localhost:8080/api/submissions
```

### Get specific submission by submission number
```
GET http://localhost:8080/api/submissions/3
```

### Create a new submission
```
POST http://localhost:8080/api/submissions
{
  		"SubmissionNumber": "3",
  		"DateSubmitted": 651387237,
  		"ModelDescription": "Baseline model",
  		"LocalCrossValidationScore": 0.6513,
  		"PublicLeaderBoardScore": 0.06423,
  		"PrivateLeaderBoardScore": 0.0622
}
```

### Modify existing submission
```
PUT http://localhost:8080/api/submissions
{
  		"SubmissionNumber": "3",
  		"DateSubmitted": 651387237,
  		"ModelDescription": "Modified baseline model",
  		"LocalCrossValidationScore": 0.6513,
  		"PublicLeaderBoardScore": 0.06423,
  		"PrivateLeaderBoardScore": 0.0622
}
```

### Delete a previous submission
```
DELETE http://localhost:8080/api/submissions/3
```