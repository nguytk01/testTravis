package datastore
import (
	"log"
	"errors"
	"time"
)
type Task struct {
  taskId []byte
  taskName string
  timeSpent float64
	taskDate time.Time
	email string
}

type TaskFromDb struct{
	TaskName string
	TimeSpent float64
	TaskDate time.Time
}

var insertTaskString = "INSERT INTO Task (taskId, taskName, timeSpent, taskDate, email) values (uuid_generate_v4(),$1, $2, $3, $4);"
var selectTasksByEmailString = "Select taskname, timespent, taskDate from Task where email = $1"
var selectCountTasksByEmail = "Select Count(*) as count from Task where email = $1"

func MakeTask( name *string, duration *float64, email *string)  *Task {
	task := new(Task)
	task.taskName = *name
	task.timeSpent = *duration
	task.email = *email
	return task
}

func AddTask( task *Task) error {
	_, stmterr := globalOrm.orm.Exec(insertTaskString, task.taskName, task.timeSpent, task.taskDate, task.email);
	if stmterr != nil {
		log.Println("statement error : " + stmterr.Error())
	}
	return stmterr
}

func SelectTasksByEmail( email string ) ([](TaskFromDb), error) {
	countRows, counterr := globalOrm.orm.Query(selectCountTasksByEmail, email)
	if (counterr != nil) {
		log.Println("statement error : " + counterr.Error())
		return nil, counterr
	}
	if countRows.Next() {
		
	} else {
		log.Println("Zero rows.")
		return nil, errors.New("No rows")
	}
	var rowsCount int
	countRows.Scan(&rowsCount)
	
	rows, stmterr := globalOrm.orm.Query(selectTasksByEmailString, email)
	if stmterr != nil {
		log.Println("statement error : " + stmterr.Error())
			return nil, stmterr
	} else {
			taskMap := make( [](TaskFromDb), rowsCount)
			taskMapCounter := 0
			var taskName string
			var timeSpent float64
			var taskDate time.Time
			var taskFromDbPtr *TaskFromDb
			for rows.Next()  {
				rows.Scan( &taskName, &timeSpent, &taskDate)
				taskFromDbPtr = new(TaskFromDb)
				taskMap[taskMapCounter] = *taskFromDbPtr 
				taskMap[taskMapCounter].TaskName = taskName
				taskMap[taskMapCounter].TimeSpent = timeSpent
				taskMap[taskMapCounter].TaskDate = taskDate
				taskMapCounter++
			}
			return taskMap[:], nil
		}
}
