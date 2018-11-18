package datastore

import (
  "log"
	"database/sql"
	_ "github.com/lib/pq"
)

type datastore struct{
  orm *sql.DB
}

var globalOrm *datastore = nil
func SetUpOrm(databaseVendor *string, 
					userName *string, password *string, 
					hostName *string, hostPort *string,
					offlineDatabasePath *string, onlineDbName *string) bool {
	if globalOrm != nil {
		return true
	}
  var err error
  globalOrm = new(datastore)
  if *databaseVendor == "sqlite"{
    //globalOrm.orm, err = gorm.Open(*databaseVendor, *offlineDatabasePath)
  } else if *databaseVendor == "mysql"{
    var connectionString = *userName + ":" + *password + "@/" + *onlineDbName + "?charset=utf8&parseTime=True&loc=Local"
    log.Println(connectionString)
    //globalOrm.orm, err = gorm.Open(*databaseVendor, connectionString)
  } else if *databaseVendor == "postgres" {
    var connectionString = "host=" + *hostName + " port=" + *hostPort + " user=" + *userName + " dbname=" + *onlineDbName + " password=" + *password + " sslmode=disable"
    log.Println(connectionString)
		globalOrm.orm, err = sql.Open( "postgres", connectionString)
  }
  
  if err != nil {
    log.Println("error while connecting to the database: " + err.Error())
    return false
  }
	
  return true
}
