package businesslogic 
import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
  "fmt"
)

func Hey ( writer http.ResponseWriter, requestPtr *http.Request ) {
	io.WriteString ( writer, "Test" )
}

func NewUnauthorizedSessionHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {	
	sessionPtr := getNewUnauthorizedSession()
	jsonByteArr, error := json.Marshal( sessionPtr )

	if ( error != nil ) {
		log.Printf( "%s\n", error.Error() )
	}
	writer.Write ( jsonByteArr )
}

func NewAuthorizedSessionHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {
	var bodyIncompleteReader io.ReadCloser = requestPtr.Body	
	body, error := ioutil.ReadAll( bodyIncompleteReader )
	var incomingSession IncomingSession
	error = json.Unmarshal( body, incomingSession )
	if ( error != nil ) {
		log.Printf( "error %s\n", error.Error() );
	}
}

func NewUserHandler ( writer http.ResponseWriter, requestPtr *http.Request ) {
		io.WriteString ( writer, "Test" )
}

func LogUserIn ( writer http.ResponseWriter, requestPtr *http.Request ) {
  fmt.Println("Log User In")
  //requestPtr.parseForm()
  
  fmt.Println(requestPtr.Body)
  fmt.Println()
  var bodyIncompleteReader io.ReadCloser = requestPtr.Body	
	body, error := ioutil.ReadAll( bodyIncompleteReader )
  fmt.Printf("len is %d\n",len(body))
  if  error != nil  {
		log.Printf( "error %s\n", error.Error() );
	}
  /*for i := 0; i < len(body); i++{
    fmt.Printf("%c", body[i])
  }
  fmt.Println("")
  var incomingUserLoginBox userLoginRequest
  //error = json.Unmarshal( body, incomingUserLoginBox )
  fmt.Println("error" + error.Error())
  log.Println("hello" + incomingUserLoginBox.UserName);*/
  
  
  loginResultPtr := new(LoginResult)
  loginResultPtr.LoginResult = "Success"
  jsonByteArr, error := json.Marshal( loginResultPtr )
  writer.Write ( jsonByteArr )
} 
