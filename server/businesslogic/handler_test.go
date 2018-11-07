package businesslogic 
import (
	"net/http"
  "net/http/httptest"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
)

func TestNewUnauthorizedSessionHandler ( t * Testing.T) {
  sessionKeyExpectedLength := 30
  server := httptest.NewServer( func( writer http.ResponseWriter, requestPtr *http.Request ) {
    sessionPtr := getNewUnauthorizedSession()
    jsonByteArr, error := json.Marshal( sessionPtr )

    if ( error != nil ) {
      log.Printf( "%s\n", error.Error() )
      t.Fail("fail")
    }
    io.WriteString ( writer, jsonByteArr )
  });
  
  res, err := http.Get(server.URL)
  if err != nil {
    body := getBody(res)
  } else {
    t.Fail("fail")
  }
  var sessionReceived sessionType
  err = json.Unmarshal( body,  sessionReceived)
  
  if ( len(sessionReceived.Session) < sessionKeyExpectedLength)  {
    t.Fail("Fail");
  }
}

func getBody( requestPtr *http.Request ) []byte {
  var bodyIncompleteReader io.ReadCloser = requestPtr.Body	
	body, error := ioutil.ReadAll( bodyIncompleteReader )
  if ( error != nil ) {
		log.Printf( "error %s\n", error.Error() )
    return nil
	}
  return body[:end]
}
