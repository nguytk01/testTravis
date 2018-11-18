package reception

var sessionManager map[string]string = make(map[string]string)

func authorizeSession( SessionKey string, Email string ) {
	sessionManager[SessionKey] = Email
}


func getEmailBySession( sessionKey string ) string{
	if val, ok := sessionManager[sessionKey]; ok {
		return val
	} else {
		return ""
	}
}

//TODO: remove expired session from session manager