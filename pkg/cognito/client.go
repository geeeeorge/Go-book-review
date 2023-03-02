package cognito

import (
	"context"
	"fmt"
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

// SignUp cognitoAPIを用いたサインアップと認証
func (c *Client) SignUp(ctx context.Context, id, password string) error {
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := cognitoidentityprovider.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
	)

	// sign up
	paramsForSignUp := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.ClientID),
		Username: aws.String(id),
		Password: aws.String(password),
	}
	res, err := client.SignUp(paramsForSignUp)
	fmt.Println(err)
	if err != nil {
		return errors.Wrap(err, "SignUp: failed to authenticate")
	}
	if res == nil {
		return errors.Wrap(ErrInvalidUserInfo, "SignUp: failed to sign up")
	}

	// 認証
	paramsForAuth := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(c.UserPoolID),
		Username:   aws.String(id),
	}
	_, err = client.AdminConfirmSignUp(paramsForAuth)
	if err != nil {
		return errors.Wrap(err, "SignUp: failed to authenticate")
	}

	return nil
}

// Login cognitoAPIを用いたログイン
func (c *Client) Login(ctx context.Context, id, password string) (string, error) {
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := cognitoidentityprovider.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
	)

	// login
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeAdminUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(id),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(c.ClientID),
		UserPoolId: aws.String(c.UserPoolID),
	}
	res, err := client.AdminInitiateAuth(params)
	if err != nil {
		return "", errors.Wrap(err, "Login: failed to authenticate")
	}
	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		return "", errors.Wrap(ErrInvalidUserInfo, "Login: failed to login")
	}

	return *res.AuthenticationResult.IdToken, nil
}
