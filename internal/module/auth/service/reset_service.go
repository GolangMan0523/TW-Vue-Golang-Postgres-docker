package service

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"encoding/hex"
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/validator"
	// "github.com/sendgrid/sendgrid-go"
  // "github.com/sendgrid/sendgrid-go/helpers/mail"
  // "github.com/HotPotatoC/twitter-clone/internal/common/config"
	// "github.com/pkg/errors"
)

type ResetInput struct {
	Email    string `json:"email" validate:"required,email"`
}

func (i ResetInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type ResetService interface {
	Execute(input ResetInput) (string, error)
}

type resetService struct {
	db database.Database
}

func NewResetService(db database.Database) ResetService {
	return resetService{db: db}
}

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func (s resetService) Execute(input ResetInput) (string, error) {

	//get user that this email belongs to

	var id int64
	var handle, email string


	//TODO: we should have a repository or at least a better DB service
	err := s.db.QueryRow("SELECT id, handle, email FROM users WHERE email = $1", input.Email).Scan(&id, &handle, &email)

	if err != nil {
		return "", errors.Wrap(err, "service.resetService.Execute")
	}

	//store a token for an email
	resetToken := GenerateSecureToken(16)

	fmt.Println(resetToken)

	//TODO move email stuff to service common class
	// m := mail.NewV3Mail()

  // address := "support@t2.social"
  // name := "T2"
  // e := mail.NewEmail(name, address)
  // m.SetFrom(e)

  // m.SetTemplateID("d-c6dcf1f72bdd4beeb15a9aa6c72fcd2c")

  // p := mail.NewPersonalization()
  // tos := []*mail.Email{
  //   mail.NewEmail("Example User", "jay.dee.force@gmail.com"),
  // }
  // p.AddTos(tos...)

	// sendgridApiKey := config.GetString("SENDGRID_API_KEY", "")

  // m.AddPersonalizations(p)  
  // request := sendgrid.GetRequest(sendgridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
  // request.Method = "POST"

  // var Body = mail.GetRequestBody(m)

  // request.Body = Body
  // response, err := sendgrid.API(request)

  // if err != nil {
  //   fmt.Println(err)
  // } else {
  //   fmt.Println(response.StatusCode)
  //   fmt.Println(response.Body)
  //   fmt.Println(response.Headers)
  // }

	return "here", nil
}
