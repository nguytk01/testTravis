package datastore
import (
	//"database/sql"
	"log"
	"errors"
)

type Account struct{
  email string 
  timeZone int
  encryptedPasswordHash string
  firstName string 
  lastName string
  middleName string
}

var insertAccountString = "INSERT INTO Account (email, timezone, encryptedpasswordhash, firstname, lastname, middlename) values ($1, $2, $3, $4, $5, $6);"
var selectAccountByEmailString = "Select * from Account where email = $1"
var matchEmailPasswordString = "Select * from Account where email = $1 and encryptedpasswordhash = $2"

func MakeAccount(email *string, timezone int, encryptedPasswordHash *string, firstName *string, lastName *string, middleName *string) *Account{
  account := new(Account)
  account.email = *email
  account.timeZone = timezone
  account.encryptedPasswordHash = *encryptedPasswordHash
  account.firstName = *firstName
  account.lastName = *lastName
  account.middleName = *middleName
  return account
}

func AddAccount( account *Account) error{
	_, stmterr := globalOrm.orm.Exec(insertAccountString, account.email, account.timeZone, account.encryptedPasswordHash, account.firstName, account.lastName, account.middleName);
	if stmterr != nil {
		log.Println("statement error : " + stmterr.Error())
	}
	return stmterr
}

func Exists( account *Account) error{
  rows, stmterr := globalOrm.orm.Query(selectAccountByEmailString, account.email);
	if stmterr != nil {
		log.Println("statement error : " + stmterr.Error())
	} else {
		rowsNext := rows.Next()
		if rowsNext == false {
			return errors.New("This account does not exist.")
		} else {
			return nil
		}
	}
	return stmterr
}

func MatchEmailPassword( account *Account) error{
  rows, stmterr := globalOrm.orm.Query(matchEmailPasswordString, account.email, account.encryptedPasswordHash);
	if stmterr != nil {
		log.Println("statement error : " + stmterr.Error())
	} else {
		rowsNext := rows.Next()
		if rowsNext == false {
			return errors.New("This account does not exist.")
		} else {
			return nil
		}
	}
	return stmterr
}