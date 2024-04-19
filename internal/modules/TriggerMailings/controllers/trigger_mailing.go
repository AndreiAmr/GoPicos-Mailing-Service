package triggerMailling

import (
	signupMailing "github.com/GoPicos-Mailing-Service/internal/modules/TriggerMailings/services/Signup"
	signupEmailDTO "github.com/GoPicos-Mailing-Service/internal/modules/TriggerMailings/services/Signup/DTO"
)

type TriggerMaillingDTO struct {
	Email string
	Token string
}

func TriggerMailling(dto TriggerMaillingDTO) {
	signupMailing.SendSignupEmail(signupEmailDTO.SignupEmailDTO{
		Email: dto.Email,
		Token: dto.Token,
	})
}
