package datastore

import (
//"net/http"
// "net/http/httptest"
  "os"
  "testing"
	"fmt"
	"log"
	"time"
)

var resetTableAccount = "DELETE from Account;"
var resetTableTask = "DELETE from Task;"

var database = os.Getenv("acs_database")
var databaseUserName = os.Getenv("acs_dbuser")
var databasePassword = os.Getenv("acs_dbpass")
var offlineDatabaseFile = os.Getenv("acs_offlineDbFile")
var onlineDatabaseName = os.Getenv("acs_onlineDbName")
var onlineDatabaseNameTest = os.Getenv("acs_onlineDbNameTest")
var hostName = os.Getenv("acs_hostname")
var hostPort = os.Getenv("acs_hostport")
	
func TestNewAccount(t *testing.T) {
	var email = "abc@abc.com"
	var timezone = 4
	var pass = "abc"
	var firstName = "Jeff"
	var lastName = "Greenwood"
	var middleName = "M."
	account := MakeAccount( &email, timezone, &pass, &firstName, &middleName, &lastName);
	addResult := AddAccount(account)
	if addResult != nil {
		log.Println("addResult failed. " + addResult.Error() )
		t.Fail()
	}
	existsResult := Exists(account)
	if  existsResult != nil {
		log.Println("Exists failed. " + existsResult.Error())
		t.Fail()
	}
}

func TestMatchEmailPassword (t *testing.T) {
	var email = "cde@cde.com"
	var timezone = 4
	var pass = "abc"
	var firstName = "Jeff"
	var lastName = "Greenwood"
	var middleName = "M."
	newAccount := MakeAccount( &email, timezone, &pass, &firstName, &middleName, &lastName);
	addResult := AddAccount(newAccount)
	if addResult != nil {
		log.Println("addResult failed. " + addResult.Error() )
		t.Fail()
	}
	
	var account Account
	account.email = "cde@cde.com"
	account.encryptedPasswordHash = "abc"
	err := MatchEmailPassword(&account)
	if  err != nil {
		log.Println("MatchEmailPassword failed. " + err.Error())
		t.Fail()
	}
	
}

func TestNewTask(t *testing.T) {
	var name = "running"
	var timespent = 0.5
	var email = "abc@abc.com"
	task := new(Task)
	task.taskName = name
	task.timeSpent = timespent
	task.taskDate = time.Now()
	task.email = email
	err := AddTask(task)
	if err != nil {
		log.Println("Add task error: " + err.Error())
		t.Fail()
	}
}

func TestGetTasks (t *testing.T) {

	var name = "running"
	var timespent = 0.5
	var email = "abc@abc.com"
	task := new(Task)
	task.taskName = name
	task.timeSpent = timespent
	task.taskDate = time.Now()
	task.email = email
	AddTask(task)

	//var mapResult *map[string]float64
	//var err error
	mapResult, err := SelectTasksByEmail( "abc@abc.com" )
	if err != nil {
		log.Println("GetTasks error " + err.Error())
		t.Fail()
	}
	if len(mapResult) <= 0 {
		log.Printf("Zero tasks returned. Expected a number greater than 0")
		t.Fail()
	}
	/*if val, ok := (*mapResult)["abc@abc.com"]; ok {
		if val - 0.5 > 0.01 {
			log.Printf("wrong time spent")
			t.Fail()
		}
	} else {
		log.Printf("There is no email field")
		t.Fail()
	}*/
}

func TestMain(m *testing.M) {
	log.SetFlags(log.Llongfile)
	if database=="postgres" && 
    (databaseUserName == "" || 
      onlineDatabaseName == "" || 
      hostName == "" || 
      hostPort == "") {
      fmt.Println("datastore_test is skipped due to insufficient environment variables.")
			os.Exit(1)
  } else if database == ""{
      fmt.Println("unsupported database vendor.")
			os.Exit(1)
  } else if database != "postgres" {
    fmt.Println("insufficient environment variables to run datastore_test. Please set acs_database, acs_dbuser, acs_dbpass, acs_offlineDbFile, acs_onlineDbName, acs_hostname, acs_hostport to the correct values");
		fmt.Printf("database is %s\n", hostName)
		os.Exit(1)
  }
	result := SetUpOrm( &database, &databaseUserName, &databasePassword, &hostName, &hostPort, &offlineDatabaseFile, &onlineDatabaseNameTest)
  if result == false {
    fmt.Println("failed to connect to database.")
		os.Exit(1)
  }
	
	
	defer globalOrm.orm.Close()
	os.Exit(m.Run())
}
