package notify

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail() {
	host := "smtp.163.com:25"
	auth := smtp.PlainAuth("", "apetonweb@163.com", "XFNKJNUQYHYRSAQC", "smtp.163.com")
	contentType := "Content-Type: text/plain; charset=UTF-8"
	to := "lbh9311@163.com"
	sendUserName := "Robot"
	user := "apetonweb@163.com"
	subject := "hello"
	body := "hhh"
	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	fmt.Println("Start Sending ...")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
