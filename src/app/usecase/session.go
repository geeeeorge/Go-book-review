package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/pkg/cognito"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func (c *Client) Signup(ctx context.Context, user model.User) error {
	// cognitoでsignup
	viper.SetEnvPrefix("AWS_COGNITO")
	viper.AutomaticEnv()
	cognitoClient := cognito.New(viper.GetString("CLIENT_ID"), viper.GetString("USER_POOL_ID"))

	if err := cognitoClient.SignUp(ctx, user.Username, user.Password); err != nil {
		return errors.Wrap(err, "usecase Signup: failed")
	}

	// Userテーブルに追加
	if err := c.repository.InsertUser(ctx, user); err != nil {
		return errors.Wrap(err, "usecase Signup: failed")
	}

	return nil
}

func (c *Client) Login(ctx context.Context, user *model.User) (string, error) {
	// cognitoでlogin
	viper.SetEnvPrefix("AWS_COGNITO")
	viper.AutomaticEnv()
	cognitoClient := cognito.New(viper.GetString("CLIENT_ID"), viper.GetString("USER_POOL_ID"))

	token, err := cognitoClient.Login(ctx, user.Username, user.Password)
	if err != nil {
		return "", errors.Wrap(err, "usecase Login: failed")
	}

	return token, nil
}
