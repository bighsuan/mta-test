package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"

	guerrilla "github.com/flashmob/go-guerrilla"
)

var daemon1, daemon2 guerrilla.Daemon

func main() {
	daemon1, daemon2 = runServer()

	http.HandleFunc("/", handler)
	http.HandleFunc("/sendemail", sendEmail)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8081", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hello!")
}

func runServer() (daemon1 guerrilla.Daemon, daemon2 guerrilla.Daemon) {

	daemon1 = createServer(&guerrilla.AppConfig{
		LogFile:      "./logs/server3",
		AllowedHosts: []string{"."},
		Servers: []guerrilla.ServerConfig{
			{
				ListenInterface: "127.0.0.3:2525",
				Hostname:        "example3.com",
				IsEnabled:       true,
			},
		},
	})

	daemon2 = createServer(&guerrilla.AppConfig{
		LogFile:      "./logs/server4",
		AllowedHosts: []string{"."},
		Servers: []guerrilla.ServerConfig{
			{
				ListenInterface: "127.0.0.4:2526",
				Hostname:        "example4.com",
				IsEnabled:       true,
			},
		},
	})

	return
}

func createServer(config *guerrilla.AppConfig) (d guerrilla.Daemon) {

	d = guerrilla.Daemon{Config: config}

	err := d.Start()
	if err != nil {
		fmt.Println("start error", err)
	}

	return
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	server := q["server"][0]
	from := q["from"][0]
	to := q["to"][0]

	conn, err := net.Dial("tcp", server)
	if err != nil {
		return
	}
	in := bufio.NewReader(conn)

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
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	daemon1.Shutdown()
	daemon2.Shutdown()
}
