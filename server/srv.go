package main

import (
	"acs560_course_project/server/reception"
  "acs560_course_project/server/datastore"
  "flag"
  "log"
  "net/http"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile)
	parseCmdFlagsSetUpOrm()
  http.HandleFunc("/", reception.Hey)
	http.HandleFunc("/newUnauthorizedSession", reception.NewUnauthorizedSessionHandler)
	http.HandleFunc("/newAuthorizedSession", reception.NewAuthorizedSessionHandler)
  http.HandleFunc("/Login", reception.LogUserIn)
  http.HandleFunc("/CreateAccount", reception.CreateAccountHandler)
  http.HandleFunc("/AddTask", reception.AddTaskHandler)
	http.HandleFunc("/GetAllTasks", reception.GetAllTasksHandler)

	
	http.ListenAndServe(":8000", nil);
  
}

func parseCmdFlagsSetUpOrm(){
  var database = flag.String("database", "postgres", "the database to use (currently supported: sqlite, mysql)")
  var databaseUserName = flag.String("dbuser", "postgres", "database's user")
  var databasePassword = flag.String("dbpassword", "123", "database's password")
	
	if value, ok := os.LookupEnv("PGPASSWORD"); ok {
		databasePassword = &value
	} 
  var offlineDatabaseFile = flag.String("sqliteDbPath", ".\\data.db", "sqlite database file's path")
  var onlineDatabaseName = flag.String("dbname", "sandbox", "database name in the online datastore")
  var hostName = flag.String("ipaddr", "127.0.0.1", "database server name")
  var hostPort = flag.String("port", "5432", "database server port")
	
  flag.Parse()
  
  var result = datastore.SetUpOrm(database, databaseUserName, databasePassword, hostName, hostPort, offlineDatabaseFile, onlineDatabaseName)
  log.Println("Setting up Orm...")
  if result == false {
    panic("Error while connecting to the database. Program exits.")
  } else {
    log.Println("Database: online.")
  }
}
