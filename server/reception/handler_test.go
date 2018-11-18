package reception 
import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewUnauthorizedSessionHandler(t *testing.T) {
	flag.Parse()
	sessionKeyExpectedLength := 30
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, requestPtr *http.Request) {
		sessionPtr := getNewUnauthorizedSession()
		jsonByteArr, error := json.Marshal(sessionPtr)

		if error != nil {
			log.Printf("%s\n", error.Error())
			t.Fail()
		}
		writer.Write(jsonByteArr)
	}))

	res, err := http.Get(server.URL)
	var body []byte
	if err == nil {
		tempArr := getBody(res)
		body = make([]byte, len(tempArr))
		copy(body[:], tempArr)
	} else {
		t.Fail()
	}
	var sessionReceived sessionType
	err = json.Unmarshal(body[:], &sessionReceived)
	if err != nil {
		log.Println(err.Error())
	}
	if len(sessionReceived.Session) < sessionKeyExpectedLength {
		log.Printf(" Key length is : %d not long enough key.", len(sessionReceived.Session))
		t.Fail()
	}
}

func getBody(requestPtr *http.Response) []byte {
	var bodyIncompleteReader io.ReadCloser = requestPtr.Body
	body, error := ioutil.ReadAll(bodyIncompleteReader)
	if error != nil {
		log.Printf("error %s\n", error.Error())
		return nil
	}
	returnedBody := make([]byte, len(body))
	copy(returnedBody, body[:])
	return returnedBody[:]
}
