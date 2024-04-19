package signupMailing

import (
	"log"

	signupEmailDTO "github.com/GoPicos-Mailing-Service/internal/modules/TriggerMailings/services/Signup/DTO"
	"gopkg.in/gomail.v2"
)

var APIKey = "mlsn.f580011bcedc8d95cf1210bed9b00b2125f5e6987d36e5613dc10a2d87474152"

func SendSignupEmail(dto signupEmailDTO.SignupEmailDTO) {
	user := "a79d1ddc680568"
	password := "fb72e39ce9b3f7"

	msg := gomail.NewMessage()

	msg.SetHeader("From", "servicesgopicos@gmail.com")
	msg.SetHeader("To", dto.Email)
	msg.SetHeader("Subject", "Testing message")
	msg.SetBody("text/plain", "Olá! Este é seu codigo para o cadastro no nosso app: "+dto.Token)

	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, user, password)

	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	} else {
		log.Println("Email sent")
	}

}
