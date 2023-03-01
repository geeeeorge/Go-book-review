package cognito

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/pkg/errors"
)

type AuthorizationKey string

const (
	ContextAuthorizationKey AuthorizationKey = "jwt"
)

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidUserInfo = errors.New("invalid user info")
)

type Interface interface {
	SignUp(ctx context.Context, id, password string) error
	Login(ctx context.Context, id, password string) (string, error)
}

type Client struct {
	ClientID   string
	UserPoolID string
}

func New(clientID, userPoolID string) Interface {
	return &Client{
		ClientID:   clientID,
		UserPoolID: userPoolID,
	}
}

// SignUp cognitoAPIを用いたサインアップ
func (c *Client) SignUp(ctx context.Context, id, password string) error {
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	params := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.ClientID),
		Username: aws.String(id),
		Password: aws.String(password),
	}
	client := cognitoidentityprovider.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
	)

	res, err := client.SignUp(params)
	if err != nil {
		return errors.Wrap(err, "SignUp: failed to authenticate")
	}
	if res == nil || !*res.UserConfirmed {
		return errors.Wrap(ErrInvalidUserInfo, "SignUp: failed to sign up")
	}
	return nil
}

// Login cognitoAPIを用いたログイン
func (c *Client) Login(ctx context.Context, id, password string) (string, error) {
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(id),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(c.ClientID),
		UserPoolId: aws.String(c.UserPoolID),
	}
	client := cognitoidentityprovider.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
	)

	res, err := client.AdminInitiateAuth(params)
	if err != nil {
		return "", errors.Wrap(err, "Login: failed to authenticate")
	}
	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		return "", errors.Wrap(ErrInvalidUserInfo, "Login: failed to login")
	}
	return *res.AuthenticationResult.IdToken, nil
}
