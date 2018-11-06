package businesslogic
import (
	"math/rand"
)

const sessionStringLength = 30

type AnonymizedSession struct{
	Session sessionType
}

type IncomingSession struct{
	Session sessionType
	UserId int
	LimitedTime float32
}


type internalSession struct{
	Session sessionType
	UserId int
	LimitedTime float32
}

// session can be string or int. Give it a type at the moment so it can be changed
// later
// The session should be sessionStringLength bytes long.
type sessionType struct {
	Session []byte
}

func getNewUnauthorizedSession() *sessionType{
	newSession := new(sessionType)
	newSession.Session = generateUniqueSessionId()
	//log.Printf( " inside session.go %s\n", string( newSession.Session ) )
	return newSession
}

func generateUniqueSessionId() []byte {
	var repository []byte  = []byte ( "abcdefghijklmopqrstuvwxyz0123456789" )
	var result [ sessionStringLength ] byte

	for i := 0; i < sessionStringLength; i++ {
		result [ i ] = repository [ rand.Intn ( 36 ) ]
	}
	return result [:]
}
