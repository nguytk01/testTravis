package businesslogic 
import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
)

func Hey ( writer http.ResponseWriter, requestPtr *http.Request ) {
	io.WriteString ( writer, "Test" )
}

func TestnewUnauthorizedSessionHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {	
	sessionPtr := getNewUnauthorizedSession()
	jsonByteArr, error := json.Marshal( sessionPtr )

	if ( error != nil ) {
		log.Printf( "%s\n", error.Error() )
	}
	writer.Write ( jsonByteArr )
}

func newAuthorizedSessionHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {
	var bodyIncompleteReader io.ReadCloser = requestPtr.Body	
	body, error := ioutil.ReadAll( bodyIncompleteReader )
	var incomingSession IncomingSession
	error = json.Unmarshal( body, incomingSession )
	if ( error != nil ) {
		log.Printf( "error %s\n", error.Error() );
	}
}

func newUserHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {
		io.WriteString ( writer, "Test" )
}
