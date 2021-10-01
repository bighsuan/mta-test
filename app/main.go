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
	daemon = runServer()

	http.HandleFunc("/", handler)
	http.HandleFunc("/sendemail", sendEmail)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8081", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hello!")
}

func runServer() (d guerrilla.Daemon) {

	d = guerrilla.Daemon{}
	_, err := d.LoadConfig("../configs/" + os.Getenv("SMTP_CONF"))
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

	in.ReadString('\n')

	fmt.Fprint(conn, "HELO example.com\r\n")
	in.ReadString('\n')

	fmt.Fprint(conn, "MAIL FROM:<"+from+">\r\n")

	in.ReadString('\n')

	fmt.Fprint(conn, "RCPT TO:<"+to+">\r\n")

	in.ReadString('\n')

	fmt.Fprint(conn, "DATA\r\n")

	in.ReadString('\n')

	fmt.Fprint(conn, "Subject: Test subject\r\n")
	fmt.Fprint(conn, "A an email body\r\n")
	fmt.Fprint(conn, ".\r\n")

	in.ReadString('\n')

	// fmt.Fprint(conn, "QUIT\r\n")
	// in.ReadString('\n')

}

func shutdown(w http.ResponseWriter, r *http.Request) {
	daemon.Shutdown()
}
