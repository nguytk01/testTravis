package businesslogic
import (
	"time"
	"acs560_course_project/server/datastore"

)
type LoginRequestData struct {
  Email string
	Password string
	TimeZone int
  SessionKey string
}

type LoginRequestResult struct{
  LoginResult string
  SessionKey string
}

type CreateAccountRequestData struct {
  FirstName string
  MiddleName string
  LastName string
  Email string
  Password string
	TimeZone int
  SessionKey string
}

type CreateAccountRequestResult struct {
  CreateAccountResult string
  SessionKey string
}

type AddTaskRequestData struct {
	TaskName string
	TimeSpent float64
	TaskDate time.Time
	SessionKey string
}

type AddTaskRequestResult struct {
	AddResult string
}

type RetrieveTaskListRequestData struct {
	SessionKey string
}

type RetrieveTaskListRequestResult struct {
	RetrieveTaskListResult string
	TaskList []datastore.TaskFromDb
}