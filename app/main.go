package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"

	guerrilla "github.com/flashmob/go-guerrilla"
)

var daemon guerrilla.Daemon

func main() {
	daemon = runMtaServer()

	// to send mail from mta server in localhost
	http.HandleFunc("/sendemail", sendEmail)
	http.ListenAndServe(":8081", nil)
}

func runMtaServer() (d guerrilla.Daemon) {
	d = guerrilla.Daemon{}
	_, err := d.LoadConfig("configs/" + os.Getenv("SMTP_CONF"))
	if err != nil {
		fmt.Println("ReadConfig error", err)

	}

	err = d.Start()
	if err != nil {
		fmt.Println("server error", err)
	}

	return
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	server := "127.0.0.1:2525"

	from := q["from"][0]
	to := q["to"][0]

	conn, err := net.Dial("tcp", server)
	if err != nil {
		return
	}
	in := bufio.NewReader(conn)

	fmt.Fprintln(conn, "MAIL FROM:<"+from+">")
	in.ReadString('\n')

	fmt.Fprintln(conn, "RCPT TO:<"+to+">")
	in.ReadString('\n')

	fmt.Fprintln(conn, "DATA")
	in.ReadString('\n')

	fmt.Fprintln(conn, "From: "+from)
	fmt.Fprintln(conn, "To: "+to)
	fmt.Fprintln(conn, "Subject: Test subject")
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, "An email body")
	fmt.Fprintln(conn, ".")
	in.ReadString('\n')

	fmt.Fprint(conn, "QUIT\r\n")
	in.ReadString('\n')
}
