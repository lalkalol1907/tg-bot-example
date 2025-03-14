package auth_service

import (
	lockservice "bot-test/pkg/lock-service"
	"bot-test/pkg/models"
	"context"
	"fmt"
	"github.com/zelenin/go-tdlib/client"
)

type ClientAuthorizer struct {
	TdlibParameters *client.SetTdlibParametersRequest
	PhoneNumber     chan string
	Code            chan string
	State           chan client.AuthorizationState
	Password        chan string
}

type IAuthService interface {
	Authorize(ctx context.Context, clientAuthorizerRaw interface{}, worker *models.Worker)
}

type AuthService struct {
	lockService *lockservice.LockService
}

func (s *AuthService) Authorize(ctx context.Context, clientAuthorizerRaw interface{}, worker *models.Worker) {
	clientAuthorizer := clientAuthorizerRaw.(ClientAuthorizer) // Костылище)

	lockKey := fmt.Sprintf("auth:%d", worker.OwnerId)

	err := s.lockService.SetLock(ctx, lockKey)
	if err != nil {
		// TODO:
	}
	defer s.lockService.RemoveLock(ctx, lockKey)

	// TODO: Тут перписать все на походы в бота (синкаем асинк, костыль, есть такое)
	for {
		select {
		case state, ok := <-clientAuthorizer.State:
			if !ok {
				return
			}

			switch state.AuthorizationStateType() {
			case client.TypeAuthorizationStateWaitPhoneNumber:
				fmt.Println("Enter phone number: ")
				var phoneNumber string
				fmt.Scanln(&phoneNumber)

				clientAuthorizer.PhoneNumber <- phoneNumber

			case client.TypeAuthorizationStateWaitCode:
				var code string

				fmt.Println("Enter code: ")
				fmt.Scanln(&code)

				clientAuthorizer.Code <- code

			case client.TypeAuthorizationStateWaitPassword:
				fmt.Println("Enter password: ")
				var password string
				fmt.Scanln(&password)

				clientAuthorizer.Password <- password

			case client.TypeAuthorizationStateReady:
				return
			}
		}
	}
}

func NewAuthService(lockService *lockservice.LockService) IAuthService {
	return &AuthService{lockService: lockService}
}
