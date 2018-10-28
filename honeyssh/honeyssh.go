package honeyssh

import (
    "github.com/gliderlabs/ssh"
    // gossh "golang.org/x/crypto/ssh"
    "io"
    "log"
    "os"
)

type LoginAttempt struct {
	User string `json:user`
	IpAddress string `json:ipaddress`
	// Command []string `json:command`
	// PubKey ssh.PublicKey `json:pubkey`
}

func (la *LoginAttempt) toCSV() string {
	line := la.User
	line += "," + la.IpAddress
	// line += "," + gossh.MarshalAuthorizedKey(la.PubKey)
	line += "\n"
	return line
}

func capture(s ssh.Session) {
	logger(s)
	io.WriteString(s, "Connection closed.\n")
}

func logger(s ssh.Session) {
	la := &LoginAttempt{
		User: s.User(),
		IpAddress: s.RemoteAddr().String(),
		// Command: s.Command(),
		// PubKey: s.PublicKey(),
	}
	
	f, err := os.OpenFile("attempts.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(la.toCSV()); err != nil {
    	panic(err)
	}

	log.Println(la.toCSV())
}

func StartServer() {
	ssh.Handle(capture)
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}