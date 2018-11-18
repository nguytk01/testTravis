package businesslogic

import (
	"acs560_course_project/server/datastore"
)

func AddTask( request *AddTaskRequestData, sessionEmail string) AddTaskRequestResult{
	task := datastore.MakeTask( &(request.TaskName), &(request.TimeSpent), &sessionEmail)
	err := datastore.AddTask(task)
	if err != nil {
		return AddTaskRequestResult{
			AddResult : "Failure",
		}
	} else {
			return AddTaskRequestResult{
			AddResult : "Success",
		}
	}
}

func GetTasks( sessionEmail string ) RetrieveTaskListRequestResult{
	result,err := datastore.SelectTasksByEmail( sessionEmail )
	if err != nil {
		return RetrieveTaskListRequestResult {
			RetrieveTaskListResult: "Failure",
			TaskList: nil,
		}
	} else {
		return RetrieveTaskListRequestResult {
			RetrieveTaskListResult: "Success",
			TaskList: result[:],
		}
	}
}