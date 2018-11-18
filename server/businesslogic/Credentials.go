package businesslogic

import (
	"acs560_course_project/server/datastore"
	"errors"
	"encoding/json"
	"log"
)

func LogUserIn( body []byte ) (*LoginRequestResult, error, string) {
	var request LoginRequestData
	err := json.Unmarshal( body[:],  &request)
  if err != nil {
    log.Println(err.Error())
		return &LoginRequestResult {
			LoginResult: "Incompatible JSON request structure.",
			SessionKey: request.SessionKey,
		}, err, ""
  }
	var space string = ""
	account := datastore.MakeAccount( &request.Email, request.TimeZone, &request.Password, &space, &space, &space)
	valid := accountValid(account)
	if valid == true {
		return &LoginRequestResult {
			LoginResult: "Success",
			SessionKey: "",
		}, nil, request.Email
	} else {
		return &LoginRequestResult{
			LoginResult: "Failure",
			SessionKey: request.SessionKey,
		}, errors.New("Invalid email/password."), ""
	}
}

func CreateAccount( body[] byte ) (*CreateAccountRequestResult, error, string) {
	var request CreateAccountRequestData
	err := json.Unmarshal( body[:],  &request)
  if err != nil {
    log.Println(err.Error())
		return &CreateAccountRequestResult {
			CreateAccountResult: "Incompatible JSON request structure.",
			SessionKey: request.SessionKey,
		}, err, ""
  }
		var space string = ""

	account := datastore.MakeAccount( &request.Email, request.TimeZone, &request.Password, &space, &space, &space)
	exists := accountExists( account )
	if  exists == true{
		return &CreateAccountRequestResult {
			CreateAccountResult: "Account already exists.",
			SessionKey: request.SessionKey,
		}, errors.New("Account already exists."), ""
	} else {
		datastore.AddAccount(account)
		return &CreateAccountRequestResult {
			CreateAccountResult: "Success",
			SessionKey: "",
		}, nil, request.Email
	}
}

func accountExists( account *datastore.Account ) bool {
	result := datastore.Exists(account)
	if result == nil {
		return true
	} else {
		return false
	}
}

func accountValid( account * datastore.Account ) bool {
	result := datastore.MatchEmailPassword(account)
	if result == nil {
		return true
	} else {
		return false
	}
}

