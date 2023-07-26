package emailService

import (
	"bytes"
	"fmt"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"gopkg.in/gomail.v2"
	"os"
	"reflect"
	"strconv"
)

func SendRegistration(toEmail string) {

	from := environment.Env.GetEnv("FROM_MAIL")
	password := environment.Env.GetEnv("PASSWORD")

	smtpHost := environment.Env.GetEnv("SMTP_HOST")
	smtpPort := environment.Env.GetEnv("SMTP_PORT")
	fmt.Println(reflect.TypeOf(password))

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		panic(err.Error())
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", "Регистрация")

	body, err := os.ReadFile(environment.Env.GetEnv("DIR_EMAIL_MESSAGES") + "/registration.html")
	if err != nil {
		panic(err.Error())
	}

	htmlMessage := bytes.NewBuffer(body).String()

	msg.SetBody("text/html", htmlMessage)

	dialer := gomail.NewDialer(smtpHost, port, from, password)

	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}

}
